package tg

import "context"

func (app App) Shutdown(_ context.Context) error {
	app.done <- struct{}{}
	return nil
}
