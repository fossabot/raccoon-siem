package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/core/globals"
)

func txMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tx, err := globals.UDBConn.NewTx(ctx)
		if err != nil {
			return
		}
		ctx.Set("tx", tx)

		ctx.Next()

		if ctx.Errors.Last() != nil {
			if err := tx.Rollback(); err != nil {
				_ = ctx.Error(err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				_ = ctx.Error(err)
			}
		}
	}
}
