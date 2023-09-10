package mocks

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
)

// WriterInterface mocks Kafka writer methods.
type WriterInterface interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close()
}

// ReaderInterface mocks Kafka reader methods.
type ReaderInterface interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close()
}

// MockWriter mocks the kafka.Writer implementation.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder records the mock invocations for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	return &MockWriter{
		ctrl:     ctrl,
		recorder: &MockWriterMockRecorder{mock: &MockWriter{ctrl: ctrl}},
	}
}

func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

func (m *MockWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessages", ctx, msgs)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockWriter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// MockReader mocks the kafka.Reader implementation.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder records the mock invocations for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

func NewMockReader(ctrl *gomock.Controller) *MockReader {
	return &MockReader{
		ctrl:     ctrl,
		recorder: &MockReaderMockRecorder{mock: &MockReader{ctrl: ctrl}},
	}
}

func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

func (m *MockReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage", ctx)
	ret0, _ := ret[0].(kafka.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockReader) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}
