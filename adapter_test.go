package goazap_test

import (
	"bytes"
	"strings"

	"encoding/json"

	"github.com/PyYoshi/goa-logging-zap"
	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type testBuffer struct {
	bytes.Buffer
}

func (b *testBuffer) Sync() error {
	return nil
}

func (b *testBuffer) Lines() []string {
	output := strings.Split(b.String(), "\n")
	return output[:len(output)-1]
}

func (b *testBuffer) Stripped() string {
	return strings.TrimRight(b.String(), "\n")
}

var _ = Describe("goa", func() {
	var logger *zap.Logger
	var adapter goa.LogAdapter

	buf := &testBuffer{}
	errBuf := &testBuffer{}

	type msgJson struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Hoge  string `json:"hoge"`
	}

	BeforeEach(func() {
		zapEncConfig := zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		}
		zapEnc := zapcore.NewJSONEncoder(zapEncConfig)
		logger = zap.New(
			zapcore.NewCore(zapEnc, buf, zap.DebugLevel),
			zap.ErrorOutput(errBuf),
		)

		adapter = goazap.New(logger)
	})

	It("adapts info messages", func() {
		adapter.Info("msg", "hoge", "fuga")
		var v msgJson
		err := json.Unmarshal(buf.Bytes(), &v)
		if err != nil {
			panic(err)
		}
		vb, err := json.Marshal(&v)
		if err != nil {
			panic(err)
		}
		Ω(string(vb)).Should(Equal(`{"level":"info","msg":"msg","hoge":"fuga"}`))
		buf.Reset()
	})

	It("adapts error messages", func() {
		adapter.Error("msg", "hoge", "fuga")
		var v msgJson
		err := json.Unmarshal([]byte(buf.Stripped()), &v)
		if err != nil {
			panic(err)
		}
		vb, err := json.Marshal(&v)
		if err != nil {
			panic(err)
		}
		Ω(string(vb)).Should(Equal(`{"level":"error","msg":"msg","hoge":"fuga"}`))
		buf.Reset()
	})
})
