package config

var (
	// RabbitURL : rabbitmq服务的入口url
	RabbitURL = "amqp://guest:guest@192.168.88.131:5672/"
)

const (
	// AsyncTransferEnable : 是否开启文件异步转移(默认同步)
	AsyncTransferEnable = true
	// TransExchangeName : 用于文件transfer的交换机
	TransExchangeName = "uploadserver.trans"
	// TransOSSQueueName : oss转移队列名
	TransOSSQueueName = "uploadserver.trans.kodo"
	// TransOSSErrQueueName : oss转移失败后写入另一个队列的队列名
	TransOSSErrQueueName = "uploadserver.trans.kodo.err"
	// TransOSSRoutingKey : routingkey
	TransOSSRoutingKey = "kodo"
)
