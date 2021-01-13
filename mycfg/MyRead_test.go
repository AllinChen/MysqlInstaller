package mycfg

import (
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		fileName string
		oc       string
		ed       string
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{"E:\\go\\src\\github.com\\AllinChen\\AutoMysql\\AutoMysql.cfg", "=", ";"}}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Read(tt.args.fileName, tt.args.oc, tt.args.ed); got == nil {
				t.Errorf("Read() = %v, want ", got)
			}
		})
	}
}
