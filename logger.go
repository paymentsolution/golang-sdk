package paymentsdk

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Logger interface {
	//если true то будут логгироваться входнящие и исходящие данные в запросах
	Enabled() bool

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type logger struct {
	enabled bool
	log     *zap.Logger
	level   string
}

func NewLogger(enabled bool, level string) (Logger, error) {

	lvl, err := zap.ParseAtomicLevel(strings.ToLower(level))
	if err != nil {
		return nil, err
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	cfg := zap.Config{
		Level:             lvl,
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	defer l.Sync()

	return &logger{
		log:     l,
		level:   strings.ToLower(level),
		enabled: enabled,
	}, nil
}

func (l *logger) Enabled() bool {
	return l.enabled
}

func (l *logger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *logger) Warn(msg string) {
	l.log.Warn(msg)
}

func (l *logger) Error(msg string) {
	l.log.Error(msg)
}
