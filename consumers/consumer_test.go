//nolint:goerr113
package consumers_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	"github.com/stretchr/testify/assert"
)

func TestConsumer(t *testing.T) {
	t.Parallel()
	type args struct {
		topic  string
		server string
		from   string
	}

	tests := []struct {
		testName string
		args     args
		wantErr  error
	}{
		{
			testName: "Test with right configs",
			args: args{
				topic: "manu", server: "localhost:9092", from: "latest",
			},
			wantErr: nil,
		},
		{
			testName: "Test with wrong auto offset topic",
			args: args{
				topic: "manu", server: "", from: "piwoi",
			},
			wantErr: errors.New("error creating consumer Invalid value \"piwoi\" for configuration property \"auto.offset.reset\""), //nolint:lll
		},
		{
			testName: "Test without any args provided",
			args: args{
				topic: "manu", server: "", from: "",
			},
			wantErr: errors.New("error creating consumer Configuration property \"auto.offset.reset\" cannot be set to empty value"), //nolint:lll
		},
		{
			testName: "Test without any args provided",
			args: args{
				topic: "", server: "", from: "latest",
			},
			wantErr: errors.New("error subscribing to topics Local: Invalid argument or configuration"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			err := consumers.Consumer(tc.args.topic, tc.args.server, tc.args.from)

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
