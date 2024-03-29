package consumer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer/mocks"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

var configs = config.ProviderConfig{
	Server: "server",
	Group:  "group",
	From:   "from",
}

func TestConsumeMessages(t *testing.T) {
	t.Parallel()

	type args struct {
		topic string
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
				topic: "test_topic",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(configs).Return(nil)

				f.provider.EXPECT().Subscribe("test_topic").Return(nil)

				f.provider.EXPECT().ReadMessage(time.Millisecond*100).Return(models.Message{}, nil)
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error/subscribe",
			args: args{
				topic: "test_topic",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(configs).Return(nil)

				f.provider.EXPECT().Subscribe(gomock.Any()).Return(assert.AnError)
			},
			wantErr: assert.Error,
		},
		{
			testName: "error/read",
			args: args{
				topic: "test_topic",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(configs).Return(nil)

				f.provider.EXPECT().Subscribe(gomock.Any()).Return(nil)

				f.provider.EXPECT().ReadMessage(gomock.Any()).Return(models.Message{}, assert.AnError)

				f.provider.EXPECT().Close().Return(nil)
			},
			wantErr: assert.Error,
		},
		{
			testName: "error/read",
			args: args{
				topic: "test_topic",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(gomock.Any()).Return(assert.AnError)

				f.provider.EXPECT().Subscribe(gomock.Any()).Return(assert.AnError)
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
				provider: mocks.NewMockProvider(ctrl),
			}

			if tc.prepare != nil {
				tc.prepare(t, &f)
			}

			consumer := consumer.NewConsumer(f.provider)
			consumer.StartConsumer()
			setupErr := consumer.SetUpProvider(configs)
			if setupErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, setupErr))
				return
			}

			err := consumer.ConsumeMessages(tc.args.topic)
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
