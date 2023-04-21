package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestSendCmd(t *testing.T) {
	t.Parallel()

	type args []string
	tt := []struct {
		testName string
		args     args
		wantErr  error
	}{
		{
			testName: "test case with no args",
			args:     nil,
			wantErr:  nil,
		},
		{
			testName: "test case with args",
			args:     args{"manu", "localhost:9092", "group1", "hey"},
			wantErr:  nil,
		},
	}
	send := &cobra.Command{Use: "send", RunE: cmd.SendCmdRun}
	cmd.SendInit()

	for _, tc := range tt {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			err := Executor(t, send, tc.args...)
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

	if err := c.Execute(); err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}

	return nil
}
