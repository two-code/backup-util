package simple

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/two-code/backup-util/pkg/plan"
)

type Driver = interface {
	Name() string

	Exec(ctx context.Context, p plan.Plan, l log.Logger) (ExecReport, error)
}

type simple struct {
}

const (
	driverName = "simple"
)

func (drv *simple) Exec(ctx context.Context, p plan.Plan, l log.Logger) (ExecReport, error) {
	panic("not implemented yet")
}

func (drv *simple) Name() string {
	return driverName
}

func New() Driver {
	return &simple{}
}
