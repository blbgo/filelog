package filelog

import (
	"github.com/blbgo/general"
)

type factory struct {
	Config
}

// NewFactory provides a general.LoggerFactory interface that creates fileLogs (returned as
// general.Logger)
func NewFactory(config Config) general.LoggerFactory {
	return &factory{Config: config}
}

// New creates a new fileLog and returns it as a general.Logger
func (r *factory) New(name string) (general.Logger, error) {
	if name == "" {
		name = "log"
	}
	return new(r.Config, name)
}
