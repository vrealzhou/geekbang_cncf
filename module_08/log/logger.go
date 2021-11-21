package log

import "go.uber.org/zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	return &Logger{logger: sugar}
}

func (l *Logger) Shutdown() {
	l.logger.Sync()
}

func (l *Logger) DebugWithFields(msg string, fields ...interface{}) {
	l.logger.Debugw(msg, fields...)
}

func (l *Logger) Debugf(msg string, params ...interface{}) {
	l.logger.Debugf(msg, params...)
}

func (l *Logger) Debug(msg interface{}) {
	l.logger.Debug(msg)
}

func (l *Logger) InfoWithFields(msg string, fields ...interface{}) {
	l.logger.Infow(msg, fields...)
}

func (l *Logger) Infof(msg string, params ...interface{}) {
	l.logger.Infof(msg, params...)
}

func (l *Logger) Info(msg interface{}) {
	l.logger.Info(msg)
}

func (l *Logger) WarnWithFields(msg string, fields ...interface{}) {
	l.logger.Warnw(msg, fields...)
}

func (l *Logger) Warnf(msg string, params ...interface{}) {
	l.logger.Warnf(msg, params...)
}

func (l *Logger) Warn(msg interface{}) {
	l.logger.Warn(msg)
}
func (l *Logger) ErrorWithFields(msg string, fields ...interface{}) {
	l.logger.Errorw(msg, fields...)
}

func (l *Logger) Errorf(msg string, params ...interface{}) {
	l.logger.Errorf(msg, params...)
}

func (l *Logger) Error(msg interface{}) {
	l.logger.Error(msg)
}
func (l *Logger) FatalWithFields(msg string, fields ...interface{}) {
	l.logger.Fatalw(msg, fields...)
}

func (l *Logger) Fatalf(msg string, params ...interface{}) {
	l.logger.Fatalf(msg, params...)
}

func (l *Logger) Fatal(msg interface{}) {
	l.logger.Fatal(msg)
}
