// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources/proto/ms/ms.proto

package ms

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order int32

const (
	Order_ASC  Order = 0
	Order_DESC Order = 1
)

var Order_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}

var Order_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x Order) String() string {
	return proto.EnumName(Order_name, int32(x))
}

func (Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{0}
}

type SearchOut_Status int32

const (
	SearchOut_OK                    SearchOut_Status = 0
	SearchOut_FORBIDDEN             SearchOut_Status = 1
	SearchOut_NOT_FOUND             SearchOut_Status = 2
	SearchOut_BAD_REQUEST           SearchOut_Status = 3
	SearchOut_SERVER_INTERNAL_ERROR SearchOut_Status = 4
)

var SearchOut_Status_name = map[int32]string{
	0: "OK",
	1: "FORBIDDEN",
	2: "NOT_FOUND",
	3: "BAD_REQUEST",
	4: "SERVER_INTERNAL_ERROR",
}

var SearchOut_Status_value = map[string]int32{
	"OK":                    0,
	"FORBIDDEN":             1,
	"NOT_FOUND":             2,
	"BAD_REQUEST":           3,
	"SERVER_INTERNAL_ERROR": 4,
}

func (x SearchOut_Status) String() string {
	return proto.EnumName(SearchOut_Status_name, int32(x))
}

func (SearchOut_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{3, 0}
}

type NewOut_Status int32

const (
	NewOut_OK                    NewOut_Status = 0
	NewOut_FORBIDDEN             NewOut_Status = 1
	NewOut_BAD_REQUEST           NewOut_Status = 2
	NewOut_SERVER_INTERNAL_ERROR NewOut_Status = 3
)

var NewOut_Status_name = map[int32]string{
	0: "OK",
	1: "FORBIDDEN",
	2: "BAD_REQUEST",
	3: "SERVER_INTERNAL_ERROR",
}

var NewOut_Status_value = map[string]int32{
	"OK":                    0,
	"FORBIDDEN":             1,
	"BAD_REQUEST":           2,
	"SERVER_INTERNAL_ERROR": 3,
}

func (x NewOut_Status) String() string {
	return proto.EnumName(NewOut_Status_name, int32(x))
}

func (NewOut_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{5, 0}
}

type CursorIn struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Cursor               string   `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CursorIn) Reset()         { *m = CursorIn{} }
func (m *CursorIn) String() string { return proto.CompactTextString(m) }
func (*CursorIn) ProtoMessage()    {}
func (*CursorIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{0}
}

func (m *CursorIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorIn.Unmarshal(m, b)
}
func (m *CursorIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorIn.Marshal(b, m, deterministic)
}
func (m *CursorIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorIn.Merge(m, src)
}
func (m *CursorIn) XXX_Size() int {
	return xxx_messageInfo_CursorIn.Size(m)
}
func (m *CursorIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorIn.DiscardUnknown(m)
}

var xxx_messageInfo_CursorIn proto.InternalMessageInfo

func (m *CursorIn) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CursorIn) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *CursorIn) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

type CursorOut struct {
	TotalCount           int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	HasNextPage          bool     `protobuf:"varint,4,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
	Cursor               string   `protobuf:"bytes,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CursorOut) Reset()         { *m = CursorOut{} }
func (m *CursorOut) String() string { return proto.CompactTextString(m) }
func (*CursorOut) ProtoMessage()    {}
func (*CursorOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{1}
}

func (m *CursorOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorOut.Unmarshal(m, b)
}
func (m *CursorOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorOut.Marshal(b, m, deterministic)
}
func (m *CursorOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorOut.Merge(m, src)
}
func (m *CursorOut) XXX_Size() int {
	return xxx_messageInfo_CursorOut.Size(m)
}
func (m *CursorOut) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorOut.DiscardUnknown(m)
}

var xxx_messageInfo_CursorOut proto.InternalMessageInfo

func (m *CursorOut) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *CursorOut) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *CursorOut) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *CursorOut) GetHasNextPage() bool {
	if m != nil {
		return m.HasNextPage
	}
	return false
}

func (m *CursorOut) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

type SearchIn struct {
	Query                string                `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Order                Order                 `protobuf:"varint,2,opt,name=order,proto3,enum=ms.Order" json:"order,omitempty"`
	Cursor               *CursorIn             `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Projection           *field_mask.FieldMask `protobuf:"bytes,4,opt,name=projection,proto3" json:"projection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SearchIn) Reset()         { *m = SearchIn{} }
func (m *SearchIn) String() string { return proto.CompactTextString(m) }
func (*SearchIn) ProtoMessage()    {}
func (*SearchIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{2}
}

func (m *SearchIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchIn.Unmarshal(m, b)
}
func (m *SearchIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchIn.Marshal(b, m, deterministic)
}
func (m *SearchIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchIn.Merge(m, src)
}
func (m *SearchIn) XXX_Size() int {
	return xxx_messageInfo_SearchIn.Size(m)
}
func (m *SearchIn) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchIn.DiscardUnknown(m)
}

var xxx_messageInfo_SearchIn proto.InternalMessageInfo

func (m *SearchIn) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchIn) GetOrder() Order {
	if m != nil {
		return m.Order
	}
	return Order_ASC
}

func (m *SearchIn) GetCursor() *CursorIn {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *SearchIn) GetProjection() *field_mask.FieldMask {
	if m != nil {
		return m.Projection
	}
	return nil
}

type SearchOut struct {
	Id                   []int64          `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Status               SearchOut_Status `protobuf:"varint,2,opt,name=status,proto3,enum=ms.SearchOut_Status" json:"status,omitempty"`
	Cursor               *CursorOut       `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchOut) Reset()         { *m = SearchOut{} }
func (m *SearchOut) String() string { return proto.CompactTextString(m) }
func (*SearchOut) ProtoMessage()    {}
func (*SearchOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{3}
}

func (m *SearchOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchOut.Unmarshal(m, b)
}
func (m *SearchOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchOut.Marshal(b, m, deterministic)
}
func (m *SearchOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchOut.Merge(m, src)
}
func (m *SearchOut) XXX_Size() int {
	return xxx_messageInfo_SearchOut.Size(m)
}
func (m *SearchOut) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchOut.DiscardUnknown(m)
}

var xxx_messageInfo_SearchOut proto.InternalMessageInfo

func (m *SearchOut) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchOut) GetStatus() SearchOut_Status {
	if m != nil {
		return m.Status
	}
	return SearchOut_OK
}

func (m *SearchOut) GetCursor() *CursorOut {
	if m != nil {
		return m.Cursor
	}
	return nil
}

type NewIn struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewIn) Reset()         { *m = NewIn{} }
func (m *NewIn) String() string { return proto.CompactTextString(m) }
func (*NewIn) ProtoMessage()    {}
func (*NewIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{4}
}

func (m *NewIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewIn.Unmarshal(m, b)
}
func (m *NewIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewIn.Marshal(b, m, deterministic)
}
func (m *NewIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewIn.Merge(m, src)
}
func (m *NewIn) XXX_Size() int {
	return xxx_messageInfo_NewIn.Size(m)
}
func (m *NewIn) XXX_DiscardUnknown() {
	xxx_messageInfo_NewIn.DiscardUnknown(m)
}

var xxx_messageInfo_NewIn proto.InternalMessageInfo

func (m *NewIn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NewOut struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               NewOut_Status `protobuf:"varint,2,opt,name=status,proto3,enum=ms.NewOut_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NewOut) Reset()         { *m = NewOut{} }
func (m *NewOut) String() string { return proto.CompactTextString(m) }
func (*NewOut) ProtoMessage()    {}
func (*NewOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_799597003239ca4c, []int{5}
}

func (m *NewOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewOut.Unmarshal(m, b)
}
func (m *NewOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewOut.Marshal(b, m, deterministic)
}
func (m *NewOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewOut.Merge(m, src)
}
func (m *NewOut) XXX_Size() int {
	return xxx_messageInfo_NewOut.Size(m)
}
func (m *NewOut) XXX_DiscardUnknown() {
	xxx_messageInfo_NewOut.DiscardUnknown(m)
}

var xxx_messageInfo_NewOut proto.InternalMessageInfo

func (m *NewOut) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NewOut) GetStatus() NewOut_Status {
	if m != nil {
		return m.Status
	}
	return NewOut_OK
}

func init() {
	proto.RegisterEnum("ms.Order", Order_name, Order_value)
	proto.RegisterEnum("ms.SearchOut_Status", SearchOut_Status_name, SearchOut_Status_value)
	proto.RegisterEnum("ms.NewOut_Status", NewOut_Status_name, NewOut_Status_value)
	proto.RegisterType((*CursorIn)(nil), "ms.CursorIn")
	proto.RegisterType((*CursorOut)(nil), "ms.CursorOut")
	proto.RegisterType((*SearchIn)(nil), "ms.SearchIn")
	proto.RegisterType((*SearchOut)(nil), "ms.SearchOut")
	proto.RegisterType((*NewIn)(nil), "ms.NewIn")
	proto.RegisterType((*NewOut)(nil), "ms.NewOut")
}

func init() { proto.RegisterFile("resources/proto/ms/ms.proto", fileDescriptor_799597003239ca4c) }

var fileDescriptor_799597003239ca4c = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0x99, 0x16, 0x2a, 0x3d, 0x2c, 0x6b, 0x9d, 0xac, 0x06, 0xd9, 0xc4, 0x25, 0x8d, 0x46,
	0x34, 0xa6, 0x24, 0x78, 0xe7, 0xdd, 0x2e, 0x74, 0x93, 0x66, 0xa5, 0x5d, 0xa7, 0xac, 0x17, 0xde,
	0x34, 0x5d, 0x18, 0xa0, 0x2e, 0xed, 0x60, 0x67, 0x1a, 0xd6, 0xb7, 0xf0, 0xca, 0x17, 0xf0, 0x75,
	0x7c, 0x28, 0xd3, 0x4e, 0x59, 0xfe, 0x44, 0xe3, 0x5d, 0xcf, 0x99, 0x6f, 0xe6, 0xfc, 0xbe, 0xef,
	0x14, 0x4e, 0x53, 0xca, 0x59, 0x96, 0x4e, 0x28, 0xef, 0xad, 0x52, 0x26, 0x58, 0x2f, 0xe6, 0xbd,
	0x98, 0x5b, 0xc5, 0x37, 0x56, 0x62, 0xde, 0xee, 0xcc, 0x19, 0x9b, 0x2f, 0xa9, 0x3c, 0xbd, 0xcd,
	0x66, 0xbd, 0x59, 0x44, 0x97, 0xd3, 0x20, 0x0e, 0xf9, 0x9d, 0x54, 0x99, 0xd7, 0x50, 0x1f, 0x64,
	0x29, 0x67, 0xa9, 0x93, 0xe0, 0x13, 0xa8, 0x2d, 0xa3, 0x38, 0x12, 0x2d, 0xd4, 0x41, 0x5d, 0x95,
	0xc8, 0x02, 0x3f, 0x03, 0x8d, 0xcd, 0x66, 0x9c, 0x8a, 0x96, 0x52, 0xb4, 0xcb, 0x2a, 0xef, 0x4f,
	0x8a, 0x9b, 0x2d, 0xb5, 0x83, 0xba, 0x3a, 0x29, 0x2b, 0xf3, 0x27, 0x02, 0x5d, 0x3e, 0xe9, 0x65,
	0x02, 0x9f, 0x41, 0x43, 0x30, 0x11, 0x2e, 0x83, 0x09, 0xcb, 0x92, 0xcd, 0xcb, 0x50, 0xb4, 0x06,
	0x79, 0x67, 0x3b, 0x54, 0xf9, 0xfb, 0x50, 0x75, 0x6f, 0xa8, 0x09, 0xcd, 0x45, 0xc8, 0x83, 0x84,
	0xde, 0x8b, 0x60, 0x15, 0xce, 0x69, 0xab, 0xda, 0x41, 0xdd, 0x3a, 0x69, 0x2c, 0x42, 0xee, 0xd2,
	0x7b, 0x71, 0x1d, 0xce, 0xe9, 0x0e, 0x58, 0x6d, 0x0f, 0xec, 0x17, 0x82, 0xba, 0x4f, 0xc3, 0x74,
	0xb2, 0x90, 0x5e, 0xbf, 0x65, 0x34, 0xfd, 0x5e, 0x10, 0xe9, 0x44, 0x16, 0xf8, 0x0c, 0x6a, 0x2c,
	0x9d, 0xd2, 0xb4, 0x80, 0x39, 0xee, 0xeb, 0x56, 0xcc, 0x2d, 0x2f, 0x6f, 0x10, 0xd9, 0xc7, 0x2f,
	0xf7, 0x4c, 0x37, 0xfa, 0x47, 0xb9, 0x62, 0x13, 0xe0, 0x66, 0x12, 0xfe, 0x00, 0xb0, 0x4a, 0xd9,
	0x57, 0x3a, 0x11, 0x11, 0x4b, 0x0a, 0xc4, 0x46, 0xbf, 0x6d, 0xc9, 0x5d, 0x58, 0x9b, 0x5d, 0x58,
	0x97, 0xf9, 0x2e, 0x46, 0x21, 0xbf, 0x23, 0x3b, 0x6a, 0xf3, 0x37, 0x02, 0x5d, 0x52, 0xe6, 0xf1,
	0x1d, 0x83, 0x12, 0x4d, 0x5b, 0xa8, 0xa3, 0x76, 0x55, 0xa2, 0x44, 0x53, 0xfc, 0x0e, 0x34, 0x2e,
	0x42, 0x91, 0xf1, 0x92, 0xf0, 0x24, 0x9f, 0xff, 0x20, 0xb7, 0xfc, 0xe2, 0x8c, 0x94, 0x1a, 0xfc,
	0xea, 0x80, 0xb6, 0xb9, 0xa5, 0xf5, 0x32, 0xf1, 0x10, 0xcc, 0x17, 0xd0, 0xe4, 0x45, 0xac, 0x81,
	0xe2, 0x5d, 0x19, 0x15, 0xdc, 0x04, 0xfd, 0xd2, 0x23, 0x17, 0xce, 0x70, 0x68, 0xbb, 0x06, 0xca,
	0x4b, 0xd7, 0x1b, 0x07, 0x97, 0xde, 0x8d, 0x3b, 0x34, 0x14, 0xfc, 0x18, 0x1a, 0x17, 0xe7, 0xc3,
	0x80, 0xd8, 0x9f, 0x6e, 0x6c, 0x7f, 0x6c, 0xa8, 0xf8, 0x39, 0x3c, 0xf5, 0x6d, 0xf2, 0xd9, 0x26,
	0x81, 0xe3, 0x8e, 0x6d, 0xe2, 0x9e, 0x7f, 0x0c, 0x6c, 0x42, 0x3c, 0x62, 0x54, 0xcd, 0x53, 0xa8,
	0xb9, 0x74, 0xed, 0x24, 0x18, 0x43, 0x35, 0x09, 0x63, 0x5a, 0xe6, 0x5d, 0x7c, 0x9b, 0x3f, 0x10,
	0x68, 0x2e, 0x5d, 0xef, 0x1a, 0x45, 0xa5, 0xd1, 0x37, 0x07, 0x46, 0x9f, 0xe4, 0xe8, 0x52, 0x7b,
	0xe0, 0xd2, 0xbc, 0xfa, 0x1f, 0xfe, 0x01, 0xaf, 0xf2, 0x6f, 0x5e, 0xf5, 0x6d, 0x1b, 0x6a, 0xc5,
	0xc2, 0xf1, 0x23, 0x50, 0xcf, 0xfd, 0x81, 0x51, 0xc1, 0x75, 0xa8, 0x0e, 0x6d, 0x7f, 0x60, 0xa0,
	0xfe, 0x08, 0x94, 0x11, 0xc7, 0xaf, 0x41, 0x93, 0x81, 0xe3, 0xa3, 0x6d, 0xf8, 0x4e, 0xd2, 0x6e,
	0xee, 0xad, 0xc2, 0xac, 0xe0, 0x17, 0xa0, 0xba, 0x74, 0x8d, 0xf5, 0x92, 0xdc, 0x49, 0xda, 0xb0,
	0x35, 0x61, 0x56, 0x6e, 0xb5, 0xe2, 0x4f, 0x78, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x56, 0xe7,
	0x80, 0xfe, 0xc6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsClient is the client API for Ms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsClient interface {
	Search(ctx context.Context, in *SearchIn, opts ...grpc.CallOption) (*SearchOut, error)
	New(ctx context.Context, in *NewIn, opts ...grpc.CallOption) (*NewOut, error)
}

type msClient struct {
	cc *grpc.ClientConn
}

func NewMsClient(cc *grpc.ClientConn) MsClient {
	return &msClient{cc}
}

func (c *msClient) Search(ctx context.Context, in *SearchIn, opts ...grpc.CallOption) (*SearchOut, error) {
	out := new(SearchOut)
	err := c.cc.Invoke(ctx, "/ms.Ms/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msClient) New(ctx context.Context, in *NewIn, opts ...grpc.CallOption) (*NewOut, error) {
	out := new(NewOut)
	err := c.cc.Invoke(ctx, "/ms.Ms/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsServer is the server API for Ms service.
type MsServer interface {
	Search(context.Context, *SearchIn) (*SearchOut, error)
	New(context.Context, *NewIn) (*NewOut, error)
}

func RegisterMsServer(s *grpc.Server, srv MsServer) {
	s.RegisterService(&_Ms_serviceDesc, srv)
}

func _Ms_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ms.Ms/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsServer).Search(ctx, req.(*SearchIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ms_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ms.Ms/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsServer).New(ctx, req.(*NewIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ms.Ms",
	HandlerType: (*MsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Ms_Search_Handler,
		},
		{
			MethodName: "New",
			Handler:    _Ms_New_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources/proto/ms/ms.proto",
}
