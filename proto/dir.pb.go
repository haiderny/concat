// Code generated by protoc-gen-gogo.
// source: dir.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	dir.proto
	node.proto
	stmt.proto

It has these top-level messages:
	PeerInfo
	PublisherInfo
	RegisterPeer
	LookupPeerRequest
	LookupPeerResponse
	ListPeersRequest
	ListPeersResponse
	StreamEnd
	StreamError
	NodeInfoRequest
	NodeInfo
	Ping
	Pong
	QueryRequest
	QueryResult
	QueryResultValue
	SimpleValue
	CompoundValue
	KeyValuePair
	DataRequest
	DataResult
	DataObject
	PushRequest
	PushResponse
	PushAccept
	PushReject
	PushValue
	PushEnd
	Statement
	StatementBody
	SimpleStatement
	CompoundStatement
	EnvelopeStatement
	ArchiveStatement
*/
package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PeerInfo struct {
	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr [][]byte `protobuf:"bytes,2,rep,name=addr" json:"addr,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto1.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{0} }

type PublisherInfo struct {
	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces" json:"namespaces,omitempty"`
}

func (m *PublisherInfo) Reset()                    { *m = PublisherInfo{} }
func (m *PublisherInfo) String() string            { return proto1.CompactTextString(m) }
func (*PublisherInfo) ProtoMessage()               {}
func (*PublisherInfo) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{1} }

// /mediachain/dir/register
type RegisterPeer struct {
	Info      *PeerInfo      `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Publisher *PublisherInfo `protobuf:"bytes,2,opt,name=publisher" json:"publisher,omitempty"`
}

func (m *RegisterPeer) Reset()                    { *m = RegisterPeer{} }
func (m *RegisterPeer) String() string            { return proto1.CompactTextString(m) }
func (*RegisterPeer) ProtoMessage()               {}
func (*RegisterPeer) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{2} }

func (m *RegisterPeer) GetInfo() *PeerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *RegisterPeer) GetPublisher() *PublisherInfo {
	if m != nil {
		return m.Publisher
	}
	return nil
}

// /mediachain/dir/lookup
type LookupPeerRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *LookupPeerRequest) Reset()                    { *m = LookupPeerRequest{} }
func (m *LookupPeerRequest) String() string            { return proto1.CompactTextString(m) }
func (*LookupPeerRequest) ProtoMessage()               {}
func (*LookupPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{3} }

type LookupPeerResponse struct {
	Peer *PeerInfo `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
}

func (m *LookupPeerResponse) Reset()                    { *m = LookupPeerResponse{} }
func (m *LookupPeerResponse) String() string            { return proto1.CompactTextString(m) }
func (*LookupPeerResponse) ProtoMessage()               {}
func (*LookupPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{4} }

func (m *LookupPeerResponse) GetPeer() *PeerInfo {
	if m != nil {
		return m.Peer
	}
	return nil
}

// /mediachain/dir/list
type ListPeersRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *ListPeersRequest) Reset()                    { *m = ListPeersRequest{} }
func (m *ListPeersRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListPeersRequest) ProtoMessage()               {}
func (*ListPeersRequest) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{5} }

type ListPeersResponse struct {
	Peers []string `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *ListPeersResponse) Reset()                    { *m = ListPeersResponse{} }
func (m *ListPeersResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListPeersResponse) ProtoMessage()               {}
func (*ListPeersResponse) Descriptor() ([]byte, []int) { return fileDescriptorDir, []int{6} }

func init() {
	proto1.RegisterType((*PeerInfo)(nil), "proto.PeerInfo")
	proto1.RegisterType((*PublisherInfo)(nil), "proto.PublisherInfo")
	proto1.RegisterType((*RegisterPeer)(nil), "proto.RegisterPeer")
	proto1.RegisterType((*LookupPeerRequest)(nil), "proto.LookupPeerRequest")
	proto1.RegisterType((*LookupPeerResponse)(nil), "proto.LookupPeerResponse")
	proto1.RegisterType((*ListPeersRequest)(nil), "proto.ListPeersRequest")
	proto1.RegisterType((*ListPeersResponse)(nil), "proto.ListPeersResponse")
}

func init() { proto1.RegisterFile("dir.proto", fileDescriptorDir) }

var fileDescriptorDir = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0xd0, 0x22, 0xfc, 0x52, 0x3e, 0x6a, 0x75, 0xc8, 0x80, 0x50, 0xe5, 0x2e, 0x65,
	0x89, 0x50, 0x99, 0x98, 0x98, 0x91, 0x3a, 0x54, 0xfe, 0x07, 0x29, 0xbe, 0x16, 0x0b, 0x88, 0x8d,
	0x2f, 0xf9, 0xff, 0xc8, 0x6e, 0x43, 0xa3, 0x4a, 0x99, 0x12, 0xdf, 0xbd, 0x1f, 0x8f, 0x0d, 0x61,
	0x6c, 0x28, 0x7d, 0x70, 0x8d, 0x93, 0xe3, 0xf4, 0x51, 0x25, 0xae, 0x36, 0x44, 0xe1, 0xbd, 0xde,
	0x39, 0x79, 0x8b, 0xdc, 0x9a, 0x22, 0x9b, 0x67, 0x4b, 0xa1, 0x73, 0x6b, 0xa4, 0xc4, 0xa8, 0x32,
	0x26, 0x14, 0xf9, 0xfc, 0x62, 0x39, 0xd1, 0xe9, 0x5f, 0xbd, 0xe1, 0x66, 0xd3, 0x6e, 0xbf, 0x2d,
	0x7f, 0x0e, 0x98, 0x1e, 0x81, 0xba, 0xfa, 0x21, 0xf6, 0xd5, 0x07, 0x71, 0xb2, 0x0a, 0xdd, 0x9b,
	0xa8, 0x3d, 0x26, 0x9a, 0xf6, 0x96, 0x1b, 0x0a, 0xb1, 0x58, 0x2e, 0x30, 0xb2, 0xf5, 0xce, 0xa5,
	0x84, 0xeb, 0xd5, 0xdd, 0x81, 0xae, 0xec, 0x98, 0x74, 0x5a, 0xca, 0x15, 0x84, 0xef, 0x5a, 0x8b,
	0x3c, 0x29, 0x67, 0x9d, 0xb2, 0x4f, 0xa3, 0x4f, 0x32, 0xb5, 0xc0, 0x74, 0xed, 0xdc, 0x57, 0xeb,
	0x63, 0x96, 0xa6, 0xdf, 0x96, 0xb8, 0x39, 0xa7, 0x55, 0xaf, 0x90, 0x7d, 0x11, 0x7b, 0x57, 0x33,
	0x45, 0x26, 0x4f, 0x14, 0x06, 0x99, 0xe2, 0x52, 0x3d, 0xe3, 0x7e, 0x6d, 0xb9, 0x89, 0x53, 0xee,
	0xe2, 0x1f, 0x20, 0xfe, 0xaf, 0x7a, 0x6c, 0x39, 0x0d, 0xd4, 0x13, 0xa6, 0x3d, 0xc7, 0xb1, 0x6b,
	0x86, 0x71, 0x8c, 0xe3, 0x22, 0x4b, 0x4f, 0x75, 0x38, 0x6c, 0x2f, 0x53, 0xe5, 0xcb, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x54, 0xb4, 0xba, 0xb1, 0x01, 0x00, 0x00,
}
