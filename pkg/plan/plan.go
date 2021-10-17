package plan

import (
	"context"

	"github.com/two-code/backup-util/pkg/plan/models"
)

type Plan interface {
	IsReadOnly() bool

	GetMetaProps() (models.MetaProps, error)
	SetMetaProps(props models.MetaProps) error

	GetParam(ctx context.Context, key models.ParamKey) (models.ParamVal, error)
	SetParam(ctx context.Context, key models.ParamKey, val models.ParamVal) error
}
