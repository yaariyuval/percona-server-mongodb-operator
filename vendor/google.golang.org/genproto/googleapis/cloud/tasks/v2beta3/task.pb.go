// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta3/task.proto

package tasks // import "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The view specifies a subset of [Task][google.cloud.tasks.v2beta3.Task] data.
//
// When a task is returned in a response, not all
// information is retrieved by default because some data, such as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that it
// contains.
type Task_View int32

const (
	// Unspecified. Defaults to BASIC.
	Task_VIEW_UNSPECIFIED Task_View = 0
	// The basic view omits fields which can be large or can contain
	// sensitive data.
	//
	// This view does not include the
	// [body in AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest.body].
	// Bodies are desirable to return only when needed, because they
	// can be large and because of the sensitivity of the data that you
	// choose to store in it.
	Task_BASIC Task_View = 1
	// All information is returned.
	//
	// Authorization for [FULL][google.cloud.tasks.v2beta3.Task.View.FULL] requires
	// `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/)
	// permission on the [Queue][google.cloud.tasks.v2beta3.Queue] resource.
	Task_FULL Task_View = 2
)

var Task_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "BASIC",
	2: "FULL",
}
var Task_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"BASIC":            1,
	"FULL":             2,
}

func (x Task_View) String() string {
	return proto.EnumName(Task_View_name, int32(x))
}
func (Task_View) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_64d7b65e49877f4b, []int{0, 0}
}

// A unit of scheduled work.
type Task struct {
	// Optionally caller-specified in [CreateTask][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//
	// The task name.
	//
	// The task name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the task's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	// * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//   hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The message to send to the worker.
	//
	// Types that are valid to be assigned to PayloadType:
	//	*Task_AppEngineHttpRequest
	//	*Task_HttpRequest
	PayloadType isTask_PayloadType `protobuf_oneof:"payload_type"`
	// The time when the task is scheduled to be attempted.
	//
	// For App Engine queues, this is when the task will be attempted or retried.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that the task was created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The deadline for requests sent to the worker. If the worker does not
	// respond by this deadline then the request is cancelled and the attempt
	// is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the
	// task according to the [RetryConfig][google.cloud.tasks.v2beta3.RetryConfig].
	//
	// Note that when the request is cancelled, Cloud Tasks will stop listing for
	// the response, but whether the worker stops processing depends on the
	// worker. For example, if the worker is stuck, it may not react to cancelled
	// requests.
	//
	// The default and maximum values depend on the type of request:
	//
	// * For [HTTP tasks][google.cloud.tasks.v2beta3.HttpRequest], the default is 10 minutes. The deadline
	//   must be in the interval [15 seconds, 30 minutes].
	//
	// * For [App Engine tasks][google.cloud.tasks.v2beta3.AppEngineHttpRequest], 0 indicates that the
	//   request has the default deadline. The default deadline depends on the
	//   [scaling
	//   type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling)
	//   of the service: 10 minutes for standard apps with automatic scaling, 24
	//   hours for standard apps with manual and basic scaling, and 60 minutes for
	//   flex apps. If the request deadline is set, it must be in the interval [15
	//   seconds, 24 hours 15 seconds]. Regardless of the task's
	//   `dispatch_deadline`, the app handler will not run for longer than than
	//   the service's timeout. We recommend setting the `dispatch_deadline` to
	//   at most a few seconds more than the app handler's timeout. For more
	//   information see
	//   [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts).
	//
	// `dispatch_deadline` will be truncated to the nearest millisecond. The
	// deadline is an approximate deadline.
	DispatchDeadline *duration.Duration `protobuf:"bytes,12,opt,name=dispatch_deadline,json=dispatchDeadline,proto3" json:"dispatch_deadline,omitempty"`
	// Output only. The number of attempts dispatched.
	//
	// This count includes attempts which have been dispatched but haven't
	// received a response.
	DispatchCount int32 `protobuf:"varint,6,opt,name=dispatch_count,json=dispatchCount,proto3" json:"dispatch_count,omitempty"`
	// Output only. The number of attempts which have received a response.
	ResponseCount int32 `protobuf:"varint,7,opt,name=response_count,json=responseCount,proto3" json:"response_count,omitempty"`
	// Output only. The status of the task's first attempt.
	//
	// Only [dispatch_time][google.cloud.tasks.v2beta3.Attempt.dispatch_time] will be set.
	// The other [Attempt][google.cloud.tasks.v2beta3.Attempt] information is not retained by Cloud Tasks.
	FirstAttempt *Attempt `protobuf:"bytes,8,opt,name=first_attempt,json=firstAttempt,proto3" json:"first_attempt,omitempty"`
	// Output only. The status of the task's last attempt.
	LastAttempt *Attempt `protobuf:"bytes,9,opt,name=last_attempt,json=lastAttempt,proto3" json:"last_attempt,omitempty"`
	// Output only. The view specifies which subset of the [Task][google.cloud.tasks.v2beta3.Task] has
	// been returned.
	View                 Task_View `protobuf:"varint,10,opt,name=view,proto3,enum=google.cloud.tasks.v2beta3.Task_View" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_64d7b65e49877f4b, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTask_PayloadType interface {
	isTask_PayloadType()
}

type Task_AppEngineHttpRequest struct {
	AppEngineHttpRequest *AppEngineHttpRequest `protobuf:"bytes,3,opt,name=app_engine_http_request,json=appEngineHttpRequest,proto3,oneof"`
}

type Task_HttpRequest struct {
	HttpRequest *HttpRequest `protobuf:"bytes,11,opt,name=http_request,json=httpRequest,proto3,oneof"`
}

func (*Task_AppEngineHttpRequest) isTask_PayloadType() {}

func (*Task_HttpRequest) isTask_PayloadType() {}

func (m *Task) GetPayloadType() isTask_PayloadType {
	if m != nil {
		return m.PayloadType
	}
	return nil
}

func (m *Task) GetAppEngineHttpRequest() *AppEngineHttpRequest {
	if x, ok := m.GetPayloadType().(*Task_AppEngineHttpRequest); ok {
		return x.AppEngineHttpRequest
	}
	return nil
}

func (m *Task) GetHttpRequest() *HttpRequest {
	if x, ok := m.GetPayloadType().(*Task_HttpRequest); ok {
		return x.HttpRequest
	}
	return nil
}

func (m *Task) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetDispatchDeadline() *duration.Duration {
	if m != nil {
		return m.DispatchDeadline
	}
	return nil
}

func (m *Task) GetDispatchCount() int32 {
	if m != nil {
		return m.DispatchCount
	}
	return 0
}

func (m *Task) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *Task) GetFirstAttempt() *Attempt {
	if m != nil {
		return m.FirstAttempt
	}
	return nil
}

func (m *Task) GetLastAttempt() *Attempt {
	if m != nil {
		return m.LastAttempt
	}
	return nil
}

func (m *Task) GetView() Task_View {
	if m != nil {
		return m.View
	}
	return Task_VIEW_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Task) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Task_OneofMarshaler, _Task_OneofUnmarshaler, _Task_OneofSizer, []interface{}{
		(*Task_AppEngineHttpRequest)(nil),
		(*Task_HttpRequest)(nil),
	}
}

func _Task_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Task)
	// payload_type
	switch x := m.PayloadType.(type) {
	case *Task_AppEngineHttpRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppEngineHttpRequest); err != nil {
			return err
		}
	case *Task_HttpRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Task.PayloadType has unexpected type %T", x)
	}
	return nil
}

func _Task_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Task)
	switch tag {
	case 3: // payload_type.app_engine_http_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppEngineHttpRequest)
		err := b.DecodeMessage(msg)
		m.PayloadType = &Task_AppEngineHttpRequest{msg}
		return true, err
	case 11: // payload_type.http_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpRequest)
		err := b.DecodeMessage(msg)
		m.PayloadType = &Task_HttpRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Task_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Task)
	// payload_type
	switch x := m.PayloadType.(type) {
	case *Task_AppEngineHttpRequest:
		s := proto.Size(x.AppEngineHttpRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Task_HttpRequest:
		s := proto.Size(x.HttpRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The status of a task attempt.
type Attempt struct {
	// Output only. The time that this attempt was scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that this attempt was dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=dispatch_time,json=dispatchTime,proto3" json:"dispatch_time,omitempty"`
	// Output only. The time that this attempt response was received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// Output only. The response from the worker for this attempt.
	//
	// If `response_time` is unset, then the task has not been attempted or is
	// currently running and the `response_status` field is meaningless.
	ResponseStatus       *status.Status `protobuf:"bytes,4,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_64d7b65e49877f4b, []int{1}
}
func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (dst *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(dst, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Attempt) GetDispatchTime() *timestamp.Timestamp {
	if m != nil {
		return m.DispatchTime
	}
	return nil
}

func (m *Attempt) GetResponseTime() *timestamp.Timestamp {
	if m != nil {
		return m.ResponseTime
	}
	return nil
}

func (m *Attempt) GetResponseStatus() *status.Status {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "google.cloud.tasks.v2beta3.Task")
	proto.RegisterType((*Attempt)(nil), "google.cloud.tasks.v2beta3.Attempt")
	proto.RegisterEnum("google.cloud.tasks.v2beta3.Task_View", Task_View_name, Task_View_value)
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta3/task.proto", fileDescriptor_task_64d7b65e49877f4b)
}

var fileDescriptor_task_64d7b65e49877f4b = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x4f, 0xdb, 0x3c,
	0x14, 0xc6, 0x09, 0xa4, 0x40, 0xdd, 0xc0, 0xdb, 0xd7, 0x42, 0x22, 0x54, 0x13, 0xab, 0x98, 0x10,
	0xbd, 0x4a, 0x36, 0xb8, 0x9a, 0xb8, 0x40, 0xf4, 0x9f, 0x5a, 0xa9, 0x9a, 0xaa, 0x14, 0x98, 0xb4,
	0x9b, 0xc8, 0x4d, 0x4c, 0x1a, 0x91, 0xda, 0x9e, 0xed, 0x80, 0xf8, 0x08, 0xfb, 0xc4, 0xbb, 0x9d,
	0xe2, 0xd8, 0x55, 0x19, 0xac, 0xdd, 0xee, 0x7a, 0xce, 0x79, 0x9e, 0x9f, 0x8f, 0xce, 0x39, 0x0d,
	0x38, 0x4d, 0x28, 0x4d, 0x32, 0xec, 0x47, 0x19, 0xcd, 0x63, 0x5f, 0x22, 0xf1, 0x20, 0xfc, 0xc7,
	0xf3, 0x29, 0x96, 0xe8, 0x42, 0x45, 0x1e, 0xe3, 0x54, 0x52, 0xd8, 0x28, 0x65, 0x9e, 0x92, 0x79,
	0x4a, 0xe6, 0x69, 0x59, 0xe3, 0x9d, 0x46, 0x20, 0x96, 0xfa, 0x88, 0x10, 0x2a, 0x91, 0x4c, 0x29,
	0x11, 0xa5, 0xb3, 0x71, 0xb4, 0x54, 0xe5, 0x58, 0xd0, 0x9c, 0x47, 0x58, 0x97, 0xce, 0x56, 0xbe,
	0xcd, 0x13, 0x2c, 0xb5, 0xf0, 0x58, 0x0b, 0x55, 0x34, 0xcd, 0xef, 0xfd, 0x38, 0xe7, 0xea, 0x11,
	0x5d, 0x7f, 0xff, 0x7b, 0x5d, 0xa6, 0x73, 0x2c, 0x24, 0x9a, 0x33, 0x2d, 0x38, 0xd4, 0x02, 0xce,
	0x22, 0x5f, 0x48, 0x24, 0x73, 0xdd, 0xdd, 0xc9, 0xcf, 0x0a, 0xb0, 0x6f, 0x90, 0x78, 0x80, 0x10,
	0xd8, 0x04, 0xcd, 0xb1, 0x6b, 0x35, 0xad, 0x56, 0x35, 0x50, 0xbf, 0x61, 0x0a, 0x0e, 0x11, 0x63,
	0x21, 0x26, 0x49, 0x4a, 0x70, 0x38, 0x93, 0x92, 0x85, 0x1c, 0x7f, 0xcf, 0xb1, 0x90, 0xee, 0x56,
	0xd3, 0x6a, 0xd5, 0xce, 0x3f, 0x7a, 0x7f, 0x1e, 0x8b, 0x77, 0xcd, 0x58, 0x4f, 0x39, 0x07, 0x52,
	0xb2, 0xa0, 0xf4, 0x0d, 0x36, 0x82, 0x03, 0xf4, 0x46, 0x1e, 0x8e, 0x80, 0xf3, 0x82, 0x5f, 0x53,
	0xfc, 0xb3, 0x55, 0xfc, 0x97, 0xd8, 0xda, 0x6c, 0x89, 0x76, 0x05, 0xf6, 0x44, 0x34, 0xc3, 0x71,
	0x9e, 0xe1, 0xb0, 0x18, 0x85, 0x6b, 0x2b, 0x5c, 0xc3, 0xe0, 0xcc, 0x9c, 0xbc, 0x1b, 0x33, 0xa7,
	0xc0, 0x31, 0x86, 0x22, 0x05, 0x2f, 0x41, 0x2d, 0xe2, 0x18, 0x49, 0x6d, 0xaf, 0xac, 0xb5, 0x83,
	0x52, 0xae, 0xcc, 0x7d, 0xf0, 0x7f, 0x9c, 0x0a, 0x86, 0x64, 0x34, 0x0b, 0x63, 0x8c, 0xe2, 0x2c,
	0x25, 0xd8, 0x75, 0x14, 0xe2, 0xe8, 0x15, 0xa2, 0xab, 0x37, 0x19, 0xd4, 0x8d, 0xa7, 0xab, 0x2d,
	0xf0, 0x14, 0xec, 0x2f, 0x38, 0x11, 0xcd, 0x89, 0x74, 0xb7, 0x9b, 0x56, 0xab, 0x12, 0xec, 0x99,
	0x6c, 0xa7, 0x48, 0x16, 0x32, 0x8e, 0x05, 0xa3, 0x44, 0x60, 0x2d, 0xdb, 0x29, 0x65, 0x26, 0x5b,
	0xca, 0x06, 0x60, 0xef, 0x3e, 0xe5, 0x42, 0x86, 0x48, 0x4a, 0x3c, 0x67, 0xd2, 0xdd, 0x55, 0x1d,
	0x7d, 0x58, 0xb9, 0xc2, 0x52, 0x1a, 0x38, 0xca, 0xa9, 0x23, 0xd8, 0x07, 0x4e, 0x86, 0x96, 0x40,
	0xd5, 0xbf, 0x07, 0xd5, 0x0a, 0xa3, 0xe1, 0x7c, 0x06, 0xf6, 0x63, 0x8a, 0x9f, 0x5c, 0xd0, 0xb4,
	0x5a, 0xfb, 0xe7, 0xa7, 0xab, 0xfc, 0xc5, 0x89, 0x7a, 0x77, 0x29, 0x7e, 0x0a, 0x94, 0xe5, 0xe4,
	0x13, 0xb0, 0x8b, 0x08, 0x1e, 0x80, 0xfa, 0xdd, 0xb0, 0xf7, 0x35, 0xbc, 0xfd, 0x32, 0x19, 0xf7,
	0x3a, 0xc3, 0xfe, 0xb0, 0xd7, 0xad, 0x6f, 0xc0, 0x2a, 0xa8, 0xb4, 0xaf, 0x27, 0xc3, 0x4e, 0xdd,
	0x82, 0xbb, 0xc0, 0xee, 0xdf, 0x8e, 0x46, 0xf5, 0xcd, 0xf6, 0x3e, 0x70, 0x18, 0x7a, 0xce, 0x28,
	0x8a, 0x43, 0xf9, 0xcc, 0xf0, 0xc9, 0x8f, 0x4d, 0xb0, 0x63, 0x3a, 0x79, 0x75, 0x2f, 0xd6, 0x3f,
	0xde, 0xcb, 0x15, 0x58, 0x2c, 0xa5, 0x04, 0x6c, 0xae, 0x07, 0x18, 0x83, 0x01, 0x2c, 0x96, 0xa8,
	0x00, 0x5b, 0xeb, 0x01, 0xc6, 0xa0, 0x2f, 0xf6, 0xbf, 0x05, 0xa0, 0xfc, 0x87, 0xeb, 0xa3, 0x87,
	0x06, 0xc1, 0x59, 0xe4, 0x4d, 0x54, 0x25, 0x58, 0x1c, 0x4c, 0x19, 0xb7, 0x09, 0x38, 0x8e, 0xe8,
	0x7c, 0xc5, 0x02, 0xda, 0xd5, 0x62, 0x03, 0xe3, 0xa2, 0x89, 0xb1, 0xf5, 0xed, 0x4a, 0x0b, 0x13,
	0x9a, 0x21, 0x92, 0x78, 0x94, 0x27, 0x7e, 0x82, 0x89, 0x6a, 0xd1, 0x2f, 0x4b, 0x88, 0xa5, 0xe2,
	0xad, 0xcf, 0xda, 0xa5, 0x8a, 0xa6, 0xdb, 0x4a, 0x7b, 0xf1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0x28, 0xef, 0x11, 0x7d, 0x05, 0x00, 0x00,
}
