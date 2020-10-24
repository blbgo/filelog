package filelog

// Config provides config values for fileLog implementations
type Config interface {
	LogDir() string
}
