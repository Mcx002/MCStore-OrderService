// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.1
// source: cart.proto

package proto_gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *UserDto          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Food      *FoodDto          `protobuf:"bytes,2,opt,name=food,proto3" json:"food,omitempty"`
	Seller    *UserDto          `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	Qty       int32             `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Variety   map[string]string `protobuf:"bytes,5,rep,name=variety,proto3" json:"variety,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Note      string            `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	Marked    bool              `protobuf:"varint,7,opt,name=marked,proto3" json:"marked,omitempty"`
	Version   int32             `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt int64             `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64             `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id        string            `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CartDto) Reset() {
	*x = CartDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartDto) ProtoMessage() {}

func (x *CartDto) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartDto.ProtoReflect.Descriptor instead.
func (*CartDto) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartDto) GetUser() *UserDto {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CartDto) GetFood() *FoodDto {
	if x != nil {
		return x.Food
	}
	return nil
}

func (x *CartDto) GetSeller() *UserDto {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *CartDto) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *CartDto) GetVariety() map[string]string {
	if x != nil {
		return x.Variety
	}
	return nil
}

func (x *CartDto) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CartDto) GetMarked() bool {
	if x != nil {
		return x.Marked
	}
	return false
}

func (x *CartDto) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CartDto) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CartDto) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *CartDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListCartDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carts    []*CartDto    `protobuf:"bytes,1,rep,name=carts,proto3" json:"carts,omitempty"`
	Metadata *ListMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ListCartDto) Reset() {
	*x = ListCartDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCartDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCartDto) ProtoMessage() {}

func (x *ListCartDto) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCartDto.ProtoReflect.Descriptor instead.
func (*ListCartDto) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *ListCartDto) GetCarts() []*CartDto {
	if x != nil {
		return x.Carts
	}
	return nil
}

func (x *ListCartDto) GetMetadata() *ListMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x07, 0x43, 0x61, 0x72,
	0x74, 0x44, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74,
	0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f,
	0x64, 0x44, 0x74, 0x6f, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x44, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x3a, 0x0a,
	0x0c, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x64, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x63, 0x61, 0x72, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x0f, 0x5a, 0x0d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cart_proto_goTypes = []interface{}{
	(*CartDto)(nil),      // 0: cart.CartDto
	(*ListCartDto)(nil),  // 1: cart.ListCartDto
	nil,                  // 2: cart.CartDto.VarietyEntry
	(*UserDto)(nil),      // 3: user.UserDto
	(*FoodDto)(nil),      // 4: food.FoodDto
	(*ListMetadata)(nil), // 5: common.ListMetadata
}
var file_cart_proto_depIdxs = []int32{
	3, // 0: cart.CartDto.user:type_name -> user.UserDto
	4, // 1: cart.CartDto.food:type_name -> food.FoodDto
	3, // 2: cart.CartDto.seller:type_name -> user.UserDto
	2, // 3: cart.CartDto.variety:type_name -> cart.CartDto.VarietyEntry
	0, // 4: cart.ListCartDto.carts:type_name -> cart.CartDto
	5, // 5: cart.ListCartDto.metadata:type_name -> common.ListMetadata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	file_user_proto_init()
	file_food_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCartDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
