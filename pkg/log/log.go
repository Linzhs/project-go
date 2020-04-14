package log

import (
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func Factory() Logger {
	return nil
}

func RegisterGinZap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		bt := time.Now()

		var (
			path  = context.Request.URL.Path
			query = context.Request.URL.RawQuery
		)

		context.Next()

		endTime := time.Now()
		latency := endTime.Sub(bt)
		if utc {
			endTime = endTime.UTC()
		}

		if len(context.Errors) > 0 {
			for _, e := range context.Errors.Errors() {
				logger.Error(e)
			}
			return
		}

		logger.Info(path,
			zap.Int("status", context.Writer.Status()),
			zap.String("method", context.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", context.ClientIP()),
			zap.String("user-agent", context.Request.UserAgent()),
			zap.String("time", endTime.Format(timeFormat)),
			zap.Duration("latency", latency))
	}
}

func RecoveryWithZap(logger *zap.Logger, printStack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequestCopy, _ := httputil.DumpRequest(context.Request, false)
				if printStack {
					logger.Error("[recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequestCopy)),
						zap.String("stack", string(debug.Stack())))
					context.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				logger.Error("[recovery from panic]",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequestCopy)))
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		context.Next()
	}
}
