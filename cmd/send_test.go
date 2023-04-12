package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestSendCmd(t *testing.T) {
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
	send := &cobra.Command{Use: "send", RunE: cmd.SendCmdRun}
	cmd.SendInit()
	for _, tc := range tt {
		out, err := executeSend(t, send, tc.args...)

		if err != nil {
			t.Fatalf("an error %v", err)
		}

		assert.Equal(t, out, tc.out)
	}

}

func executeSend(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}
