// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: app_identity_service.proto

package app_identity

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

type AppIdentityServiceError_ErrorCode int32

const (
	AppIdentityServiceError_SUCCESS           AppIdentityServiceError_ErrorCode = 0
	AppIdentityServiceError_UNKNOWN_SCOPE     AppIdentityServiceError_ErrorCode = 9
	AppIdentityServiceError_BLOB_TOO_LARGE    AppIdentityServiceError_ErrorCode = 1000
	AppIdentityServiceError_DEADLINE_EXCEEDED AppIdentityServiceError_ErrorCode = 1001
	AppIdentityServiceError_NOT_A_VALID_APP   AppIdentityServiceError_ErrorCode = 1002
	AppIdentityServiceError_UNKNOWN_ERROR     AppIdentityServiceError_ErrorCode = 1003
	AppIdentityServiceError_NOT_ALLOWED       AppIdentityServiceError_ErrorCode = 1005
	AppIdentityServiceError_NOT_IMPLEMENTED   AppIdentityServiceError_ErrorCode = 1006
)

// Enum value maps for AppIdentityServiceError_ErrorCode.
var (
	AppIdentityServiceError_ErrorCode_name = map[int32]string{
		0:    "SUCCESS",
		9:    "UNKNOWN_SCOPE",
		1000: "BLOB_TOO_LARGE",
		1001: "DEADLINE_EXCEEDED",
		1002: "NOT_A_VALID_APP",
		1003: "UNKNOWN_ERROR",
		1005: "NOT_ALLOWED",
		1006: "NOT_IMPLEMENTED",
	}
	AppIdentityServiceError_ErrorCode_value = map[string]int32{
		"SUCCESS":           0,
		"UNKNOWN_SCOPE":     9,
		"BLOB_TOO_LARGE":    1000,
		"DEADLINE_EXCEEDED": 1001,
		"NOT_A_VALID_APP":   1002,
		"UNKNOWN_ERROR":     1003,
		"NOT_ALLOWED":       1005,
		"NOT_IMPLEMENTED":   1006,
	}
)

func (x AppIdentityServiceError_ErrorCode) Enum() *AppIdentityServiceError_ErrorCode {
	p := new(AppIdentityServiceError_ErrorCode)
	*p = x
	return p
}

func (x AppIdentityServiceError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppIdentityServiceError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_app_identity_service_proto_enumTypes[0].Descriptor()
}

func (AppIdentityServiceError_ErrorCode) Type() protoreflect.EnumType {
	return &file_app_identity_service_proto_enumTypes[0]
}

func (x AppIdentityServiceError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppIdentityServiceError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppIdentityServiceError_ErrorCode(num)
	return nil
}

// Deprecated: Use AppIdentityServiceError_ErrorCode.Descriptor instead.
func (AppIdentityServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{0, 0}
}

type AppIdentityServiceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppIdentityServiceError) Reset() {
	*x = AppIdentityServiceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppIdentityServiceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppIdentityServiceError) ProtoMessage() {}

func (x *AppIdentityServiceError) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppIdentityServiceError.ProtoReflect.Descriptor instead.
func (*AppIdentityServiceError) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{0}
}

type SignForAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BytesToSign []byte `protobuf:"bytes,1,opt,name=bytes_to_sign,json=bytesToSign" json:"bytes_to_sign,omitempty"`
}

func (x *SignForAppRequest) Reset() {
	*x = SignForAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignForAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignForAppRequest) ProtoMessage() {}

func (x *SignForAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignForAppRequest.ProtoReflect.Descriptor instead.
func (*SignForAppRequest) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{1}
}

func (x *SignForAppRequest) GetBytesToSign() []byte {
	if x != nil {
		return x.BytesToSign
	}
	return nil
}

type SignForAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName        *string `protobuf:"bytes,1,opt,name=key_name,json=keyName" json:"key_name,omitempty"`
	SignatureBytes []byte  `protobuf:"bytes,2,opt,name=signature_bytes,json=signatureBytes" json:"signature_bytes,omitempty"`
}

func (x *SignForAppResponse) Reset() {
	*x = SignForAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignForAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignForAppResponse) ProtoMessage() {}

func (x *SignForAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignForAppResponse.ProtoReflect.Descriptor instead.
func (*SignForAppResponse) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignForAppResponse) GetKeyName() string {
	if x != nil && x.KeyName != nil {
		return *x.KeyName
	}
	return ""
}

func (x *SignForAppResponse) GetSignatureBytes() []byte {
	if x != nil {
		return x.SignatureBytes
	}
	return nil
}

type GetPublicCertificateForAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPublicCertificateForAppRequest) Reset() {
	*x = GetPublicCertificateForAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicCertificateForAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicCertificateForAppRequest) ProtoMessage() {}

func (x *GetPublicCertificateForAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicCertificateForAppRequest.ProtoReflect.Descriptor instead.
func (*GetPublicCertificateForAppRequest) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{3}
}

type PublicCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName            *string `protobuf:"bytes,1,opt,name=key_name,json=keyName" json:"key_name,omitempty"`
	X509CertificatePem *string `protobuf:"bytes,2,opt,name=x509_certificate_pem,json=x509CertificatePem" json:"x509_certificate_pem,omitempty"`
}

func (x *PublicCertificate) Reset() {
	*x = PublicCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCertificate) ProtoMessage() {}

func (x *PublicCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCertificate.ProtoReflect.Descriptor instead.
func (*PublicCertificate) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{4}
}

func (x *PublicCertificate) GetKeyName() string {
	if x != nil && x.KeyName != nil {
		return *x.KeyName
	}
	return ""
}

func (x *PublicCertificate) GetX509CertificatePem() string {
	if x != nil && x.X509CertificatePem != nil {
		return *x.X509CertificatePem
	}
	return ""
}

type GetPublicCertificateForAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicCertificateList      []*PublicCertificate `protobuf:"bytes,1,rep,name=public_certificate_list,json=publicCertificateList" json:"public_certificate_list,omitempty"`
	MaxClientCacheTimeInSecond *int64               `protobuf:"varint,2,opt,name=max_client_cache_time_in_second,json=maxClientCacheTimeInSecond" json:"max_client_cache_time_in_second,omitempty"`
}

func (x *GetPublicCertificateForAppResponse) Reset() {
	*x = GetPublicCertificateForAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicCertificateForAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicCertificateForAppResponse) ProtoMessage() {}

func (x *GetPublicCertificateForAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicCertificateForAppResponse.ProtoReflect.Descriptor instead.
func (*GetPublicCertificateForAppResponse) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPublicCertificateForAppResponse) GetPublicCertificateList() []*PublicCertificate {
	if x != nil {
		return x.PublicCertificateList
	}
	return nil
}

func (x *GetPublicCertificateForAppResponse) GetMaxClientCacheTimeInSecond() int64 {
	if x != nil && x.MaxClientCacheTimeInSecond != nil {
		return *x.MaxClientCacheTimeInSecond
	}
	return 0
}

type GetServiceAccountNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServiceAccountNameRequest) Reset() {
	*x = GetServiceAccountNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceAccountNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAccountNameRequest) ProtoMessage() {}

func (x *GetServiceAccountNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAccountNameRequest.ProtoReflect.Descriptor instead.
func (*GetServiceAccountNameRequest) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{6}
}

type GetServiceAccountNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAccountName *string `protobuf:"bytes,1,opt,name=service_account_name,json=serviceAccountName" json:"service_account_name,omitempty"`
}

func (x *GetServiceAccountNameResponse) Reset() {
	*x = GetServiceAccountNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceAccountNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAccountNameResponse) ProtoMessage() {}

func (x *GetServiceAccountNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAccountNameResponse.ProtoReflect.Descriptor instead.
func (*GetServiceAccountNameResponse) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetServiceAccountNameResponse) GetServiceAccountName() string {
	if x != nil && x.ServiceAccountName != nil {
		return *x.ServiceAccountName
	}
	return ""
}

type GetAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope              []string `protobuf:"bytes,1,rep,name=scope" json:"scope,omitempty"`
	ServiceAccountId   *int64   `protobuf:"varint,2,opt,name=service_account_id,json=serviceAccountId" json:"service_account_id,omitempty"`
	ServiceAccountName *string  `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName" json:"service_account_name,omitempty"`
}

func (x *GetAccessTokenRequest) Reset() {
	*x = GetAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenRequest) ProtoMessage() {}

func (x *GetAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetAccessTokenRequest) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *GetAccessTokenRequest) GetServiceAccountId() int64 {
	if x != nil && x.ServiceAccountId != nil {
		return *x.ServiceAccountId
	}
	return 0
}

func (x *GetAccessTokenRequest) GetServiceAccountName() string {
	if x != nil && x.ServiceAccountName != nil {
		return *x.ServiceAccountName
	}
	return ""
}

type GetAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken    *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	ExpirationTime *int64  `protobuf:"varint,2,opt,name=expiration_time,json=expirationTime" json:"expiration_time,omitempty"`
}

func (x *GetAccessTokenResponse) Reset() {
	*x = GetAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenResponse) ProtoMessage() {}

func (x *GetAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetAccessTokenResponse) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

func (x *GetAccessTokenResponse) GetExpirationTime() int64 {
	if x != nil && x.ExpirationTime != nil {
		return *x.ExpirationTime
	}
	return 0
}

type GetDefaultGcsBucketNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDefaultGcsBucketNameRequest) Reset() {
	*x = GetDefaultGcsBucketNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultGcsBucketNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultGcsBucketNameRequest) ProtoMessage() {}

func (x *GetDefaultGcsBucketNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultGcsBucketNameRequest.ProtoReflect.Descriptor instead.
func (*GetDefaultGcsBucketNameRequest) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{10}
}

type GetDefaultGcsBucketNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultGcsBucketName *string `protobuf:"bytes,1,opt,name=default_gcs_bucket_name,json=defaultGcsBucketName" json:"default_gcs_bucket_name,omitempty"`
}

func (x *GetDefaultGcsBucketNameResponse) Reset() {
	*x = GetDefaultGcsBucketNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_identity_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultGcsBucketNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultGcsBucketNameResponse) ProtoMessage() {}

func (x *GetDefaultGcsBucketNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_identity_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultGcsBucketNameResponse.ProtoReflect.Descriptor instead.
func (*GetDefaultGcsBucketNameResponse) Descriptor() ([]byte, []int) {
	return file_app_identity_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetDefaultGcsBucketNameResponse) GetDefaultGcsBucketName() string {
	if x != nil && x.DefaultGcsBucketName != nil {
		return *x.DefaultGcsBucketName
	}
	return ""
}

var File_app_identity_service_proto protoreflect.FileDescriptor

var file_app_identity_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x22, 0xc6, 0x01, 0x0a, 0x17, 0x41,
	0x70, 0x70, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xaa, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x43, 0x4f,
	0x50, 0x45, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0e, 0x42, 0x4c, 0x4f, 0x42, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0xe8, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x44, 0x45, 0x41,
	0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0xe9,
	0x07, 0x12, 0x14, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0xea, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xeb, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0xed, 0x07, 0x12, 0x14, 0x0a,
	0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44,
	0x10, 0xee, 0x07, 0x22, 0x37, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x46, 0x6f, 0x72, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x22, 0x58, 0x0a, 0x12,
	0x53, 0x69, 0x67, 0x6e, 0x46, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x11, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x78,
	0x35, 0x30, 0x39, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x78, 0x35, 0x30, 0x39, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6d, 0x22, 0xc2, 0x01,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x17, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x15, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a,
	0x1f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x6d, 0x61, 0x78, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x51, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x1e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x63, 0x73, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x67, 0x63, 0x73, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3b,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
}

var (
	file_app_identity_service_proto_rawDescOnce sync.Once
	file_app_identity_service_proto_rawDescData = file_app_identity_service_proto_rawDesc
)

func file_app_identity_service_proto_rawDescGZIP() []byte {
	file_app_identity_service_proto_rawDescOnce.Do(func() {
		file_app_identity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_identity_service_proto_rawDescData)
	})
	return file_app_identity_service_proto_rawDescData
}

var file_app_identity_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_identity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_app_identity_service_proto_goTypes = []interface{}{
	(AppIdentityServiceError_ErrorCode)(0),     // 0: appengine.v2.AppIdentityServiceError.ErrorCode
	(*AppIdentityServiceError)(nil),            // 1: appengine.v2.AppIdentityServiceError
	(*SignForAppRequest)(nil),                  // 2: appengine.v2.SignForAppRequest
	(*SignForAppResponse)(nil),                 // 3: appengine.v2.SignForAppResponse
	(*GetPublicCertificateForAppRequest)(nil),  // 4: appengine.v2.GetPublicCertificateForAppRequest
	(*PublicCertificate)(nil),                  // 5: appengine.v2.PublicCertificate
	(*GetPublicCertificateForAppResponse)(nil), // 6: appengine.v2.GetPublicCertificateForAppResponse
	(*GetServiceAccountNameRequest)(nil),       // 7: appengine.v2.GetServiceAccountNameRequest
	(*GetServiceAccountNameResponse)(nil),      // 8: appengine.v2.GetServiceAccountNameResponse
	(*GetAccessTokenRequest)(nil),              // 9: appengine.v2.GetAccessTokenRequest
	(*GetAccessTokenResponse)(nil),             // 10: appengine.v2.GetAccessTokenResponse
	(*GetDefaultGcsBucketNameRequest)(nil),     // 11: appengine.v2.GetDefaultGcsBucketNameRequest
	(*GetDefaultGcsBucketNameResponse)(nil),    // 12: appengine.v2.GetDefaultGcsBucketNameResponse
}
var file_app_identity_service_proto_depIdxs = []int32{
	5, // 0: appengine.v2.GetPublicCertificateForAppResponse.public_certificate_list:type_name -> appengine.v2.PublicCertificate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_identity_service_proto_init() }
func file_app_identity_service_proto_init() {
	if File_app_identity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_identity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppIdentityServiceError); i {
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
		file_app_identity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignForAppRequest); i {
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
		file_app_identity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignForAppResponse); i {
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
		file_app_identity_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicCertificateForAppRequest); i {
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
		file_app_identity_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicCertificate); i {
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
		file_app_identity_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicCertificateForAppResponse); i {
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
		file_app_identity_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceAccountNameRequest); i {
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
		file_app_identity_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceAccountNameResponse); i {
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
		file_app_identity_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenRequest); i {
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
		file_app_identity_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenResponse); i {
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
		file_app_identity_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultGcsBucketNameRequest); i {
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
		file_app_identity_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultGcsBucketNameResponse); i {
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
			RawDescriptor: file_app_identity_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_identity_service_proto_goTypes,
		DependencyIndexes: file_app_identity_service_proto_depIdxs,
		EnumInfos:         file_app_identity_service_proto_enumTypes,
		MessageInfos:      file_app_identity_service_proto_msgTypes,
	}.Build()
	File_app_identity_service_proto = out.File
	file_app_identity_service_proto_rawDesc = nil
	file_app_identity_service_proto_goTypes = nil
	file_app_identity_service_proto_depIdxs = nil
}