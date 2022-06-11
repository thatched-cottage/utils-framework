package monitor

type monitorFlag int

const (
	Close    monitorFlag       = 1 << iota //	关闭监控，不采集任何信息。
	Func                                   //	开启监控，输出主要信息（调用时间，和调用方信息）
	Periodic                               //	开启监控，每分钟输出平均调用时间，以及调用次数
	All      = Func | Periodic             //	全部信息，打印所有信息
)
