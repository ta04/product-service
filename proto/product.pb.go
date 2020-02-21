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

// Product message
type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

// GetAllProductsRequest message
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

// CreateProductsResponse message
type CreateProductResponse struct {
	CreatedProduct       *Product `protobuf:"bytes,1,opt,name=createdProduct,proto3" json:"createdProduct,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductResponse) Reset()         { *m = CreateProductResponse{} }
func (m *CreateProductResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductResponse) ProtoMessage()    {}
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *CreateProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductResponse.Unmarshal(m, b)
}
func (m *CreateProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductResponse.Marshal(b, m, deterministic)
}
func (m *CreateProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductResponse.Merge(m, src)
}
func (m *CreateProductResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductResponse.Size(m)
}
func (m *CreateProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductResponse proto.InternalMessageInfo

func (m *CreateProductResponse) GetCreatedProduct() *Product {
	if m != nil {
		return m.CreatedProduct
	}
	return nil
}

func (m *CreateProductResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetAllProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllProductsResponse) Reset()         { *m = GetAllProductsResponse{} }
func (m *GetAllProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllProductsResponse) ProtoMessage()    {}
func (*GetAllProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *GetAllProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllProductsResponse.Unmarshal(m, b)
}
func (m *GetAllProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllProductsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllProductsResponse.Merge(m, src)
}
func (m *GetAllProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllProductsResponse.Size(m)
}
func (m *GetAllProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllProductsResponse proto.InternalMessageInfo

func (m *GetAllProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

// Error message
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
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
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
	proto.RegisterType((*CreateProductResponse)(nil), "product.CreateProductResponse")
	proto.RegisterType((*GetAllProductsResponse)(nil), "product.GetAllProductsResponse")
	proto.RegisterType((*Error)(nil), "product.Error")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0xfd, 0x0a, 0x2c, 0xf0, 0x0d, 0x61, 0x63, 0x26, 0x82, 0x0d, 0x07, 0xdd, 0x6c, 0x3c, 0x70,
	0x30, 0x1c, 0x30, 0x26, 0x5e, 0x0d, 0x51, 0xaf, 0x66, 0xf9, 0x05, 0xeb, 0x76, 0x62, 0x9a, 0x00,
	0xad, 0x6d, 0x57, 0x7f, 0x8a, 0xff, 0xc1, 0x5f, 0x69, 0xb6, 0xb4, 0x28, 0x88, 0xb7, 0x79, 0xf3,
	0xfa, 0x66, 0xe6, 0xbd, 0x5d, 0x18, 0x6a, 0xa3, 0x44, 0x5d, 0xb9, 0x99, 0x36, 0xca, 0x29, 0xec,
	0x05, 0x98, 0x7f, 0x30, 0xe8, 0x3d, 0x6d, 0x6b, 0x4c, 0xa1, 0x25, 0x05, 0x67, 0x19, 0x9b, 0xfe,
	0x2f, 0x5a, 0x52, 0x20, 0x42, 0x67, 0x53, 0xae, 0x89, 0xb7, 0x7c, 0xc7, 0xd7, 0x98, 0xc1, 0x40,
	0x90, 0xad, 0x8c, 0xd4, 0x4e, 0xaa, 0x0d, 0x6f, 0x7b, 0xea, 0x67, 0x0b, 0x4f, 0x21, 0xd1, 0x46,
	0x56, 0xc4, 0x3b, 0x19, 0x9b, 0xb2, 0x62, 0x0b, 0x90, 0x43, 0x4f, 0xcb, 0xca, 0xd5, 0x86, 0x78,
	0xe2, 0x35, 0x11, 0xe2, 0x18, 0xba, 0xd6, 0x95, 0xae, 0xb6, 0xbc, 0x9b, 0xb1, 0x69, 0xbf, 0x08,
	0x28, 0x3f, 0x83, 0xd1, 0x23, 0xb9, 0xbb, 0xd5, 0x2a, 0x9c, 0x67, 0x0b, 0x7a, 0xad, 0xc9, 0xba,
	0xfc, 0x1d, 0x46, 0x0b, 0x43, 0xa5, 0xa3, 0x40, 0x14, 0x64, 0xb5, 0xda, 0x58, 0xc2, 0x5b, 0x48,
	0x2b, 0x4f, 0x88, 0xc0, 0x78, 0x2f, 0x83, 0xf9, 0xc9, 0x2c, 0x9a, 0x8f, 0x8a, 0x83, 0x77, 0x78,
	0x09, 0x09, 0x19, 0xa3, 0x8c, 0xb7, 0x3a, 0x98, 0xa7, 0x3b, 0xc1, 0x7d, 0xd3, 0x2d, 0xb6, 0x64,
	0xfe, 0x00, 0xe3, 0xc3, 0x8b, 0xc2, 0xe6, 0x2b, 0xe8, 0x07, 0x85, 0xe5, 0x2c, 0x6b, 0x1f, 0xdd,
	0xb9, 0x7b, 0x91, 0xdf, 0x40, 0xe2, 0xe7, 0x36, 0x01, 0x57, 0x4a, 0x50, 0x88, 0xdc, 0xd7, 0x4d,
	0x50, 0x6b, 0xb2, 0xb6, 0x7c, 0x89, 0xb9, 0x47, 0x38, 0xff, 0x64, 0x90, 0x86, 0x61, 0x4b, 0x32,
	0x6f, 0x4d, 0xaa, 0x0b, 0x18, 0xee, 0x45, 0x81, 0xbf, 0xd6, 0x4e, 0xce, 0x77, 0x9d, 0xa3, 0xa1,
	0xe5, 0xff, 0x70, 0x09, 0xe9, 0xbe, 0x2d, 0xfc, 0xd6, 0x1c, 0xfd, 0x02, 0x93, 0x8b, 0x3f, 0xf9,
	0x38, 0xf4, 0xb9, 0xeb, 0xff, 0xb3, 0xeb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xc1, 0x58,
	0x21, 0x78, 0x02, 0x00, 0x00,
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
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*CreateProductResponse, error)
	// Method to get all products
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Method to create a product
	CreateProduct(context.Context, *Product) (*CreateProductResponse, error)
	// Method to get all products
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) CreateProduct(ctx context.Context, req *Product) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedProductServiceServer) GetAllProducts(ctx context.Context, req *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
