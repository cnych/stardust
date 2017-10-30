package logx

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"reflect"
	"testing"
)

func TestDebugMode(t *testing.T) {
	WithField("st", fmt.Sprintf("%d->%d", 100, 200)).Debug("Changed")
}

func TestWithFields(t *testing.T) {
	type args struct {
		fields Fields
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
