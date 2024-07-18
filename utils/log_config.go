package utils

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type LogConfig struct {
	// Stdout redirect logs to stdout
	Stdout bool
	// WriteFS write logs to files
	WriteFS bool
	// Syncer
	Syncer io.Writer
	// Encoder
	Encoder zapcore.TimeEncoder

	Level zapcore.Level
}

type LogConfigBuilder struct {
	config *LogConfig
}

func NewConcreteLogConfigBuilder() *LogConfigBuilder {
	return &LogConfigBuilder{config: &LogConfig{
		Stdout:  true,
		WriteFS: false,
		Syncer:  os.Stdout,
		Encoder: primitiveArrayEncoder,
		Level:   zapcore.InfoLevel,
	}}
}

func (b *LogConfigBuilder) WithTimeEncoder(enc zapcore.TimeEncoder) *LogConfigBuilder {
	b.config.Encoder = enc
	return b
}

func (b *LogConfigBuilder) WithStdout() *LogConfigBuilder {
	b.config.Stdout = true
	return b
}

func (b *LogConfigBuilder) WithFile() *LogConfigBuilder {
	b.config.WriteFS = true
	b.config.Stdout = false
	return b
}

func (b *LogConfigBuilder) WithLevel(lvl zapcore.Level) *LogConfigBuilder {
	b.config.Level = lvl
	return b
}

func (b *LogConfigBuilder) Build() *LogConfig {
	return b.config
}

var _ = DefaultSyncer

var DefaultSyncer = &lumberjack.Logger{
	Compress:  true,
	MaxSize:   20 * (1 << 20),
	Filename:  "MICRO_CMDB.INFO",
	MaxAge:    30,
	LocalTime: true,
}
