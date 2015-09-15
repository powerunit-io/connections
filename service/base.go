package service

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/managers"
)

// BaseService -
type BaseService struct {
	*logging.Logger

	Workers     managers.Manager
	Connections managers.Manager
	Devices     managers.Manager

	Config *config.Config

	done chan bool
}

// Version - Service version
func (bs *BaseService) Version() float64 {
	return bs.Config.Get("service_version").(float64)
}

// Name - Service name
func (bs *BaseService) Name() string {
	return bs.Config.Get("service_name").(string)
}
