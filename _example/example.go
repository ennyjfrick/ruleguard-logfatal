package main

import (
	"errors"
	"fmt"
	"io"
	stdlog "log"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func alwaysErr() error {
	return errors.New("")
}

func main() {
	zapLogger := zap.Must(zap.NewProduction())
	logrusLogger := logrus.New()
	zerologLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	privateMethodLogger := NewCustomStdOutLogger()

	// noLogFatal rules group examples

	// calling Fatal methods

	// stdlib logger examples
	if err := alwaysErr(); err != nil {
		stdlog.Fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		stdlog.Fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		stdlog.Fatalln("encountered fatal error", err)
	}

	// Zap examples
	if err := alwaysErr(); err != nil {
		zap.L().Fatal("encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Fatal("encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zap.S().Fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Fatalln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Fatalw("encountered fatal error", "error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Fatalln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Fatalw("encountered fatal error", "error", err)
	}

	// logrus examples
	if err := alwaysErr(); err != nil {
		logrus.Fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.Fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.Fatalln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.FatalFn(func() []interface{} {
			return []interface{}{"error", err}
		})
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Fatalln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.FatalFn(func() []interface{} {
			return []interface{}{"error", err}
		})
	}

	// zerolog examples
	if err := alwaysErr(); err != nil {
		zlog.Fatal().Err(err).Msg("encountered fatal error")
	}
	if err := alwaysErr(); err != nil {
		zerologLogger.Fatal().Err(err).Msg("encountered fatal error")
	}

	// logger with private methods examples
	if err := alwaysErr(); err != nil {
		privateMethodLogger.fatal("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.fatalf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.fatalln("encountered fatal error", err)
	}

	// logging at the Fatal level rather than calling a method

	// Zap examples
	if err := alwaysErr(); err != nil {
		zap.L().Log(zap.FatalLevel, "encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Log(zap.FatalLevel, "encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zap.S().Log(zap.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logf(zap.FatalLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logln(zap.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logw(zap.FatalLevel, "encountered fatal error", "error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Log(zap.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logf(zap.FatalLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logln(zap.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logw(zap.FatalLevel, "encountered fatal error", "error", err)
	}

	// logrus examples
	if err := alwaysErr(); err != nil {
		logrusLogger.Log(logrus.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Logf(logrus.FatalLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Logln(logrus.FatalLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.LogFn(logrus.FatalLevel, func() []interface{} {
			return []interface{}{"error", err}
		})
	}

	// zerolog examples
	if err := alwaysErr(); err != nil {
		zlog.WithLevel(zerolog.FatalLevel).Err(err).Msg("encountered fatal error")
	}
	if err := alwaysErr(); err != nil {
		zerologLogger.WithLevel(zerolog.FatalLevel).Err(err).Msg("encountered fatal error")
	}

	// private method logger examples
	if err := alwaysErr(); err != nil {
		privateMethodLogger.log(CustomLogLevelFatal, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.logf(CustomLogLevelFatal, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.logln(CustomLogLevelFatal, "encountered fatal error", err)
	}

	// explicit exits are ok
	if err := alwaysErr(); err != nil {
		zap.L().Error("encountered fatal error", zap.Error(err))
		os.Exit(1)
	}

	// noLogPanic rules group examples

	// Zap examples
	if err := alwaysErr(); err != nil {
		zap.L().Panic("encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Panic("encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zap.S().Panic("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Panicf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Panicln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Panicw("encountered fatal error", "error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Panic("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Panicf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Panicln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Panicw("encountered fatal error", "error", err)
	}

	// logrus examples
	if err := alwaysErr(); err != nil {
		logrus.Panic("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.Panicf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.Panicln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrus.PanicFn(func() []interface{} {
			return []interface{}{"error", err}
		})
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Panic("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Panicf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Panicln("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.PanicFn(func() []interface{} {
			return []interface{}{"error", err}
		})
	}

	// zerolog examples
	if err := alwaysErr(); err != nil {
		zlog.Fatal().Err(err).Msg("encountered fatal error")
	}
	if err := alwaysErr(); err != nil {
		zerologLogger.Fatal().Err(err).Msg("encountered fatal error")
	}

	// logger with private methods examples
	if err := alwaysErr(); err != nil {
		privateMethodLogger.panic("encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.panicf("encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.panicln("encountered fatal error", err)
	}

	// logging at the Panic level rather than calling a method\

	// Zap examples
	if err := alwaysErr(); err != nil {
		zap.L().Log(zap.PanicLevel, "encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Log(zap.PanicLevel, "encountered fatal error", zap.Error(err))
	}
	if err := alwaysErr(); err != nil {
		zap.S().Log(zap.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logf(zap.PanicLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logln(zap.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zap.S().Logw(zap.PanicLevel, "encountered fatal error", "error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Log(zap.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logf(zap.PanicLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logln(zap.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		zapLogger.Sugar().Logw(zap.PanicLevel, "encountered fatal error", "error", err)
	}

	// logrus examples
	if err := alwaysErr(); err != nil {
		logrusLogger.Log(logrus.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Logf(logrus.PanicLevel, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.Logln(logrus.PanicLevel, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		logrusLogger.LogFn(logrus.PanicLevel, func() []interface{} {
			return []interface{}{"error", err}
		})
	}

	// zerolog examples
	if err := alwaysErr(); err != nil {
		zlog.WithLevel(zerolog.PanicLevel).Err(err).Msg("encountered fatal error")
	}
	if err := alwaysErr(); err != nil {
		zerologLogger.WithLevel(zerolog.PanicLevel).Err(err).Msg("encountered fatal error")
	}

	// private method logger examples
	if err := alwaysErr(); err != nil {
		privateMethodLogger.log(CustomLogLevelPanic, "encountered fatal error", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.logf(CustomLogLevelPanic, "encountered fatal error: %v", err)
	}
	if err := alwaysErr(); err != nil {
		privateMethodLogger.logln(CustomLogLevelPanic, "encountered fatal error", err)
	}

	// explicit panics are ok
	if err := alwaysErr(); err != nil {
		zap.L().Error("encountered fatal error", zap.Error(err))
		panic(err)
	}

}

type CustomLogLevel int

const (
	CustomLogLevelFatal CustomLogLevel = iota
	CustomLogLevelPanic
)

type CustomLogger interface {
	fatal(args ...interface{})
	fatalf(format string, args ...interface{})
	fatalln(args ...interface{})
	panic(args ...interface{})
	panicf(format string, args ...interface{})
	panicln(args ...interface{})
	log(lvl CustomLogLevel, args ...interface{})
	logf(lvl CustomLogLevel, format string, args ...interface{})
	logln(lvl CustomLogLevel, args ...interface{})
}

type customLogger struct {
	out io.Writer
}

func (c *customLogger) log(lvl CustomLogLevel, args ...interface{}) {
	switch lvl {
	case CustomLogLevelFatal:
		c.fatal(args...)
	case CustomLogLevelPanic:
		c.panic(args...)
	}
}

func (c *customLogger) logf(lvl CustomLogLevel, format string, args ...interface{}) {
	switch lvl {
	case CustomLogLevelFatal:
		c.fatalf(format, args...)
	case CustomLogLevelPanic:
		c.panicf(format, args...)
	}
}

func (c *customLogger) logln(lvl CustomLogLevel, args ...interface{}) {
	switch lvl {
	case CustomLogLevelFatal:
		c.fatalln(args...)
	case CustomLogLevelPanic:
		c.panicln(args...)
	}
}

func (c *customLogger) fatal(args ...interface{}) {
	_, _ = fmt.Fprint(c.out, args...)
	os.Exit(1)
}

func (c *customLogger) fatalf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(c.out, format, args...)
	os.Exit(1)
}

func (c *customLogger) fatalln(args ...interface{}) {
	_, _ = fmt.Fprintln(c.out, args...)
	os.Exit(1)
}

func (c *customLogger) panic(args ...interface{}) {
	panic(fmt.Sprint(args...))
}

func (c *customLogger) panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func (c *customLogger) panicln(args ...interface{}) {
	panic(fmt.Sprintln(args...))
}

func NewCustomStdOutLogger() CustomLogger {
	return &customLogger{out: os.Stdout}
}
