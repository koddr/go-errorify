package errorify

import (
	"errors"
	"testing"
)

func TestCreateFile(t *testing.T) {
	type args struct {
		err  error
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{"no error", args{err: nil, name: "filename.go"}},
		{"error", args{err: errors.New("error"), name: "./folder/filename.go"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateFile(tt.args.err, tt.args.name)
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
		{"no error", args{err: nil}},
		{"error", args{err: errors.New("error")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Simple(tt.args.err)
		})
	}
}
