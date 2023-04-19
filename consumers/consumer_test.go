//nolint:goerr113
package consumers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers/mock"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

func Test_Consume(t *testing.T) {
	t.Parallel()

	type args struct {
		topic string
		from  string
	}

	type fields struct {
		provider *mock.MockProvider
	}

	tests := []struct {
		testName string
		prepare  func(t *testing.T, f *fields)
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SubscribeTopics([]string{"test-topic"}).Return(nil)

				f.provider.EXPECT().ReadMessage(time.Millisecond*100).Return(models.Message{}, nil)
			},
			args: args{
				topic: "test-topic", from: "latest",
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error/subscribe",
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SubscribeTopics(gomock.Any()).Return(assert.AnError)
			},
			args: args{
				topic: "test-topic", from: "latest",
			},
			wantErr: assert.Error,
		},
		{
			testName: "error/read",
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SubscribeTopics(gomock.Any()).Return(nil)

				f.provider.EXPECT().ReadMessage(gomock.Any()).Return(models.Message{}, assert.AnError)

				f.provider.EXPECT().Close().Return(nil)
			},
			args: args{
				topic: "test-topic", from: "latest",
			},
			wantErr: assert.Error,
		},
		{
			testName: "error/read then close error",
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SubscribeTopics(gomock.Any()).Return(nil)

				f.provider.EXPECT().ReadMessage(gomock.Any()).Return(models.Message{}, assert.AnError)

				f.provider.EXPECT().Close().Return(assert.AnError)
			},
			args: args{
				topic: "test-topic", from: "latest",
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			f := fields{
				provider: mock.NewMockProvider(ctrl),
			}

			if tc.prepare != nil {
				tc.prepare(t, &f)
			}

			consumer := consumers.NewConsumer(f.provider)
			consumer.StartConsumer()

			err := consumer.Consume(tc.args.topic, tc.args.from)

			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			// if tc.wantErr != nil {
			// 	assert.EqualError(t, err, tc.wantErr.Error(), tc.testName)
			// 	return
			// }

			if !tc.wantErr(t, err, fmt.Sprintf("Consume(%v, %v)", tc.args.topic, &tc.args.from)) {
				return
			}
		})
	}
}
