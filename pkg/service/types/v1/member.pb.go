// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: types/v1/member.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    string                 `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	TeamSlug  string                 `protobuf:"bytes,2,opt,name=team_slug,json=teamSlug,proto3" json:"team_slug,omitempty"`
	TeamName  string                 `protobuf:"bytes,3,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	UserId    string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserSlug  string                 `protobuf:"bytes,5,opt,name=user_slug,json=userSlug,proto3" json:"user_slug,omitempty"`
	UserName  string                 `protobuf:"bytes,6,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_types_v1_member_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Member) GetTeamSlug() string {
	if x != nil {
		return x.TeamSlug
	}
	return ""
}

func (x *Member) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *Member) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Member) GetUserSlug() string {
	if x != nil {
		return x.UserSlug
	}
	return ""
}

func (x *Member) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Member) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Member) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_types_v1_member_proto protoreflect.FileDescriptor

var file_types_v1_member_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x53,
	0x6c, 0x75, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x9b, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6c, 0x65, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x6b,
	0x6c, 0x65, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_member_proto_rawDescOnce sync.Once
	file_types_v1_member_proto_rawDescData = file_types_v1_member_proto_rawDesc
)

func file_types_v1_member_proto_rawDescGZIP() []byte {
	file_types_v1_member_proto_rawDescOnce.Do(func() {
		file_types_v1_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_member_proto_rawDescData)
	})
	return file_types_v1_member_proto_rawDescData
}

var file_types_v1_member_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_v1_member_proto_goTypes = []interface{}{
	(*Member)(nil),                // 0: types.v1.Member
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_types_v1_member_proto_depIdxs = []int32{
	1, // 0: types.v1.Member.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: types.v1.Member.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_v1_member_proto_init() }
func file_types_v1_member_proto_init() {
	if File_types_v1_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_v1_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_types_v1_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_member_proto_goTypes,
		DependencyIndexes: file_types_v1_member_proto_depIdxs,
		MessageInfos:      file_types_v1_member_proto_msgTypes,
	}.Build()
	File_types_v1_member_proto = out.File
	file_types_v1_member_proto_rawDesc = nil
	file_types_v1_member_proto_goTypes = nil
	file_types_v1_member_proto_depIdxs = nil
}