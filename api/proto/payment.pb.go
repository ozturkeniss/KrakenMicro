// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc2
// source: api/proto/payment.proto

package proto

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

type ProcessPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	PaymentMethod string                 `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	mi := &file_api_proto_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessPaymentRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ProcessPaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ProcessPaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ProcessPaymentRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

type GetPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     uint32                 `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentRequest) Reset() {
	*x = GetPaymentRequest{}
	mi := &file_api_proto_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentRequest) ProtoMessage() {}

func (x *GetPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentRequest) GetPaymentId() uint32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     uint32                 `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_api_proto_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentResponse) GetPaymentId() uint32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *PaymentResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PaymentResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_api_proto_payment_proto protoreflect.FileDescriptor

const file_api_proto_payment_proto_rawDesc = "" +
	"\n" +
	"\x17api/proto/payment.proto\x12\apayment\"\x8b\x01\n" +
	"\x15ProcessPaymentRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12%\n" +
	"\x0epayment_method\x18\x04 \x01(\tR\rpaymentMethod\"2\n" +
	"\x11GetPaymentRequest\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\rR\tpaymentId\"\xb4\x01\n" +
	"\x0fPaymentResponse\x12\x1d\n" +
	"\n" +
	"payment_id\x18\x01 \x01(\rR\tpaymentId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\rR\x06userId\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt2\xa4\x01\n" +
	"\x0ePaymentService\x12L\n" +
	"\x0eProcessPayment\x12\x1e.payment.ProcessPaymentRequest\x1a\x18.payment.PaymentResponse\"\x00\x12D\n" +
	"\n" +
	"GetPayment\x12\x1a.payment.GetPaymentRequest\x1a\x18.payment.PaymentResponse\"\x00B\x13Z\x11gomicro/api/protob\x06proto3"

var (
	file_api_proto_payment_proto_rawDescOnce sync.Once
	file_api_proto_payment_proto_rawDescData []byte
)

func file_api_proto_payment_proto_rawDescGZIP() []byte {
	file_api_proto_payment_proto_rawDescOnce.Do(func() {
		file_api_proto_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_payment_proto_rawDesc), len(file_api_proto_payment_proto_rawDesc)))
	})
	return file_api_proto_payment_proto_rawDescData
}

var file_api_proto_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_payment_proto_goTypes = []any{
	(*ProcessPaymentRequest)(nil), // 0: payment.ProcessPaymentRequest
	(*GetPaymentRequest)(nil),     // 1: payment.GetPaymentRequest
	(*PaymentResponse)(nil),       // 2: payment.PaymentResponse
}
var file_api_proto_payment_proto_depIdxs = []int32{
	0, // 0: payment.PaymentService.ProcessPayment:input_type -> payment.ProcessPaymentRequest
	1, // 1: payment.PaymentService.GetPayment:input_type -> payment.GetPaymentRequest
	2, // 2: payment.PaymentService.ProcessPayment:output_type -> payment.PaymentResponse
	2, // 3: payment.PaymentService.GetPayment:output_type -> payment.PaymentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_payment_proto_init() }
func file_api_proto_payment_proto_init() {
	if File_api_proto_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_payment_proto_rawDesc), len(file_api_proto_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_payment_proto_goTypes,
		DependencyIndexes: file_api_proto_payment_proto_depIdxs,
		MessageInfos:      file_api_proto_payment_proto_msgTypes,
	}.Build()
	File_api_proto_payment_proto = out.File
	file_api_proto_payment_proto_goTypes = nil
	file_api_proto_payment_proto_depIdxs = nil
}
