package kafka

// import (
// 	"context"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/segmentio/kafka-go"
// 	"placio-api/mocks" // Ensure this is the correct import path for your mocks.
// )

// func TestPublishMessage(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockWriter := mocks.NewMockWriter(ctrl)

// 	p := &Producer{
// 		writer: mockWriter,
// 	}

// 	key := []byte("testKey")
// 	value := []byte("testValue")

// 	// Set expectation on mockWriter to be called with specific parameters
// 	mockWriter.EXPECT().WriteMessages(
// 		gomock.Any(),
// 		kafka.Message{
// 			Key:   key,
// 			Value: value,
// 		},
// 	).Return(nil)

// 	// Execute the method
// 	err := p.PublishMessage(context.Background(), key, value)

// 	// Check for unexpected errors
// 	if err != nil {
// 		t.Fatalf("PublishMessage failed: %v", err)
// 	}
// }
