// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: gRPCServer/gRPCServer.proto

package team00v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64                `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"` // Норм распределение чего-то там
	Time      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPCServer_gRPCServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPCServer_gRPCServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_gRPCServer_gRPCServer_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ConnectResponse) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *ConnectResponse) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_gRPCServer_gRPCServer_proto protoreflect.FileDescriptor

var file_gRPCServer_gRPCServer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x52, 0x50,
	0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65,
	0x78, 0x30, 0x30, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7e, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x32, 0x40, 0x0a, 0x04, 0x45, 0x78, 0x30, 0x30, 0x12, 0x38, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x65,
	0x78, 0x30, 0x30, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x6f, 0x74, 0x65, 0x61, 0x6d, 0x31, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x30, 0x30, 0x2e, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x30, 0x30, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPCServer_gRPCServer_proto_rawDescOnce sync.Once
	file_gRPCServer_gRPCServer_proto_rawDescData = file_gRPCServer_gRPCServer_proto_rawDesc
)

func file_gRPCServer_gRPCServer_proto_rawDescGZIP() []byte {
	file_gRPCServer_gRPCServer_proto_rawDescOnce.Do(func() {
		file_gRPCServer_gRPCServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPCServer_gRPCServer_proto_rawDescData)
	})
	return file_gRPCServer_gRPCServer_proto_rawDescData
}

var file_gRPCServer_gRPCServer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gRPCServer_gRPCServer_proto_goTypes = []any{
	(*ConnectResponse)(nil),       // 0: ex00.ConnectResponse
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_gRPCServer_gRPCServer_proto_depIdxs = []int32{
	1, // 0: ex00.ConnectResponse.time:type_name -> google.protobuf.Timestamp
	2, // 1: ex00.Ex00.Connect:input_type -> google.protobuf.Empty
	0, // 2: ex00.Ex00.Connect:output_type -> ex00.ConnectResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gRPCServer_gRPCServer_proto_init() }
func file_gRPCServer_gRPCServer_proto_init() {
	if File_gRPCServer_gRPCServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPCServer_gRPCServer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectResponse); i {
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
			RawDescriptor: file_gRPCServer_gRPCServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPCServer_gRPCServer_proto_goTypes,
		DependencyIndexes: file_gRPCServer_gRPCServer_proto_depIdxs,
		MessageInfos:      file_gRPCServer_gRPCServer_proto_msgTypes,
	}.Build()
	File_gRPCServer_gRPCServer_proto = out.File
	file_gRPCServer_gRPCServer_proto_rawDesc = nil
	file_gRPCServer_gRPCServer_proto_goTypes = nil
	file_gRPCServer_gRPCServer_proto_depIdxs = nil
}
