/*
Package goazap contains an adapter that makes it possible to configure goa so it uses zap
as logger backend.
Usage:
    logger := zap.New(
		zap.NewJSONEncoder(),
	)
    // Initialize logger handler using zap package
    service.WithLogger(goazap.New(logger))
    // ... Proceed with configuring and starting the goa service

    // In handlers:
    goazap.Logger(ctx).Info("foo")
*/
package goazap

import (
	"fmt"
	"time"

	"github.com/goadesign/goa"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

type adapter struct {
	zap.Logger
}

// New wraps a zap logger into a goa logger adapter.
func New(logger zap.Logger) goa.LogAdapter {
	return &adapter{Logger: logger}
}

// Logger returns the zap logger stored in the given context if any, nil otherwise.
func Logger(ctx context.Context) zap.Logger {
	logger := goa.ContextLogger(ctx)
	if a, ok := logger.(*adapter); ok {
		return a.Logger
	}
	return nil
}

// Info logs informational messages using zap.
func (a *adapter) Info(msg string, data ...interface{}) {
	fields := data2fields(data)
	a.Logger.Info(msg, *fields...)
}

// Error logs error messages using zap.
func (a *adapter) Error(msg string, data ...interface{}) {
	fields := data2fields(data)
	a.Logger.Error(msg, *fields...)
}

// New creates a new logger given a context.
func (a *adapter) New(data ...interface{}) goa.LogAdapter {
	fields := data2fields(data)
	return &adapter{Logger: a.Logger.With(*fields...)}
}

func data2fields(keyvals []interface{}) *[]zap.Field {
	n := (len(keyvals) + 1) / 2
	fields := make([]zap.Field, n)

	fi := 0
	for i := 0; i < len(keyvals); i += 2 {
		if key, ok := keyvals[i].(string); ok {
			if i+1 < len(keyvals) {
				v := keyvals[i+1]
				if val, ok := v.([]byte); ok {
					fields[fi] = zap.Base64(key, val)
				} else if val, ok := v.(bool); ok {
					fields[fi] = zap.Bool(key, val)
				} else if val, ok := v.(float32); ok {
					fields[fi] = zap.Float64(key, float64(val))
				} else if val, ok := v.(float64); ok {
					fields[fi] = zap.Float64(key, val)
				} else if val, ok := v.(int); ok {
					fields[fi] = zap.Int(key, val)
				} else if val, ok := v.(int8); ok {
					fields[fi] = zap.Int64(key, int64(val))
				} else if val, ok := v.(int16); ok {
					fields[fi] = zap.Int64(key, int64(val))
				} else if val, ok := v.(int32); ok {
					fields[fi] = zap.Int64(key, int64(val))
				} else if val, ok := v.(int64); ok {
					fields[fi] = zap.Int64(key, val)
				} else if val, ok := v.(uint); ok {
					fields[fi] = zap.Uint(key, val)
				} else if val, ok := v.(uint8); ok {
					fields[fi] = zap.Uint64(key, uint64(val))
				} else if val, ok := v.(uint16); ok {
					fields[fi] = zap.Uint64(key, uint64(val))
				} else if val, ok := v.(uint32); ok {
					fields[fi] = zap.Uint64(key, uint64(val))
				} else if val, ok := v.(uint64); ok {
					fields[fi] = zap.Uint64(key, val)
				} else if val, ok := v.(string); ok {
					fields[fi] = zap.String(key, val)
				} else if val, ok := v.(fmt.Stringer); ok {
					fields[fi] = zap.Stringer(key, val)
				} else if val, ok := v.(time.Time); ok {
					fields[fi] = zap.Time(key, val)
				} else if val, ok := v.(error); ok {
					fields[fi] = zap.Error(val)
				} else if val, ok := v.(time.Duration); ok {
					fields[fi] = zap.Duration(key, val)
				} else if val, ok := v.(zap.LogMarshaler); ok {
					fields[fi] = zap.Marshaler(key, val)
				} else {
					fields[fi] = zap.Object(key, v)
				}
			}
		} else {
			fields[fi] = zap.Skip()
		}
		fi = fi + 1
	}
	return &fields
}
