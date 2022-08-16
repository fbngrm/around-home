// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: match/v1/match.proto

package matchv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/type/datetime"
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

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Materials string  `protobuf:"bytes,2,opt,name=materials,proto3" json:"materials,omitempty"`
	Rating    uint32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Location  string  `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Radius    float64 `protobuf:"fixed64,5,opt,name=radius,proto3" json:"radius,omitempty"`
	Distance  float64 `protobuf:"fixed64,6,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_match_v1_match_proto_rawDescGZIP(), []int{0}
}

func (x *Partner) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Partner) GetMaterials() string {
	if x != nil {
		return x.Materials
	}
	return ""
}

func (x *Partner) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Partner) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Partner) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *Partner) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type MatchPartnersWithRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Material string `protobuf:"bytes,1,opt,name=material,proto3" json:"material,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"` // unused: Square meters of the floor
}

func (x *MatchPartnersWithRequestInput) Reset() {
	*x = MatchPartnersWithRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPartnersWithRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPartnersWithRequestInput) ProtoMessage() {}

func (x *MatchPartnersWithRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPartnersWithRequestInput.ProtoReflect.Descriptor instead.
func (*MatchPartnersWithRequestInput) Descriptor() ([]byte, []int) {
	return file_match_v1_match_proto_rawDescGZIP(), []int{1}
}

func (x *MatchPartnersWithRequestInput) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

func (x *MatchPartnersWithRequestInput) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type MatchPartnersWithRequestOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partner []*Partner `protobuf:"bytes,1,rep,name=partner,proto3" json:"partner,omitempty"`
}

func (x *MatchPartnersWithRequestOutput) Reset() {
	*x = MatchPartnersWithRequestOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_v1_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPartnersWithRequestOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPartnersWithRequestOutput) ProtoMessage() {}

func (x *MatchPartnersWithRequestOutput) ProtoReflect() protoreflect.Message {
	mi := &file_match_v1_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPartnersWithRequestOutput.ProtoReflect.Descriptor instead.
func (*MatchPartnersWithRequestOutput) Descriptor() ([]byte, []int) {
	return file_match_v1_match_proto_rawDescGZIP(), []int{2}
}

func (x *MatchPartnersWithRequestOutput) GetPartner() []*Partner {
	if x != nil {
		return x.Partner
	}
	return nil
}

var File_match_v1_match_proto protoreflect.FileDescriptor

var file_match_v1_match_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x57, 0x0a, 0x1d,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x1e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x32, 0xaa, 0x01, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x18, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x7d, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x62, 0x6e,
	0x67, 0x72, 0x6d, 0x2f, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x68, 0x6f, 0x6d, 0x65, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_v1_match_proto_rawDescOnce sync.Once
	file_match_v1_match_proto_rawDescData = file_match_v1_match_proto_rawDesc
)

func file_match_v1_match_proto_rawDescGZIP() []byte {
	file_match_v1_match_proto_rawDescOnce.Do(func() {
		file_match_v1_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_v1_match_proto_rawDescData)
	})
	return file_match_v1_match_proto_rawDescData
}

var file_match_v1_match_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_match_v1_match_proto_goTypes = []interface{}{
	(*Partner)(nil),                        // 0: match.v1.Partner
	(*MatchPartnersWithRequestInput)(nil),  // 1: match.v1.MatchPartnersWithRequestInput
	(*MatchPartnersWithRequestOutput)(nil), // 2: match.v1.MatchPartnersWithRequestOutput
}
var file_match_v1_match_proto_depIdxs = []int32{
	0, // 0: match.v1.MatchPartnersWithRequestOutput.partner:type_name -> match.v1.Partner
	1, // 1: match.v1.matchService.MatchPartnersWithRequest:input_type -> match.v1.MatchPartnersWithRequestInput
	2, // 2: match.v1.matchService.MatchPartnersWithRequest:output_type -> match.v1.MatchPartnersWithRequestOutput
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_match_v1_match_proto_init() }
func file_match_v1_match_proto_init() {
	if File_match_v1_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_v1_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
		file_match_v1_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchPartnersWithRequestInput); i {
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
		file_match_v1_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchPartnersWithRequestOutput); i {
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
			RawDescriptor: file_match_v1_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_v1_match_proto_goTypes,
		DependencyIndexes: file_match_v1_match_proto_depIdxs,
		MessageInfos:      file_match_v1_match_proto_msgTypes,
	}.Build()
	File_match_v1_match_proto = out.File
	file_match_v1_match_proto_rawDesc = nil
	file_match_v1_match_proto_goTypes = nil
	file_match_v1_match_proto_depIdxs = nil
}