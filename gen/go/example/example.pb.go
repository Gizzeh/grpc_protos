// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.6.1
// source: example/example.proto

package example

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

type GreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetingRequest) Reset() {
	*x = GreetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRequest) ProtoMessage() {}

func (x *GreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRequest.ProtoReflect.Descriptor instead.
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetNumbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetNumbersRequest) Reset() {
	*x = GetNumbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumbersRequest) ProtoMessage() {}

func (x *GetNumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumbersRequest.ProtoReflect.Descriptor instead.
func (*GetNumbersRequest) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2}
}

func (x *GetNumbersRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{3}
}

func (x *Number) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SendNumbersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendNumbersResponse) Reset() {
	*x = SendNumbersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNumbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNumbersResponse) ProtoMessage() {}

func (x *SendNumbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNumbersResponse.ProtoReflect.Descriptor instead.
func (*SendNumbersResponse) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{4}
}

func (x *SendNumbersResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x22, 0x25, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x20, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x86, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x12, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_example_proto_goTypes = []interface{}{
	(*GreetingRequest)(nil),     // 0: example.GreetingRequest
	(*GreetingResponse)(nil),    // 1: example.GreetingResponse
	(*GetNumbersRequest)(nil),   // 2: example.GetNumbersRequest
	(*Number)(nil),              // 3: example.Number
	(*SendNumbersResponse)(nil), // 4: example.SendNumbersResponse
}
var file_example_example_proto_depIdxs = []int32{
	0, // 0: example.ExampleService.Greeting:input_type -> example.GreetingRequest
	2, // 1: example.ExampleService.GetNumbers:input_type -> example.GetNumbersRequest
	3, // 2: example.ExampleService.SendNumbers:input_type -> example.Number
	3, // 3: example.ExampleService.Square:input_type -> example.Number
	1, // 4: example.ExampleService.Greeting:output_type -> example.GreetingResponse
	3, // 5: example.ExampleService.GetNumbers:output_type -> example.Number
	4, // 6: example.ExampleService.SendNumbers:output_type -> example.SendNumbersResponse
	3, // 7: example.ExampleService.Square:output_type -> example.Number
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingRequest); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetingResponse); i {
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
		file_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumbersRequest); i {
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
		file_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_example_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNumbersResponse); i {
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
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}
