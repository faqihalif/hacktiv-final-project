// generate file
// protoc --go_out=. ./domain/patient/proto/*.proto --go-grpc_out=.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: domain/patient/proto/dto.proto

package patient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PatientProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth string `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	UserCreate  string `protobuf:"bytes,7,opt,name=user_create,json=userCreate,proto3" json:"user_create,omitempty"`
}

func (x *PatientProto) Reset() {
	*x = PatientProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_patient_proto_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientProto) ProtoMessage() {}

func (x *PatientProto) ProtoReflect() protoreflect.Message {
	mi := &file_domain_patient_proto_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientProto.ProtoReflect.Descriptor instead.
func (*PatientProto) Descriptor() ([]byte, []int) {
	return file_domain_patient_proto_dto_proto_rawDescGZIP(), []int{0}
}

func (x *PatientProto) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PatientProto) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PatientProto) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *PatientProto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PatientProto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientProto) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PatientProto) GetUserCreate() string {
	if x != nil {
		return x.UserCreate
	}
	return ""
}

type PatientListProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*PatientResponseProto `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PatientListProto) Reset() {
	*x = PatientListProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_patient_proto_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientListProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientListProto) ProtoMessage() {}

func (x *PatientListProto) ProtoReflect() protoreflect.Message {
	mi := &file_domain_patient_proto_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientListProto.ProtoReflect.Descriptor instead.
func (*PatientListProto) Descriptor() ([]byte, []int) {
	return file_domain_patient_proto_dto_proto_rawDescGZIP(), []int{1}
}

func (x *PatientListProto) GetList() []*PatientResponseProto {
	if x != nil {
		return x.List
	}
	return nil
}

type PatientFilterProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PatientFilterProto) Reset() {
	*x = PatientFilterProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_patient_proto_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientFilterProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientFilterProto) ProtoMessage() {}

func (x *PatientFilterProto) ProtoReflect() protoreflect.Message {
	mi := &file_domain_patient_proto_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientFilterProto.ProtoReflect.Descriptor instead.
func (*PatientFilterProto) Descriptor() ([]byte, []int) {
	return file_domain_patient_proto_dto_proto_rawDescGZIP(), []int{2}
}

func (x *PatientFilterProto) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *PatientFilterProto) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PatientFilterProto) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PatientFilterProto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PatientResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId   string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	UserCreate  string `protobuf:"bytes,8,opt,name=user_create,json=userCreate,proto3" json:"user_create,omitempty"`
}

func (x *PatientResponseProto) Reset() {
	*x = PatientResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_patient_proto_dto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientResponseProto) ProtoMessage() {}

func (x *PatientResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_domain_patient_proto_dto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientResponseProto.ProtoReflect.Descriptor instead.
func (*PatientResponseProto) Descriptor() ([]byte, []int) {
	return file_domain_patient_proto_dto_proto_rawDescGZIP(), []int{3}
}

func (x *PatientResponseProto) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *PatientResponseProto) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PatientResponseProto) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PatientResponseProto) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *PatientResponseProto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PatientResponseProto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientResponseProto) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PatientResponseProto) GetUserCreate() string {
	if x != nil {
		return x.UserCreate
	}
	return ""
}

var File_domain_patient_proto_dto_proto protoreflect.FileDescriptor

var file_domain_patient_proto_dto_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x10, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x89, 0x02, 0x0a, 0x14, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x32, 0xc9, 0x01, 0x0a, 0x13, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_patient_proto_dto_proto_rawDescOnce sync.Once
	file_domain_patient_proto_dto_proto_rawDescData = file_domain_patient_proto_dto_proto_rawDesc
)

func file_domain_patient_proto_dto_proto_rawDescGZIP() []byte {
	file_domain_patient_proto_dto_proto_rawDescOnce.Do(func() {
		file_domain_patient_proto_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_patient_proto_dto_proto_rawDescData)
	})
	return file_domain_patient_proto_dto_proto_rawDescData
}

var file_domain_patient_proto_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_domain_patient_proto_dto_proto_goTypes = []interface{}{
	(*PatientProto)(nil),         // 0: patient.PatientProto
	(*PatientListProto)(nil),     // 1: patient.PatientListProto
	(*PatientFilterProto)(nil),   // 2: patient.PatientFilterProto
	(*PatientResponseProto)(nil), // 3: patient.PatientResponseProto
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_domain_patient_proto_dto_proto_depIdxs = []int32{
	3, // 0: patient.PatientListProto.list:type_name -> patient.PatientResponseProto
	0, // 1: patient.PatientProtoService.Create:input_type -> patient.PatientProto
	4, // 2: patient.PatientProtoService.List:input_type -> google.protobuf.Empty
	2, // 3: patient.PatientProtoService.Find:input_type -> patient.PatientFilterProto
	4, // 4: patient.PatientProtoService.Create:output_type -> google.protobuf.Empty
	1, // 5: patient.PatientProtoService.List:output_type -> patient.PatientListProto
	1, // 6: patient.PatientProtoService.Find:output_type -> patient.PatientListProto
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_patient_proto_dto_proto_init() }
func file_domain_patient_proto_dto_proto_init() {
	if File_domain_patient_proto_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_patient_proto_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientProto); i {
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
		file_domain_patient_proto_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientListProto); i {
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
		file_domain_patient_proto_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientFilterProto); i {
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
		file_domain_patient_proto_dto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientResponseProto); i {
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
			RawDescriptor: file_domain_patient_proto_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_patient_proto_dto_proto_goTypes,
		DependencyIndexes: file_domain_patient_proto_dto_proto_depIdxs,
		MessageInfos:      file_domain_patient_proto_dto_proto_msgTypes,
	}.Build()
	File_domain_patient_proto_dto_proto = out.File
	file_domain_patient_proto_dto_proto_rawDesc = nil
	file_domain_patient_proto_dto_proto_goTypes = nil
	file_domain_patient_proto_dto_proto_depIdxs = nil
}
