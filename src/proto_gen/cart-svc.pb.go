// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.1
// source: cart-svc.proto

package proto_gen

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cart_svc_proto protoreflect.FileDescriptor

var file_cart_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x61, 0x72, 0x74, 0x53, 0x76, 0x63, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x97, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x0d,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f,
	0x1a, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44,
	0x74, 0x6f, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x72,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_cart_svc_proto_goTypes = []interface{}{
	(*empty.Empty)(nil),        // 0: google.protobuf.Empty
	(*CartDto)(nil),            // 1: cart.CartDto
	(*ListOptions)(nil),        // 2: common.ListOptions
	(*Health)(nil),             // 3: common.Health
	(*wrappers.BoolValue)(nil), // 4: google.protobuf.BoolValue
	(*ListCartDto)(nil),        // 5: cart.ListCartDto
}
var file_cart_svc_proto_depIdxs = []int32{
	0, // 0: cartSvc.Cart.GetHealth:input_type -> google.protobuf.Empty
	1, // 1: cartSvc.Cart.AddCartItem:input_type -> cart.CartDto
	1, // 2: cartSvc.Cart.UpdateCartItem:input_type -> cart.CartDto
	1, // 3: cartSvc.Cart.DeleteCartItem:input_type -> cart.CartDto
	2, // 4: cartSvc.Cart.ListCartItem:input_type -> common.ListOptions
	3, // 5: cartSvc.Cart.GetHealth:output_type -> common.Health
	1, // 6: cartSvc.Cart.AddCartItem:output_type -> cart.CartDto
	1, // 7: cartSvc.Cart.UpdateCartItem:output_type -> cart.CartDto
	4, // 8: cartSvc.Cart.DeleteCartItem:output_type -> google.protobuf.BoolValue
	5, // 9: cartSvc.Cart.ListCartItem:output_type -> cart.ListCartDto
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cart_svc_proto_init() }
func file_cart_svc_proto_init() {
	if File_cart_svc_proto != nil {
		return
	}
	file_common_proto_init()
	file_cart_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cart_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_svc_proto_goTypes,
		DependencyIndexes: file_cart_svc_proto_depIdxs,
	}.Build()
	File_cart_svc_proto = out.File
	file_cart_svc_proto_rawDesc = nil
	file_cart_svc_proto_goTypes = nil
	file_cart_svc_proto_depIdxs = nil
}
