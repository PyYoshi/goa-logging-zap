package goazap_test

import (
	"bytes"
	"strings"

	"github.com/PyYoshi/goa-logging-zap"
	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/uber-go/zap"
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
	var logger zap.Logger
	var adapter goa.LogAdapter

	buf := &testBuffer{}
	// errBuf := &testBuffer{}

	BeforeEach(func() {
		logger = zap.New(
			zap.NewJSONEncoder(zap.NoTime()),
			zap.Output(buf),
			// zap.ErrorOutput(errBuf),
		)
		adapter = goazap.New(logger)
	})

	It("adapts info messages", func() {
		msg := "msg"
		adapter.Info(msg, "hoge", "fuga")
		Ω(buf.Stripped()).Should(Equal(`{"level":"info","msg":"msg","hoge":"fuga"}`))
	})
})
