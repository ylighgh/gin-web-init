package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type swift struct {
	core *zap.SugaredLogger
}

func NewLogger(cb *LogConfigBuilder) Logger {
	cnf := cb.Build()
	syncWriter := zapcore.AddSync(cnf.Syncer)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder.EncodeTime = cnf.Encoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(cnf.Level))
	l := zap.New(zapcore.NewTee(core), zap.AddCaller(), zap.AddCallerSkip(1))
	return &swift{core: l.Sugar()}
}

func primitiveArrayEncoder(c time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(c.Format("2006-01-02 15:04:05"))
}

func (l *swift) Fatal(args ...interface{}) {
	l.core.Fatal(args)
}

func (l *swift) Fatalf(format string, args ...interface{}) {
	l.core.Fatalf(format, args)
}

func (l *swift) Fatalln(args ...interface{}) {
	l.core.Fatalln(args)
}

func (l *swift) Panic(args ...interface{}) {
	l.core.Panic(args)
}

func (l *swift) Panicf(format string, args ...interface{}) {
	l.core.Panicf(format, args)
}

func (l *swift) Panicln(args ...interface{}) {
	l.core.Panicln(args)
}

func (l *swift) Print(args ...interface{}) {
	l.core.Info(args)
}

func (l *swift) Printf(format string, args ...interface{}) {
	l.core.Infof(format, args)
}

func (l *swift) Println(args ...interface{}) {
	l.core.Infoln(args)
}

func (l *swift) Debug(args ...interface{}) {
	l.core.Debug(args)
}

func (l *swift) Debugf(format string, args ...interface{}) {
	l.core.Debugf(format, args)
}

func (l *swift) Debugln(args ...interface{}) {
	l.core.Debugln(args)
}

func (l *swift) Error(args ...interface{}) {
	l.core.Error(args)
}

func (l *swift) Errorf(format string, args ...interface{}) {
	l.core.Errorf(format, args)
}

func (l *swift) Errorln(args ...interface{}) {
	l.core.Errorln(args)
}

func (l *swift) Info(args ...interface{}) {
	l.core.Info(args)
}

func (l *swift) Infof(format string, args ...interface{}) {
	l.core.Infof(format, args)
}

func (l *swift) Infoln(args ...interface{}) {
	l.core.Infoln(args)
}

func (l *swift) Warn(args ...interface{}) {
	l.core.Warn(args)
}

func (l *swift) Warnf(format string, args ...interface{}) {
	l.core.Warnf(format, args)
}

func (l *swift) Warnln(args ...interface{}) {
	l.core.Warnln(args)
}
