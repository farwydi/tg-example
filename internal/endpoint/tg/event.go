package tg

import (
	"context"
	"go.uber.org/zap"
	"time"
)

func (app App) Event() {
	for {
		select {
		case <-app.done:
			app.logger.Info("Tg done")
			return
		case update := <-app.updates:
			go func() {
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				if err := app.Update(ctx, update); err != nil {
					app.logger.Error("Fail message processing", zap.Error(err))
				}
			}()
		}
	}
}
