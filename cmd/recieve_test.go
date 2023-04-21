package cmd_test

import (
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestReceiveCmd(t *testing.T) {
	t.Parallel()
	type args []string

	tests := []struct {
		testName string
		args     args
		wantErr  error
	}{
		{
			testName: "test with no args",
			args:     nil,
			wantErr:  nil,
		},
		{
			testName: "test with args",
			args:     args{"manu", "localhost:9092", "latest", "group1"},
			wantErr:  nil,
		},
	}
	receive := &cobra.Command{Use: "receive", RunE: cmd.ReceiveCmdRun}
	cmd.ReceiveInit()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			err := Executor(t, receive, tc.args...)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error(), tc.testName)

				return
			}
		})
	}
}
