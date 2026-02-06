package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Key        = "Logger"
	Request_id = "Req_id"
)

type Logger struct {
	l *zap.Logger
}

func GetContextWithNewLogger(ctx context.Context, dev bool) (context.Context, error) {
	var config zap.Config
	if dev {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("NewCtxWithLogger: %w", err)
	}
	ctx = context.WithValue(ctx, Key, &Logger{logger})
	return ctx, nil
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(Request_id) != nil {
		fields = append(fields, zap.String(Request_id, ctx.Value(Request_id).(string)))
	}
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(Request_id) != nil {
		fields = append(fields, zap.String(Request_id, ctx.Value(Request_id).(string)))
	}
	l.l.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(Request_id) != nil {
		fields = append(fields, zap.String(Request_id, ctx.Value(Request_id).(string)))
	}
	l.l.Error(msg, fields...)
}
