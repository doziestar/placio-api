package kafka

// import (
// 	"context"
// 	"placio-api/mocks"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/segmentio/kafka-go"
// )

// func TestStart(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockReader := mocks.NewMockReader(ctrl)

// 	kc := &KafkaConsumer{
// 		reader: mockReader,
// 	}

// 	message := []byte("testMessage")

// 	// Set an expectation that ReadMessage will be called and return the mock message.
// 	mockReader.EXPECT().ReadMessage(gomock.Any()).Return(kafka.Message{
// 		Value: message,
// 	}, nil).Times(1)  // this means we expect only one message to be read

// 	errCount := 0
// 	callback := func(msg []byte) error {
// 		if string(msg) != string(message) {
// 			t.Fatalf("Expected message %v but got %v", string(message), string(msg))
// 		}
// 		errCount++
// 		return nil
// 	}

// 	// This will break after one message due to the .Times(1) setting above
// 	kc.Start(callback)

// 	if errCount != 1 {
// 		t.Fatalf("Callback was called %d times, expected 1", errCount)
// 	}
// }