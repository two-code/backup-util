package driver

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/two-code/backup-util/pkg/plan"
)

type Driver interface {
	Name() string

	Exec(ctx context.Context, p plan.Plan, l log.Logger) (ExecReport, error)
}
