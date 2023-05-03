package cmd_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer/mocks"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producer"
	pmocks "github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producer/mocks"
)

func TestRecieve(t *testing.T) {
	t.Parallel()
	type args []string

	type fields struct {
		provider *mocks.MockProvider
	}

	tt := []struct {
		testName string
		args     args
		prepare  func(t *testing.T, f *fields)
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			args:     args{"manu", "localhost:9092", "latest", "group1"},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(gomock.Any()).Return(nil)

				f.provider.EXPECT().Subscribe(gomock.Any()).Return(assert.AnError)
			},
			wantErr: assert.Error,
		},
		{
			testName: "error",
			args:     args{"manu", "localhost:9092", "latest", "group1"},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(gomock.Any()).Return(assert.AnError)
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			cmdInstance, provider, _ := createCMD(t)
			f := fields{
				provider: provider,
			}

			tc.prepare(t, &f)
			receive := &cobra.Command{Use: "receive"}
			cmdInstance.ReceiveInit()
			err := cmdInstance.ReceiveCmdRun(receive, tc.args)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}
		})
	}
}

func createCMD(t *testing.T) (*cmd.CMD, *mocks.MockProvider, *pmocks.MockProvider) {
	ctrl := gomock.NewController(t)

	consumerProvider := mocks.NewMockProvider(ctrl)
	producerProvider := pmocks.NewMockProvider(ctrl)

	c := consumer.NewConsumer(consumerProvider)
	c.StartConsumer()

	p := producer.NewProducer(producerProvider)

	cmdMock := cmd.NewCMD(*c, *p)

	return cmdMock, consumerProvider, producerProvider
}
