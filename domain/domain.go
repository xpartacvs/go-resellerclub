package domain

import (
	"fmt"

	"github.com/xpartacvs/go-resellerclub/core"
)

type domain struct {
	core core.Core
}

type Domain interface {
	Test()
}

func New(c core.Core) Domain {
	return &domain{
		core: c,
	}
}

func (d *domain) Test() {
	fmt.Println("Hello")
}
