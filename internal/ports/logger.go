package ports

type ZapLogger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}
