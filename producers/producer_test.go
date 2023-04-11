//nolint:goerr113
package producers_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producers"
	"github.com/stretchr/testify/assert"
)

func TestProducer(t *testing.T) {
	t.Parallel()
	type args struct {
		topic   string
		server  string
		message string
	}

	tests := []struct {
		testName string
		args     args
		wantErr  error
	}{
		{
			testName: "Test with right configs",
			args: args{
				topic: "manu", server: "localhost:9092", message: "Hey",
			},
			wantErr: nil,
		},
		{
			testName: "Test with wrong configs",
			args: args{
				topic: "", server: "", message: "",
			},
			wantErr: errors.New("an error occurred producing messages Local: Invalid argument or configuration"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			err := producers.Producer(tc.args.topic, tc.args.server, tc.args.message)
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
