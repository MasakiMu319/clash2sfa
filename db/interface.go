package db

import (
	"context"

	"github.com/xmdhs/clash2sfa/modle"
)

type DB interface {
	GetArg(cxt context.Context, blake3 string) (modle.ConvertArg, error)
	PutArg(cxt context.Context, blake3 string, arg modle.ConvertArg) error
}
