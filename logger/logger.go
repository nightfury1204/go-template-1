package logger

type StructLogger interface {
	Println(fn, tid string, msg string)
	Printf(fn, tid string, format string, args ...interface{})
	Warnln(fn, tid string, msg string)
	Errorln(fn, tid string, msg string)
	Errorf(fn, tid string, format string, args ...interface{})
	Print(level LogLevel, fn, tid string, msg string)
}

var (
	DefaultOutStructLogger StructLogger
)

func init() {
	DefaultOutStructLogger = NewZeroLevelLogger()
}
