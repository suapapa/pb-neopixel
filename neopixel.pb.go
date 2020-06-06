// A very simple protocol definition, consisting of only
// one message.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: neopixel.proto

package main

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LED struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index *uint32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Color *uint32 `protobuf:"varint,2,req,name=color" json:"color,omitempty"`
}

func (x *LED) Reset() {
	*x = LED{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neopixel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LED) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LED) ProtoMessage() {}

func (x *LED) ProtoReflect() protoreflect.Message {
	mi := &file_neopixel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LED.ProtoReflect.Descriptor instead.
func (*LED) Descriptor() ([]byte, []int) {
	return file_neopixel_proto_rawDescGZIP(), []int{0}
}

func (x *LED) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *LED) GetColor() uint32 {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return 0
}

type NeoPixel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clear *bool  `protobuf:"varint,1,req,name=clear" json:"clear,omitempty"`
	Strip []*LED `protobuf:"bytes,2,rep,name=strip" json:"strip,omitempty"` // a will be ignored
}

func (x *NeoPixel) Reset() {
	*x = NeoPixel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neopixel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeoPixel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeoPixel) ProtoMessage() {}

func (x *NeoPixel) ProtoReflect() protoreflect.Message {
	mi := &file_neopixel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeoPixel.ProtoReflect.Descriptor instead.
func (*NeoPixel) Descriptor() ([]byte, []int) {
	return file_neopixel_proto_rawDescGZIP(), []int{1}
}

func (x *NeoPixel) GetClear() bool {
	if x != nil && x.Clear != nil {
		return *x.Clear
	}
	return false
}

func (x *NeoPixel) GetStrip() []*LED {
	if x != nil {
		return x.Strip
	}
	return nil
}

var File_neopixel_proto protoreflect.FileDescriptor

var file_neopixel_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x6f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x03, 0x4c, 0x45, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x08, 0x4e, 0x65, 0x6f, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x05,
	0x63, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x05, 0x73, 0x74, 0x72, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x45, 0x44, 0x52, 0x05, 0x73, 0x74, 0x72, 0x69,
	0x70,
}

var (
	file_neopixel_proto_rawDescOnce sync.Once
	file_neopixel_proto_rawDescData = file_neopixel_proto_rawDesc
)

func file_neopixel_proto_rawDescGZIP() []byte {
	file_neopixel_proto_rawDescOnce.Do(func() {
		file_neopixel_proto_rawDescData = protoimpl.X.CompressGZIP(file_neopixel_proto_rawDescData)
	})
	return file_neopixel_proto_rawDescData
}

var file_neopixel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_neopixel_proto_goTypes = []interface{}{
	(*LED)(nil),      // 0: LED
	(*NeoPixel)(nil), // 1: NeoPixel
}
var file_neopixel_proto_depIdxs = []int32{
	0, // 0: NeoPixel.strip:type_name -> LED
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_neopixel_proto_init() }
func file_neopixel_proto_init() {
	if File_neopixel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neopixel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LED); i {
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
		file_neopixel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeoPixel); i {
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
			RawDescriptor: file_neopixel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_neopixel_proto_goTypes,
		DependencyIndexes: file_neopixel_proto_depIdxs,
		MessageInfos:      file_neopixel_proto_msgTypes,
	}.Build()
	File_neopixel_proto = out.File
	file_neopixel_proto_rawDesc = nil
	file_neopixel_proto_goTypes = nil
	file_neopixel_proto_depIdxs = nil
}