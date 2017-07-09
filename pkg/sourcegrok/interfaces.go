package sourcegrok

const (
	//LogLevelDebug set loglevel to debug
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelError
)

type Logger interface {
	Log(level int, messages ...interface{})
}
