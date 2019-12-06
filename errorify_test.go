package errorify

import (
	"errors"
	"testing"
)

func TestCreateFile(t *testing.T) {
	type args struct {
		err    error
		params []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"nil error without params", args{err: nil}},
		{"error without params", args{err: errors.New("error")}},
		{"error 1 param", args{err: errors.New("error"), params: []string{"file.txt"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateFile(tt.args.err, tt.args.params...)
		})
	}
}

func TestSimple(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"nil error without params", args{err: nil}},
		{"error", args{err: errors.New("error")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Simple(tt.args.err)
		})
	}
}
