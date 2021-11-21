package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/vrealzhou/geekbang_cncf/module_08/log"
)

type SignalService struct {
	ch  chan os.Signal
	log *log.Logger
}

func NewSignalService(logger *log.Logger) *SignalService {
	return &SignalService{
		ch:  make(chan os.Signal, 1),
		log: logger,
	}
}

func (srv *SignalService) Start(ctx context.Context) error {
	srv.log.Info("Waiting for signal")
	signal.Notify(srv.ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-srv.ch: // 收到指定的信号
		srv.log.Infof("signal service: got singnal %v", s)
		return normalShutdown
	case <-ctx.Done(): // 其他service终止
		return errors.Wrap(ctx.Err(), "signal service")
	}
}

func (s *SignalService) Shutdown(ctx context.Context) error {
	s.log.Info("Stop listening signals")
	close(s.ch)
	return nil
}
