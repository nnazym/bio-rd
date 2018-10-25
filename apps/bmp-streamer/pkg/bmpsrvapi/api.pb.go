// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package bmpsrvapi is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	AdjRIBInStreamRequest
	RIBUpdate
	Route
	Prefix
	IP
	Path
	BGPPath
	ASPathSegment
	LargeCommunity
	UnknownAttribute
*/
package bmpsrvapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdjRIBInStreamRequest struct {
	Router *IP `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
}

func (m *AdjRIBInStreamRequest) Reset()                    { *m = AdjRIBInStreamRequest{} }
func (m *AdjRIBInStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*AdjRIBInStreamRequest) ProtoMessage()               {}
func (*AdjRIBInStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AdjRIBInStreamRequest) GetRouter() *IP {
	if m != nil {
		return m.Router
	}
	return nil
}

type RIBUpdate struct {
	Peer          *IP    `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	Advertisement bool   `protobuf:"varint,2,opt,name=advertisement" json:"advertisement,omitempty"`
	Route         *Route `protobuf:"bytes,3,opt,name=route" json:"route,omitempty"`
}

func (m *RIBUpdate) Reset()                    { *m = RIBUpdate{} }
func (m *RIBUpdate) String() string            { return proto.CompactTextString(m) }
func (*RIBUpdate) ProtoMessage()               {}
func (*RIBUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RIBUpdate) GetPeer() *IP {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *RIBUpdate) GetAdvertisement() bool {
	if m != nil {
		return m.Advertisement
	}
	return false
}

func (m *RIBUpdate) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type Route struct {
	Pfx  *Prefix `protobuf:"bytes,1,opt,name=pfx" json:"pfx,omitempty"`
	Path *Path   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Route) GetPfx() *Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

func (m *Route) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type Prefix struct {
	Address *IP    `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Pfxlen  uint32 `protobuf:"varint,2,opt,name=pfxlen" json:"pfxlen,omitempty"`
}

func (m *Prefix) Reset()                    { *m = Prefix{} }
func (m *Prefix) String() string            { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()               {}
func (*Prefix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Prefix) GetAddress() *IP {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Prefix) GetPfxlen() uint32 {
	if m != nil {
		return m.Pfxlen
	}
	return 0
}

type IP struct {
	Higher    uint64 `protobuf:"varint,1,opt,name=higher" json:"higher,omitempty"`
	Lower     uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	IPVersion uint32 `protobuf:"varint,3,opt,name=IP_version,json=IPVersion" json:"IP_version,omitempty"`
}

func (m *IP) Reset()                    { *m = IP{} }
func (m *IP) String() string            { return proto.CompactTextString(m) }
func (*IP) ProtoMessage()               {}
func (*IP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IP) GetHigher() uint64 {
	if m != nil {
		return m.Higher
	}
	return 0
}

func (m *IP) GetLower() uint64 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *IP) GetIPVersion() uint32 {
	if m != nil {
		return m.IPVersion
	}
	return 0
}

type Path struct {
	Type    uint32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	BGPPath *BGPPath `protobuf:"bytes,2,opt,name=BGP_path,json=BGPPath" json:"BGP_path,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Path) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Path) GetBGPPath() *BGPPath {
	if m != nil {
		return m.BGPPath
	}
	return nil
}

type BGPPath struct {
	PathIdentifier    uint32              `protobuf:"varint,1,opt,name=path_identifier,json=pathIdentifier" json:"path_identifier,omitempty"`
	NextHop           *IP                 `protobuf:"bytes,2,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
	LocalPref         uint32              `protobuf:"varint,3,opt,name=local_pref,json=localPref" json:"local_pref,omitempty"`
	ASPath            []*ASPathSegment    `protobuf:"bytes,4,rep,name=AS_path,json=ASPath" json:"AS_path,omitempty"`
	Origin            uint32              `protobuf:"varint,5,opt,name=origin" json:"origin,omitempty"`
	MED               uint32              `protobuf:"varint,6,opt,name=MED" json:"MED,omitempty"`
	EBGP              bool                `protobuf:"varint,7,opt,name=EBGP" json:"EBGP,omitempty"`
	BGPIdentifier     uint32              `protobuf:"varint,8,opt,name=BGP_identifier,json=BGPIdentifier" json:"BGP_identifier,omitempty"`
	Source            *IP                 `protobuf:"bytes,9,opt,name=source" json:"source,omitempty"`
	Communities       []uint32            `protobuf:"varint,10,rep,packed,name=communities" json:"communities,omitempty"`
	LargeCommunities  []*LargeCommunity   `protobuf:"bytes,11,rep,name=large_communities,json=largeCommunities" json:"large_communities,omitempty"`
	OriginatorId      uint32              `protobuf:"varint,12,opt,name=originator_id,json=originatorId" json:"originator_id,omitempty"`
	ClusterList       []uint32            `protobuf:"varint,13,rep,packed,name=cluster_list,json=clusterList" json:"cluster_list,omitempty"`
	UnknownAttributes []*UnknownAttribute `protobuf:"bytes,14,rep,name=unknown_attributes,json=unknownAttributes" json:"unknown_attributes,omitempty"`
}

func (m *BGPPath) Reset()                    { *m = BGPPath{} }
func (m *BGPPath) String() string            { return proto.CompactTextString(m) }
func (*BGPPath) ProtoMessage()               {}
func (*BGPPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BGPPath) GetPathIdentifier() uint32 {
	if m != nil {
		return m.PathIdentifier
	}
	return 0
}

func (m *BGPPath) GetNextHop() *IP {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *BGPPath) GetLocalPref() uint32 {
	if m != nil {
		return m.LocalPref
	}
	return 0
}

func (m *BGPPath) GetASPath() []*ASPathSegment {
	if m != nil {
		return m.ASPath
	}
	return nil
}

func (m *BGPPath) GetOrigin() uint32 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *BGPPath) GetMED() uint32 {
	if m != nil {
		return m.MED
	}
	return 0
}

func (m *BGPPath) GetEBGP() bool {
	if m != nil {
		return m.EBGP
	}
	return false
}

func (m *BGPPath) GetBGPIdentifier() uint32 {
	if m != nil {
		return m.BGPIdentifier
	}
	return 0
}

func (m *BGPPath) GetSource() *IP {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *BGPPath) GetCommunities() []uint32 {
	if m != nil {
		return m.Communities
	}
	return nil
}

func (m *BGPPath) GetLargeCommunities() []*LargeCommunity {
	if m != nil {
		return m.LargeCommunities
	}
	return nil
}

func (m *BGPPath) GetOriginatorId() uint32 {
	if m != nil {
		return m.OriginatorId
	}
	return 0
}

func (m *BGPPath) GetClusterList() []uint32 {
	if m != nil {
		return m.ClusterList
	}
	return nil
}

func (m *BGPPath) GetUnknownAttributes() []*UnknownAttribute {
	if m != nil {
		return m.UnknownAttributes
	}
	return nil
}

type ASPathSegment struct {
	ASSequence bool     `protobuf:"varint,1,opt,name=AS_sequence,json=ASSequence" json:"AS_sequence,omitempty"`
	ASNs       []uint32 `protobuf:"varint,2,rep,packed,name=ASNs" json:"ASNs,omitempty"`
}

func (m *ASPathSegment) Reset()                    { *m = ASPathSegment{} }
func (m *ASPathSegment) String() string            { return proto.CompactTextString(m) }
func (*ASPathSegment) ProtoMessage()               {}
func (*ASPathSegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ASPathSegment) GetASSequence() bool {
	if m != nil {
		return m.ASSequence
	}
	return false
}

func (m *ASPathSegment) GetASNs() []uint32 {
	if m != nil {
		return m.ASNs
	}
	return nil
}

type LargeCommunity struct {
	GlobalAdministrator uint32 `protobuf:"varint,1,opt,name=global_administrator,json=globalAdministrator" json:"global_administrator,omitempty"`
	DataPart1           uint32 `protobuf:"varint,2,opt,name=data_part1,json=dataPart1" json:"data_part1,omitempty"`
	DataPart2           uint32 `protobuf:"varint,3,opt,name=data_part2,json=dataPart2" json:"data_part2,omitempty"`
}

func (m *LargeCommunity) Reset()                    { *m = LargeCommunity{} }
func (m *LargeCommunity) String() string            { return proto.CompactTextString(m) }
func (*LargeCommunity) ProtoMessage()               {}
func (*LargeCommunity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LargeCommunity) GetGlobalAdministrator() uint32 {
	if m != nil {
		return m.GlobalAdministrator
	}
	return 0
}

func (m *LargeCommunity) GetDataPart1() uint32 {
	if m != nil {
		return m.DataPart1
	}
	return 0
}

func (m *LargeCommunity) GetDataPart2() uint32 {
	if m != nil {
		return m.DataPart2
	}
	return 0
}

type UnknownAttribute struct {
	Optional   bool   `protobuf:"varint,1,opt,name=optional" json:"optional,omitempty"`
	Transitive bool   `protobuf:"varint,2,opt,name=transitive" json:"transitive,omitempty"`
	Partial    bool   `protobuf:"varint,3,opt,name=partial" json:"partial,omitempty"`
	TypeCode   uint32 `protobuf:"varint,4,opt,name=type_code,json=typeCode" json:"type_code,omitempty"`
	Value      []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UnknownAttribute) Reset()                    { *m = UnknownAttribute{} }
func (m *UnknownAttribute) String() string            { return proto.CompactTextString(m) }
func (*UnknownAttribute) ProtoMessage()               {}
func (*UnknownAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UnknownAttribute) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *UnknownAttribute) GetTransitive() bool {
	if m != nil {
		return m.Transitive
	}
	return false
}

func (m *UnknownAttribute) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *UnknownAttribute) GetTypeCode() uint32 {
	if m != nil {
		return m.TypeCode
	}
	return 0
}

func (m *UnknownAttribute) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*AdjRIBInStreamRequest)(nil), "bmpsrvapi.AdjRIBInStreamRequest")
	proto.RegisterType((*RIBUpdate)(nil), "bmpsrvapi.RIBUpdate")
	proto.RegisterType((*Route)(nil), "bmpsrvapi.Route")
	proto.RegisterType((*Prefix)(nil), "bmpsrvapi.Prefix")
	proto.RegisterType((*IP)(nil), "bmpsrvapi.IP")
	proto.RegisterType((*Path)(nil), "bmpsrvapi.Path")
	proto.RegisterType((*BGPPath)(nil), "bmpsrvapi.BGPPath")
	proto.RegisterType((*ASPathSegment)(nil), "bmpsrvapi.ASPathSegment")
	proto.RegisterType((*LargeCommunity)(nil), "bmpsrvapi.LargeCommunity")
	proto.RegisterType((*UnknownAttribute)(nil), "bmpsrvapi.UnknownAttribute")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RIBService service

type RIBServiceClient interface {
	AdjRIBInStream(ctx context.Context, in *AdjRIBInStreamRequest, opts ...grpc.CallOption) (RIBService_AdjRIBInStreamClient, error)
}

type rIBServiceClient struct {
	cc *grpc.ClientConn
}

func NewRIBServiceClient(cc *grpc.ClientConn) RIBServiceClient {
	return &rIBServiceClient{cc}
}

func (c *rIBServiceClient) AdjRIBInStream(ctx context.Context, in *AdjRIBInStreamRequest, opts ...grpc.CallOption) (RIBService_AdjRIBInStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RIBService_serviceDesc.Streams[0], c.cc, "/bmpsrvapi.RIBService/adjRIBInStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBServiceAdjRIBInStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBService_AdjRIBInStreamClient interface {
	Recv() (*RIBUpdate, error)
	grpc.ClientStream
}

type rIBServiceAdjRIBInStreamClient struct {
	grpc.ClientStream
}

func (x *rIBServiceAdjRIBInStreamClient) Recv() (*RIBUpdate, error) {
	m := new(RIBUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RIBService service

type RIBServiceServer interface {
	AdjRIBInStream(*AdjRIBInStreamRequest, RIBService_AdjRIBInStreamServer) error
}

func RegisterRIBServiceServer(s *grpc.Server, srv RIBServiceServer) {
	s.RegisterService(&_RIBService_serviceDesc, srv)
}

func _RIBService_AdjRIBInStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AdjRIBInStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBServiceServer).AdjRIBInStream(m, &rIBServiceAdjRIBInStreamServer{stream})
}

type RIBService_AdjRIBInStreamServer interface {
	Send(*RIBUpdate) error
	grpc.ServerStream
}

type rIBServiceAdjRIBInStreamServer struct {
	grpc.ServerStream
}

func (x *rIBServiceAdjRIBInStreamServer) Send(m *RIBUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _RIBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bmpsrvapi.RIBService",
	HandlerType: (*RIBServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "adjRIBInStream",
			Handler:       _RIBService_AdjRIBInStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x57, 0x1a, 0x37, 0x7f, 0x26, 0x75, 0xae, 0x1d, 0x0a, 0x32, 0x77, 0x02, 0x72, 0x2e, 0xe5,
	0xf2, 0x42, 0x45, 0xc3, 0x3b, 0x52, 0x72, 0x77, 0x14, 0x9f, 0x0e, 0xe4, 0x5b, 0xeb, 0x10, 0x6f,
	0xd6, 0x36, 0xde, 0x24, 0x0b, 0x8e, 0xd7, 0xec, 0xae, 0x73, 0xa9, 0xc4, 0x0b, 0x9f, 0x83, 0x2f,
	0xc2, 0xc7, 0x43, 0xbb, 0xb6, 0xcb, 0xa6, 0xea, 0xbd, 0xcd, 0xfc, 0x7e, 0x33, 0xb3, 0xf3, 0x9b,
	0x8c, 0x27, 0x30, 0xa4, 0x25, 0xbf, 0x2a, 0xa5, 0xd0, 0x02, 0x87, 0xb7, 0xdb, 0x52, 0xc9, 0x1d,
	0x2d, 0x79, 0xf8, 0x03, 0x7c, 0x3a, 0xcf, 0x7e, 0x27, 0xd1, 0x22, 0x2a, 0x12, 0x2d, 0x19, 0xdd,
	0x12, 0xf6, 0x67, 0xc5, 0x94, 0xc6, 0x4b, 0xe8, 0x49, 0x51, 0x69, 0x26, 0x83, 0xce, 0xa4, 0x33,
	0x1d, 0xcd, 0xfc, 0xab, 0xfb, 0xa4, 0xab, 0x28, 0x26, 0x0d, 0x19, 0xfe, 0x05, 0x43, 0x12, 0x2d,
	0xde, 0x97, 0x19, 0xd5, 0x0c, 0x9f, 0x83, 0x57, 0xb2, 0x8f, 0x65, 0x58, 0x0a, 0xbf, 0x06, 0x9f,
	0x66, 0x3b, 0x26, 0x35, 0x57, 0x6c, 0xcb, 0x0a, 0x1d, 0x1c, 0x4d, 0x3a, 0xd3, 0x01, 0x39, 0x04,
	0xf1, 0x1b, 0x38, 0xb6, 0xf5, 0x83, 0xae, 0xad, 0x74, 0xea, 0x54, 0x22, 0x06, 0x27, 0x35, 0x1d,
	0xbe, 0x83, 0x63, 0xeb, 0xe3, 0x05, 0x74, 0xcb, 0xd5, 0xbe, 0x79, 0xf8, 0xcc, 0x09, 0x8f, 0x25,
	0x5b, 0xf1, 0x3d, 0x31, 0x2c, 0x5e, 0x80, 0x57, 0x52, 0xbd, 0xb1, 0x4f, 0x8e, 0x66, 0x4f, 0xdc,
	0x28, 0xaa, 0x37, 0xc4, 0x92, 0x61, 0x04, 0xbd, 0x3a, 0x07, 0x5f, 0x40, 0x9f, 0x66, 0x99, 0x64,
	0x4a, 0x3d, 0x2e, 0xa8, 0x65, 0xf1, 0x33, 0xe8, 0x95, 0xab, 0x7d, 0xce, 0x0a, 0x5b, 0xd9, 0x27,
	0x8d, 0x17, 0xbe, 0x83, 0xa3, 0x28, 0x36, 0xec, 0x86, 0xaf, 0x37, 0xcd, 0x58, 0x3c, 0xd2, 0x78,
	0x78, 0x0e, 0xc7, 0xb9, 0xf8, 0xc0, 0xa4, 0x4d, 0xf2, 0x48, 0xed, 0xe0, 0x17, 0x00, 0x51, 0x9c,
	0xee, 0x98, 0x54, 0x5c, 0x14, 0x56, 0xbe, 0x4f, 0x86, 0x51, 0xfc, 0x6b, 0x0d, 0x84, 0x11, 0x78,
	0xa6, 0x57, 0x44, 0xf0, 0xf4, 0x5d, 0xc9, 0x6c, 0x49, 0x9f, 0x58, 0x1b, 0xbf, 0x85, 0xc1, 0xe2,
	0x26, 0x4e, 0x1d, 0x89, 0xe8, 0x34, 0xbc, 0xb8, 0x89, 0xad, 0xca, 0x7e, 0x63, 0x84, 0xff, 0x7a,
	0xd0, 0xda, 0xf8, 0x02, 0x9e, 0x98, 0xb4, 0x94, 0x67, 0xac, 0xd0, 0x7c, 0xc5, 0x9b, 0x66, 0x7d,
	0x32, 0x36, 0x70, 0x74, 0x8f, 0xe2, 0x14, 0x06, 0x05, 0xdb, 0xeb, 0x74, 0x23, 0xca, 0xe6, 0x8d,
	0x87, 0x43, 0x31, 0xf4, 0x4f, 0xa2, 0x34, 0x42, 0x72, 0xb1, 0xa4, 0x79, 0x5a, 0x4a, 0xb6, 0x6a,
	0x85, 0x58, 0xc4, 0x8c, 0x17, 0xaf, 0xa1, 0x3f, 0x4f, 0xea, 0x5e, 0xbd, 0x49, 0x77, 0x3a, 0x9a,
	0x05, 0x4e, 0x9d, 0x79, 0x62, 0xba, 0x4a, 0xd8, 0xda, 0x2c, 0x03, 0xe9, 0xd5, 0xae, 0x19, 0xa4,
	0x90, 0x7c, 0xcd, 0x8b, 0xe0, 0xb8, 0x1e, 0x73, 0xed, 0xe1, 0x29, 0x74, 0x7f, 0x7e, 0xfd, 0x2a,
	0xe8, 0x59, 0xd0, 0x98, 0x66, 0x3a, 0xaf, 0x17, 0x37, 0x71, 0xd0, 0xb7, 0xbb, 0x65, 0x6d, 0xbc,
	0x84, 0xb1, 0x99, 0x8e, 0xa3, 0x70, 0x60, 0x13, 0xfc, 0xc5, 0x4d, 0xec, 0x08, 0xbc, 0x84, 0x9e,
	0x12, 0x95, 0x5c, 0xb2, 0x60, 0xf8, 0xe8, 0xda, 0xd7, 0x24, 0x4e, 0x60, 0xb4, 0x14, 0xdb, 0x6d,
	0x55, 0x70, 0xcd, 0x99, 0x0a, 0x60, 0xd2, 0x9d, 0xfa, 0xc4, 0x85, 0xf0, 0x47, 0x38, 0xcb, 0xa9,
	0x5c, 0xb3, 0xd4, 0x8d, 0x1b, 0x59, 0xa9, 0x9f, 0x3b, 0x35, 0xdf, 0x9a, 0x98, 0x97, 0x4d, 0xc8,
	0x1d, 0x39, 0xcd, 0x5d, 0xdf, 0xd4, 0xb9, 0x00, 0xbf, 0xd6, 0x49, 0xb5, 0x90, 0x29, 0xcf, 0x82,
	0x13, 0xdb, 0xf6, 0xc9, 0xff, 0x60, 0x94, 0xe1, 0x73, 0x38, 0x59, 0xe6, 0x95, 0xd2, 0x4c, 0xa6,
	0x39, 0x57, 0x3a, 0xf0, 0x9b, 0x7e, 0x6a, 0xec, 0x2d, 0x57, 0x1a, 0xdf, 0x00, 0x56, 0xc5, 0x1f,
	0x85, 0xf8, 0x50, 0xa4, 0x54, 0x6b, 0xc9, 0x6f, 0x2b, 0xcd, 0x54, 0x30, 0xb6, 0x0d, 0x3d, 0x73,
	0x1a, 0x7a, 0x5f, 0x07, 0xcd, 0xdb, 0x18, 0x72, 0x56, 0x3d, 0x40, 0x54, 0xf8, 0x0a, 0xfc, 0x83,
	0x9f, 0x08, 0xbf, 0x82, 0xd1, 0x3c, 0x49, 0x95, 0x39, 0x1d, 0xc5, 0xb2, 0xde, 0xca, 0x01, 0x81,
	0x79, 0x92, 0x34, 0x88, 0xf9, 0x45, 0xe6, 0xc9, 0x2f, 0x2a, 0x38, 0xb2, 0x8d, 0x59, 0x3b, 0xfc,
	0xbb, 0x03, 0xe3, 0x43, 0xf9, 0x78, 0x0d, 0xe7, 0xeb, 0x5c, 0xdc, 0xd2, 0x3c, 0xa5, 0xd9, 0x96,
	0x17, 0x5c, 0x69, 0x69, 0x14, 0x36, 0xcb, 0xf8, 0x49, 0xcd, 0xcd, 0x5d, 0xca, 0xec, 0x59, 0x46,
	0x35, 0x4d, 0x4b, 0x2a, 0xf5, 0x75, 0xf3, 0x01, 0x0e, 0x0d, 0x12, 0x1b, 0xe0, 0x80, 0x9e, 0xb5,
	0x6b, 0xd8, 0xd2, 0xb3, 0xf0, 0x9f, 0x0e, 0x9c, 0x3e, 0x54, 0x8c, 0x4f, 0x61, 0x20, 0x4a, 0xcd,
	0x45, 0x41, 0xf3, 0x46, 0xca, 0xbd, 0x8f, 0x5f, 0x02, 0x68, 0x49, 0x0b, 0xc5, 0x35, 0xdf, 0xb1,
	0xe6, 0x78, 0x39, 0x08, 0x06, 0xd0, 0x37, 0x4f, 0x71, 0x9a, 0xdb, 0xc7, 0x06, 0xa4, 0x75, 0xf1,
	0x19, 0x0c, 0xcd, 0x67, 0x9a, 0x2e, 0x45, 0xc6, 0x02, 0xcf, 0x36, 0x32, 0x30, 0xc0, 0x4b, 0x91,
	0x31, 0x73, 0x0c, 0x76, 0x34, 0xaf, 0x98, 0x5d, 0xed, 0x13, 0x52, 0x3b, 0xb3, 0xdf, 0x00, 0x48,
	0xb4, 0x48, 0x98, 0xdc, 0xf1, 0x25, 0xc3, 0x37, 0x30, 0xa6, 0x07, 0xa7, 0x1a, 0x27, 0xee, 0x37,
	0xf3, 0xd8, 0x15, 0x7f, 0x7a, 0xee, 0x5e, 0xce, 0xf6, 0x4e, 0x7f, 0xd7, 0xb9, 0xed, 0xd9, 0x3f,
	0x82, 0xef, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xea, 0xf0, 0x86, 0x15, 0x06, 0x00, 0x00,
}