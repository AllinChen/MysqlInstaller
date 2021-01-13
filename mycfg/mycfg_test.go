package mycfg

import (
	"fmt"
	"testing"
)

func TestStringRead(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{"E:\\go\\src\\github.com\\AutoMysql\\代码规范.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringRead(tt.args.fileName); got == "" {
				t.Errorf("StringRead() = %v", got)
			}
		})
	}
}

func TestGetCfg(t *testing.T) {
	type args struct {
		cfgfile string
	}
	tests := []struct {
		name string
		args args
		//wantCfg *Cfg
	}{
		// TODO: Add test cases.
		{"t1", args{"E:\\go\\src\\github.com\\AutoMysql\\AutoMysql.cfg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCfg := GetCfg(tt.args.cfgfile); 0 < 1 {
				fmt.Println(gotCfg)
			}
		})
	}
}
