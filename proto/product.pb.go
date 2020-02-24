// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Product is a message for product
type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Picture              string   `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Status               bool     `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Product) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// GetAllProductsRequest is a message for requesting all products
type GetAllProductsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllProductsRequest) Reset()         { *m = GetAllProductsRequest{} }
func (m *GetAllProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllProductsRequest) ProtoMessage()    {}
func (*GetAllProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *GetAllProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllProductsRequest.Unmarshal(m, b)
}
func (m *GetAllProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllProductsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllProductsRequest.Merge(m, src)
}
func (m *GetAllProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllProductsRequest.Size(m)
}
func (m *GetAllProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllProductsRequest proto.InternalMessageInfo

// DeleteProductRequest is a message for deleting a product
type DeleteProductRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductRequest) Reset()         { *m = DeleteProductRequest{} }
func (m *DeleteProductRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductRequest) ProtoMessage()    {}
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *DeleteProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductRequest.Unmarshal(m, b)
}
func (m *DeleteProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductRequest.Merge(m, src)
}
func (m *DeleteProductRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductRequest.Size(m)
}
func (m *DeleteProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductRequest proto.InternalMessageInfo

func (m *DeleteProductRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// SingleJobResponse is a message for a single job response
type SingleJobResponse struct {
	CreatedProduct       *Product `protobuf:"bytes,1,opt,name=createdProduct,proto3" json:"createdProduct,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleJobResponse) Reset()         { *m = SingleJobResponse{} }
func (m *SingleJobResponse) String() string { return proto.CompactTextString(m) }
func (*SingleJobResponse) ProtoMessage()    {}
func (*SingleJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *SingleJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleJobResponse.Unmarshal(m, b)
}
func (m *SingleJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleJobResponse.Marshal(b, m, deterministic)
}
func (m *SingleJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleJobResponse.Merge(m, src)
}
func (m *SingleJobResponse) XXX_Size() int {
	return xxx_messageInfo_SingleJobResponse.Size(m)
}
func (m *SingleJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SingleJobResponse proto.InternalMessageInfo

func (m *SingleJobResponse) GetCreatedProduct() *Product {
	if m != nil {
		return m.CreatedProduct
	}
	return nil
}

func (m *SingleJobResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// MultipleJobResponse is a message for a multiple job response
type MultipleJobResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MultipleJobResponse) Reset()         { *m = MultipleJobResponse{} }
func (m *MultipleJobResponse) String() string { return proto.CompactTextString(m) }
func (*MultipleJobResponse) ProtoMessage()    {}
func (*MultipleJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
}

func (m *MultipleJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipleJobResponse.Unmarshal(m, b)
}
func (m *MultipleJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipleJobResponse.Marshal(b, m, deterministic)
}
func (m *MultipleJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipleJobResponse.Merge(m, src)
}
func (m *MultipleJobResponse) XXX_Size() int {
	return xxx_messageInfo_MultipleJobResponse.Size(m)
}
func (m *MultipleJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipleJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultipleJobResponse proto.InternalMessageInfo

func (m *MultipleJobResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

// Error is a message for errors
type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Product)(nil), "product.Product")
	proto.RegisterType((*GetAllProductsRequest)(nil), "product.GetAllProductsRequest")
	proto.RegisterType((*DeleteProductRequest)(nil), "product.DeleteProductRequest")
	proto.RegisterType((*SingleJobResponse)(nil), "product.SingleJobResponse")
	proto.RegisterType((*MultipleJobResponse)(nil), "product.MultipleJobResponse")
	proto.RegisterType((*Error)(nil), "product.Error")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdb, 0x8a, 0xdb, 0x30,
	0x10, 0x8d, 0x92, 0x38, 0x97, 0x09, 0x31, 0xad, 0x9a, 0xb6, 0x22, 0xb4, 0xc5, 0x88, 0x52, 0xfc,
	0x50, 0xf2, 0x90, 0x52, 0xe8, 0x5b, 0x29, 0x69, 0x29, 0x2c, 0xbb, 0x10, 0x94, 0x2f, 0x70, 0xec,
	0x21, 0x08, 0x1c, 0xcb, 0x2b, 0xc9, 0xfb, 0x2b, 0xfb, 0x9d, 0xfb, 0x07, 0x8b, 0x15, 0xd9, 0x6c,
	0x12, 0xb3, 0x6f, 0x73, 0x66, 0xe6, 0xcc, 0x8c, 0xce, 0x11, 0xcc, 0x4b, 0xad, 0xb2, 0x2a, 0xb5,
	0xab, 0x52, 0x2b, 0xab, 0xe8, 0xd8, 0x43, 0xfe, 0x48, 0x60, 0xbc, 0x3d, 0xc5, 0x34, 0x84, 0xbe,
	0xcc, 0x18, 0x89, 0x48, 0x1c, 0x88, 0xbe, 0xcc, 0x28, 0x85, 0x61, 0x91, 0x1c, 0x91, 0xf5, 0x23,
	0x12, 0x4f, 0x85, 0x8b, 0x69, 0x04, 0xb3, 0x0c, 0x4d, 0xaa, 0x65, 0x69, 0xa5, 0x2a, 0xd8, 0xc0,
	0x95, 0x5e, 0xa6, 0xe8, 0x02, 0x82, 0x52, 0xcb, 0x14, 0xd9, 0x30, 0x22, 0x31, 0x11, 0x27, 0x40,
	0x19, 0x8c, 0x4b, 0x99, 0xda, 0x4a, 0x23, 0x0b, 0x1c, 0xa7, 0x81, 0xf4, 0x03, 0x8c, 0x8c, 0x4d,
	0x6c, 0x65, 0xd8, 0x28, 0x22, 0xf1, 0x44, 0x78, 0xc4, 0x3f, 0xc2, 0xfb, 0xff, 0x68, 0xff, 0xe4,
	0xb9, 0x3f, 0xcf, 0x08, 0xbc, 0xaf, 0xd0, 0x58, 0xfe, 0x0d, 0x16, 0x7f, 0x31, 0x47, 0x8b, 0xbe,
	0xe0, 0xf3, 0x97, 0xe7, 0x73, 0x03, 0x6f, 0x77, 0xb2, 0x38, 0xe4, 0x78, 0xa3, 0xf6, 0x02, 0x4d,
	0xa9, 0x0a, 0x83, 0xf4, 0x17, 0x84, 0xa9, 0xc6, 0xc4, 0x62, 0xe6, 0xd9, 0x8e, 0x30, 0x5b, 0xbf,
	0x59, 0x35, 0x02, 0x35, 0x53, 0x2f, 0xfa, 0xe8, 0x57, 0x08, 0x50, 0x6b, 0xa5, 0x9d, 0x1c, 0xb3,
	0x75, 0xd8, 0x12, 0xfe, 0xd5, 0x59, 0x71, 0x2a, 0xf2, 0x0d, 0xbc, 0xbb, 0xab, 0x72, 0x2b, 0xcb,
	0xf3, 0xb5, 0xdf, 0x61, 0xe2, 0xdb, 0x0d, 0x23, 0xd1, 0xa0, 0x73, 0x61, 0xdb, 0xc1, 0x7f, 0x42,
	0xe0, 0x86, 0xd6, 0x0e, 0xa4, 0x2a, 0x43, 0x77, 0xe3, 0x54, 0xb8, 0xb8, 0x56, 0xf2, 0x88, 0xc6,
	0x24, 0x87, 0xc6, 0x98, 0x06, 0xae, 0x9f, 0x08, 0x84, 0x7e, 0xd8, 0x0e, 0xf5, 0x43, 0x2d, 0xfb,
	0x6f, 0x98, 0x6f, 0xdc, 0x33, 0x9a, 0x57, 0x5c, 0xad, 0x5d, 0x2e, 0xdb, 0xcc, 0x95, 0x5a, 0xbc,
	0x47, 0xb7, 0x10, 0x9e, 0xbb, 0x40, 0xbf, 0xb4, 0xfd, 0x9d, 0xf6, 0x2c, 0x3f, 0xb5, 0xf5, 0x0e,
	0x21, 0x78, 0x8f, 0xde, 0xc2, 0xfc, 0xcc, 0x3e, 0xfa, 0xb9, 0x25, 0x74, 0xd9, 0xfa, 0xfa, 0x7d,
	0xfb, 0x91, 0xfb, 0xcf, 0x3f, 0x9e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc8, 0x5e, 0x10, 0xe0,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Method to create a product
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*SingleJobResponse, error)
	// Method to get all products
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*MultipleJobResponse, error)
	// Method to delete a product
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*SingleJobResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*SingleJobResponse, error) {
	out := new(SingleJobResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*MultipleJobResponse, error) {
	out := new(MultipleJobResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*SingleJobResponse, error) {
	out := new(SingleJobResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Method to create a product
	CreateProduct(context.Context, *Product) (*SingleJobResponse, error)
	// Method to get all products
	GetAllProducts(context.Context, *GetAllProductsRequest) (*MultipleJobResponse, error)
	// Method to delete a product
	DeleteProduct(context.Context, *DeleteProductRequest) (*SingleJobResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) CreateProduct(ctx context.Context, req *Product) (*SingleJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedProductServiceServer) GetAllProducts(ctx context.Context, req *GetAllProductsRequest) (*MultipleJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (*UnimplementedProductServiceServer) DeleteProduct(ctx context.Context, req *DeleteProductRequest) (*SingleJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _ProductService_GetAllProducts_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
