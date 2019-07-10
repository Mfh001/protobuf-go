// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/weak2/test_weak.proto

package weak2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type WeakImportMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
}

func (x *WeakImportMessage2) Reset() {
	*x = WeakImportMessage2{}
}

func (x *WeakImportMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeakImportMessage2) ProtoMessage() {}

func (x *WeakImportMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_test_weak2_test_weak_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeakImportMessage2.ProtoReflect.Descriptor instead.
func (*WeakImportMessage2) Descriptor() ([]byte, []int) {
	return file_test_weak2_test_weak_proto_rawDescGZIP(), []int{0}
}

func (x *WeakImportMessage2) GetA() int32 {
	if x != nil && x.A != nil {
		return *x.A
	}
	return 0
}

var File_test_weak2_test_weak_proto protoreflect.FileDescriptor

var file_test_weak2_test_weak_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x77, 0x65, 0x61, 0x6b, 0x32, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x77, 0x65, 0x61, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x77, 0x65, 0x61, 0x6b, 0x22, 0x22, 0x0a, 0x12, 0x57, 0x65, 0x61, 0x6b, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x77, 0x65, 0x61, 0x6b, 0x32,
}

var (
	file_test_weak2_test_weak_proto_rawDescOnce sync.Once
	file_test_weak2_test_weak_proto_rawDescData = file_test_weak2_test_weak_proto_rawDesc
)

func file_test_weak2_test_weak_proto_rawDescGZIP() []byte {
	file_test_weak2_test_weak_proto_rawDescOnce.Do(func() {
		file_test_weak2_test_weak_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_weak2_test_weak_proto_rawDescData)
	})
	return file_test_weak2_test_weak_proto_rawDescData
}

var file_test_weak2_test_weak_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_weak2_test_weak_proto_goTypes = []interface{}{
	(*WeakImportMessage2)(nil), // 0: goproto.proto.test.weak.WeakImportMessage2
}
var file_test_weak2_test_weak_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_test_weak2_test_weak_proto_init() }
func file_test_weak2_test_weak_proto_init() {
	if File_test_weak2_test_weak_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_weak2_test_weak_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeakImportMessage2); i {
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
			RawDescriptor: file_test_weak2_test_weak_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_weak2_test_weak_proto_goTypes,
		DependencyIndexes: file_test_weak2_test_weak_proto_depIdxs,
		MessageInfos:      file_test_weak2_test_weak_proto_msgTypes,
	}.Build()
	File_test_weak2_test_weak_proto = out.File
	file_test_weak2_test_weak_proto_rawDesc = nil
	file_test_weak2_test_weak_proto_goTypes = nil
	file_test_weak2_test_weak_proto_depIdxs = nil
}
