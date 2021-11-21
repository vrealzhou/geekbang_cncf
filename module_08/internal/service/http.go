package service

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/vrealzhou/geekbang_cncf/module_08/log"
)

type HTTPService struct {
	stopped bool
	srv     *http.Server
	log     *log.Logger
}

func NewHTTPService(port int, handler http.Handler, logger *log.Logger) *HTTPService {
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler,
	}
	s := &HTTPService{
		srv: srv,
		log: logger,
	}
	return s
}

func (s *HTTPService) Start(ctx context.Context) (err error) {
	innerCtx, cancel := context.WithCancel(ctx)
	go func() {
		s.log.Infof("Start HTTP server: %s", s.srv.Addr)
		err = errors.Wrap(s.srv.ListenAndServe(), "HTTP service")
		cancel()
	}()
	<-innerCtx.Done()
	return
}

func (s *HTTPService) Shutdown(ctx context.Context) error {
	if s.stopped {
		return nil
	}
	s.log.Info("Shutdown HTTP server")
	s.stopped = true
	ctxWithDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(shutdownTimeout(ctx)))
	defer cancel()
	return s.srv.Shutdown(ctxWithDeadline)
}
