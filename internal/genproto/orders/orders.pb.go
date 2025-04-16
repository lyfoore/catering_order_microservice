// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: api/proto/orders.proto

package orders

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []int64                `protobuf:"varint,3,rep,packed,name=items,proto3" json:"items,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_api_proto_orders_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetItems() []int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []int64                `protobuf:"varint,2,rep,packed,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_api_proto_orders_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_api_proto_orders_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []int64                `protobuf:"varint,3,rep,packed,name=items,proto3" json:"items,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	mi := &file_api_proto_orders_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrderRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateOrderRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateOrderRequest) GetItems() []int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *UpdateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	mi := &file_api_proto_orders_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOrderRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []int64                `protobuf:"varint,3,rep,packed,name=items,proto3" json:"items,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	mi := &file_api_proto_orders_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{5}
}

func (x *OrderResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderResponse) GetItems() []int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Deleted       bool                   `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrderResponse) Reset() {
	*x = DeleteOrderResponse{}
	mi := &file_api_proto_orders_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderResponse) ProtoMessage() {}

func (x *DeleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_orders_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_orders_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOrderResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_api_proto_orders_proto protoreflect.FileDescriptor

const file_api_proto_orders_proto_rawDesc = "" +
	"\n" +
	"\x16api/proto/orders.proto\"^\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05items\x18\x03 \x03(\x03R\x05items\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"C\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05items\x18\x02 \x03(\x03R\x05items\"!\n" +
	"\x0fGetOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"k\n" +
	"\x12UpdateOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05items\x18\x03 \x03(\x03R\x05items\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"$\n" +
	"\x12DeleteOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\"f\n" +
	"\rOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x04R\x06userId\x12\x14\n" +
	"\x05items\x18\x03 \x03(\x03R\x05items\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"/\n" +
	"\x13DeleteOrderResponse\x12\x18\n" +
	"\adeleted\x18\x01 \x01(\bR\adeleted2\xde\x01\n" +
	"\fOrderService\x122\n" +
	"\vCreateOrder\x12\x13.CreateOrderRequest\x1a\x0e.OrderResponse\x12,\n" +
	"\bGetOrder\x12\x10.GetOrderRequest\x1a\x0e.OrderResponse\x122\n" +
	"\vUpdateOrder\x12\x13.UpdateOrderRequest\x1a\x0e.OrderResponse\x128\n" +
	"\vDeleteOrder\x12\x13.DeleteOrderRequest\x1a\x14.DeleteOrderResponseB\tZ\aorders/b\x06proto3"

var (
	file_api_proto_orders_proto_rawDescOnce sync.Once
	file_api_proto_orders_proto_rawDescData []byte
)

func file_api_proto_orders_proto_rawDescGZIP() []byte {
	file_api_proto_orders_proto_rawDescOnce.Do(func() {
		file_api_proto_orders_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_orders_proto_rawDesc), len(file_api_proto_orders_proto_rawDesc)))
	})
	return file_api_proto_orders_proto_rawDescData
}

var file_api_proto_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_orders_proto_goTypes = []any{
	(*Order)(nil),               // 0: Order
	(*CreateOrderRequest)(nil),  // 1: CreateOrderRequest
	(*GetOrderRequest)(nil),     // 2: GetOrderRequest
	(*UpdateOrderRequest)(nil),  // 3: UpdateOrderRequest
	(*DeleteOrderRequest)(nil),  // 4: DeleteOrderRequest
	(*OrderResponse)(nil),       // 5: OrderResponse
	(*DeleteOrderResponse)(nil), // 6: DeleteOrderResponse
}
var file_api_proto_orders_proto_depIdxs = []int32{
	1, // 0: OrderService.CreateOrder:input_type -> CreateOrderRequest
	2, // 1: OrderService.GetOrder:input_type -> GetOrderRequest
	3, // 2: OrderService.UpdateOrder:input_type -> UpdateOrderRequest
	4, // 3: OrderService.DeleteOrder:input_type -> DeleteOrderRequest
	5, // 4: OrderService.CreateOrder:output_type -> OrderResponse
	5, // 5: OrderService.GetOrder:output_type -> OrderResponse
	5, // 6: OrderService.UpdateOrder:output_type -> OrderResponse
	6, // 7: OrderService.DeleteOrder:output_type -> DeleteOrderResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_orders_proto_init() }
func file_api_proto_orders_proto_init() {
	if File_api_proto_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_orders_proto_rawDesc), len(file_api_proto_orders_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_orders_proto_goTypes,
		DependencyIndexes: file_api_proto_orders_proto_depIdxs,
		MessageInfos:      file_api_proto_orders_proto_msgTypes,
	}.Build()
	File_api_proto_orders_proto = out.File
	file_api_proto_orders_proto_goTypes = nil
	file_api_proto_orders_proto_depIdxs = nil
}
