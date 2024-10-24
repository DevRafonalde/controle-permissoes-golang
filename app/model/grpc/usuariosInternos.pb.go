// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: usuariosInternos.proto

package grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListaUsuariosInternos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsuariosInternos []*UsuarioInterno `protobuf:"bytes,1,rep,name=usuariosInternos,proto3" json:"usuariosInternos,omitempty"`
	Meta             *Meta             `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListaUsuariosInternos) Reset() {
	*x = ListaUsuariosInternos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usuariosInternos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaUsuariosInternos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaUsuariosInternos) ProtoMessage() {}

func (x *ListaUsuariosInternos) ProtoReflect() protoreflect.Message {
	mi := &file_usuariosInternos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaUsuariosInternos.ProtoReflect.Descriptor instead.
func (*ListaUsuariosInternos) Descriptor() ([]byte, []int) {
	return file_usuariosInternos_proto_rawDescGZIP(), []int{0}
}

func (x *ListaUsuariosInternos) GetUsuariosInternos() []*UsuarioInterno {
	if x != nil {
		return x.UsuariosInternos
	}
	return nil
}

func (x *ListaUsuariosInternos) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ResponsePerfisVinculados struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perfis []*Perfil `protobuf:"bytes,1,rep,name=perfis,proto3" json:"perfis,omitempty"`
}

func (x *ResponsePerfisVinculados) Reset() {
	*x = ResponsePerfisVinculados{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usuariosInternos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePerfisVinculados) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePerfisVinculados) ProtoMessage() {}

func (x *ResponsePerfisVinculados) ProtoReflect() protoreflect.Message {
	mi := &file_usuariosInternos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePerfisVinculados.ProtoReflect.Descriptor instead.
func (*ResponsePerfisVinculados) Descriptor() ([]byte, []int) {
	return file_usuariosInternos_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePerfisVinculados) GetPerfis() []*Perfil {
	if x != nil {
		return x.Perfis
	}
	return nil
}

var File_usuariosInternos_proto protoreflect.FileDescriptor

var file_usuariosInternos_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x79, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x55, 0x73, 0x75, 0x61, 0x72,
	0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x12, 0x40, 0x0a, 0x10, 0x75,
	0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x52, 0x10, 0x75, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x40, 0x0a,
	0x18, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x56,
	0x69, 0x6e, 0x63, 0x75, 0x6c, 0x61, 0x64, 0x6f, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x66, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x52, 0x06, 0x70, 0x65, 0x72, 0x66, 0x69, 0x73, 0x32,
	0xd6, 0x0c, 0x0a, 0x10, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x6f, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73,
	0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x64, 0x6f, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x12,
	0x38, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73,
	0x2f, 0x7b, 0x74, 0x61, 0x6d, 0x61, 0x6e, 0x68, 0x6f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x7d,
	0x2f, 0x7b, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x7d, 0x12, 0x6c, 0x0a, 0x16, 0x46, 0x69, 0x6e,
	0x64, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x75, 0x61,
	0x72, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x12, 0x24, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x66, 0x69, 0x73, 0x56, 0x69, 0x6e, 0x63, 0x75, 0x6c, 0x61, 0x64, 0x6f, 0x73, 0x12,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x56, 0x69, 0x6e, 0x63, 0x75, 0x6c, 0x61, 0x64, 0x6f, 0x73,
	0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x69, 0x73,
	0x2d, 0x76, 0x69, 0x6e, 0x63, 0x75, 0x6c, 0x61, 0x64, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x6c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69,
	0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x1a, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x66,
	0x69, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f,
	0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x12, 0x69,
	0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x75, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x22, 0x2c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x26, 0x22, 0x24, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x6f, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f,
	0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x75, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x65, 0x72, 0x66, 0x69, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x1a, 0x1f, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x11, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x61, 0x72, 0x53, 0x65, 0x6e, 0x68, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x61, 0x72, 0x53, 0x65, 0x6e, 0x68, 0x61, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x1a, 0x33, 0x2f, 0x73,
	0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61,
	0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x61, 0x6c,
	0x74, 0x65, 0x72, 0x61, 0x72, 0x2d, 0x73, 0x65, 0x6e, 0x68, 0x61, 0x2d, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x94, 0x01, 0x0a, 0x1a, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x72, 0x53, 0x65, 0x6e,
	0x68, 0x61, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f,
	0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x61, 0x72, 0x53, 0x65, 0x6e, 0x68, 0x61, 0x55, 0x73, 0x75, 0x61, 0x72,
	0x69, 0x6f, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x3a, 0x01,
	0x2a, 0x1a, 0x35, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x6f, 0x73, 0x2f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x72, 0x2d, 0x73, 0x65, 0x6e, 0x68, 0x61,
	0x2d, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x76, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x72, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30,
	0x1a, 0x2e, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x76, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x61, 0x74, 0x69, 0x76, 0x61, 0x72, 0x55, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x1a, 0x2e, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73,
	0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x73, 0x61, 0x74, 0x69,
	0x76, 0x61, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x75, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74,
	0x6f, 0x72, 0x6e, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f,
	0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x73, 0x69,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72,
	0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x6e, 0x68, 0x61, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x6e, 0x68, 0x61, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a,
	0x01, 0x2a, 0x22, 0x31, 0x2f, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d,
	0x73, 0x65, 0x6e, 0x68, 0x61, 0x12, 0x71, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x6e, 0x68, 0x61, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x68, 0x61, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x1a, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x1a, 0x2b, 0x2f, 0x73, 0x69,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72,
	0x69, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2d, 0x73, 0x65, 0x6e, 0x68, 0x61, 0x42, 0x13, 0x5a, 0x11, 0x73, 0x69, 0x2d, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usuariosInternos_proto_rawDescOnce sync.Once
	file_usuariosInternos_proto_rawDescData = file_usuariosInternos_proto_rawDesc
)

func file_usuariosInternos_proto_rawDescGZIP() []byte {
	file_usuariosInternos_proto_rawDescOnce.Do(func() {
		file_usuariosInternos_proto_rawDescData = protoimpl.X.CompressGZIP(file_usuariosInternos_proto_rawDescData)
	})
	return file_usuariosInternos_proto_rawDescData
}

var file_usuariosInternos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_usuariosInternos_proto_goTypes = []any{
	(*ListaUsuariosInternos)(nil),      // 0: grpc.ListaUsuariosInternos
	(*ResponsePerfisVinculados)(nil),   // 1: grpc.ResponsePerfisVinculados
	(*UsuarioInterno)(nil),             // 2: grpc.UsuarioInterno
	(*Meta)(nil),                       // 3: grpc.Meta
	(*Perfil)(nil),                     // 4: grpc.Perfil
	(*RequestAllPaginado)(nil),         // 5: grpc.RequestAllPaginado
	(*RequestId)(nil),                  // 6: grpc.RequestId
	(*UsuarioPerfis)(nil),              // 7: grpc.UsuarioPerfis
	(*RequestAlterarSenhaAdmin)(nil),   // 8: grpc.RequestAlterarSenhaAdmin
	(*RequestAlterarSenhaUsuario)(nil), // 9: grpc.RequestAlterarSenhaUsuario
	(*LoginUsuario)(nil),               // 10: grpc.LoginUsuario
	(*EmailReset)(nil),                 // 11: grpc.EmailReset
	(*ResetSenhaUsuario)(nil),          // 12: grpc.ResetSenhaUsuario
	(*ResponseBool)(nil),               // 13: grpc.ResponseBool
	(*RetornoLoginUsuario)(nil),        // 14: grpc.RetornoLoginUsuario
	(*ResponseTokenResetSenha)(nil),    // 15: grpc.ResponseTokenResetSenha
}
var file_usuariosInternos_proto_depIdxs = []int32{
	2,  // 0: grpc.ListaUsuariosInternos.usuariosInternos:type_name -> grpc.UsuarioInterno
	3,  // 1: grpc.ListaUsuariosInternos.meta:type_name -> grpc.Meta
	4,  // 2: grpc.ResponsePerfisVinculados.perfis:type_name -> grpc.Perfil
	5,  // 3: grpc.UsuariosInternos.FindAllUsuariosInternos:input_type -> grpc.RequestAllPaginado
	6,  // 4: grpc.UsuariosInternos.FindUsuarioInternoById:input_type -> grpc.RequestId
	6,  // 5: grpc.UsuariosInternos.GetPerfisVinculados:input_type -> grpc.RequestId
	7,  // 6: grpc.UsuariosInternos.CreateUsuarioInterno:input_type -> grpc.UsuarioPerfis
	6,  // 7: grpc.UsuariosInternos.CloneUsuarioInterno:input_type -> grpc.RequestId
	7,  // 8: grpc.UsuariosInternos.UpdateUsuarioInterno:input_type -> grpc.UsuarioPerfis
	8,  // 9: grpc.UsuariosInternos.AlterarSenhaAdmin:input_type -> grpc.RequestAlterarSenhaAdmin
	9,  // 10: grpc.UsuariosInternos.AlterarSenhaUsuarioInterno:input_type -> grpc.RequestAlterarSenhaUsuario
	6,  // 11: grpc.UsuariosInternos.RestaurarUsuarioInterno:input_type -> grpc.RequestId
	6,  // 12: grpc.UsuariosInternos.DesativarUsuarioInterno:input_type -> grpc.RequestId
	10, // 13: grpc.UsuariosInternos.Login:input_type -> grpc.LoginUsuario
	11, // 14: grpc.UsuariosInternos.TokenResetSenha:input_type -> grpc.EmailReset
	12, // 15: grpc.UsuariosInternos.ResetSenha:input_type -> grpc.ResetSenhaUsuario
	0,  // 16: grpc.UsuariosInternos.FindAllUsuariosInternos:output_type -> grpc.ListaUsuariosInternos
	7,  // 17: grpc.UsuariosInternos.FindUsuarioInternoById:output_type -> grpc.UsuarioPerfis
	1,  // 18: grpc.UsuariosInternos.GetPerfisVinculados:output_type -> grpc.ResponsePerfisVinculados
	7,  // 19: grpc.UsuariosInternos.CreateUsuarioInterno:output_type -> grpc.UsuarioPerfis
	7,  // 20: grpc.UsuariosInternos.CloneUsuarioInterno:output_type -> grpc.UsuarioPerfis
	7,  // 21: grpc.UsuariosInternos.UpdateUsuarioInterno:output_type -> grpc.UsuarioPerfis
	13, // 22: grpc.UsuariosInternos.AlterarSenhaAdmin:output_type -> grpc.ResponseBool
	13, // 23: grpc.UsuariosInternos.AlterarSenhaUsuarioInterno:output_type -> grpc.ResponseBool
	13, // 24: grpc.UsuariosInternos.RestaurarUsuarioInterno:output_type -> grpc.ResponseBool
	13, // 25: grpc.UsuariosInternos.DesativarUsuarioInterno:output_type -> grpc.ResponseBool
	14, // 26: grpc.UsuariosInternos.Login:output_type -> grpc.RetornoLoginUsuario
	15, // 27: grpc.UsuariosInternos.TokenResetSenha:output_type -> grpc.ResponseTokenResetSenha
	13, // 28: grpc.UsuariosInternos.ResetSenha:output_type -> grpc.ResponseBool
	16, // [16:29] is the sub-list for method output_type
	3,  // [3:16] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_usuariosInternos_proto_init() }
func file_usuariosInternos_proto_init() {
	if File_usuariosInternos_proto != nil {
		return
	}
	file_modelos_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_usuariosInternos_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListaUsuariosInternos); i {
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
		file_usuariosInternos_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResponsePerfisVinculados); i {
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
			RawDescriptor: file_usuariosInternos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usuariosInternos_proto_goTypes,
		DependencyIndexes: file_usuariosInternos_proto_depIdxs,
		MessageInfos:      file_usuariosInternos_proto_msgTypes,
	}.Build()
	File_usuariosInternos_proto = out.File
	file_usuariosInternos_proto_rawDesc = nil
	file_usuariosInternos_proto_goTypes = nil
	file_usuariosInternos_proto_depIdxs = nil
}