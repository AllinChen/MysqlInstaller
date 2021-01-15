package myssh

import "testing"

func TestInstall(t *testing.T) {
	type args struct {
		mycfgfile string
		ip        string
		port      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{"./AutoMysql.cfg", "192.168.171.149", "3320"}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Install(tt.args.mycfgfile, tt.args.ip, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
