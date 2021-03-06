package mc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	p2p_crypto "github.com/libp2p/go-libp2p-crypto"
	multihash "github.com/multiformats/go-multihash"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

var (
	MalformedEntityId = errors.New("Malformed entity id")
	UnknownIdProvider = errors.New("Unknwon identity provider")
)

type EntityId struct {
	KeyId string `json:"keyId"` // public key multihash
	Key   []byte `json:"key"`   // marshalled public key
}

type EntityKeyNotFound struct {
	err string
}

func (e EntityKeyNotFound) Error() string {
	return e.err
}

func entityKeyNotFound(what string) error {
	err := fmt.Sprintf("Entity key not found: %s", what)
	return EntityKeyNotFound{err}
}

func CheckEntityId(entity string) error {
	ix := strings.Index(entity, ":")
	if ix < 0 {
		return MalformedEntityId
	}

	prov := entity[:ix]
	user := entity[ix+1:]

	if !urx.Match([]byte(user)) {
		return MalformedEntityId
	}

	_, ok := idProviders[prov]
	if !ok {
		return UnknownIdProvider
	}

	return nil
}

var urx *regexp.Regexp

func init() {
	rx, err := regexp.Compile("^[a-zA-Z0-9.-]+$")
	if err != nil {
		log.Fatal(err)
	}
	urx = rx
}

func LookupEntityKey(entity string, keyId string) (p2p_crypto.PubKey, error) {
	err := CheckEntityId(entity)
	if err != nil {
		return nil, err
	}

	ix := strings.Index(entity, ":")
	prov := entity[:ix]
	user := entity[ix+1:]

	lookup := idProviders[prov]
	return lookup(user, keyId)
}

type LookupKeyFunc func(user, keyId string) (p2p_crypto.PubKey, error)

func lookupBlockstack(user, keyId string) (p2p_crypto.PubKey, error) {
	khash, err := multihash.FromB58String(keyId)
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("blockstack", "lookup", user).Output()
	if err != nil {
		xerr, ok := err.(*exec.ExitError)
		if ok {
			return nil, fmt.Errorf("blockstack error: %s", string(xerr.Stderr))
		}
		return nil, err
	}

	var res map[string]interface{}
	err = json.Unmarshal(out, &res)
	if err != nil {
		return nil, err
	}

	prof, ok := res["profile"].(map[string]interface{})
	if !ok {
		return nil, entityKeyNotFound("Missing or malformed blockstack profile")
	}

	accts, ok := prof["account"].([]interface{})
	if !ok {
		return nil, entityKeyNotFound("Missing or malformed blockstack account list")
	}

	for _, acct := range accts {
		xacct, ok := acct.(map[string]interface{})
		if !ok {
			continue
		}

		svc, ok := xacct["service"].(string)
		if !ok {
			continue
		}

		if svc != "mediachain" {
			continue
		}

		key, ok := xacct["identifier"].(string)
		if !ok {
			break
		}

		return unmarshalEntityKey(key, khash)
	}

	return nil, entityKeyNotFound("No mediachain account in blockstack profile")
}

func lookupKeybase(user, keyId string) (p2p_crypto.PubKey, error) {
	khash, err := multihash.FromB58String(keyId)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.keybase.pub/mediachain.json", user)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch {
	case res.StatusCode == 404:
		return nil, entityKeyNotFound("Error retrieving mediachain id from keybase: 404 Not Found")

	case res.StatusCode != 200:
		return nil, fmt.Errorf("keybase error: %d %s", res.StatusCode, res.Status)
	}

	var pub EntityId
	err = json.NewDecoder(res.Body).Decode(&pub)
	if err != nil {
		return nil, err
	}

	if pub.KeyId != keyId {
		return nil, entityKeyNotFound("Key id mismatch")
	}

	return unmarshalEntityKeyBytes(pub.Key, khash)
}

var idProviders = map[string]LookupKeyFunc{
	"blockstack": lookupBlockstack,
	"keybase":    lookupKeybase,
}

func unmarshalEntityKey(key string, khash multihash.Multihash) (p2p_crypto.PubKey, error) {
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return unmarshalEntityKeyBytes(data, khash)
}

func unmarshalEntityKeyBytes(key []byte, khash multihash.Multihash) (p2p_crypto.PubKey, error) {
	hash, err := multihash.Sum(key, int(khash[0]), -1)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(hash, khash) {
		return nil, entityKeyNotFound("Key hash mismatch")
	}

	return p2p_crypto.UnmarshalPublicKey(key)
}
