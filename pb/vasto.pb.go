// Code generated by protoc-gen-go.
// source: vasto.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	vasto.proto

It has these top-level messages:
	BalanceRequest
	StoreMessage
	ClientMessage
	Topology
	Ring
	StoreResource
	Location
	StoreHeartbeat
	ClientHeartbeat
	Empty
	Requests
	Responses
	Request
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	GetRequest
	GetResponse
	RangeGetRequest
	RangeGetResponse
	Response
	KeyValue
*/
package pb

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

// ////////////////////////////////////////////////
// 1. master received request to balance the data
type BalanceRequest struct {
	DataCenter string `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	StoreGroup string `protobuf:"bytes,2,opt,name=store_group,json=storeGroup" json:"store_group,omitempty"`
	StoreCount uint32 `protobuf:"varint,3,opt,name=store_count,json=storeCount" json:"store_count,omitempty"`
}

func (m *BalanceRequest) Reset()                    { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()               {}
func (*BalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BalanceRequest) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *BalanceRequest) GetStoreGroup() string {
	if m != nil {
		return m.StoreGroup
	}
	return ""
}

func (m *BalanceRequest) GetStoreCount() uint32 {
	if m != nil {
		return m.StoreCount
	}
	return 0
}

type StoreMessage struct {
}

func (m *StoreMessage) Reset()                    { *m = StoreMessage{} }
func (m *StoreMessage) String() string            { return proto.CompactTextString(m) }
func (*StoreMessage) ProtoMessage()               {}
func (*StoreMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClientMessage struct {
	Stores   []*StoreResource `protobuf:"bytes,1,rep,name=stores" json:"stores,omitempty"`
	IsDelete bool             `protobuf:"varint,2,opt,name=is_delete,json=isDelete" json:"is_delete,omitempty"`
}

func (m *ClientMessage) Reset()                    { *m = ClientMessage{} }
func (m *ClientMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()               {}
func (*ClientMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientMessage) GetStores() []*StoreResource {
	if m != nil {
		return m.Stores
	}
	return nil
}

func (m *ClientMessage) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

type Topology struct {
	Rings []*Ring `protobuf:"bytes,1,rep,name=rings" json:"rings,omitempty"`
}

func (m *Topology) Reset()                    { *m = Topology{} }
func (m *Topology) String() string            { return proto.CompactTextString(m) }
func (*Topology) ProtoMessage()               {}
func (*Topology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Topology) GetRings() []*Ring {
	if m != nil {
		return m.Rings
	}
	return nil
}

type Ring struct {
	DataCenter string           `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Stores     []*StoreResource `protobuf:"bytes,2,rep,name=stores" json:"stores,omitempty"`
}

func (m *Ring) Reset()                    { *m = Ring{} }
func (m *Ring) String() string            { return proto.CompactTextString(m) }
func (*Ring) ProtoMessage()               {}
func (*Ring) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ring) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Ring) GetStores() []*StoreResource {
	if m != nil {
		return m.Stores
	}
	return nil
}

type StoreResource struct {
	Id       int32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *StoreResource) Reset()                    { *m = StoreResource{} }
func (m *StoreResource) String() string            { return proto.CompactTextString(m) }
func (*StoreResource) ProtoMessage()               {}
func (*StoreResource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StoreResource) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StoreResource) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	DataCenter string `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Location) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Location) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// ////////////////////////////////////////////////
type StoreHeartbeat struct {
	DataCenter string         `protobuf:"bytes,1,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	Store      *StoreResource `protobuf:"bytes,2,opt,name=store" json:"store,omitempty"`
}

func (m *StoreHeartbeat) Reset()                    { *m = StoreHeartbeat{} }
func (m *StoreHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*StoreHeartbeat) ProtoMessage()               {}
func (*StoreHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StoreHeartbeat) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *StoreHeartbeat) GetStore() *StoreResource {
	if m != nil {
		return m.Store
	}
	return nil
}

type ClientHeartbeat struct {
	Location *Location `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *ClientHeartbeat) Reset()                    { *m = ClientHeartbeat{} }
func (m *ClientHeartbeat) String() string            { return proto.CompactTextString(m) }
func (*ClientHeartbeat) ProtoMessage()               {}
func (*ClientHeartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ClientHeartbeat) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// ////////////////////////////////////////////////
type Requests struct {
	Requests []*Request `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *Requests) Reset()                    { *m = Requests{} }
func (m *Requests) String() string            { return proto.CompactTextString(m) }
func (*Requests) ProtoMessage()               {}
func (*Requests) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Requests) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

type Responses struct {
	Responses []*Response `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *Responses) Reset()                    { *m = Responses{} }
func (m *Responses) String() string            { return proto.CompactTextString(m) }
func (*Responses) ProtoMessage()               {}
func (*Responses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Responses) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Request struct {
	Put      *PutRequest      `protobuf:"bytes,1,opt,name=put" json:"put,omitempty"`
	Get      *GetRequest      `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	RangeGet *RangeGetRequest `protobuf:"bytes,3,opt,name=range_get,json=rangeGet" json:"range_get,omitempty"`
	Delete   *DeleteRequest   `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Request) GetPut() *PutRequest {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Request) GetGet() *GetRequest {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetRangeGet() *RangeGetRequest {
	if m != nil {
		return m.RangeGet
	}
	return nil
}

func (m *Request) GetDelete() *DeleteRequest {
	if m != nil {
		return m.Delete
	}
	return nil
}

type PutRequest struct {
	KeyValue    *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
	TimestampNs int64     `protobuf:"varint,2,opt,name=timestamp_ns,json=timestampNs" json:"timestamp_ns,omitempty"`
	TtlMs       uint32    `protobuf:"varint,3,opt,name=ttl_ms,json=ttlMs" json:"ttl_ms,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PutRequest) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *PutRequest) GetTimestampNs() int64 {
	if m != nil {
		return m.TimestampNs
	}
	return 0
}

func (m *PutRequest) GetTtlMs() uint32 {
	if m != nil {
		return m.TtlMs
	}
	return 0
}

type PutResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PutResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *PutResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type DeleteResponse struct {
	Ok     bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *DeleteResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Ok       bool      `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	KeyValue *KeyValue `protobuf:"bytes,2,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GetResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetResponse) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type RangeGetRequest struct {
	Prefix   []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	StartKey []byte `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *RangeGetRequest) Reset()                    { *m = RangeGetRequest{} }
func (m *RangeGetRequest) String() string            { return proto.CompactTextString(m) }
func (*RangeGetRequest) ProtoMessage()               {}
func (*RangeGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *RangeGetRequest) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *RangeGetRequest) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *RangeGetRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RangeGetResponse struct {
	KeyValues []*KeyValue `protobuf:"bytes,1,rep,name=key_values,json=keyValues" json:"key_values,omitempty"`
}

func (m *RangeGetResponse) Reset()                    { *m = RangeGetResponse{} }
func (m *RangeGetResponse) String() string            { return proto.CompactTextString(m) }
func (*RangeGetResponse) ProtoMessage()               {}
func (*RangeGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RangeGetResponse) GetKeyValues() []*KeyValue {
	if m != nil {
		return m.KeyValues
	}
	return nil
}

type Response struct {
	Put      *PutResponse      `protobuf:"bytes,1,opt,name=put" json:"put,omitempty"`
	Get      *GetResponse      `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	RangeGet *RangeGetResponse `protobuf:"bytes,3,opt,name=range_get,json=rangeGet" json:"range_get,omitempty"`
	Delete   *DeleteResponse   `protobuf:"bytes,4,opt,name=delete" json:"delete,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *Response) GetPut() *PutResponse {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Response) GetGet() *GetResponse {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Response) GetRangeGet() *RangeGetResponse {
	if m != nil {
		return m.RangeGet
	}
	return nil
}

func (m *Response) GetDelete() *DeleteResponse {
	if m != nil {
		return m.Delete
	}
	return nil
}

type KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*BalanceRequest)(nil), "pb.BalanceRequest")
	proto.RegisterType((*StoreMessage)(nil), "pb.StoreMessage")
	proto.RegisterType((*ClientMessage)(nil), "pb.ClientMessage")
	proto.RegisterType((*Topology)(nil), "pb.Topology")
	proto.RegisterType((*Ring)(nil), "pb.Ring")
	proto.RegisterType((*StoreResource)(nil), "pb.StoreResource")
	proto.RegisterType((*Location)(nil), "pb.Location")
	proto.RegisterType((*StoreHeartbeat)(nil), "pb.StoreHeartbeat")
	proto.RegisterType((*ClientHeartbeat)(nil), "pb.ClientHeartbeat")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Requests)(nil), "pb.Requests")
	proto.RegisterType((*Responses)(nil), "pb.Responses")
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*PutRequest)(nil), "pb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "pb.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
	proto.RegisterType((*RangeGetRequest)(nil), "pb.RangeGetRequest")
	proto.RegisterType((*RangeGetResponse)(nil), "pb.RangeGetResponse")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*KeyValue)(nil), "pb.KeyValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VastoMaster service

type VastoMasterClient interface {
	RegisterStore(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterStoreClient, error)
	RegisterClient(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterClientClient, error)
}

type vastoMasterClient struct {
	cc *grpc.ClientConn
}

func NewVastoMasterClient(cc *grpc.ClientConn) VastoMasterClient {
	return &vastoMasterClient{cc}
}

func (c *vastoMasterClient) RegisterStore(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterStoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VastoMaster_serviceDesc.Streams[0], c.cc, "/pb.VastoMaster/RegisterStore", opts...)
	if err != nil {
		return nil, err
	}
	x := &vastoMasterRegisterStoreClient{stream}
	return x, nil
}

type VastoMaster_RegisterStoreClient interface {
	Send(*StoreHeartbeat) error
	Recv() (*StoreMessage, error)
	grpc.ClientStream
}

type vastoMasterRegisterStoreClient struct {
	grpc.ClientStream
}

func (x *vastoMasterRegisterStoreClient) Send(m *StoreHeartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vastoMasterRegisterStoreClient) Recv() (*StoreMessage, error) {
	m := new(StoreMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vastoMasterClient) RegisterClient(ctx context.Context, opts ...grpc.CallOption) (VastoMaster_RegisterClientClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VastoMaster_serviceDesc.Streams[1], c.cc, "/pb.VastoMaster/RegisterClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &vastoMasterRegisterClientClient{stream}
	return x, nil
}

type VastoMaster_RegisterClientClient interface {
	Send(*ClientHeartbeat) error
	Recv() (*ClientMessage, error)
	grpc.ClientStream
}

type vastoMasterRegisterClientClient struct {
	grpc.ClientStream
}

func (x *vastoMasterRegisterClientClient) Send(m *ClientHeartbeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vastoMasterRegisterClientClient) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for VastoMaster service

type VastoMasterServer interface {
	RegisterStore(VastoMaster_RegisterStoreServer) error
	RegisterClient(VastoMaster_RegisterClientServer) error
}

func RegisterVastoMasterServer(s *grpc.Server, srv VastoMasterServer) {
	s.RegisterService(&_VastoMaster_serviceDesc, srv)
}

func _VastoMaster_RegisterStore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VastoMasterServer).RegisterStore(&vastoMasterRegisterStoreServer{stream})
}

type VastoMaster_RegisterStoreServer interface {
	Send(*StoreMessage) error
	Recv() (*StoreHeartbeat, error)
	grpc.ServerStream
}

type vastoMasterRegisterStoreServer struct {
	grpc.ServerStream
}

func (x *vastoMasterRegisterStoreServer) Send(m *StoreMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vastoMasterRegisterStoreServer) Recv() (*StoreHeartbeat, error) {
	m := new(StoreHeartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VastoMaster_RegisterClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VastoMasterServer).RegisterClient(&vastoMasterRegisterClientServer{stream})
}

type VastoMaster_RegisterClientServer interface {
	Send(*ClientMessage) error
	Recv() (*ClientHeartbeat, error)
	grpc.ServerStream
}

type vastoMasterRegisterClientServer struct {
	grpc.ServerStream
}

func (x *vastoMasterRegisterClientServer) Send(m *ClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vastoMasterRegisterClientServer) Recv() (*ClientHeartbeat, error) {
	m := new(ClientHeartbeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VastoMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VastoMaster",
	HandlerType: (*VastoMasterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterStore",
			Handler:       _VastoMaster_RegisterStore_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RegisterClient",
			Handler:       _VastoMaster_RegisterClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vasto.proto",
}

// Client API for VastoStore service

type VastoStoreClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type vastoStoreClient struct {
	cc *grpc.ClientConn
}

func NewVastoStoreClient(cc *grpc.ClientConn) VastoStoreClient {
	return &vastoStoreClient{cc}
}

func (c *vastoStoreClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/pb.VastoStore/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VastoStore service

type VastoStoreServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterVastoStoreServer(s *grpc.Server, srv VastoStoreServer) {
	s.RegisterService(&_VastoStore_serviceDesc, srv)
}

func _VastoStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VastoStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VastoStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VastoStoreServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VastoStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VastoStore",
	HandlerType: (*VastoStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _VastoStore_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vasto.proto",
}

func init() { proto.RegisterFile("vasto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xed, 0x4e, 0xdb, 0x30,
	0x14, 0x25, 0x29, 0x85, 0xf4, 0xa6, 0x1f, 0xe0, 0xb1, 0xa9, 0x62, 0x12, 0x6b, 0xfd, 0x87, 0xc2,
	0x24, 0xc4, 0x8a, 0xf6, 0x21, 0x21, 0x6d, 0xd2, 0x18, 0x82, 0x89, 0x31, 0x21, 0x6f, 0x62, 0xd2,
	0x34, 0x29, 0x72, 0x5b, 0x2f, 0x8a, 0x9a, 0x26, 0x21, 0x76, 0xd0, 0xfa, 0x0a, 0x7b, 0x94, 0xfd,
	0xd9, 0x2b, 0x4e, 0x76, 0xec, 0xa4, 0x2d, 0x0c, 0xb6, 0x7f, 0xb9, 0xc7, 0xc7, 0xd7, 0xc7, 0xf7,
	0x1e, 0xdf, 0x80, 0x7b, 0x4d, 0xb9, 0x88, 0xf7, 0x92, 0x34, 0x16, 0x31, 0xb2, 0x93, 0x01, 0xe6,
	0xd0, 0x7c, 0x4b, 0x43, 0x1a, 0x0d, 0x19, 0x61, 0x57, 0x19, 0xe3, 0x02, 0x3d, 0x01, 0x77, 0x44,
	0x05, 0xf5, 0x86, 0x2c, 0x12, 0x2c, 0x6d, 0x5b, 0x1d, 0xab, 0x57, 0x23, 0x20, 0xa1, 0x23, 0x85,
	0x48, 0x02, 0x17, 0x71, 0xca, 0x3c, 0x3f, 0x8d, 0xb3, 0xa4, 0x6d, 0xe7, 0x04, 0x05, 0x9d, 0x48,
	0xa4, 0x24, 0x0c, 0xe3, 0x2c, 0x12, 0xed, 0x4a, 0xc7, 0xea, 0x35, 0x34, 0xe1, 0x48, 0x22, 0xb8,
	0x09, 0xf5, 0x4f, 0x32, 0x3a, 0x67, 0x9c, 0x53, 0x9f, 0xe1, 0x2f, 0xd0, 0x38, 0x0a, 0x03, 0x16,
	0x09, 0x0d, 0xa0, 0x1d, 0x58, 0x51, 0x74, 0xde, 0xb6, 0x3a, 0x95, 0x9e, 0xdb, 0x5f, 0xdf, 0x4b,
	0x06, 0x7b, 0x6a, 0x0b, 0x61, 0x3c, 0xce, 0xd2, 0x21, 0x23, 0x9a, 0x80, 0x1e, 0x43, 0x2d, 0xe0,
	0xde, 0x88, 0x85, 0x4c, 0x30, 0xa5, 0xc5, 0x21, 0x4e, 0xc0, 0xdf, 0xa9, 0x18, 0xef, 0x82, 0xf3,
	0x39, 0x4e, 0xe2, 0x30, 0xf6, 0xa7, 0x68, 0x0b, 0xaa, 0x69, 0x10, 0xf9, 0x26, 0xa5, 0x23, 0x53,
	0x92, 0x20, 0xf2, 0x49, 0x0e, 0x63, 0x02, 0xcb, 0x32, 0xbc, 0xff, 0xfe, 0xa5, 0x38, 0xfb, 0x1e,
	0x71, 0xf8, 0x3d, 0x34, 0xe6, 0x16, 0x50, 0x13, 0xec, 0x60, 0xa4, 0x72, 0x56, 0x89, 0x1d, 0x8c,
	0x50, 0x0f, 0x9c, 0x30, 0x1e, 0x52, 0x11, 0xc4, 0x91, 0x12, 0xef, 0xf6, 0xeb, 0x32, 0xdb, 0x07,
	0x8d, 0x91, 0x62, 0x15, 0x1f, 0x83, 0x63, 0xd0, 0xfb, 0x25, 0xb6, 0x61, 0x95, 0x8e, 0x46, 0x29,
	0xe3, 0x5c, 0xb7, 0xc7, 0x84, 0xf8, 0x2b, 0x34, 0x95, 0xa2, 0x53, 0x46, 0x53, 0x31, 0x60, 0xf4,
	0x1f, 0xfa, 0xbd, 0x0d, 0x55, 0x75, 0x1d, 0x2d, 0xf0, 0x96, 0xeb, 0xe6, 0xeb, 0xf8, 0x10, 0x5a,
	0x79, 0x1b, 0xcb, 0xe4, 0xb3, 0xf7, 0xb3, 0xee, 0xbc, 0xdf, 0x2a, 0x54, 0x8f, 0x27, 0x89, 0x98,
	0xe2, 0x03, 0x70, 0xb4, 0x15, 0x39, 0xda, 0x06, 0x27, 0xd5, 0xdf, 0xba, 0x6d, 0xae, 0x6a, 0x5b,
	0x8e, 0x91, 0x62, 0x11, 0xbf, 0x84, 0x1a, 0x61, 0x3c, 0x89, 0x23, 0xce, 0x38, 0xda, 0x85, 0x5a,
	0x6a, 0x02, 0xbd, 0xad, 0x9e, 0x6f, 0xcb, 0x41, 0x52, 0x2e, 0xe3, 0x5f, 0x16, 0xac, 0x1a, 0xe7,
	0x77, 0xa0, 0x92, 0x64, 0x42, 0xeb, 0x6c, 0xca, 0x1d, 0x17, 0x99, 0x30, 0x67, 0xc9, 0x25, 0xc9,
	0xf0, 0x99, 0xd0, 0x85, 0x50, 0x8c, 0x13, 0x56, 0x32, 0x7c, 0x26, 0xd0, 0x3e, 0xd4, 0x52, 0x1a,
	0xf9, 0xcc, 0x93, 0xbc, 0x8a, 0xe2, 0x3d, 0x50, 0x67, 0x4b, 0x70, 0x86, 0xec, 0xa4, 0x1a, 0x90,
	0x76, 0xd2, 0xee, 0x5d, 0x2e, 0xeb, 0x9b, 0xfb, 0xd7, 0x90, 0x35, 0x01, 0x5f, 0x01, 0x94, 0x8a,
	0xd0, 0x0e, 0xd4, 0xc6, 0x6c, 0xea, 0x5d, 0xd3, 0x30, 0x63, 0xb3, 0xc5, 0x3d, 0x63, 0xd3, 0x4b,
	0x89, 0x11, 0x67, 0xac, 0xbf, 0x50, 0x17, 0xea, 0x22, 0x98, 0x30, 0x2e, 0xe8, 0x24, 0xf1, 0xa2,
	0xdc, 0x14, 0x15, 0xe2, 0x16, 0xd8, 0x47, 0x8e, 0x1e, 0xc2, 0x8a, 0x10, 0xa1, 0x37, 0xe1, 0xfa,
	0xbd, 0x56, 0x85, 0x08, 0xcf, 0x39, 0x7e, 0x0e, 0xae, 0x3a, 0x32, 0xaf, 0x97, 0xf4, 0x6f, 0x3c,
	0x56, 0x87, 0x39, 0xc4, 0x8e, 0xc7, 0xe8, 0x91, 0x7c, 0x0b, 0x54, 0x64, 0xc6, 0x67, 0x3a, 0xc2,
	0x5d, 0x68, 0xcc, 0x5d, 0x01, 0xad, 0x41, 0x65, 0xcc, 0xa6, 0x6a, 0x67, 0x9d, 0xc8, 0x4f, 0xfc,
	0x0a, 0x9a, 0x86, 0xf2, 0x9f, 0xc9, 0xb7, 0x00, 0xca, 0x4a, 0xde, 0x92, 0xf9, 0x14, 0x5c, 0xb5,
	0xfe, 0x97, 0xb4, 0x73, 0x75, 0xb3, 0xef, 0xaa, 0x1b, 0xfe, 0x06, 0xad, 0x85, 0xc6, 0x49, 0x51,
	0x49, 0xca, 0xbe, 0x07, 0x3f, 0xf4, 0x89, 0x3a, 0x92, 0x73, 0x88, 0x0b, 0x9a, 0x0a, 0x4f, 0x8a,
	0xb1, 0xd5, 0x92, 0xa3, 0x80, 0x33, 0x36, 0x45, 0x1b, 0x50, 0x0d, 0x83, 0x49, 0x60, 0x66, 0x61,
	0x1e, 0xe0, 0x37, 0xb0, 0x56, 0x66, 0xd7, 0x62, 0x9f, 0x02, 0x14, 0xe2, 0xe6, 0xcc, 0x5b, 0xa8,
	0xab, 0x19, 0x75, 0x1c, 0xff, 0xb6, 0xe4, 0x5b, 0xd1, 0x3b, 0xbb, 0xb3, 0xee, 0x6d, 0x15, 0xee,
	0xd5, 0x96, 0x57, 0xf6, 0xed, 0xce, 0xda, 0xb7, 0x55, 0xd8, 0xd7, 0x50, 0xa4, 0x7f, 0x9f, 0xdd,
	0xf4, 0xef, 0xc6, 0xbc, 0x7f, 0x35, 0xbb, 0x34, 0xf0, 0xee, 0x82, 0x81, 0xd1, 0xac, 0x81, 0x35,
	0xdb, 0x38, 0xb8, 0x0f, 0x8e, 0xb9, 0xc8, 0xcd, 0xc6, 0xc9, 0x32, 0x95, 0x5d, 0xa9, 0x93, 0x3c,
	0xe8, 0xff, 0xb4, 0xc0, 0xbd, 0x94, 0xbf, 0xad, 0x73, 0xca, 0xe5, 0x3c, 0x3a, 0x84, 0x06, 0x61,
	0x7e, 0x20, 0xbf, 0xd5, 0x18, 0x42, 0xa8, 0x98, 0x48, 0xc5, 0xe0, 0xd9, 0x5c, 0x2b, 0x30, 0xf3,
	0x93, 0x59, 0xea, 0x59, 0xfb, 0x16, 0x7a, 0x0d, 0x4d, 0xb3, 0x39, 0x9f, 0x55, 0x48, 0x3d, 0xcf,
	0x85, 0xb9, 0xb5, 0xb9, 0x5e, 0x82, 0x73, 0xfb, 0xfb, 0x2f, 0x00, 0x94, 0x96, 0xfc, 0xe4, 0x1e,
	0x54, 0x2e, 0x32, 0x81, 0x16, 0x66, 0xc5, 0xe6, 0x62, 0xf5, 0xf1, 0xd2, 0x60, 0x45, 0xfd, 0x72,
	0x0f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x86, 0xad, 0x4e, 0x81, 0x07, 0x00, 0x00,
}
