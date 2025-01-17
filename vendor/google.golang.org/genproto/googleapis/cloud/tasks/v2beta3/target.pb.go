// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta3/target.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The HTTP method used to execute the task.
type HttpMethod int32

const (
	// HTTP method unspecified
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
	// HTTP PATCH
	HttpMethod_PATCH HttpMethod = 6
	// HTTP OPTIONS
	HttpMethod_OPTIONS HttpMethod = 7
)

var HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNSPECIFIED",
	1: "POST",
	2: "GET",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
	6: "PATCH",
	7: "OPTIONS",
}

var HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNSPECIFIED": 0,
	"POST":                    1,
	"GET":                     2,
	"HEAD":                    3,
	"PUT":                     4,
	"DELETE":                  5,
	"PATCH":                   6,
	"OPTIONS":                 7,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}

func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{0}
}

// HTTP request.
//
// Warning: This is an [alpha](https://cloud.google.com/terms/launch-stages)
// feature. If you haven't already joined, you can [use this form to sign
// up](https://docs.google.com/forms/d/e/1FAIpQLSfc4uEy9CBHKYUSdnY1hdhKDCX7julVZHy3imOiR-XrU7bUNQ/viewform).
//
// The task will be pushed to the worker as an HTTP request. If the worker
// or the redirected worker acknowledges the task by returning a successful HTTP
// response code ([`200` - `299`]), the task will removed from the queue. If
// any other HTTP response code is returned or no response is received, the
// task will be retried according to the following:
//
// * User-specified throttling: [retry configuration][Queue.RetryConfig],
//   [rate limits][Queue.RateLimits], and the [queue's state][google.cloud.tasks.v2beta3.Queue.state].
//
// * System throttling: To prevent the worker from overloading, Cloud Tasks may
//   temporarily reduce the queue's effective rate. User-specified settings
//   will not be changed.
//
//  System throttling happens because:
//
//   * Cloud Tasks backoffs on all errors. Normally the backoff specified in
//     [rate limits][Queue.RateLimits] will be used. But if the worker returns
//     `429` (Too Many Requests), `503` (Service Unavailable), or the rate of
//     errors is high, Cloud Tasks will use a higher backoff rate. The retry
//     specified in the `Retry-After` HTTP response header is considered.
//
//   * To prevent traffic spikes and to smooth sudden large traffic spikes,
//     dispatches ramp up slowly when the queue is newly created or idle and
//     if large numbers of tasks suddenly become available to dispatch (due to
//     spikes in create task rates, the queue being unpaused, or many tasks
//     that are scheduled at the same time).
type HttpRequest struct {
	// Required. The full url path that the request will be sent to.
	//
	// This string must begin with either "http://" or "https://". Some examples
	// are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Tasks will
	// encode some characters for safety and compatibility. The maximum allowed
	// URL length is 2083 characters after encoding.
	//
	// The `Location` header response from a redirect response [`300` - `399`]
	// may be followed. The redirect is not counted as a separate attempt.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The HTTP method to use for the request. The default is POST.
	HttpMethod HttpMethod `protobuf:"varint,2,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.tasks.v2beta3.HttpMethod" json:"http_method,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//
	// These headers represent a subset of the headers that will accompany the
	// task's HTTP request. Some HTTP request headers will be ignored or replaced.
	//
	// A partial list of headers that will be ignored or replaced is:
	//
	// * Host: This will be computed by Cloud Tasks and derived from
	//   [HttpRequest.url][google.cloud.tasks.v2beta3.HttpRequest.url].
	// * Content-Length: This will be computed by Cloud Tasks.
	// * User-Agent: This will be set to `"Google-Cloud-Tasks"`.
	// * X-Google-*: Google use only.
	// * X-AppEngine-*: Google use only.
	//
	// `Content-Type` won't be set by Cloud Tasks. You can explicitly set
	// `Content-Type` to a media type when the
	//  [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//  For example, `Content-Type` can be set to `"application/octet-stream"` or
	//  `"application/json"`.
	//
	// Headers which can have multiple values (according to RFC2616) can be
	// specified using comma-separated values.
	//
	// The size of the headers must be less than 80KB.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request body.
	//
	// A request body is allowed only if the
	// [HTTP method][google.cloud.tasks.v2beta3.HttpRequest.http_method] is POST, PUT, or PATCH. It is an
	// error to set body on a task with an incompatible [HttpMethod][google.cloud.tasks.v2beta3.HttpMethod].
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// The mode for generating an `Authorization` header for HTTP requests.
	//
	// If specified, all `Authorization` headers in the [HttpTarget.headers][]
	// field will be overridden.
	//
	// Types that are valid to be assigned to AuthorizationHeader:
	//	*HttpRequest_OauthToken
	//	*HttpRequest_OidcToken
	AuthorizationHeader  isHttpRequest_AuthorizationHeader `protobuf_oneof:"authorization_header"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HttpRequest) Reset()         { *m = HttpRequest{} }
func (m *HttpRequest) String() string { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()    {}
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{0}
}

func (m *HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRequest.Unmarshal(m, b)
}
func (m *HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRequest.Marshal(b, m, deterministic)
}
func (m *HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRequest.Merge(m, src)
}
func (m *HttpRequest) XXX_Size() int {
	return xxx_messageInfo_HttpRequest.Size(m)
}
func (m *HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRequest proto.InternalMessageInfo

func (m *HttpRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpRequest) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type isHttpRequest_AuthorizationHeader interface {
	isHttpRequest_AuthorizationHeader()
}

type HttpRequest_OauthToken struct {
	OauthToken *OAuthToken `protobuf:"bytes,5,opt,name=oauth_token,json=oauthToken,proto3,oneof"`
}

type HttpRequest_OidcToken struct {
	OidcToken *OidcToken `protobuf:"bytes,6,opt,name=oidc_token,json=oidcToken,proto3,oneof"`
}

func (*HttpRequest_OauthToken) isHttpRequest_AuthorizationHeader() {}

func (*HttpRequest_OidcToken) isHttpRequest_AuthorizationHeader() {}

func (m *HttpRequest) GetAuthorizationHeader() isHttpRequest_AuthorizationHeader {
	if m != nil {
		return m.AuthorizationHeader
	}
	return nil
}

func (m *HttpRequest) GetOauthToken() *OAuthToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpRequest_OauthToken); ok {
		return x.OauthToken
	}
	return nil
}

func (m *HttpRequest) GetOidcToken() *OidcToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpRequest_OidcToken); ok {
		return x.OidcToken
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRequest_OauthToken)(nil),
		(*HttpRequest_OidcToken)(nil),
	}
}

// App Engine HTTP queue.
//
// The task will be delivered to the App Engine application hostname
// specified by its [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] and [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest].
// The documentation for [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest] explains how the
// task's host URL is constructed.
//
// Using [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
type AppEngineHttpQueue struct {
	// Overrides for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	//
	// If set, `app_engine_routing_override` is used for all tasks in
	// the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,1,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *AppEngineHttpQueue) Reset()         { *m = AppEngineHttpQueue{} }
func (m *AppEngineHttpQueue) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpQueue) ProtoMessage()    {}
func (*AppEngineHttpQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{1}
}

func (m *AppEngineHttpQueue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpQueue.Unmarshal(m, b)
}
func (m *AppEngineHttpQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpQueue.Marshal(b, m, deterministic)
}
func (m *AppEngineHttpQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpQueue.Merge(m, src)
}
func (m *AppEngineHttpQueue) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpQueue.Size(m)
}
func (m *AppEngineHttpQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpQueue.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpQueue proto.InternalMessageInfo

func (m *AppEngineHttpQueue) GetAppEngineRoutingOverride() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRoutingOverride
	}
	return nil
}

// App Engine HTTP request.
//
// The message defines the HTTP request that is sent to an App Engine app when
// the task is dispatched.
//
// This proto can only be used for tasks in a queue which has
// [app_engine_http_queue][google.cloud.tasks.v2beta3.Queue.app_engine_http_queue] set.
//
// Using [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
//
// The task will be delivered to the App Engine app which belongs to the same
// project as the queue. For more information, see
// [How Requests are
// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by
// [dispatch
// files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
// Traffic is encrypted during transport and never leaves Google datacenters.
// Because this traffic is carried over a communication mechanism internal to
// Google, you cannot explicitly set the protocol (for example, HTTP or HTTPS).
// The request to the handler, however, will appear to have used the HTTP
// protocol.
//
// The [AppEngineRouting][google.cloud.tasks.v2beta3.AppEngineRouting] used to construct the URL that the task is
// delivered to can be set at the queue-level or task-level:
//
// * If set,
//    [app_engine_routing_override][google.cloud.tasks.v2beta3.AppEngineHttpQueue.app_engine_routing_override]
//    is used for all tasks in the queue, no matter what the setting
//    is for the
//    [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
//
//
// The `url` that the task will be sent to is:
//
// * `url =` [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] `+`
//   [relative_uri][google.cloud.tasks.v2beta3.AppEngineHttpRequest.relative_uri]
//
// Tasks can be dispatched to secure app handlers, unsecure app handlers, and
// URIs restricted with
// [`login:
// admin`](https://cloud.google.com/appengine/docs/standard/python/config/appref).
// Because tasks are not run as any user, they cannot be dispatched to URIs
// restricted with
// [`login:
// required`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// Task dispatches also do not follow redirects.
//
// The task attempt has succeeded if the app's request handler returns
// an HTTP response code in the range [`200` - `299`]. `503` is
// considered an App Engine system error instead of an application
// error. Requests returning error `503` will be retried regardless of
// retry configuration and not counted against retry counts.
// Any other response code or a failure to receive a response before the
// deadline is a failed attempt.
type AppEngineHttpRequest struct {
	// The HTTP method to use for the request. The default is POST.
	//
	// The app's request handler for the task's target URL must be able to handle
	// HTTP requests with this http_method, otherwise the task attempt will fail
	// with error code 405 (Method Not Allowed). See
	// [Writing a push task request
	// handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler)
	// and the documentation for the request handlers in the language your app is
	// written in e.g.
	// [Python Request
	// Handler](https://cloud.google.com/appengine/docs/python/tools/webapp/requesthandlerclass).
	HttpMethod HttpMethod `protobuf:"varint,1,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.tasks.v2beta3.HttpMethod" json:"http_method,omitempty"`
	// Task-level setting for App Engine routing.
	//
	// If set,
	// [app_engine_routing_override][google.cloud.tasks.v2beta3.AppEngineHttpQueue.app_engine_routing_override]
	// is used for all tasks in the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	AppEngineRouting *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing,json=appEngineRouting,proto3" json:"app_engine_routing,omitempty"`
	// The relative URI.
	//
	// The relative URI must begin with "/" and must be a valid HTTP relative URI.
	// It can contain a path and query string arguments.
	// If the relative URI is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters.
	RelativeUri string `protobuf:"bytes,3,opt,name=relative_uri,json=relativeUri,proto3" json:"relative_uri,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	// Repeated headers are not supported but a header value can contain commas.
	//
	// Cloud Tasks sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//   This header can be modified, but Cloud Tasks will append
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//   modified `User-Agent`.
	//
	// If the task has a [body][google.cloud.tasks.v2beta3.AppEngineHttpRequest.body], Cloud
	// Tasks sets the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   `"application/octet-stream"`. The default can be overridden by explicitly
	//   setting `Content-Type` to a particular media type when the
	//   [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//   For example, `Content-Type` can be set to `"application/json"`.
	// * `Content-Length`: This is computed by Cloud Tasks. This value is
	//   output only.   It cannot be changed.
	//
	// The headers below cannot be set or overridden:
	//
	// * `Host`
	// * `X-Google-*`
	// * `X-AppEngine-*`
	//
	// In addition, Cloud Tasks sets some headers when the task is dispatched,
	// such as headers containing information about the task; see
	// [request
	// headers](https://cloud.google.com/appengine/docs/python/taskqueue/push/creating-handlers#reading_request_headers).
	// These headers are set only when the task is dispatched, so they are not
	// visible when the task is returned in a Cloud Tasks response.
	//
	// Although there is no specific limit for the maximum number of headers or
	// the size, there is a limit on the maximum size of the [Task][google.cloud.tasks.v2beta3.Task]. For more
	// information, see the [CreateTask][google.cloud.tasks.v2beta3.CloudTasks.CreateTask] documentation.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request body.
	//
	// A request body is allowed only if the HTTP method is POST or PUT. It is
	// an error to set a body on a task with an incompatible [HttpMethod][google.cloud.tasks.v2beta3.HttpMethod].
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineHttpRequest) Reset()         { *m = AppEngineHttpRequest{} }
func (m *AppEngineHttpRequest) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpRequest) ProtoMessage()    {}
func (*AppEngineHttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{2}
}

func (m *AppEngineHttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpRequest.Unmarshal(m, b)
}
func (m *AppEngineHttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpRequest.Marshal(b, m, deterministic)
}
func (m *AppEngineHttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpRequest.Merge(m, src)
}
func (m *AppEngineHttpRequest) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpRequest.Size(m)
}
func (m *AppEngineHttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpRequest proto.InternalMessageInfo

func (m *AppEngineHttpRequest) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *AppEngineHttpRequest) GetAppEngineRouting() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRouting
	}
	return nil
}

func (m *AppEngineHttpRequest) GetRelativeUri() string {
	if m != nil {
		return m.RelativeUri
	}
	return ""
}

func (m *AppEngineHttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AppEngineHttpRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// App Engine Routing.
//
// Defines routing characteristics specific to App Engine - service, version,
// and instance.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	// App service.
	//
	// By default, the task is sent to the service which is the default
	// service when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance] are the empty string.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// App version.
	//
	// By default, the task is sent to the version which is the default
	// version when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance] are the empty string.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// App instance.
	//
	// By default, the task is sent to an instance which is available when
	// the task is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information, see
	// [App Engine Standard request
	// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request
	// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. The host that the task is sent to.
	//
	// The host is constructed from the domain name of the app associated with
	// the queue's project ID (for example <app-id>.appspot.com), and the
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service], [version][google.cloud.tasks.v2beta3.AppEngineRouting.version],
	// and [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. Tasks which were created using
	// the App Engine SDK might have a custom domain name.
	//
	// For more information, see
	// [How Requests are
	// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineRouting) Reset()         { *m = AppEngineRouting{} }
func (m *AppEngineRouting) String() string { return proto.CompactTextString(m) }
func (*AppEngineRouting) ProtoMessage()    {}
func (*AppEngineRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{3}
}

func (m *AppEngineRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineRouting.Unmarshal(m, b)
}
func (m *AppEngineRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineRouting.Marshal(b, m, deterministic)
}
func (m *AppEngineRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineRouting.Merge(m, src)
}
func (m *AppEngineRouting) XXX_Size() int {
	return xxx_messageInfo_AppEngineRouting.Size(m)
}
func (m *AppEngineRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineRouting proto.InternalMessageInfo

func (m *AppEngineRouting) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AppEngineRouting) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppEngineRouting) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *AppEngineRouting) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// Contains information needed for generating an
// [OAuth token](https://developers.google.com/identity/protocols/OAuth2).
// This type of authorization should be used when sending requests to a GCP
// endpoint.
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OAuth token.
	// The service account must be within the same project as the queue. The
	// caller must have iam.serviceAccounts.actAs permission for the service
	// account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// OAuth scope to be used for generating OAuth access token.
	// If not specified, "https://www.googleapis.com/auth/cloud-platform"
	// will be used.
	Scope                string   `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{4}
}

func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (m *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(m, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OAuthToken) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

// Contains information needed for generating an
// [OpenID Connect
// token](https://developers.google.com/identity/protocols/OpenIDConnect). This
// type of authorization should be used when sending requests to third party
// endpoints.
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OIDC token.
	// The service account must be within the same project as the queue. The
	// caller must have iam.serviceAccounts.actAs permission for the service
	// account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// Audience to be used when generating OIDC token. If not specified, the URI
	// specified in target will be used.
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OidcToken) Reset()         { *m = OidcToken{} }
func (m *OidcToken) String() string { return proto.CompactTextString(m) }
func (*OidcToken) ProtoMessage()    {}
func (*OidcToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{5}
}

func (m *OidcToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OidcToken.Unmarshal(m, b)
}
func (m *OidcToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OidcToken.Marshal(b, m, deterministic)
}
func (m *OidcToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OidcToken.Merge(m, src)
}
func (m *OidcToken) XXX_Size() int {
	return xxx_messageInfo_OidcToken.Size(m)
}
func (m *OidcToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OidcToken.DiscardUnknown(m)
}

var xxx_messageInfo_OidcToken proto.InternalMessageInfo

func (m *OidcToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OidcToken) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2beta3.HttpMethod", HttpMethod_name, HttpMethod_value)
	proto.RegisterType((*HttpRequest)(nil), "google.cloud.tasks.v2beta3.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.tasks.v2beta3.HttpRequest.HeadersEntry")
	proto.RegisterType((*AppEngineHttpQueue)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpQueue")
	proto.RegisterType((*AppEngineHttpRequest)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpRequest.HeadersEntry")
	proto.RegisterType((*AppEngineRouting)(nil), "google.cloud.tasks.v2beta3.AppEngineRouting")
	proto.RegisterType((*OAuthToken)(nil), "google.cloud.tasks.v2beta3.OAuthToken")
	proto.RegisterType((*OidcToken)(nil), "google.cloud.tasks.v2beta3.OidcToken")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta3/target.proto", fileDescriptor_595de6119aed6d9e)
}

var fileDescriptor_595de6119aed6d9e = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x6f, 0xd3, 0x3a,
	0x14, 0x5e, 0x96, 0xfe, 0x58, 0x4f, 0xa6, 0xab, 0xc8, 0x77, 0xf7, 0xde, 0xa8, 0xbb, 0x42, 0xa5,
	0x12, 0x50, 0x21, 0x94, 0x4a, 0x1d, 0x0f, 0x68, 0x08, 0xa1, 0x6e, 0xcd, 0xd6, 0x4a, 0x6c, 0x0d,
	0x59, 0x06, 0xd2, 0x78, 0x88, 0xbc, 0xc4, 0x4a, 0xad, 0x76, 0x71, 0x70, 0x9c, 0x4a, 0xe5, 0x8d,
	0x27, 0xfe, 0x20, 0xfe, 0x41, 0x14, 0x27, 0x69, 0xa1, 0x40, 0x19, 0xe3, 0xcd, 0xdf, 0x39, 0xc7,
	0x9f, 0xcf, 0x77, 0xf2, 0xd9, 0x81, 0x47, 0x21, 0x63, 0xe1, 0x8c, 0x74, 0xfd, 0x19, 0x4b, 0x83,
	0xae, 0xc0, 0xc9, 0x34, 0xe9, 0xce, 0x7b, 0xd7, 0x44, 0xe0, 0x83, 0xae, 0xc0, 0x3c, 0x24, 0xc2,
	0x8c, 0x39, 0x13, 0x0c, 0x35, 0xf3, 0x42, 0x53, 0x16, 0x9a, 0xb2, 0xd0, 0x2c, 0x0a, 0x9b, 0xff,
	0x17, 0x24, 0x38, 0xa6, 0x5d, 0x1c, 0x45, 0x4c, 0x60, 0x41, 0x59, 0x94, 0xe4, 0x3b, 0xdb, 0x9f,
	0x55, 0xd0, 0x86, 0x42, 0xc4, 0x0e, 0x79, 0x9f, 0x92, 0x44, 0x20, 0x1d, 0xd4, 0x94, 0xcf, 0x0c,
	0xa5, 0xa5, 0x74, 0x1a, 0x4e, 0xb6, 0x44, 0xa7, 0xa0, 0x4d, 0x84, 0x88, 0xbd, 0x1b, 0x22, 0x26,
	0x2c, 0x30, 0xb6, 0x5b, 0x4a, 0xe7, 0xaf, 0xde, 0x43, 0xf3, 0xe7, 0x27, 0x9a, 0x19, 0xdf, 0x99,
	0xac, 0x76, 0x60, 0xb2, 0x5c, 0xa3, 0x73, 0xa8, 0x4f, 0x08, 0x0e, 0x08, 0x4f, 0x0c, 0xb5, 0xa5,
	0x76, 0xb4, 0xde, 0xd3, 0x5f, 0x91, 0x14, 0x4d, 0x99, 0xc3, 0x7c, 0x9b, 0x15, 0x09, 0xbe, 0x70,
	0x4a, 0x12, 0x84, 0xa0, 0x72, 0xcd, 0x82, 0x85, 0x51, 0x69, 0x29, 0x9d, 0x5d, 0x47, 0xae, 0xd1,
	0x08, 0x34, 0x86, 0x53, 0x31, 0xf1, 0x04, 0x9b, 0x92, 0xc8, 0xa8, 0xb6, 0x94, 0x8e, 0xb6, 0xb9,
	0xd9, 0x71, 0x3f, 0x15, 0x13, 0x37, 0xab, 0x1e, 0x6e, 0x39, 0x20, 0x37, 0x4b, 0x84, 0x4e, 0x00,
	0x18, 0x0d, 0xfc, 0x82, 0xa9, 0x26, 0x99, 0x1e, 0x6c, 0x64, 0xa2, 0x81, 0x5f, 0x12, 0x35, 0x58,
	0x09, 0x9a, 0x87, 0xb0, 0xfb, 0x75, 0xff, 0xd9, 0x84, 0xa7, 0x64, 0x51, 0x4e, 0x78, 0x4a, 0x16,
	0x68, 0x0f, 0xaa, 0x73, 0x3c, 0x4b, 0x89, 0x9c, 0x6d, 0xc3, 0xc9, 0xc1, 0xe1, 0xf6, 0x33, 0xe5,
	0xe8, 0x5f, 0xd8, 0xcb, 0x1a, 0x62, 0x9c, 0x7e, 0x90, 0x5f, 0xcd, 0xcb, 0xb5, 0xb7, 0x3f, 0x2a,
	0x80, 0xfa, 0x71, 0x6c, 0x45, 0x21, 0x8d, 0x48, 0x36, 0xa9, 0xd7, 0x29, 0x49, 0x09, 0x9a, 0xc2,
	0x3e, 0x8e, 0x63, 0x8f, 0xc8, 0xb0, 0xc7, 0x59, 0x2a, 0x68, 0x14, 0x7a, 0x6c, 0x4e, 0x38, 0xa7,
	0x01, 0x91, 0x47, 0x6a, 0xbd, 0x27, 0x9b, 0x34, 0x2c, 0x49, 0x9d, 0x7c, 0xb3, 0x63, 0xe0, 0xb5,
	0xc8, 0xb8, 0x60, 0x6b, 0x7f, 0x52, 0x61, 0xef, 0x9b, 0x1e, 0x4a, 0x0b, 0xad, 0x19, 0x46, 0xb9,
	0xb3, 0x61, 0xae, 0x00, 0x7d, 0x2f, 0x47, 0x0e, 0xe9, 0x77, 0x55, 0xe8, 0xeb, 0x2a, 0xd0, 0x7d,
	0xd8, 0xe5, 0x64, 0x86, 0x05, 0x9d, 0x13, 0x2f, 0xe5, 0xd4, 0x50, 0xe5, 0xe8, 0xb5, 0x32, 0x76,
	0xc9, 0x29, 0x7a, 0xbb, 0xf2, 0x6b, 0x45, 0xfa, 0xf5, 0xc5, 0xad, 0xce, 0xbc, 0xbd, 0x71, 0xab,
	0x2b, 0xe3, 0xfe, 0x89, 0x4b, 0xda, 0x73, 0xd0, 0xd7, 0x15, 0x23, 0x03, 0xea, 0x09, 0xe1, 0x73,
	0xea, 0x93, 0x82, 0xa3, 0x84, 0x59, 0x66, 0x4e, 0x78, 0x42, 0x59, 0x54, 0x30, 0x95, 0x10, 0x35,
	0x61, 0x87, 0x46, 0x89, 0xc0, 0x91, 0x4f, 0x8a, 0x79, 0x2c, 0x71, 0xd6, 0xf3, 0x84, 0x25, 0x42,
	0x5e, 0xb6, 0x86, 0x23, 0xd7, 0xed, 0x37, 0x00, 0xab, 0xdb, 0x83, 0x7a, 0xf0, 0x4f, 0x71, 0x84,
	0x87, 0x7d, 0x9f, 0xa5, 0x91, 0xf0, 0xc8, 0x0d, 0xa6, 0xe5, 0x5b, 0xf2, 0x77, 0x91, 0xec, 0xe7,
	0x39, 0x2b, 0x4b, 0x65, 0x9a, 0x12, 0x9f, 0xc5, 0x4b, 0x4d, 0x12, 0xb4, 0xdf, 0x41, 0x63, 0x79,
	0x97, 0xee, 0x44, 0xdb, 0x84, 0x1d, 0x9c, 0x06, 0x94, 0x64, 0x42, 0x72, 0xe6, 0x25, 0x7e, 0x9c,
	0x00, 0xac, 0xec, 0x86, 0xf6, 0xe1, 0xbf, 0xa1, 0xeb, 0xda, 0xde, 0x99, 0xe5, 0x0e, 0xc7, 0x03,
	0xef, 0xf2, 0xfc, 0xc2, 0xb6, 0x8e, 0x47, 0x27, 0x23, 0x6b, 0xa0, 0x6f, 0xa1, 0x1d, 0xa8, 0xd8,
	0xe3, 0x0b, 0x57, 0x57, 0x50, 0x1d, 0xd4, 0x53, 0xcb, 0xd5, 0xb7, 0xb3, 0xd0, 0xd0, 0xea, 0x0f,
	0x74, 0x35, 0x0b, 0xd9, 0x97, 0xae, 0x5e, 0x41, 0x00, 0xb5, 0x81, 0xf5, 0xca, 0x72, 0x2d, 0xbd,
	0x8a, 0x1a, 0x50, 0xb5, 0xfb, 0xee, 0xf1, 0x50, 0xaf, 0x21, 0x0d, 0xea, 0x63, 0xdb, 0x1d, 0x8d,
	0xcf, 0x2f, 0xf4, 0xfa, 0x51, 0x0c, 0xf7, 0x7c, 0x76, 0xb3, 0xc1, 0x3e, 0x47, 0x9a, 0x2b, 0xdf,
	0x73, 0x3b, 0x7b, 0x94, 0x6d, 0xe5, 0xea, 0x65, 0x51, 0x1a, 0xb2, 0x19, 0x8e, 0x42, 0x93, 0xf1,
	0xb0, 0x1b, 0x92, 0x48, 0x3e, 0xd9, 0xdd, 0x3c, 0x85, 0x63, 0x9a, 0xfc, 0xe8, 0xc7, 0xf0, 0x5c,
	0xa2, 0xeb, 0x9a, 0xac, 0x3d, 0xf8, 0x12, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xca, 0x64, 0x96, 0x43,
	0x06, 0x00, 0x00,
}
