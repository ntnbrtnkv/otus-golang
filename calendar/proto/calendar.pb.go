// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: calendar.proto

package grpc_gen

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Result) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	TimeStart   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=TimeStart,proto3" json:"TimeStart,omitempty"`
	TimeEnd     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=TimeEnd,proto3" json:"TimeEnd,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetTimeStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStart
	}
	return nil
}

func (x *Event) GetTimeEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeEnd
	}
	return nil
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type IDEventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDEventMessage) Reset() {
	*x = IDEventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDEventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDEventMessage) ProtoMessage() {}

func (x *IDEventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDEventMessage.ProtoReflect.Descriptor instead.
func (*IDEventMessage) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *IDEventMessage) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_calendar_proto protoreflect.FileDescriptor

var file_calendar_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xbf,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x54, 0x69,
	0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64,
	0x22, 0x28, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x49, 0x44,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x9f, 0x01, 0x0a,
	0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x07, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x49, 0x44, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1e, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6e,
	0x62, 0x72, 0x74, 0x6e, 0x6b, 0x76, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calendar_proto_rawDescOnce sync.Once
	file_calendar_proto_rawDescData = file_calendar_proto_rawDesc
)

func file_calendar_proto_rawDescGZIP() []byte {
	file_calendar_proto_rawDescOnce.Do(func() {
		file_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendar_proto_rawDescData)
	})
	return file_calendar_proto_rawDescData
}

var file_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calendar_proto_goTypes = []interface{}{
	(*Result)(nil),                // 0: Result
	(*Event)(nil),                 // 1: Event
	(*Events)(nil),                // 2: Events
	(*IDEventMessage)(nil),        // 3: IDEventMessage
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_calendar_proto_depIdxs = []int32{
	4, // 0: Event.TimeStart:type_name -> google.protobuf.Timestamp
	4, // 1: Event.TimeEnd:type_name -> google.protobuf.Timestamp
	1, // 2: Events.events:type_name -> Event
	1, // 3: Calendar.AddEvent:input_type -> Event
	3, // 4: Calendar.RemoveEvent:input_type -> IDEventMessage
	1, // 5: Calendar.ChangeEvent:input_type -> Event
	5, // 6: Calendar.ListEvents:input_type -> google.protobuf.Empty
	0, // 7: Calendar.AddEvent:output_type -> Result
	0, // 8: Calendar.RemoveEvent:output_type -> Result
	0, // 9: Calendar.ChangeEvent:output_type -> Result
	2, // 10: Calendar.ListEvents:output_type -> Events
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_calendar_proto_init() }
func file_calendar_proto_init() {
	if File_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
		file_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDEventMessage); i {
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
			RawDescriptor: file_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendar_proto_goTypes,
		DependencyIndexes: file_calendar_proto_depIdxs,
		MessageInfos:      file_calendar_proto_msgTypes,
	}.Build()
	File_calendar_proto = out.File
	file_calendar_proto_rawDesc = nil
	file_calendar_proto_goTypes = nil
	file_calendar_proto_depIdxs = nil
}
