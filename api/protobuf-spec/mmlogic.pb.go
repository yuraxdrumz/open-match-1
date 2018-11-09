// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mmlogic.proto

/*
Package mmlogic is a generated protocol buffer package.

It is generated from these files:
	mmlogic.proto

It has these top-level messages:
	Profile
	MatchObject
	Roster
	Filter
	Stats
	PlayerPool
	Player
	Result
	IlInput
	Timestamp
	ConnectionInfo
	Assignments
*/
package mmlogic

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

type Profile struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties string `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// When you send a Profile to the backendAPI, it looks to see if you populated
	// this field with protobuf-encoded PlayerPool objects containing valid the filters
	// objects.  If you did, they are used by OM.  If you didn't, the backendAPI
	// next looks in your properties blob at the key specified in the 'jsonkeys.pools'
	// config value from config/matchmaker_config.json - If it finds valid player
	// pool definitions at that key, it will try to unmarshal them into this field.
	// If you didn't specify valid player pools in either place, OM assumes you
	// know what you're doing and just leaves this unpopulatd.
	Pools []*PlayerPool `protobuf:"bytes,4,rep,name=pools" json:"pools,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetPools() []*PlayerPool {
	if m != nil {
		return m.Pools
	}
	return nil
}

// A MMF takes the Profile object above, and generates a MatchObject.
type MatchObject struct {
	Id         string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties string        `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Rosters    []*Roster     `protobuf:"bytes,3,rep,name=rosters" json:"rosters,omitempty"`
	Pools      []*PlayerPool `protobuf:"bytes,4,rep,name=pools" json:"pools,omitempty"`
}

func (m *MatchObject) Reset()                    { *m = MatchObject{} }
func (m *MatchObject) String() string            { return proto.CompactTextString(m) }
func (*MatchObject) ProtoMessage()               {}
func (*MatchObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MatchObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MatchObject) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *MatchObject) GetRosters() []*Roster {
	if m != nil {
		return m.Rosters
	}
	return nil
}

func (m *MatchObject) GetPools() []*PlayerPool {
	if m != nil {
		return m.Pools
	}
	return nil
}

// Data structure to hold a list of players in a match.
type Roster struct {
	Name    string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Players []*Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
}

func (m *Roster) Reset()                    { *m = Roster{} }
func (m *Roster) String() string            { return proto.CompactTextString(m) }
func (*Roster) ProtoMessage()               {}
func (*Roster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Roster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Roster) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

// A filter to apply to the player pool.
type Filter struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Attribute string `protobuf:"bytes,2,opt,name=attribute" json:"attribute,omitempty"`
	Maxv      int64  `protobuf:"varint,3,opt,name=maxv" json:"maxv,omitempty"`
	Minv      int64  `protobuf:"varint,4,opt,name=minv" json:"minv,omitempty"`
	Stats     *Stats `protobuf:"bytes,5,opt,name=stats" json:"stats,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetAttribute() string {
	if m != nil {
		return m.Attribute
	}
	return ""
}

func (m *Filter) GetMaxv() int64 {
	if m != nil {
		return m.Maxv
	}
	return 0
}

func (m *Filter) GetMinv() int64 {
	if m != nil {
		return m.Minv
	}
	return 0
}

func (m *Filter) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Stats struct {
	Count   int64   `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Elapsed float64 `protobuf:"fixed64,2,opt,name=elapsed" json:"elapsed,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Stats) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Stats) GetElapsed() float64 {
	if m != nil {
		return m.Elapsed
	}
	return 0
}

type PlayerPool struct {
	Name    string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
	Roster  *Roster   `protobuf:"bytes,3,opt,name=roster" json:"roster,omitempty"`
	Stats   *Stats    `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
}

func (m *PlayerPool) Reset()                    { *m = PlayerPool{} }
func (m *PlayerPool) String() string            { return proto.CompactTextString(m) }
func (*PlayerPool) ProtoMessage()               {}
func (*PlayerPool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlayerPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlayerPool) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *PlayerPool) GetRoster() *Roster {
	if m != nil {
		return m.Roster
	}
	return nil
}

func (m *PlayerPool) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Data structure for a profile to pass to the matchmaking function.
type Player struct {
	Id         string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties string              `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
	Pool       string              `protobuf:"bytes,3,opt,name=pool" json:"pool,omitempty"`
	Attributes []*Player_Attribute `protobuf:"bytes,4,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *Player) GetPool() string {
	if m != nil {
		return m.Pool
	}
	return ""
}

func (m *Player) GetAttributes() []*Player_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Player_Attribute struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *Player_Attribute) Reset()                    { *m = Player_Attribute{} }
func (m *Player_Attribute) String() string            { return proto.CompactTextString(m) }
func (*Player_Attribute) ProtoMessage()               {}
func (*Player_Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *Player_Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player_Attribute) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Simple message to return success/failure and error status.
type Result struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// IlInput is an empty message reserved for future use.
type IlInput struct {
}

func (m *IlInput) Reset()                    { *m = IlInput{} }
func (m *IlInput) String() string            { return proto.CompactTextString(m) }
func (*IlInput) ProtoMessage()               {}
func (*IlInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Epoch timestamp in seconds.
type Timestamp struct {
	Ts int64 `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Timestamp) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

// Simple message used to pass the connection string for the DGS to the player.
type ConnectionInfo struct {
	ConnectionString string `protobuf:"bytes,1,opt,name=connection_string,json=connectionString" json:"connection_string,omitempty"`
}

func (m *ConnectionInfo) Reset()                    { *m = ConnectionInfo{} }
func (m *ConnectionInfo) String() string            { return proto.CompactTextString(m) }
func (*ConnectionInfo) ProtoMessage()               {}
func (*ConnectionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ConnectionInfo) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type Assignments struct {
	Rosters        []*Roster       `protobuf:"bytes,1,rep,name=rosters" json:"rosters,omitempty"`
	ConnectionInfo *ConnectionInfo `protobuf:"bytes,2,opt,name=connection_info,json=connectionInfo" json:"connection_info,omitempty"`
}

func (m *Assignments) Reset()                    { *m = Assignments{} }
func (m *Assignments) String() string            { return proto.CompactTextString(m) }
func (*Assignments) ProtoMessage()               {}
func (*Assignments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Assignments) GetRosters() []*Roster {
	if m != nil {
		return m.Rosters
	}
	return nil
}

func (m *Assignments) GetConnectionInfo() *ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*MatchObject)(nil), "MatchObject")
	proto.RegisterType((*Roster)(nil), "Roster")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*Stats)(nil), "Stats")
	proto.RegisterType((*PlayerPool)(nil), "PlayerPool")
	proto.RegisterType((*Player)(nil), "Player")
	proto.RegisterType((*Player_Attribute)(nil), "Player.Attribute")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*IlInput)(nil), "IlInput")
	proto.RegisterType((*Timestamp)(nil), "Timestamp")
	proto.RegisterType((*ConnectionInfo)(nil), "ConnectionInfo")
	proto.RegisterType((*Assignments)(nil), "Assignments")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	//  Send GetProfile an profile with the ID field populated, it will return a
	//  'filled' one.
	//  Note: filters are assumed to have been checked for validity by the
	//  backendapi  when accepting a profile
	GetProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error)
	// CreateProposal  does the following:
	//  - adds all players in the Roster to the proposed player ignore list
	//  - writes the proposed match to the provided key
	//  - adds that key to the list of proposals to be considered
	CreateProposal(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error)
	ReturnError(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Result, error)
	// Player listing and filtering functions
	//
	// RetrievePlayerPool gets the list of players for every Filter in the
	// PlayerPool, and then removes all players it finds in the ignore list.  It
	// combines the results, and returns the resulting player pool.
	GetPlayerPool(ctx context.Context, in *PlayerPool, opts ...grpc.CallOption) (API_GetPlayerPoolClient, error)
	// Ignore List functions
	//
	// IlInput is an empty message reserved for future use.
	GetAllIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error)
	// RetrieveIgnoreList retrieves players from the
	// 'ignoreLists.proposedPlayers' ignore list specified in the config file
	// that were placed on the list before the provided timestamp.  To retrieve
	// all players on the list, provide the current time as the timestamp.
	ListIgnoredPlayers(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Roster, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/API/GetProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateProposal(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/API/CreateProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ReturnError(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/API/ReturnError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPlayerPool(ctx context.Context, in *PlayerPool, opts ...grpc.CallOption) (API_GetPlayerPoolClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/API/GetPlayerPool", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIGetPlayerPoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_GetPlayerPoolClient interface {
	Recv() (*PlayerPool, error)
	grpc.ClientStream
}

type aPIGetPlayerPoolClient struct {
	grpc.ClientStream
}

func (x *aPIGetPlayerPoolClient) Recv() (*PlayerPool, error) {
	m := new(PlayerPool)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) GetAllIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error) {
	out := new(Roster)
	err := grpc.Invoke(ctx, "/API/GetAllIgnoredPlayers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListIgnoredPlayers(ctx context.Context, in *Timestamp, opts ...grpc.CallOption) (*Roster, error) {
	out := new(Roster)
	err := grpc.Invoke(ctx, "/API/ListIgnoredPlayers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	//  Send GetProfile an profile with the ID field populated, it will return a
	//  'filled' one.
	//  Note: filters are assumed to have been checked for validity by the
	//  backendapi  when accepting a profile
	GetProfile(context.Context, *Profile) (*Profile, error)
	// CreateProposal  does the following:
	//  - adds all players in the Roster to the proposed player ignore list
	//  - writes the proposed match to the provided key
	//  - adds that key to the list of proposals to be considered
	CreateProposal(context.Context, *MatchObject) (*Result, error)
	ReturnError(context.Context, *Result) (*Result, error)
	// Player listing and filtering functions
	//
	// RetrievePlayerPool gets the list of players for every Filter in the
	// PlayerPool, and then removes all players it finds in the ignore list.  It
	// combines the results, and returns the resulting player pool.
	GetPlayerPool(*PlayerPool, API_GetPlayerPoolServer) error
	// Ignore List functions
	//
	// IlInput is an empty message reserved for future use.
	GetAllIgnoredPlayers(context.Context, *IlInput) (*Roster, error)
	// RetrieveIgnoreList retrieves players from the
	// 'ignoreLists.proposedPlayers' ignore list specified in the config file
	// that were placed on the list before the provided timestamp.  To retrieve
	// all players on the list, provide the current time as the timestamp.
	ListIgnoredPlayers(context.Context, *Timestamp) (*Roster, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/CreateProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateProposal(ctx, req.(*MatchObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ReturnError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Result)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ReturnError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ReturnError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ReturnError(ctx, req.(*Result))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetPlayerPool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayerPool)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).GetPlayerPool(m, &aPIGetPlayerPoolServer{stream})
}

type API_GetPlayerPoolServer interface {
	Send(*PlayerPool) error
	grpc.ServerStream
}

type aPIGetPlayerPoolServer struct {
	grpc.ServerStream
}

func (x *aPIGetPlayerPoolServer) Send(m *PlayerPool) error {
	return x.ServerStream.SendMsg(m)
}

func _API_GetAllIgnoredPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAllIgnoredPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetAllIgnoredPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAllIgnoredPlayers(ctx, req.(*IlInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListIgnoredPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timestamp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListIgnoredPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListIgnoredPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListIgnoredPlayers(ctx, req.(*Timestamp))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _API_GetProfile_Handler,
		},
		{
			MethodName: "CreateProposal",
			Handler:    _API_CreateProposal_Handler,
		},
		{
			MethodName: "ReturnError",
			Handler:    _API_ReturnError_Handler,
		},
		{
			MethodName: "GetAllIgnoredPlayers",
			Handler:    _API_GetAllIgnoredPlayers_Handler,
		},
		{
			MethodName: "ListIgnoredPlayers",
			Handler:    _API_ListIgnoredPlayers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPlayerPool",
			Handler:       _API_GetPlayerPool_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mmlogic.proto",
}

func init() { proto.RegisterFile("mmlogic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x9a, 0x36, 0x59, 0x5f, 0x58, 0xc7, 0xac, 0x1d, 0xa2, 0x32, 0xc1, 0xe6, 0xd3, 0xd0,
	0x44, 0x04, 0x45, 0x88, 0x5d, 0x10, 0xaa, 0x26, 0x98, 0x2a, 0x81, 0xa8, 0x3c, 0xee, 0x28, 0x4b,
	0xdd, 0xe2, 0xc9, 0xb1, 0x23, 0xdb, 0xa9, 0xe0, 0x36, 0xc1, 0x7f, 0xe2, 0xf7, 0xa1, 0xd8, 0x49,
	0x9b, 0x02, 0x95, 0x80, 0xdb, 0xfb, 0xbe, 0xe7, 0x67, 0x7f, 0xef, 0x7b, 0xb6, 0x61, 0x3f, 0xcf,
	0xb9, 0x5c, 0xb2, 0x2c, 0x29, 0x94, 0x34, 0x12, 0x17, 0x10, 0xce, 0x94, 0x5c, 0x30, 0x4e, 0xd1,
	0x10, 0xba, 0x6c, 0x1e, 0x7b, 0x27, 0xde, 0xd9, 0x80, 0x74, 0xd9, 0x1c, 0x3d, 0x04, 0x28, 0x94,
	0x2c, 0xa8, 0x32, 0x8c, 0xea, 0xb8, 0x6b, 0xf9, 0x16, 0x83, 0x10, 0xf4, 0x44, 0x9a, 0xd3, 0xd8,
	0xb7, 0x19, 0x1b, 0xa3, 0x53, 0xe8, 0x17, 0x52, 0x72, 0x1d, 0xf7, 0x4e, 0xfc, 0xb3, 0x68, 0x1c,
	0x25, 0x33, 0x9e, 0x7e, 0xa5, 0x6a, 0x26, 0x25, 0x27, 0x2e, 0x83, 0xbf, 0x7b, 0x10, 0xbd, 0x4f,
	0x4d, 0xf6, 0xf9, 0xc3, 0xcd, 0x2d, 0xcd, 0xcc, 0x3f, 0x1f, 0x7b, 0x0a, 0xa1, 0x92, 0xda, 0x50,
	0xa5, 0x63, 0xdf, 0x1e, 0x12, 0x26, 0xc4, 0x62, 0xd2, 0xf0, 0x7f, 0xa3, 0xe2, 0x35, 0x04, 0xae,
	0x6a, 0xdd, 0x86, 0xb7, 0xd5, 0x46, 0x58, 0xd8, 0x92, 0x4a, 0x80, 0x3b, 0xc3, 0x6d, 0x41, 0x1a,
	0x1e, 0xdf, 0x79, 0x10, 0xbc, 0x65, 0x7c, 0xd7, 0x0e, 0xc7, 0x30, 0x48, 0x8d, 0x51, 0xec, 0xa6,
	0x34, 0xb4, 0x6e, 0x62, 0x43, 0x54, 0x15, 0x79, 0xfa, 0x65, 0x65, 0xad, 0xf3, 0x89, 0x8d, 0x2d,
	0xc7, 0xc4, 0x2a, 0xee, 0xd5, 0x1c, 0x13, 0x2b, 0x74, 0x0c, 0x7d, 0x6d, 0x52, 0xa3, 0xe3, 0xfe,
	0x89, 0x77, 0x16, 0x8d, 0x83, 0xe4, 0xba, 0x42, 0xc4, 0x91, 0xf8, 0x25, 0xf4, 0x2d, 0x46, 0x47,
	0xd0, 0xcf, 0x64, 0x29, 0x8c, 0x55, 0xe0, 0x13, 0x07, 0x50, 0x0c, 0x21, 0xe5, 0x69, 0xa1, 0xe9,
	0xdc, 0x0a, 0xf0, 0x48, 0x03, 0xf1, 0x37, 0x0f, 0x60, 0x63, 0xc9, 0x2e, 0x07, 0x16, 0xb6, 0xbb,
	0x8d, 0x03, 0xae, 0x5b, 0xd2, 0xf0, 0xe8, 0x11, 0x04, 0xce, 0x70, 0xdb, 0x46, 0x6b, 0x0e, 0x35,
	0xbd, 0x51, 0xdf, 0xfb, 0x93, 0xfa, 0x1f, 0x1e, 0x04, 0x4e, 0xc4, 0xff, 0xdc, 0xbc, 0x6a, 0x8a,
	0xcd, 0xcd, 0xab, 0x62, 0xf4, 0x0c, 0x60, 0xed, 0x6f, 0x33, 0xf8, 0xc3, 0x7a, 0x6a, 0xc9, 0xa4,
	0xc9, 0x90, 0xd6, 0xa2, 0xd1, 0x0b, 0x18, 0x4c, 0xda, 0x23, 0xf9, 0xcd, 0x84, 0x23, 0xe8, 0xaf,
	0x52, 0x5e, 0xba, 0x01, 0xfa, 0xc4, 0x01, 0x7c, 0x01, 0x01, 0xa1, 0xba, 0xe4, 0xd6, 0x61, 0x5d,
	0x66, 0x19, 0xd5, 0xda, 0x96, 0xed, 0x91, 0x06, 0x56, 0x95, 0x54, 0x29, 0xa9, 0x6a, 0xf1, 0x0e,
	0xe0, 0x01, 0x84, 0x53, 0x3e, 0x15, 0x45, 0x69, 0xf0, 0x03, 0x18, 0x7c, 0x64, 0x39, 0xd5, 0x26,
	0xcd, 0x8b, 0xaa, 0x7f, 0xa3, 0xeb, 0xe1, 0x75, 0x8d, 0xc6, 0xaf, 0x60, 0x78, 0x29, 0x85, 0xa0,
	0x99, 0x61, 0x52, 0x4c, 0xc5, 0x42, 0xa2, 0x73, 0x38, 0xcc, 0xd6, 0xcc, 0x27, 0x6d, 0x14, 0x13,
	0xcb, 0x5a, 0xea, 0xfd, 0x4d, 0xe2, 0xda, 0xf2, 0xf8, 0x16, 0xa2, 0x89, 0xd6, 0x6c, 0x29, 0x72,
	0x2a, 0xcc, 0xd6, 0x83, 0xf1, 0x76, 0x3c, 0x98, 0x0b, 0x38, 0x68, 0x6d, 0xcf, 0xc4, 0x42, 0x5a,
	0xe1, 0xd1, 0xf8, 0x20, 0xd9, 0x16, 0x42, 0x86, 0xd9, 0x16, 0x1e, 0xdf, 0x75, 0xc1, 0x9f, 0xcc,
	0xa6, 0x08, 0x03, 0x5c, 0x51, 0xd3, 0x7c, 0x25, 0x7b, 0x49, 0x1d, 0x8d, 0xd6, 0x11, 0xee, 0xa0,
	0xc7, 0x30, 0xbc, 0x54, 0x34, 0x35, 0x74, 0xa6, 0x64, 0x21, 0x75, 0xca, 0xd1, 0xbd, 0xa4, 0xf5,
	0x13, 0x8c, 0xc2, 0xc4, 0xf9, 0x8a, 0x3b, 0xe8, 0x14, 0x22, 0x42, 0x4d, 0xa9, 0xc4, 0x9b, 0xca,
	0x38, 0xd4, 0x64, 0xda, 0x4b, 0x9e, 0xc0, 0x7e, 0x75, 0xe2, 0xe6, 0x1a, 0xb7, 0x9f, 0xf9, 0xa8,
	0x0d, 0x70, 0xe7, 0xa9, 0x87, 0xce, 0xe1, 0xe8, 0x8a, 0x9a, 0x09, 0xe7, 0xd3, 0xa5, 0x90, 0x8a,
	0xce, 0x5d, 0x5a, 0xa3, 0xbd, 0xa4, 0x1e, 0xc9, 0xa8, 0xb1, 0x05, 0x77, 0xd0, 0x39, 0xa0, 0x77,
	0x4c, 0x9b, 0x5f, 0x96, 0x42, 0xb2, 0x1e, 0x59, 0x6b, 0xf1, 0x4d, 0x60, 0x7f, 0xd2, 0xe7, 0x3f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x76, 0x53, 0x60, 0xd1, 0x5a, 0x05, 0x00, 0x00,
}
