package di

import (
	"secbank.api/internal/interfaces"
	"sync"
)

type CustomerDI struct {
	interfaces.GenericDI
}

func (customerDI *CustomerDI) Initialize() {

}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
