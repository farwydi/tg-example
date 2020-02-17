package tg

import (
	"context"
	"net"
	"net/http"
	"sync"
	"tgex/domain"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Register(cnf *domain.Config, lifecycle fx.Lifecycle, down fx.Shutdowner, logger *zap.Logger, p Params) error {
	_app = &App{
		Params:   p,
		logger:   logger,
		start:    time.Now(),
		done:     make(chan struct{}),
		endpoint: cnf.Endpoint,
	}

	srv := &http.Server{
		Handler: nil,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Start tg...")

			ln, err := net.Listen("tcp", ":8080")
			if err != nil {
				return err
			}

			err = _app.Init("793120747:AAEuZmW8tn-xqold391QimysCes3hjP6Dns")
			if err != nil {
				return err
			}

			go func() {
				if err := srv.Serve(ln); err != nil {
					logger.Error("Fail start http server", zap.Error(err))
				}
			}()

			go func() {
				time.Sleep(time.Second)
				if err := _app.Setup(); err != nil {
					logger.Error("Fail setup tg", zap.Error(err))

					if err := down.Shutdown(); err != nil {
						logger.Error("Fail shutdown", zap.Error(err))
					}
				}

				go restart(logger, _app.Event)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stop tg...")

			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				if err := srv.Shutdown(ctx); err != nil {
					logger.Error("Fail stop tg app", zap.Error(err))
				}
			}()
			go func() {
				defer wg.Done()
				if err := _app.Shutdown(ctx); err != nil {
					logger.Error("Fail stop http server", zap.Error(err))
				}
			}()
			wg.Wait()

			return nil
		},
	})

	return nil
}
