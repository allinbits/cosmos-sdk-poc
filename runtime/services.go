package runtime

import (
	"fmt"

	"github.com/fdymylja/tmos/runtime/module"
	"github.com/hashicorp/go-multierror"
	"k8s.io/klog/v2"
)

// ExtensionService wraps one module.ExtensionService and extends its name to contain the module name too
type ExtensionService struct {
	ModuleName string
	module.ExtensionService
}

func (s ExtensionService) Name() string {
	return fmt.Sprintf("%s.%s", s.ModuleName, s.ExtensionService.Name())
}

func NewServiceOrchestrator() *ServiceGroup {
	return &ServiceGroup{}
}

// ServiceGroup contains a group of multiple extension services of modules.
type ServiceGroup struct {
	initialized bool
	services    []ExtensionService
}

func (s *ServiceGroup) AddServices(moduleName string, xts ...module.ExtensionService) {
	for _, xt := range xts {
		s.services = append(s.services, ExtensionService{
			ModuleName:       moduleName,
			ExtensionService: xt,
		})
	}
}

func (s *ServiceGroup) Start() error {
	for _, svc := range s.services {
		klog.Infof("starting service %s", svc.Name())
		err := svc.Start()
		if err != nil {
			return fmt.Errorf("unable to start service %s: %w", svc.Name(), err)
		}
	}
	return nil
}

func (s *ServiceGroup) Stop() error {
	errGroup := new(multierror.Error)
	for _, svc := range s.services {
		klog.Infof("stopping service %s", svc.Name())
		err := svc.Stop()
		if err != nil {
			klog.Infof("unable to stop service %s", svc.Name())
			errGroup = multierror.Append(errGroup, err)
		}
	}
	return errGroup.ErrorOrNil()
}
