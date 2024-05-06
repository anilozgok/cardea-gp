package utils

import (
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"time"
)

type ForgotPasswordCtx struct {
	Passcode  int
	CreatedAt time.Time
	Verified  bool
	Expired   bool
	User      entity.User
}
