package core

import "github.com/gin-gonic/gin"

func txMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tx, err := UDBConn.NewTx(ctx)
		if err != nil {
			return
		}
		setTx(ctx, tx)

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