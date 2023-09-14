package common

const (
	LabelMethod              = "method"
	LabelErr                 = "err"
	LabelTook                = "took"
	LabelError               = "error"
	LabelAuthorization       = "Authorization"
	LabelContentType         = "Content-Type"
	LabelRequest             = "request"
	LabelResponse            = "response"
	LabelApplicationJsonUtf8 = "application/json; charset=utf-8"
	LabelUser                = "user"
	LabelId                  = "Id"
	LabelUserId              = "UserId"
	AllowedOrigins           = "*"
	StrategyLabel            = "Strategy"
)

const (
	ServerStartingMsg = "Server starting on"
	ServerFailedMsg   = "Server failed to start"
	LabelMsg          = "msg"
	LabelAddress      = "address"
	LabelTimeStamp    = "timestamp"
)

const (
	RequestCountName   = "request_count"
	RequestCountHelp   = "Number of requests received"
	RequestLatencyName = "request_latency_microseconds"
	RequestLatencyHelp = "Total duration of request in microseconds"
	CountResultName    = "count_result"
	CountResultHelp    = "The result of count method"
)
