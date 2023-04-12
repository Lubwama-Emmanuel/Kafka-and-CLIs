package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestReceiveCmd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		testName string
		args     []string
		wantErr  error
	}{
		{
			testName: "test with no args",
			args:     nil,
			wantErr:  nil,
		},
		{
			testName: "test with args",
			args:     []string{"manu", "localhost:9092", "latest", "group1"},
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

func Executor(t *testing.T, c *cobra.Command, args ...string) error {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()

	return fmt.Errorf("an error %w", err)
}
