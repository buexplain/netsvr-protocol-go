//*
// Copyright 2023 buexplain@qq.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: registerReq.proto

package netsvr

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

// business向worker请求，注册自己
// 注册逻辑：检查注册条件后，会给business的连接异步写入注册成功的信息、将business的连接注册到管理器，让business的连接接收网关转发的客户数据，如果注册失败，会返回失败的信息
// 如果不想接收来自客户的信息，只是与网关交互，可以不发起注册指令
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 表示business进程可以接收的事件，如果要接收多个，则按位或，具体的枚举值可以查看event.proto
	Events int32 `protobuf:"varint,1,opt,name=events,proto3" json:"events,omitempty"`
	// 该参数表示接下来，需要worker服务器开启多少协程来处理本business的请求
	// 如果本business，非常频繁的与worker交互,并且是那种组播、广播的耗时操作
	// 可以考虑开大一点，但是也不能无限大，开太多也许不能解决问题，因为发送消息到客户连接是会被阻塞的，建议5~100条左右即可
	// 请根据业务，实际压测一下试试，找到最佳的数量
	// 请注意worker默认已经开启了一条协程来处理本business的请求，所以该值只有在大于1的时候才会开启更多协程
	ProcessCmdGoroutineNum uint32 `protobuf:"varint,2,opt,name=processCmdGoroutineNum,proto3" json:"processCmdGoroutineNum,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registerReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_registerReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_registerReq_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetEvents() int32 {
	if x != nil {
		return x.Events
	}
	return 0
}

func (x *RegisterReq) GetProcessCmdGoroutineNum() uint32 {
	if x != nil {
		return x.ProcessCmdGoroutineNum
	}
	return 0
}

var File_registerReq_proto protoreflect.FileDescriptor

var file_registerReq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x74, 0x73, 0x76, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36,
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6d, 0x64, 0x47, 0x6f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6d, 0x64, 0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x42, 0x27, 0x5a, 0x07, 0x6e, 0x65, 0x74, 0x73, 0x76, 0x72,
	0x2f, 0xca, 0x02, 0x06, 0x4e, 0x65, 0x74, 0x73, 0x76, 0x72, 0xe2, 0x02, 0x12, 0x4e, 0x65, 0x74,
	0x73, 0x76, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registerReq_proto_rawDescOnce sync.Once
	file_registerReq_proto_rawDescData = file_registerReq_proto_rawDesc
)

func file_registerReq_proto_rawDescGZIP() []byte {
	file_registerReq_proto_rawDescOnce.Do(func() {
		file_registerReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_registerReq_proto_rawDescData)
	})
	return file_registerReq_proto_rawDescData
}

var file_registerReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_registerReq_proto_goTypes = []interface{}{
	(*RegisterReq)(nil), // 0: netsvr.registerReq.RegisterReq
}
var file_registerReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_registerReq_proto_init() }
func file_registerReq_proto_init() {
	if File_registerReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registerReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
			RawDescriptor: file_registerReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registerReq_proto_goTypes,
		DependencyIndexes: file_registerReq_proto_depIdxs,
		MessageInfos:      file_registerReq_proto_msgTypes,
	}.Build()
	File_registerReq_proto = out.File
	file_registerReq_proto_rawDesc = nil
	file_registerReq_proto_goTypes = nil
	file_registerReq_proto_depIdxs = nil
}
