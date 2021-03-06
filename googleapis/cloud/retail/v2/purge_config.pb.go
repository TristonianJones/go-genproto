// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/retail/v2/purge_config.proto

package retail

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Metadata related to the progress of the Purge operation.
// This will be returned by the google.longrunning.Operation.metadata field.
type PurgeMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the request / operation.
	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *PurgeMetadata) Reset() {
	*x = PurgeMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeMetadata) ProtoMessage() {}

func (x *PurgeMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeMetadata.ProtoReflect.Descriptor instead.
func (*PurgeMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP(), []int{0}
}

func (x *PurgeMetadata) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *PurgeMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// Request message for PurgeProducts method.
type PurgeProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the catalog under which the products are
	// created. The format is
	// "projects/${projectId}/locations/global/catalogs/${catalogId}"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Filter matching the products to be purged. Only supported value
	// at the moment is "*" (all items).
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The default value is false. Override this flag to true to
	// actually perform the purge. If the field is not set to true, a sampling of
	// products to be deleted will be returned.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *PurgeProductsRequest) Reset() {
	*x = PurgeProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeProductsRequest) ProtoMessage() {}

func (x *PurgeProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeProductsRequest.ProtoReflect.Descriptor instead.
func (*PurgeProductsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP(), []int{1}
}

func (x *PurgeProductsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PurgeProductsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PurgeProductsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Response of the PurgeProductsRequest. If the long running operation is
// successfully done, then this message is returned by the
// google.longrunning.Operations.response field.
type PurgeProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total count of products purged as a result of the operation.
	PurgedCount int64 `protobuf:"varint,1,opt,name=purged_count,json=purgedCount,proto3" json:"purged_count,omitempty"`
	// A random sampling of products deleted (or will be deleted) depending
	// on the `force` property in the request. Max of 500 items will be returned.
	// Currently, this is only populated if force=false.
	ProductsSample []*Product `protobuf:"bytes,2,rep,name=products_sample,json=productsSample,proto3" json:"products_sample,omitempty"`
}

func (x *PurgeProductsResponse) Reset() {
	*x = PurgeProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeProductsResponse) ProtoMessage() {}

func (x *PurgeProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeProductsResponse.ProtoReflect.Descriptor instead.
func (*PurgeProductsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP(), []int{2}
}

func (x *PurgeProductsResponse) GetPurgedCount() int64 {
	if x != nil {
		return x.PurgedCount
	}
	return 0
}

func (x *PurgeProductsResponse) GetProductsSample() []*Product {
	if x != nil {
		return x.ProductsSample
	}
	return nil
}

// Request message for PurgeUserEvents method.
type PurgeUserEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the event_store under which the events are
	// created. The format is
	//
	// "projects/${projectId}/locations/global/catalogs/${catalogId}/eventStores/${eventStoreId}"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The filter string to specify the events to be deleted. Empty
	// string filter is not allowed. The eligible fields
	// for filtering are:
	//
	// * `eventType`: UserEvent.eventType field of type string.
	// * `eventTime`: in ISO 8601 "zulu" format.
	// * `visitorId`: field of type string. Specifying this will delete all
	//   events associated with a visitor.
	// * `userId`: field of type string. Specifying this will delete all events
	//   associated with a user.
	//
	// Examples:
	//
	// * Deleting all events in a time range:
	//   `eventTime > "2012-04-23T18:25:43.511Z"
	//   eventTime < "2012-04-23T18:30:43.511Z"`
	// * Deleting specific eventType in time range:
	//   `eventTime > "2012-04-23T18:25:43.511Z" eventType = "detail-page-view"`
	// * Deleting all events for a specific visitor:
	//   `visitorId = "visitor1024"`
	//
	// The filtering fields are assumed to have an implicit AND.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The default value is false. Override this flag to true to
	// actually perform the purge. If the field is not set to true, a sampling of
	// events to be deleted will be returned.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *PurgeUserEventsRequest) Reset() {
	*x = PurgeUserEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeUserEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeUserEventsRequest) ProtoMessage() {}

func (x *PurgeUserEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeUserEventsRequest.ProtoReflect.Descriptor instead.
func (*PurgeUserEventsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP(), []int{3}
}

func (x *PurgeUserEventsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PurgeUserEventsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PurgeUserEventsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Response of the PurgeUserEventsRequest. If the long running operation is
// successfully done, then this message is returned by the
// google.longrunning.Operations.response field.
type PurgeUserEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total count of events purged as a result of the operation.
	PurgedEventsCount int64 `protobuf:"varint,1,opt,name=purged_events_count,json=purgedEventsCount,proto3" json:"purged_events_count,omitempty"`
	// A sampling of events deleted (or will be deleted) depending on the `force`
	// property in the request. Max of 500 items will be returned.
	UserEventsSample []*UserEvent `protobuf:"bytes,2,rep,name=user_events_sample,json=userEventsSample,proto3" json:"user_events_sample,omitempty"`
}

func (x *PurgeUserEventsResponse) Reset() {
	*x = PurgeUserEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeUserEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeUserEventsResponse) ProtoMessage() {}

func (x *PurgeUserEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_purge_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeUserEventsResponse.ProtoReflect.Descriptor instead.
func (*PurgeUserEventsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP(), []int{4}
}

func (x *PurgeUserEventsResponse) GetPurgedEventsCount() int64 {
	if x != nil {
		return x.PurgedEventsCount
	}
	return 0
}

func (x *PurgeUserEventsResponse) GetUserEventsSample() []*UserEvent {
	if x != nil {
		return x.UserEventsSample
	}
	return nil
}

var File_google_cloud_retail_v2_purge_config_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2_purge_config_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0d, 0x50, 0x75, 0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x66,
	0x0a, 0x14, 0x50, 0x75, 0x72, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x72, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x68, 0x0a,
	0x16, 0x50, 0x75, 0x72, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x17, 0x50, 0x75, 0x72, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x42, 0xc5, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x42, 0x10, 0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x3b, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02,
	0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32,
	0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2_purge_config_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2_purge_config_proto_rawDescData = file_google_cloud_retail_v2_purge_config_proto_rawDesc
)

func file_google_cloud_retail_v2_purge_config_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2_purge_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2_purge_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2_purge_config_proto_rawDescData)
	})
	return file_google_cloud_retail_v2_purge_config_proto_rawDescData
}

var file_google_cloud_retail_v2_purge_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_retail_v2_purge_config_proto_goTypes = []interface{}{
	(*PurgeMetadata)(nil),           // 0: google.cloud.retail.v2.PurgeMetadata
	(*PurgeProductsRequest)(nil),    // 1: google.cloud.retail.v2.PurgeProductsRequest
	(*PurgeProductsResponse)(nil),   // 2: google.cloud.retail.v2.PurgeProductsResponse
	(*PurgeUserEventsRequest)(nil),  // 3: google.cloud.retail.v2.PurgeUserEventsRequest
	(*PurgeUserEventsResponse)(nil), // 4: google.cloud.retail.v2.PurgeUserEventsResponse
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
	(*Product)(nil),                 // 6: google.cloud.retail.v2.Product
	(*UserEvent)(nil),               // 7: google.cloud.retail.v2.UserEvent
}
var file_google_cloud_retail_v2_purge_config_proto_depIdxs = []int32{
	5, // 0: google.cloud.retail.v2.PurgeMetadata.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: google.cloud.retail.v2.PurgeProductsResponse.products_sample:type_name -> google.cloud.retail.v2.Product
	7, // 2: google.cloud.retail.v2.PurgeUserEventsResponse.user_events_sample:type_name -> google.cloud.retail.v2.UserEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2_purge_config_proto_init() }
func file_google_cloud_retail_v2_purge_config_proto_init() {
	if File_google_cloud_retail_v2_purge_config_proto != nil {
		return
	}
	file_google_cloud_retail_v2_product_proto_init()
	file_google_cloud_retail_v2_user_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2_purge_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeMetadata); i {
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
		file_google_cloud_retail_v2_purge_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeProductsRequest); i {
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
		file_google_cloud_retail_v2_purge_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeProductsResponse); i {
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
		file_google_cloud_retail_v2_purge_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeUserEventsRequest); i {
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
		file_google_cloud_retail_v2_purge_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeUserEventsResponse); i {
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
			RawDescriptor: file_google_cloud_retail_v2_purge_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2_purge_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2_purge_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2_purge_config_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2_purge_config_proto = out.File
	file_google_cloud_retail_v2_purge_config_proto_rawDesc = nil
	file_google_cloud_retail_v2_purge_config_proto_goTypes = nil
	file_google_cloud_retail_v2_purge_config_proto_depIdxs = nil
}
