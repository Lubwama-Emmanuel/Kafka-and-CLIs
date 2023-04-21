package producers_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producers"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producers/mocks"
)

func TestProduceMessages(t *testing.T) {
	t.Parallel()

	type args struct {
		topic   string
		message string
	}

	type fields struct {
		provider *mocks.MockProvider
	}

	tests := []struct {
		testName string
		args     args
		prepare  func(t *testing.T, f *fields)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			args: args{
				topic:   "test_topic",
				message: "message 1",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().Produce("test_topic", "message 1").Return(nil)

				f.provider.EXPECT().KafkaMessage().Return(nil).AnyTimes()

				f.provider.EXPECT().Flush(15 * 1000)
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error/produce-kafkamessage",
			args: args{
				topic:   "test_topic",
				message: "message 1",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(assert.AnError)

				f.provider.EXPECT().KafkaMessage().Return(assert.AnError).AnyTimes()
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			f := fields{
				provider: mocks.NewMockProvider(ctrl),
			}

			if tc.prepare != nil {
				tc.prepare(t, &f)
			}

			producer := producers.NewProducer(f.provider)

			err := producer.ProduceMessages(tc.args.topic, tc.args.message)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			if !tc.wantErr(t, err, fmt.Sprintf("Consume(%v)", tc.args.topic)) {
				return
			}
		})
	}
}
