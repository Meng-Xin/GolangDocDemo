// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc3
// source: deviceCollection.proto

// 指定包名

package pb

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

type DeviceStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId      string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StatusMessage string `protobuf:"bytes,3,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"` // 1:正常 2:故障 3:关闭
}

func (x *DeviceStatusUpdate) Reset() {
	*x = DeviceStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceCollection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatusUpdate) ProtoMessage() {}

func (x *DeviceStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_deviceCollection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatusUpdate.ProtoReflect.Descriptor instead.
func (*DeviceStatusUpdate) Descriptor() ([]byte, []int) {
	return file_deviceCollection_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceStatusUpdate) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceStatusUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceStatusUpdate) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type DeviceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeviceStatusResponse) Reset() {
	*x = DeviceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceCollection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatusResponse) ProtoMessage() {}

func (x *DeviceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deviceCollection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatusResponse.ProtoReflect.Descriptor instead.
func (*DeviceStatusResponse) Descriptor() ([]byte, []int) {
	return file_deviceCollection_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_deviceCollection_proto protoreflect.FileDescriptor

var file_deviceCollection_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x6c, 0x0a, 0x12,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x66, 0x0a, 0x22,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deviceCollection_proto_rawDescOnce sync.Once
	file_deviceCollection_proto_rawDescData = file_deviceCollection_proto_rawDesc
)

func file_deviceCollection_proto_rawDescGZIP() []byte {
	file_deviceCollection_proto_rawDescOnce.Do(func() {
		file_deviceCollection_proto_rawDescData = protoimpl.X.CompressGZIP(file_deviceCollection_proto_rawDescData)
	})
	return file_deviceCollection_proto_rawDescData
}

var file_deviceCollection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deviceCollection_proto_goTypes = []interface{}{
	(*DeviceStatusUpdate)(nil),   // 0: pb.DeviceStatusUpdate
	(*DeviceStatusResponse)(nil), // 1: pb.DeviceStatusResponse
}
var file_deviceCollection_proto_depIdxs = []int32{
	0, // 0: pb.DeviceInformationCollectionService.SendStatus:input_type -> pb.DeviceStatusUpdate
	1, // 1: pb.DeviceInformationCollectionService.SendStatus:output_type -> pb.DeviceStatusResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deviceCollection_proto_init() }
func file_deviceCollection_proto_init() {
	if File_deviceCollection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deviceCollection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceStatusUpdate); i {
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
		file_deviceCollection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceStatusResponse); i {
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
			RawDescriptor: file_deviceCollection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deviceCollection_proto_goTypes,
		DependencyIndexes: file_deviceCollection_proto_depIdxs,
		MessageInfos:      file_deviceCollection_proto_msgTypes,
	}.Build()
	File_deviceCollection_proto = out.File
	file_deviceCollection_proto_rawDesc = nil
	file_deviceCollection_proto_goTypes = nil
	file_deviceCollection_proto_depIdxs = nil
}