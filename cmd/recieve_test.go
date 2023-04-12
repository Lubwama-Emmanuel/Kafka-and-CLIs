package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestReceiveCmd(t *testing.T) {
	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: nil,
			err:  assert.AnError,
			out:  "",
		},
		{
			args: []string{"manu", "localhost:9092", "latest", "group1"},
			err:  nil,
			out:  "",
		},
	}
	receive := &cobra.Command{Use: "receive", RunE: cmd.ReceiveCmdRun}
	cmd.ReceiveInit()
	for _, tc := range tt {
		out, err := execute(t, receive, tc.args...)

		if err != nil {
			t.Fatalf("an error %v", err)
		}

		assert.Equal(t, out, tc.out)
	}

}

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}
