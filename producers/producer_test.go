package producers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producers/pmocks"
)

func TestProduceMessages(t *testing.T) {
	t.Parallel()

	type args struct {
		topic   string
		message string
	}

	type fields struct {
		provider *pmocks.MockProvider
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
				topic:   "test_topic",
				message: "message 1",
			},
			prepare: func(t *testing.T, f *fields) {
				f.provider.E
			},
		},
	}
}
