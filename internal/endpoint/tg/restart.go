package tg

import (
	"time"

	"go.uber.org/zap"
)

func restart(logger *zap.Logger, callback func()) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic on event", zap.Any("recover", r))

			time.Sleep(5 * time.Second)
		}

		go restart(logger, callback)
	}()

	callback()
}
