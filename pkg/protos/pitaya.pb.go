// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pitaya.proto

package protos

import (
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

var File_pitaya_proto protoreflect.FileDescriptor

var file_pitaya_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x69, 0x74, 0x61, 0x79, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6b,
	0x69, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd2, 0x01, 0x0a, 0x06, 0x50, 0x69,
	0x74, 0x61, 0x79, 0x61, 0x12, 0x2b, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x4b,
	0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x42, 0x3c,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x70,
	0x66, 0x72, 0x65, 0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x74, 0x61, 0x79, 0x61,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xaa, 0x02, 0x0e, 0x4e, 0x50,
	0x69, 0x74, 0x61, 0x79, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_pitaya_proto_goTypes = []interface{}{
	(*Request)(nil),    // 0: protos.Request
	(*Push)(nil),       // 1: protos.Push
	(*BindMsg)(nil),    // 2: protos.BindMsg
	(*KickMsg)(nil),    // 3: protos.KickMsg
	(*Response)(nil),   // 4: protos.Response
	(*KickAnswer)(nil), // 5: protos.KickAnswer
}
var file_pitaya_proto_depIdxs = []int32{
	0, // 0: protos.Pitaya.Call:input_type -> protos.Request
	1, // 1: protos.Pitaya.PushToUser:input_type -> protos.Push
	2, // 2: protos.Pitaya.SessionBindRemote:input_type -> protos.BindMsg
	3, // 3: protos.Pitaya.KickUser:input_type -> protos.KickMsg
	4, // 4: protos.Pitaya.Call:output_type -> protos.Response
	4, // 5: protos.Pitaya.PushToUser:output_type -> protos.Response
	4, // 6: protos.Pitaya.SessionBindRemote:output_type -> protos.Response
	5, // 7: protos.Pitaya.KickUser:output_type -> protos.KickAnswer
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pitaya_proto_init() }
func file_pitaya_proto_init() {
	if File_pitaya_proto != nil {
		return
	}
	file_request_proto_init()
	file_response_proto_init()
	file_push_proto_init()
	file_bind_proto_init()
	file_kick_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pitaya_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pitaya_proto_goTypes,
		DependencyIndexes: file_pitaya_proto_depIdxs,
	}.Build()
	File_pitaya_proto = out.File
	file_pitaya_proto_rawDesc = nil
	file_pitaya_proto_goTypes = nil
	file_pitaya_proto_depIdxs = nil
}
