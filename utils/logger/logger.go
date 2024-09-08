package logger

import (
	"log"
	"log/slog"
	"os"

	"github.com/go-resty/resty/v2"
)

var logger *slog.Logger

func InitLogger() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger = slog.New(logHandler)
	slog.SetDefault(logger)
}

func Info(msg string, args ...any) {
	if len(args) == 0 {
		logger.Error(msg)
		return
	}

	logger.Info(msg, genSlogAttrs(args)...)
}

func Error(err error, args ...any) {
	if len(args) == 0 {
		logger.Error(err.Error())
		return
	}

	logger.Error(err.Error(), genSlogAttrs(args)...)
}

func ApiError(resp *resty.Response) {
	// Extract request data
	requestData := map[string]interface{}{
		"method":  resp.Request.Method,
		"url":     resp.Request.URL,
		"headers": resp.Request.Header,
		"body":    resp.Request.Body,
	}

	// Log request and response
	logger.Error(resp.String(),
		slog.String("status", resp.Status()),
		slog.Any("request", requestData),
		slog.Any("response", resp.Error()),
	)
}

func genSlogAttrs(args []any) (retval []any) {
	argsLen := len(args)
	slogAttrLen := argsLen / 2
	const orphanKey = "NO-KEY"

	var orphan any
	if argsLen%2 != 0 {
		log.Println("uneven number of arguments, placing last provided value under the key '" + orphanKey + "'")
		orphan = args[argsLen-1]
		args = args[:argsLen-1]
		slogAttrLen += 1
	}

	attrs := make([]slog.Attr, slogAttrLen)
	for i := 0; i < argsLen-1; i += 2 {
		attrs = append(attrs, slog.Any(args[i].(string), args[i+1]))
	}

	if orphan != nil {
		attrs = append(attrs, slog.Any(orphanKey, orphan))
	}

	for i := range attrs {
		retval = append(retval, attrs[i].Key, attrs[i].Value)
	}

	return
}
