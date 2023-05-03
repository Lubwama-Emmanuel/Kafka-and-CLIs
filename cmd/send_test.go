package cmd_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producer/mocks"
)

func TestSend(t *testing.T) {
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
			testName: "error",
			args:     args{"manu", "localhost:9092", "latest", "group1"},
			prepare: func(t *testing.T, f *fields) {
				f.provider.EXPECT().SetUp(gomock.Any()).Return(nil)

				f.provider.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(assert.AnError).AnyTimes()

				f.provider.EXPECT().DeliveryReport().Return(assert.AnError).AnyTimes()
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

			cmdInstance, _, provider := createCMD(t)
			f := fields{
				provider: provider,
			}

			tc.prepare(t, &f)
			send := &cobra.Command{Use: "send", RunE: cmdInstance.SendCmdRun}
			cmdInstance.SendInit()
			err := cmdInstance.SendCmdRun(send, tc.args)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}
		})
	}
}
