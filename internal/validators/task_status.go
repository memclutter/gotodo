package validators

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/memclutter/gotodo/internal/models"
)

func TaskStatus(ctx context.Context, fl validator.FieldLevel) bool {
	return coreslices.StringIn(fl.Field().String(), models.TaskStatuses)
}
