// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: core.proto

package humanize_protobuf

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

type ConversationEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AskerName           string `protobuf:"bytes,1,opt,name=askerName,proto3" json:"askerName,omitempty"`
	AskerId             string `protobuf:"bytes,2,opt,name=askerId,proto3" json:"askerId,omitempty"`
	ResponderName       string `protobuf:"bytes,3,opt,name=responderName,proto3" json:"responderName,omitempty"`
	ResponderId         string `protobuf:"bytes,4,opt,name=responderId,proto3" json:"responderId,omitempty"`
	MessageEmotion      string `protobuf:"bytes,5,opt,name=message_emotion,json=messageEmotion,proto3" json:"message_emotion,omitempty"`
	Message             string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Response            string `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`
	ConversationEntryId int32  `protobuf:"varint,10,opt,name=conversation_entry_id,json=conversationEntryId,proto3" json:"conversation_entry_id,omitempty"`
	ResponseEmotion     string `protobuf:"bytes,11,opt,name=response_emotion,json=responseEmotion,proto3" json:"response_emotion,omitempty"`
}

func (x *ConversationEntry) Reset() {
	*x = ConversationEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationEntry) ProtoMessage() {}

func (x *ConversationEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationEntry.ProtoReflect.Descriptor instead.
func (*ConversationEntry) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationEntry) GetAskerName() string {
	if x != nil {
		return x.AskerName
	}
	return ""
}

func (x *ConversationEntry) GetAskerId() string {
	if x != nil {
		return x.AskerId
	}
	return ""
}

func (x *ConversationEntry) GetResponderName() string {
	if x != nil {
		return x.ResponderName
	}
	return ""
}

func (x *ConversationEntry) GetResponderId() string {
	if x != nil {
		return x.ResponderId
	}
	return ""
}

func (x *ConversationEntry) GetMessageEmotion() string {
	if x != nil {
		return x.MessageEmotion
	}
	return ""
}

func (x *ConversationEntry) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConversationEntry) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ConversationEntry) GetConversationEntryId() int32 {
	if x != nil {
		return x.ConversationEntryId
	}
	return 0
}

func (x *ConversationEntry) GetResponseEmotion() string {
	if x != nil {
		return x.ResponseEmotion
	}
	return ""
}

type MemoryLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log []*ConversationEntry `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
}

func (x *MemoryLog) Reset() {
	*x = MemoryLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLog) ProtoMessage() {}

func (x *MemoryLog) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLog.ProtoReflect.Descriptor instead.
func (*MemoryLog) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *MemoryLog) GetLog() []*ConversationEntry {
	if x != nil {
		return x.Log
	}
	return nil
}

type GenerationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenerationConfigId     string   `protobuf:"bytes,1,opt,name=generation_config_id,json=generationConfigId,proto3" json:"generation_config_id,omitempty"`
	Temperature            float32  `protobuf:"fixed32,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	TopP                   float32  `protobuf:"fixed32,5,opt,name=top_p,json=topP,proto3" json:"top_p,omitempty"`
	TopK                   int64    `protobuf:"varint,6,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	RepetitionPenalty      float32  `protobuf:"fixed32,7,opt,name=repetition_penalty,json=repetitionPenalty,proto3" json:"repetition_penalty,omitempty"`
	Length                 int64    `protobuf:"varint,8,opt,name=length,proto3" json:"length,omitempty"`
	StopSequences          []string `protobuf:"bytes,9,rep,name=stop_sequences,json=stopSequences,proto3" json:"stop_sequences,omitempty"`
	CharacterPrompt        string   `protobuf:"bytes,10,opt,name=character_prompt,json=characterPrompt,proto3" json:"character_prompt,omitempty"`
	WordsInEmotionalPrimer int32    `protobuf:"varint,11,opt,name=words_in_emotional_primer,json=wordsInEmotionalPrimer,proto3" json:"words_in_emotional_primer,omitempty"`
	MaxMemoryLogEntries    int32    `protobuf:"varint,12,opt,name=max_memory_log_entries,json=maxMemoryLogEntries,proto3" json:"max_memory_log_entries,omitempty"`
}

func (x *GenerationConfig) Reset() {
	*x = GenerationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationConfig) ProtoMessage() {}

func (x *GenerationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationConfig.ProtoReflect.Descriptor instead.
func (*GenerationConfig) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *GenerationConfig) GetGenerationConfigId() string {
	if x != nil {
		return x.GenerationConfigId
	}
	return ""
}

func (x *GenerationConfig) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GenerationConfig) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

func (x *GenerationConfig) GetTopK() int64 {
	if x != nil {
		return x.TopK
	}
	return 0
}

func (x *GenerationConfig) GetRepetitionPenalty() float32 {
	if x != nil {
		return x.RepetitionPenalty
	}
	return 0
}

func (x *GenerationConfig) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GenerationConfig) GetStopSequences() []string {
	if x != nil {
		return x.StopSequences
	}
	return nil
}

func (x *GenerationConfig) GetCharacterPrompt() string {
	if x != nil {
		return x.CharacterPrompt
	}
	return ""
}

func (x *GenerationConfig) GetWordsInEmotionalPrimer() int32 {
	if x != nil {
		return x.WordsInEmotionalPrimer
	}
	return 0
}

func (x *GenerationConfig) GetMaxMemoryLogEntries() int32 {
	if x != nil {
		return x.MaxMemoryLogEntries
	}
	return 0
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x6b,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x09, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x4c, 0x6f, 0x67, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x99,
	0x03, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x12, 0x13, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x70,
	0x4b, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x72,
	0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x70,
	0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_proto_goTypes = []interface{}{
	(*ConversationEntry)(nil), // 0: core.ConversationEntry
	(*MemoryLog)(nil),         // 1: core.MemoryLog
	(*GenerationConfig)(nil),  // 2: core.GenerationConfig
}
var file_core_proto_depIdxs = []int32{
	0, // 0: core.MemoryLog.log:type_name -> core.ConversationEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationEntry); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLog); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerationConfig); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}
