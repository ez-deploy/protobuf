// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: project/project.proto

package project

import (
	model "github.com/ez-deploy/protobuf/model"
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

type CreateProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *model.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *CreateProjectReq) Reset() {
	*x = CreateProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectReq) ProtoMessage() {}

func (x *CreateProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectReq.ProtoReflect.Descriptor instead.
func (*CreateProjectReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectReq) GetProject() *model.Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type DeleteProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *DeleteProjectReq) Reset() {
	*x = DeleteProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectReq) ProtoMessage() {}

func (x *DeleteProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectReq.ProtoReflect.Descriptor instead.
func (*DeleteProjectReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteProjectReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

// create if not exist, else update service.
type SetServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string         `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Service     *model.Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *SetServiceReq) Reset() {
	*x = SetServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetServiceReq) ProtoMessage() {}

func (x *SetServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetServiceReq.ProtoReflect.Descriptor instead.
func (*SetServiceReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{2}
}

func (x *SetServiceReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *SetServiceReq) GetService() *model.Service {
	if x != nil {
		return x.Service
	}
	return nil
}

type GetServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *GetServiceReq) Reset() {
	*x = GetServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceReq) ProtoMessage() {}

func (x *GetServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceReq.ProtoReflect.Descriptor instead.
func (*GetServiceReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *GetServiceReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GetServiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   *model.Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Service *model.Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetServiceResp) Reset() {
	*x = GetServiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceResp) ProtoMessage() {}

func (x *GetServiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceResp.ProtoReflect.Descriptor instead.
func (*GetServiceResp) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{4}
}

func (x *GetServiceResp) GetError() *model.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetServiceResp) GetService() *model.Service {
	if x != nil {
		return x.Service
	}
	return nil
}

// ListService by project_id, list all service under project.
type ListServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *ListServiceReq) Reset() {
	*x = ListServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceReq) ProtoMessage() {}

func (x *ListServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceReq.ProtoReflect.Descriptor instead.
func (*ListServiceReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{5}
}

func (x *ListServiceReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type ListServiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   *model.Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Service []*model.Service `protobuf:"bytes,2,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *ListServiceResp) Reset() {
	*x = ListServiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResp) ProtoMessage() {}

func (x *ListServiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResp.ProtoReflect.Descriptor instead.
func (*ListServiceResp) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{6}
}

func (x *ListServiceResp) GetError() *model.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ListServiceResp) GetService() []*model.Service {
	if x != nil {
		return x.Service
	}
	return nil
}

// delete service by service_id.
type DeleteServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *DeleteServiceReq) Reset() {
	*x = DeleteServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceReq) ProtoMessage() {}

func (x *DeleteServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceReq.ProtoReflect.Descriptor instead.
func (*DeleteServiceReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteServiceReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *DeleteServiceReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type ListPodsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *ListPodsReq) Reset() {
	*x = ListPodsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPodsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPodsReq) ProtoMessage() {}

func (x *ListPodsReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPodsReq.ProtoReflect.Descriptor instead.
func (*ListPodsReq) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{8}
}

func (x *ListPodsReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ListPodsReq) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type ListPodsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *model.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Pods  []*model.Pod `protobuf:"bytes,2,rep,name=pods,proto3" json:"pods,omitempty"`
}

func (x *ListPodsResp) Reset() {
	*x = ListPodsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_project_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPodsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPodsResp) ProtoMessage() {}

func (x *ListPodsResp) ProtoReflect() protoreflect.Message {
	mi := &file_project_project_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPodsResp.ProtoReflect.Descriptor instead.
func (*ListPodsResp) Descriptor() ([]byte, []int) {
	return file_project_project_proto_rawDescGZIP(), []int{9}
}

func (x *ListPodsResp) GetError() *model.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ListPodsResp) GetPods() []*model.Pod {
	if x != nil {
		return x.Pods
	}
	return nil
}

var File_project_project_proto protoreflect.FileDescriptor

var file_project_project_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x35, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x33, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x58,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x70, 0x6f, 0x64,
	0x73, 0x32, 0xc3, 0x03, 0x0a, 0x03, 0x4f, 0x70, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x7a, 0x2d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_project_proto_rawDescOnce sync.Once
	file_project_project_proto_rawDescData = file_project_project_proto_rawDesc
)

func file_project_project_proto_rawDescGZIP() []byte {
	file_project_project_proto_rawDescOnce.Do(func() {
		file_project_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_project_proto_rawDescData)
	})
	return file_project_project_proto_rawDescData
}

var file_project_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_project_project_proto_goTypes = []interface{}{
	(*CreateProjectReq)(nil), // 0: project.CreateProjectReq
	(*DeleteProjectReq)(nil), // 1: project.DeleteProjectReq
	(*SetServiceReq)(nil),    // 2: project.SetServiceReq
	(*GetServiceReq)(nil),    // 3: project.GetServiceReq
	(*GetServiceResp)(nil),   // 4: project.GetServiceResp
	(*ListServiceReq)(nil),   // 5: project.ListServiceReq
	(*ListServiceResp)(nil),  // 6: project.ListServiceResp
	(*DeleteServiceReq)(nil), // 7: project.DeleteServiceReq
	(*ListPodsReq)(nil),      // 8: project.ListPodsReq
	(*ListPodsResp)(nil),     // 9: project.ListPodsResp
	(*model.Project)(nil),    // 10: model.Project
	(*model.Service)(nil),    // 11: model.Service
	(*model.Error)(nil),      // 12: model.Error
	(*model.Pod)(nil),        // 13: model.Pod
	(*model.CommonResp)(nil), // 14: model.CommonResp
}
var file_project_project_proto_depIdxs = []int32{
	10, // 0: project.CreateProjectReq.project:type_name -> model.Project
	11, // 1: project.SetServiceReq.service:type_name -> model.Service
	12, // 2: project.GetServiceResp.error:type_name -> model.Error
	11, // 3: project.GetServiceResp.service:type_name -> model.Service
	12, // 4: project.ListServiceResp.error:type_name -> model.Error
	11, // 5: project.ListServiceResp.service:type_name -> model.Service
	12, // 6: project.ListPodsResp.error:type_name -> model.Error
	13, // 7: project.ListPodsResp.pods:type_name -> model.Pod
	0,  // 8: project.Ops.CreateProject:input_type -> project.CreateProjectReq
	1,  // 9: project.Ops.DeleteProject:input_type -> project.DeleteProjectReq
	2,  // 10: project.Ops.SetService:input_type -> project.SetServiceReq
	3,  // 11: project.Ops.GetService:input_type -> project.GetServiceReq
	5,  // 12: project.Ops.ListService:input_type -> project.ListServiceReq
	7,  // 13: project.Ops.DeleteService:input_type -> project.DeleteServiceReq
	8,  // 14: project.Ops.ListPods:input_type -> project.ListPodsReq
	14, // 15: project.Ops.CreateProject:output_type -> model.CommonResp
	14, // 16: project.Ops.DeleteProject:output_type -> model.CommonResp
	14, // 17: project.Ops.SetService:output_type -> model.CommonResp
	4,  // 18: project.Ops.GetService:output_type -> project.GetServiceResp
	6,  // 19: project.Ops.ListService:output_type -> project.ListServiceResp
	14, // 20: project.Ops.DeleteService:output_type -> model.CommonResp
	9,  // 21: project.Ops.ListPods:output_type -> project.ListPodsResp
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_project_project_proto_init() }
func file_project_project_proto_init() {
	if File_project_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectReq); i {
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
		file_project_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectReq); i {
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
		file_project_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetServiceReq); i {
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
		file_project_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceReq); i {
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
		file_project_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceResp); i {
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
		file_project_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceReq); i {
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
		file_project_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceResp); i {
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
		file_project_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceReq); i {
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
		file_project_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPodsReq); i {
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
		file_project_project_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPodsResp); i {
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
			RawDescriptor: file_project_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_project_proto_goTypes,
		DependencyIndexes: file_project_project_proto_depIdxs,
		MessageInfos:      file_project_project_proto_msgTypes,
	}.Build()
	File_project_project_proto = out.File
	file_project_project_proto_rawDesc = nil
	file_project_project_proto_goTypes = nil
	file_project_project_proto_depIdxs = nil
}
