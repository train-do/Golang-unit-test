package main

import (
	"be-golang-chapter-36-implem/infra"
	"be-golang-chapter-36-implem/router"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("Error", zap.Error(err))
		return
	}

	router.SetupReouter(ctx)
}
