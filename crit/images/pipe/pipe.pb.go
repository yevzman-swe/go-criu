// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pipe.proto

package pipe

import (
	fown "github.com/checkpoint-restore/go-criu/v6/crit/images/fown"
	_ "github.com/checkpoint-restore/go-criu/v6/crit/images/opts"
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

type PipeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *uint32         `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	PipeId *uint32         `protobuf:"varint,2,req,name=pipe_id,json=pipeId" json:"pipe_id,omitempty"`
	Flags  *uint32         `protobuf:"varint,3,req,name=flags" json:"flags,omitempty"`
	Fown   *fown.FownEntry `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
}

func (x *PipeEntry) Reset() {
	*x = PipeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipeEntry) ProtoMessage() {}

func (x *PipeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipeEntry.ProtoReflect.Descriptor instead.
func (*PipeEntry) Descriptor() ([]byte, []int) {
	return file_pipe_proto_rawDescGZIP(), []int{0}
}

func (x *PipeEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PipeEntry) GetPipeId() uint32 {
	if x != nil && x.PipeId != nil {
		return *x.PipeId
	}
	return 0
}

func (x *PipeEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *PipeEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

var File_pipe_proto protoreflect.FileDescriptor

var file_pipe_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f, 0x70,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08,
	0x01, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x66, 0x6f, 0x77, 0x6e,
}

var (
	file_pipe_proto_rawDescOnce sync.Once
	file_pipe_proto_rawDescData = file_pipe_proto_rawDesc
)

func file_pipe_proto_rawDescGZIP() []byte {
	file_pipe_proto_rawDescOnce.Do(func() {
		file_pipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipe_proto_rawDescData)
	})
	return file_pipe_proto_rawDescData
}

var file_pipe_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pipe_proto_goTypes = []interface{}{
	(*PipeEntry)(nil),      // 0: pipe_entry
	(*fown.FownEntry)(nil), // 1: fown_entry
}
var file_pipe_proto_depIdxs = []int32{
	1, // 0: pipe_entry.fown:type_name -> fown_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pipe_proto_init() }
func file_pipe_proto_init() {
	if File_pipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipeEntry); i {
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
			RawDescriptor: file_pipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pipe_proto_goTypes,
		DependencyIndexes: file_pipe_proto_depIdxs,
		MessageInfos:      file_pipe_proto_msgTypes,
	}.Build()
	File_pipe_proto = out.File
	file_pipe_proto_rawDesc = nil
	file_pipe_proto_goTypes = nil
	file_pipe_proto_depIdxs = nil
}