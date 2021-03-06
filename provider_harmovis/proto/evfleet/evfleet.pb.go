// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: evfleet.proto

package proto

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EvFleetSupply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId        uint32  `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	TimeStamp      string  `protobuf:"bytes,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	VehicleId      uint64  `protobuf:"varint,3,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	Latitude       float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude      float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Soc            float32 `protobuf:"fixed32,6,opt,name=soc,proto3" json:"soc,omitempty"`
	Soh            float32 `protobuf:"fixed32,7,opt,name=soh,proto3" json:"soh,omitempty"`
	AirConditioner uint32  `protobuf:"varint,8,opt,name=air_conditioner,json=airConditioner,proto3" json:"air_conditioner,omitempty"`
}

func (x *EvFleetSupply) Reset() {
	*x = EvFleetSupply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evfleet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvFleetSupply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvFleetSupply) ProtoMessage() {}

func (x *EvFleetSupply) ProtoReflect() protoreflect.Message {
	mi := &file_evfleet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvFleetSupply.ProtoReflect.Descriptor instead.
func (*EvFleetSupply) Descriptor() ([]byte, []int) {
	return file_evfleet_proto_rawDescGZIP(), []int{0}
}

func (x *EvFleetSupply) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *EvFleetSupply) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *EvFleetSupply) GetVehicleId() uint64 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

func (x *EvFleetSupply) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EvFleetSupply) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *EvFleetSupply) GetSoc() float32 {
	if x != nil {
		return x.Soc
	}
	return 0
}

func (x *EvFleetSupply) GetSoh() float32 {
	if x != nil {
		return x.Soh
	}
	return 0
}

func (x *EvFleetSupply) GetAirConditioner() uint32 {
	if x != nil {
		return x.AirConditioner
	}
	return 0
}

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId      uint64  `protobuf:"varint,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	Soc            float32 `protobuf:"fixed32,2,opt,name=soc,proto3" json:"soc,omitempty"`
	Soh            float32 `protobuf:"fixed32,3,opt,name=soh,proto3" json:"soh,omitempty"`
	AirConditioner uint32  `protobuf:"varint,4,opt,name=air_conditioner,json=airConditioner,proto3" json:"air_conditioner,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evfleet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_evfleet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_evfleet_proto_rawDescGZIP(), []int{1}
}

func (x *Vehicle) GetVehicleId() uint64 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

func (x *Vehicle) GetSoc() float32 {
	if x != nil {
		return x.Soc
	}
	return 0
}

func (x *Vehicle) GetSoh() float32 {
	if x != nil {
		return x.Soh
	}
	return 0
}

func (x *Vehicle) GetAirConditioner() uint32 {
	if x != nil {
		return x.AirConditioner
	}
	return 0
}

type VehicleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     uint32     `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	TimeStamp   string     `protobuf:"bytes,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	ModuleId    uint64     `protobuf:"varint,3,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	ProvideId   string     `protobuf:"bytes,4,opt,name=provide_id,json=provideId,proto3" json:"provide_id,omitempty"`
	VehicleList []*Vehicle `protobuf:"bytes,5,rep,name=vehicle_list,json=vehicleList,proto3" json:"vehicle_list,omitempty"`
}

func (x *VehicleList) Reset() {
	*x = VehicleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evfleet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleList) ProtoMessage() {}

func (x *VehicleList) ProtoReflect() protoreflect.Message {
	mi := &file_evfleet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleList.ProtoReflect.Descriptor instead.
func (*VehicleList) Descriptor() ([]byte, []int) {
	return file_evfleet_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleList) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *VehicleList) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *VehicleList) GetModuleId() uint64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *VehicleList) GetProvideId() string {
	if x != nil {
		return x.ProvideId
	}
	return ""
}

func (x *VehicleList) GetVehicleList() []*Vehicle {
	if x != nil {
		return x.VehicleList
	}
	return nil
}

type EvFleetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *EvFleetResponse) Reset() {
	*x = EvFleetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evfleet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvFleetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvFleetResponse) ProtoMessage() {}

func (x *EvFleetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evfleet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvFleetResponse.ProtoReflect.Descriptor instead.
func (*EvFleetResponse) Descriptor() ([]byte, []int) {
	return file_evfleet_proto_rawDescGZIP(), []int{3}
}

func (x *EvFleetResponse) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_evfleet_proto protoreflect.FileDescriptor

var file_evfleet_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x76, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x46, 0x6c, 0x65, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x63,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x6f, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x6f, 0x68, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0x75, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73,
	0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x73, 0x6f, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x61,
	0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x22, 0xb6, 0x01,
	0x0a, 0x0b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x45, 0x76, 0x46, 0x6c, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xc3, 0x01, 0x0a, 0x07, 0x45, 0x76, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x12, 0x5b, 0x0a,
	0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x46, 0x6c, 0x65, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x76, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x65, 0x76, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x46, 0x6c, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x65, 0x76, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x65, 0x76, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evfleet_proto_rawDescOnce sync.Once
	file_evfleet_proto_rawDescData = file_evfleet_proto_rawDesc
)

func file_evfleet_proto_rawDescGZIP() []byte {
	file_evfleet_proto_rawDescOnce.Do(func() {
		file_evfleet_proto_rawDescData = protoimpl.X.CompressGZIP(file_evfleet_proto_rawDescData)
	})
	return file_evfleet_proto_rawDescData
}

var file_evfleet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_evfleet_proto_goTypes = []interface{}{
	(*EvFleetSupply)(nil),   // 0: proto.EvFleetSupply
	(*Vehicle)(nil),         // 1: proto.Vehicle
	(*VehicleList)(nil),     // 2: proto.VehicleList
	(*EvFleetResponse)(nil), // 3: proto.EvFleetResponse
}
var file_evfleet_proto_depIdxs = []int32{
	1, // 0: proto.VehicleList.vehicle_list:type_name -> proto.Vehicle
	0, // 1: proto.EvFleet.NotifyFleetInfo:input_type -> proto.EvFleetSupply
	2, // 2: proto.EvFleet.ProvideFleetInfo:input_type -> proto.VehicleList
	3, // 3: proto.EvFleet.NotifyFleetInfo:output_type -> proto.EvFleetResponse
	3, // 4: proto.EvFleet.ProvideFleetInfo:output_type -> proto.EvFleetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_evfleet_proto_init() }
func file_evfleet_proto_init() {
	if File_evfleet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evfleet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvFleetSupply); i {
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
		file_evfleet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
		file_evfleet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleList); i {
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
		file_evfleet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvFleetResponse); i {
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
			RawDescriptor: file_evfleet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evfleet_proto_goTypes,
		DependencyIndexes: file_evfleet_proto_depIdxs,
		MessageInfos:      file_evfleet_proto_msgTypes,
	}.Build()
	File_evfleet_proto = out.File
	file_evfleet_proto_rawDesc = nil
	file_evfleet_proto_goTypes = nil
	file_evfleet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EvFleetClient is the client API for EvFleet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvFleetClient interface {
	NotifyFleetInfo(ctx context.Context, in *EvFleetSupply, opts ...grpc.CallOption) (*EvFleetResponse, error)
	ProvideFleetInfo(ctx context.Context, in *VehicleList, opts ...grpc.CallOption) (*EvFleetResponse, error)
}

type evFleetClient struct {
	cc grpc.ClientConnInterface
}

func NewEvFleetClient(cc grpc.ClientConnInterface) EvFleetClient {
	return &evFleetClient{cc}
}

func (c *evFleetClient) NotifyFleetInfo(ctx context.Context, in *EvFleetSupply, opts ...grpc.CallOption) (*EvFleetResponse, error) {
	out := new(EvFleetResponse)
	err := c.cc.Invoke(ctx, "/proto.EvFleet/NotifyFleetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evFleetClient) ProvideFleetInfo(ctx context.Context, in *VehicleList, opts ...grpc.CallOption) (*EvFleetResponse, error) {
	out := new(EvFleetResponse)
	err := c.cc.Invoke(ctx, "/proto.EvFleet/ProvideFleetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvFleetServer is the server API for EvFleet service.
type EvFleetServer interface {
	NotifyFleetInfo(context.Context, *EvFleetSupply) (*EvFleetResponse, error)
	ProvideFleetInfo(context.Context, *VehicleList) (*EvFleetResponse, error)
}

// UnimplementedEvFleetServer can be embedded to have forward compatible implementations.
type UnimplementedEvFleetServer struct {
}

func (*UnimplementedEvFleetServer) NotifyFleetInfo(context.Context, *EvFleetSupply) (*EvFleetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyFleetInfo not implemented")
}
func (*UnimplementedEvFleetServer) ProvideFleetInfo(context.Context, *VehicleList) (*EvFleetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvideFleetInfo not implemented")
}

func RegisterEvFleetServer(s *grpc.Server, srv EvFleetServer) {
	s.RegisterService(&_EvFleet_serviceDesc, srv)
}

func _EvFleet_NotifyFleetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvFleetSupply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvFleetServer).NotifyFleetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EvFleet/NotifyFleetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvFleetServer).NotifyFleetInfo(ctx, req.(*EvFleetSupply))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvFleet_ProvideFleetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehicleList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvFleetServer).ProvideFleetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EvFleet/ProvideFleetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvFleetServer).ProvideFleetInfo(ctx, req.(*VehicleList))
	}
	return interceptor(ctx, in, info, handler)
}

var _EvFleet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EvFleet",
	HandlerType: (*EvFleetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyFleetInfo",
			Handler:    _EvFleet_NotifyFleetInfo_Handler,
		},
		{
			MethodName: "ProvideFleetInfo",
			Handler:    _EvFleet_ProvideFleetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evfleet.proto",
}
