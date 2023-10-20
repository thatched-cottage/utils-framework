package codes

import (
	"fmt"
	"strconv"
)

// Code 是 规范中定义的无符号32位错误码。
type Code uint32

const (
	// OK 表示成功。
	OK Code = 0

	// Canceled 表示操作被取消（通常是由调用方发起的）。
	//
	// 当请求取消时， 框架会生成此错误码。
	Canceled Code = 1

	// Unknown 错误。如果来自另一个地址空间的 Status 值属于本地址空间中未知的错误空间，则可能返回此错误。
	// 如果 API 不返回足够的错误信息，可能会转换为此错误。
	//
	// 在上述两种情况下， 框架会生成此错误码。
	Unknown Code = 2

	// InvalidArgument 表示客户端指定了无效参数。
	// 注意，与 FailedPrecondition 不同。它表示问题参数不管系统状态如何都有问题（例如，格式错误的文件名）。
	//
	//  框架不会生成此错误码。
	InvalidArgument Code = 3

	// DeadlineExceeded 表示操作在完成之前已过期。
	// 对于改变系统状态的操作，即使操作已成功完成，也可能返回此错误。
	// 例如，服务器的成功响应可能被延迟到超过截止时间。
	//
	// 当超出截止时间时， 框架会生成此错误码。
	DeadlineExceeded Code = 4

	// NotFound 表示找不到某个请求的实体（例如，文件或目录）。
	//
	//  框架不会生成此错误码。
	NotFound Code = 5

	// AlreadyExists 表示尝试创建一个已存在的实体失败。
	//
	//  框架不会生成此错误码。
	AlreadyExists Code = 6

	// PermissionDenied 表示调用方没有执行指定操作的权限。它不能用于资源耗尽导致的拒绝访问
	// （对于这些错误，请使用 ResourceExhausted）。如果无法识别调用方身份，
	// 则不应使用此错误码（对于这些错误，请使用 Unauthenticated）。
	//
	//  核心框架不会生成此错误码，但是可以期望认证中间件生成它。
	PermissionDenied Code = 7

	// ResourceExhausted 表示某个资源已耗尽，可能是每个用户的配额或整个文件系统的空间。
	//
	// 当内存不足、服务器负载过重或消息超过配置的最大大小时， 框架会生成此错误码。
	ResourceExhausted Code = 8

	// FailedPrecondition 表示由于系统未达到操作所需的状态而拒绝了操作。
	// 例如，要删除的目录可能非空，将 rmdir 操作应用于非目录等等。
	//
	// 可以帮助服务实施者在 FailedPrecondition、Aborted 和 Unavailable 之间做决策的验收测试：
	// （a）如果客户端可以重试失败的调用，请使用 Unavailable。
	// （b）如果客户端应在更高级别重试（例如，重新启动读-修改-写序列），请使用 Aborted。
	// （c）如果客户端在系统状态被明确修复之前不应该重试，请使用 FailedPrecondition。例如，
	// 如果因为目录非空而导致 "rmdir" 失败，则应返回 FailedPrecondition，
	// 因为客户端在首先删除其中的文件之前不应重试。
	// （d）如果客户端对资源执行有条件的 REST Get/Update/Delete 操作，而服务器上的资源与条件不匹配，
	// 则应返回 FailedPrecondition。例如，在同一资源上发生冲突的读-修改-写。
	//

	FailedPrecondition Code = 9

	// Aborted 表示操作被取消，通常由并发冲突引起。
	//
	// 对于并发更新失败的操作（例如，存储在任何给定资源上的条件竞争），
	// 使用 Aborted。与 FailedPrecondition 不同，它确保客户端重试之前先处理其他并发操作。
	//
	//  框架不会生成此错误码。
	Aborted Code = 10

	// OutOfRange 表示请求超出了有效范围。
	// 定义了 "-inf" 和 "+inf" 之间的方法参数和字段，如果指定值超出此范围，则返回此错误。
	//
	// 注意，与 InvalidArgument 不同。InvalidArgment 是指参数类型错误或缺少必需的参数等，
	// 而 OutOfRange 是指参数值超出允许的范围。
	//
	//  框架不会生成此错误码。
	OutOfRange Code = 11

	// Unimplemented 表示服务器不支持所请求的操作。
	//
	//  框架不会生成此错误码。
	Unimplemented Code = 12

	// Internal 表示服务器内部错误。这可能是由于  实现或底层系统中的其他低级问题引起的。
	//
	//  框架不会生成此错误码。
	Internal Code = 13

	// Unavailable 表示服务当前无法使用。这可能是由于暂时的服务器过载（返回此错误的最常见原因），
	// 或者服务器已关闭维护。
	//
	//  框架不会生成此错误码。
	Unavailable Code = 14

	// DataLoss 表示由于不可恢复的数据丢失或损坏，操作失败。
	//
	//  框架不会生成此错误码。
	DataLoss Code = 15

	// Unauthenticated 表示请求没有经过身份验证。客户端未提供有效的身份验证凭据。
	//
	//  核心框架不会生成此错误码，但是可以期望认证中间件生成它。
	Unauthenticated Code = 16

	_maxCode = 17
)

var strToCode = map[string]Code{
	`"OK"`: OK,
	`"CANCELLED"`:/* [sic] */ Canceled,
	`"UNKNOWN"`:             Unknown,
	`"INVALID_ARGUMENT"`:    InvalidArgument,
	`"DEADLINE_EXCEEDED"`:   DeadlineExceeded,
	`"NOT_FOUND"`:           NotFound,
	`"ALREADY_EXISTS"`:      AlreadyExists,
	`"PERMISSION_DENIED"`:   PermissionDenied,
	`"RESOURCE_EXHAUSTED"`:  ResourceExhausted,
	`"FAILED_PRECONDITION"`: FailedPrecondition,
	`"ABORTED"`:             Aborted,
	`"OUT_OF_RANGE"`:        OutOfRange,
	`"UNIMPLEMENTED"`:       Unimplemented,
	`"INTERNAL"`:            Internal,
	`"UNAVAILABLE"`:         Unavailable,
	`"DATA_LOSS"`:           DataLoss,
	`"UNAUTHENTICATED"`:     Unauthenticated,
}

// UnmarshalJSON unmarshals b into the Code.
func (c *Code) UnmarshalJSON(b []byte) error {
	// From json.Unmarshaler: By convention, to approximate the behavior of
	// Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as
	// a no-op.
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 32); err == nil {
		if ci >= _maxCode {
			return fmt.Errorf("invalid code: %q", ci)
		}

		*c = Code(ci)
		return nil
	}

	if jc, ok := strToCode[string(b)]; ok {
		*c = jc
		return nil
	}
	return fmt.Errorf("invalid code: %q", string(b))
}
