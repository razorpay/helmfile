package helmexec

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logWriterGenerator struct {
	log        *zap.SugaredLogger
	level      zapcore.Level
	skipPrefix bool
}

func (g logWriterGenerator) Writer(prefix string) *logWriter {
	if g.skipPrefix {
		prefix = ""
	}
	return &logWriter{
		log:    g.log,
		prefix: prefix,
		level:  g.level,
	}
}

type logWriter struct {
	log    *zap.SugaredLogger
	prefix string
	level  zapcore.Level
}

func (w *logWriter) Write(p []byte) (int, error) {
	msg := fmt.Sprintf("%s%s", w.prefix, strings.TrimSpace(string(p)))
	if ce := w.log.Desugar().Check(w.level, msg); ce != nil {
		ce.Write()
	}
	return len(p), nil
}
