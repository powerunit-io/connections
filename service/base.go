package service

import (
	"github.com/powerunit-io/platform/logging"
)

// BaseService -
type BaseService struct {
	*logging.Logger
}

// Start - Generic service start function
func (bs *BaseService) Start() error {
	return nil
}

// Stop - Generic service stop function
func (bs *BaseService) Stop() error {
	return nil
}
