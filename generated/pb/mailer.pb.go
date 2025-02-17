// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: mailer.proto

package pb

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

type SendHTMLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        *Sender                `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To            []string               `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Html          string                 `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3,oneof" json:"scheduled_time,omitempty"`
}

func (x *SendHTMLReq) Reset() {
	*x = SendHTMLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHTMLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHTMLReq) ProtoMessage() {}

func (x *SendHTMLReq) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHTMLReq.ProtoReflect.Descriptor instead.
func (*SendHTMLReq) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *SendHTMLReq) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendHTMLReq) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendHTMLReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendHTMLReq) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

func (x *SendHTMLReq) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

type SendTemplateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        *Sender                `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To            []string               `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	TemplateId    string                 `protobuf:"bytes,4,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3,oneof" json:"scheduled_time,omitempty"`
}

func (x *SendTemplateReq) Reset() {
	*x = SendTemplateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTemplateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTemplateReq) ProtoMessage() {}

func (x *SendTemplateReq) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTemplateReq.ProtoReflect.Descriptor instead.
func (*SendTemplateReq) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *SendTemplateReq) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendTemplateReq) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendTemplateReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendTemplateReq) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendTemplateReq) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

type SendRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId     string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	TemplateId    string                 `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
}

func (x *SendRes) Reset() {
	*x = SendRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRes) ProtoMessage() {}

func (x *SendRes) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRes.ProtoReflect.Descriptor instead.
func (*SendRes) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{2}
}

func (x *SendRes) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SendRes) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendRes) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

type Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *Sender) Reset() {
	*x = Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sender) ProtoMessage() {}

func (x *Sender) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sender.ProtoReflect.Descriptor instead.
func (*Sender) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{3}
}

func (x *Sender) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Sender) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

var File_mailer_proto protoreflect.FileDescriptor

var file_mailer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x12, 0x46, 0x0a,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b,
	0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x46, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x07, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x32,
	0x78, 0x0a, 0x06, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x48, 0x54, 0x4d, 0x4c, 0x12, 0x13, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6b, 0x61, 0x6e,
	0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mailer_proto_rawDescOnce sync.Once
	file_mailer_proto_rawDescData = file_mailer_proto_rawDesc
)

func file_mailer_proto_rawDescGZIP() []byte {
	file_mailer_proto_rawDescOnce.Do(func() {
		file_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailer_proto_rawDescData)
	})
	return file_mailer_proto_rawDescData
}

var file_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mailer_proto_goTypes = []interface{}{
	(*SendHTMLReq)(nil),           // 0: kannon.SendHTMLReq
	(*SendTemplateReq)(nil),       // 1: kannon.SendTemplateReq
	(*SendRes)(nil),               // 2: kannon.SendRes
	(*Sender)(nil),                // 3: kannon.Sender
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_mailer_proto_depIdxs = []int32{
	3, // 0: kannon.SendHTMLReq.sender:type_name -> kannon.Sender
	4, // 1: kannon.SendHTMLReq.scheduled_time:type_name -> google.protobuf.Timestamp
	3, // 2: kannon.SendTemplateReq.sender:type_name -> kannon.Sender
	4, // 3: kannon.SendTemplateReq.scheduled_time:type_name -> google.protobuf.Timestamp
	4, // 4: kannon.SendRes.scheduled_time:type_name -> google.protobuf.Timestamp
	0, // 5: kannon.Mailer.SendHTML:input_type -> kannon.SendHTMLReq
	1, // 6: kannon.Mailer.SendTemplate:input_type -> kannon.SendTemplateReq
	2, // 7: kannon.Mailer.SendHTML:output_type -> kannon.SendRes
	2, // 8: kannon.Mailer.SendTemplate:output_type -> kannon.SendRes
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mailer_proto_init() }
func file_mailer_proto_init() {
	if File_mailer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHTMLReq); i {
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
		file_mailer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTemplateReq); i {
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
		file_mailer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRes); i {
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
		file_mailer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sender); i {
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
	file_mailer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_mailer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailer_proto_goTypes,
		DependencyIndexes: file_mailer_proto_depIdxs,
		MessageInfos:      file_mailer_proto_msgTypes,
	}.Build()
	File_mailer_proto = out.File
	file_mailer_proto_rawDesc = nil
	file_mailer_proto_goTypes = nil
	file_mailer_proto_depIdxs = nil
}
