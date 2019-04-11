// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart_service.proto

package cart_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateCartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCartRequest) Reset()         { *m = CreateCartRequest{} }
func (m *CreateCartRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCartRequest) ProtoMessage()    {}
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{0}
}

func (m *CreateCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCartRequest.Unmarshal(m, b)
}
func (m *CreateCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCartRequest.Marshal(b, m, deterministic)
}
func (m *CreateCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCartRequest.Merge(m, src)
}
func (m *CreateCartRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCartRequest.Size(m)
}
func (m *CreateCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCartRequest proto.InternalMessageInfo

type CreateCartResponse struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCartResponse) Reset()         { *m = CreateCartResponse{} }
func (m *CreateCartResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCartResponse) ProtoMessage()    {}
func (*CreateCartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{1}
}

func (m *CreateCartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCartResponse.Unmarshal(m, b)
}
func (m *CreateCartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCartResponse.Marshal(b, m, deterministic)
}
func (m *CreateCartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCartResponse.Merge(m, src)
}
func (m *CreateCartResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCartResponse.Size(m)
}
func (m *CreateCartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCartResponse proto.InternalMessageInfo

func (m *CreateCartResponse) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

type AddLineItemRequest struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity             int64    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLineItemRequest) Reset()         { *m = AddLineItemRequest{} }
func (m *AddLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*AddLineItemRequest) ProtoMessage()    {}
func (*AddLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{2}
}

func (m *AddLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLineItemRequest.Unmarshal(m, b)
}
func (m *AddLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLineItemRequest.Marshal(b, m, deterministic)
}
func (m *AddLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLineItemRequest.Merge(m, src)
}
func (m *AddLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_AddLineItemRequest.Size(m)
}
func (m *AddLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddLineItemRequest proto.InternalMessageInfo

func (m *AddLineItemRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *AddLineItemRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *AddLineItemRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type AddLineItemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLineItemResponse) Reset()         { *m = AddLineItemResponse{} }
func (m *AddLineItemResponse) String() string { return proto.CompactTextString(m) }
func (*AddLineItemResponse) ProtoMessage()    {}
func (*AddLineItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{3}
}

func (m *AddLineItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLineItemResponse.Unmarshal(m, b)
}
func (m *AddLineItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLineItemResponse.Marshal(b, m, deterministic)
}
func (m *AddLineItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLineItemResponse.Merge(m, src)
}
func (m *AddLineItemResponse) XXX_Size() int {
	return xxx_messageInfo_AddLineItemResponse.Size(m)
}
func (m *AddLineItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLineItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddLineItemResponse proto.InternalMessageInfo

type RemoveLineItemRequest struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity             int64    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveLineItemRequest) Reset()         { *m = RemoveLineItemRequest{} }
func (m *RemoveLineItemRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveLineItemRequest) ProtoMessage()    {}
func (*RemoveLineItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{4}
}

func (m *RemoveLineItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveLineItemRequest.Unmarshal(m, b)
}
func (m *RemoveLineItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveLineItemRequest.Marshal(b, m, deterministic)
}
func (m *RemoveLineItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveLineItemRequest.Merge(m, src)
}
func (m *RemoveLineItemRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveLineItemRequest.Size(m)
}
func (m *RemoveLineItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveLineItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveLineItemRequest proto.InternalMessageInfo

func (m *RemoveLineItemRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *RemoveLineItemRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *RemoveLineItemRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type RemoveLineItemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveLineItemResponse) Reset()         { *m = RemoveLineItemResponse{} }
func (m *RemoveLineItemResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveLineItemResponse) ProtoMessage()    {}
func (*RemoveLineItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{5}
}

func (m *RemoveLineItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveLineItemResponse.Unmarshal(m, b)
}
func (m *RemoveLineItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveLineItemResponse.Marshal(b, m, deterministic)
}
func (m *RemoveLineItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveLineItemResponse.Merge(m, src)
}
func (m *RemoveLineItemResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveLineItemResponse.Size(m)
}
func (m *RemoveLineItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveLineItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveLineItemResponse proto.InternalMessageInfo

type EmptyCartRequest struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyCartRequest) Reset()         { *m = EmptyCartRequest{} }
func (m *EmptyCartRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyCartRequest) ProtoMessage()    {}
func (*EmptyCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{6}
}

func (m *EmptyCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyCartRequest.Unmarshal(m, b)
}
func (m *EmptyCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyCartRequest.Marshal(b, m, deterministic)
}
func (m *EmptyCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyCartRequest.Merge(m, src)
}
func (m *EmptyCartRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyCartRequest.Size(m)
}
func (m *EmptyCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyCartRequest proto.InternalMessageInfo

func (m *EmptyCartRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

type EmptyCartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyCartResponse) Reset()         { *m = EmptyCartResponse{} }
func (m *EmptyCartResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyCartResponse) ProtoMessage()    {}
func (*EmptyCartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{7}
}

func (m *EmptyCartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyCartResponse.Unmarshal(m, b)
}
func (m *EmptyCartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyCartResponse.Marshal(b, m, deterministic)
}
func (m *EmptyCartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyCartResponse.Merge(m, src)
}
func (m *EmptyCartResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyCartResponse.Size(m)
}
func (m *EmptyCartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyCartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyCartResponse proto.InternalMessageInfo

type GetLineItemsRequest struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLineItemsRequest) Reset()         { *m = GetLineItemsRequest{} }
func (m *GetLineItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetLineItemsRequest) ProtoMessage()    {}
func (*GetLineItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{8}
}

func (m *GetLineItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLineItemsRequest.Unmarshal(m, b)
}
func (m *GetLineItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLineItemsRequest.Marshal(b, m, deterministic)
}
func (m *GetLineItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLineItemsRequest.Merge(m, src)
}
func (m *GetLineItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetLineItemsRequest.Size(m)
}
func (m *GetLineItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLineItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLineItemsRequest proto.InternalMessageInfo

func (m *GetLineItemsRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

type LineItem struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity             int64                `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalPrice           float32              `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{9}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LineItem) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *LineItem) GetTotalPrice() float32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *LineItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GetLineItemsResponse struct {
	LineItem             []*LineItem          `protobuf:"bytes,1,rep,name=line_item,json=lineItem,proto3" json:"line_item,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetLineItemsResponse) Reset()         { *m = GetLineItemsResponse{} }
func (m *GetLineItemsResponse) String() string { return proto.CompactTextString(m) }
func (*GetLineItemsResponse) ProtoMessage()    {}
func (*GetLineItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{10}
}

func (m *GetLineItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLineItemsResponse.Unmarshal(m, b)
}
func (m *GetLineItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLineItemsResponse.Marshal(b, m, deterministic)
}
func (m *GetLineItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLineItemsResponse.Merge(m, src)
}
func (m *GetLineItemsResponse) XXX_Size() int {
	return xxx_messageInfo_GetLineItemsResponse.Size(m)
}
func (m *GetLineItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLineItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLineItemsResponse proto.InternalMessageInfo

func (m *GetLineItemsResponse) GetLineItem() []*LineItem {
	if m != nil {
		return m.LineItem
	}
	return nil
}

func (m *GetLineItemsResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCartRequest)(nil), "CreateCartRequest")
	proto.RegisterType((*CreateCartResponse)(nil), "CreateCartResponse")
	proto.RegisterType((*AddLineItemRequest)(nil), "AddLineItemRequest")
	proto.RegisterType((*AddLineItemResponse)(nil), "AddLineItemResponse")
	proto.RegisterType((*RemoveLineItemRequest)(nil), "RemoveLineItemRequest")
	proto.RegisterType((*RemoveLineItemResponse)(nil), "RemoveLineItemResponse")
	proto.RegisterType((*EmptyCartRequest)(nil), "EmptyCartRequest")
	proto.RegisterType((*EmptyCartResponse)(nil), "EmptyCartResponse")
	proto.RegisterType((*GetLineItemsRequest)(nil), "GetLineItemsRequest")
	proto.RegisterType((*LineItem)(nil), "LineItem")
	proto.RegisterType((*GetLineItemsResponse)(nil), "GetLineItemsResponse")
}

func init() { proto.RegisterFile("cart_service.proto", fileDescriptor_c9a99120c5507bc1) }

var fileDescriptor_c9a99120c5507bc1 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8b, 0xd3, 0x50,
	0x10, 0xc7, 0x93, 0xd6, 0x5d, 0x9b, 0x89, 0x88, 0x9d, 0xb4, 0xbb, 0x21, 0x20, 0x1b, 0x72, 0x90,
	0x82, 0xf8, 0x16, 0xaa, 0x20, 0x0a, 0x1e, 0x96, 0x22, 0x52, 0xf0, 0x20, 0xd1, 0x7b, 0xc9, 0x26,
	0xe3, 0xfa, 0x30, 0xc9, 0xcb, 0x26, 0x93, 0x85, 0xfe, 0x6b, 0xe2, 0x1f, 0x27, 0x7d, 0x49, 0x6a,
	0xba, 0x6d, 0x15, 0x2f, 0x1e, 0x67, 0xde, 0x97, 0xf9, 0xf1, 0x7d, 0x9f, 0x01, 0x8c, 0xa3, 0x92,
	0x57, 0x15, 0x95, 0x77, 0x32, 0x26, 0x51, 0x94, 0x8a, 0x95, 0x77, 0x71, 0xa3, 0xd4, 0x4d, 0x4a,
	0x97, 0x3a, 0xba, 0xae, 0xbf, 0x5e, 0xb2, 0xcc, 0xa8, 0xe2, 0x28, 0x2b, 0x1a, 0x41, 0xe0, 0xc0,
	0x78, 0x51, 0x52, 0xc4, 0xb4, 0x88, 0x4a, 0x0e, 0xe9, 0xb6, 0xa6, 0x8a, 0x83, 0x17, 0x80, 0xfd,
	0x64, 0x55, 0xa8, 0xbc, 0x22, 0x3c, 0x87, 0x87, 0xba, 0x83, 0x4c, 0x5c, 0xd3, 0x37, 0x67, 0xc3,
	0xf0, 0x74, 0x13, 0x2e, 0x93, 0xe0, 0x1b, 0xe0, 0x55, 0x92, 0x7c, 0x94, 0x39, 0x2d, 0x99, 0xb2,
	0xb6, 0xc8, 0x51, 0x39, 0x3e, 0x05, 0x28, 0x4a, 0x95, 0xd4, 0xb1, 0x7e, 0x1b, 0xe8, 0x37, 0xab,
	0xcd, 0x2c, 0x13, 0xf4, 0x60, 0x74, 0x5b, 0x47, 0x39, 0x4b, 0x5e, 0xbb, 0x43, 0xfd, 0xb8, 0x8d,
	0x83, 0x29, 0x38, 0x3b, 0x9d, 0x9a, 0xc9, 0x82, 0xef, 0x30, 0x0d, 0x29, 0x53, 0x77, 0xf4, 0x3f,
	0x66, 0x70, 0xe1, 0xec, 0x7e, 0xb3, 0x76, 0x8c, 0xe7, 0xf0, 0xe4, 0x7d, 0x56, 0xf0, 0xba, 0x67,
	0xe5, 0x71, 0xd3, 0x1c, 0x18, 0xf7, 0xc4, 0x6d, 0x05, 0x01, 0xce, 0x07, 0xe2, 0xae, 0x70, 0xf5,
	0xd7, 0x22, 0x3f, 0x4c, 0x18, 0x75, 0x6a, 0x9c, 0xc0, 0x09, 0x4b, 0x4e, 0x49, 0x6b, 0xac, 0xb0,
	0x09, 0xd0, 0x07, 0x3b, 0xa1, 0x2a, 0x2e, 0x65, 0xc1, 0x52, 0xe5, 0x7a, 0x55, 0x2b, 0xec, 0xa7,
	0xfe, 0xb4, 0x2c, 0x5e, 0x80, 0xcd, 0x8a, 0xa3, 0x74, 0x55, 0x94, 0x32, 0x26, 0xf7, 0x81, 0x6f,
	0xce, 0x06, 0x21, 0xe8, 0xd4, 0xa7, 0x4d, 0x06, 0xdf, 0x00, 0xd4, 0x45, 0x12, 0x31, 0x25, 0xab,
	0x88, 0xdd, 0x13, 0xdf, 0x9c, 0xd9, 0x73, 0x4f, 0x34, 0xd4, 0x89, 0x8e, 0x3a, 0xf1, 0xa5, 0xa3,
	0x2e, 0xb4, 0x5a, 0xf5, 0x15, 0x07, 0x6b, 0x98, 0xec, 0x2e, 0xdb, 0x72, 0xf6, 0x0c, 0xac, 0x54,
	0xe6, 0xb4, 0x92, 0x4c, 0x99, 0x6b, 0xfa, 0xc3, 0x99, 0x3d, 0xb7, 0xc4, 0xd6, 0xec, 0x51, 0xda,
	0xed, 0xbb, 0xdb, 0x7a, 0xf0, 0x0f, 0xad, 0xe7, 0x3f, 0x07, 0x60, 0x6f, 0x8c, 0xff, 0xdc, 0x1c,
	0x0b, 0xbe, 0x06, 0xf8, 0x0d, 0x3c, 0xa2, 0xd8, 0x3b, 0x09, 0xcf, 0x11, 0xfb, 0x17, 0x11, 0x18,
	0xf8, 0x16, 0xec, 0x1e, 0x90, 0xe8, 0x88, 0xfd, 0x43, 0xf0, 0x26, 0xe2, 0x10, 0xb3, 0x06, 0x2e,
	0xe0, 0xf1, 0x2e, 0x48, 0x78, 0x26, 0x0e, 0x62, 0xec, 0x9d, 0x8b, 0x23, 0xc4, 0x19, 0xf8, 0x0a,
	0xac, 0x2d, 0x46, 0x38, 0x16, 0xf7, 0xf9, 0xf3, 0x50, 0xec, 0x53, 0x66, 0xe0, 0x3b, 0x78, 0xd4,
	0xb7, 0x1e, 0x27, 0xe2, 0x00, 0x76, 0xde, 0x54, 0x1c, 0xfa, 0x9f, 0xc0, 0xb8, 0x3e, 0xd5, 0xee,
	0xbe, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x72, 0x67, 0xa6, 0x52, 0x72, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error)
	AddLineItem(ctx context.Context, in *AddLineItemRequest, opts ...grpc.CallOption) (*AddLineItemResponse, error)
	RemoveLineItem(ctx context.Context, in *RemoveLineItemRequest, opts ...grpc.CallOption) (*RemoveLineItemResponse, error)
	EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*EmptyCartResponse, error)
	GetLineItems(ctx context.Context, in *GetLineItemsRequest, opts ...grpc.CallOption) (*GetLineItemsResponse, error)
}

type cartServiceClient struct {
	cc *grpc.ClientConn
}

func NewCartServiceClient(cc *grpc.ClientConn) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error) {
	out := new(CreateCartResponse)
	err := c.cc.Invoke(ctx, "/CartService/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddLineItem(ctx context.Context, in *AddLineItemRequest, opts ...grpc.CallOption) (*AddLineItemResponse, error) {
	out := new(AddLineItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/AddLineItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveLineItem(ctx context.Context, in *RemoveLineItemRequest, opts ...grpc.CallOption) (*RemoveLineItemResponse, error) {
	out := new(RemoveLineItemResponse)
	err := c.cc.Invoke(ctx, "/CartService/RemoveLineItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*EmptyCartResponse, error) {
	out := new(EmptyCartResponse)
	err := c.cc.Invoke(ctx, "/CartService/EmptyCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetLineItems(ctx context.Context, in *GetLineItemsRequest, opts ...grpc.CallOption) (*GetLineItemsResponse, error) {
	out := new(GetLineItemsResponse)
	err := c.cc.Invoke(ctx, "/CartService/GetLineItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error)
	AddLineItem(context.Context, *AddLineItemRequest) (*AddLineItemResponse, error)
	RemoveLineItem(context.Context, *RemoveLineItemRequest) (*RemoveLineItemResponse, error)
	EmptyCart(context.Context, *EmptyCartRequest) (*EmptyCartResponse, error)
	GetLineItems(context.Context, *GetLineItemsRequest) (*GetLineItemsResponse, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) CreateCart(ctx context.Context, req *CreateCartRequest) (*CreateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (*UnimplementedCartServiceServer) AddLineItem(ctx context.Context, req *AddLineItemRequest) (*AddLineItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLineItem not implemented")
}
func (*UnimplementedCartServiceServer) RemoveLineItem(ctx context.Context, req *RemoveLineItemRequest) (*RemoveLineItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLineItem not implemented")
}
func (*UnimplementedCartServiceServer) EmptyCart(ctx context.Context, req *EmptyCartRequest) (*EmptyCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCart not implemented")
}
func (*UnimplementedCartServiceServer) GetLineItems(ctx context.Context, req *GetLineItemsRequest) (*GetLineItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLineItems not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddLineItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLineItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddLineItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/AddLineItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddLineItem(ctx, req.(*AddLineItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveLineItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLineItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveLineItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/RemoveLineItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveLineItem(ctx, req.(*RemoveLineItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_EmptyCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).EmptyCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/EmptyCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).EmptyCart(ctx, req.(*EmptyCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetLineItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLineItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetLineItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CartService/GetLineItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetLineItems(ctx, req.(*GetLineItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "AddLineItem",
			Handler:    _CartService_AddLineItem_Handler,
		},
		{
			MethodName: "RemoveLineItem",
			Handler:    _CartService_RemoveLineItem_Handler,
		},
		{
			MethodName: "EmptyCart",
			Handler:    _CartService_EmptyCart_Handler,
		},
		{
			MethodName: "GetLineItems",
			Handler:    _CartService_GetLineItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}
