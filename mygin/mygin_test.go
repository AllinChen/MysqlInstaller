package mygin

import "testing"

func TestReipport(t *testing.T) {
	type args struct {
		ipport string
	}
	tests := []struct {
		name     string
		args     args
		wantIp   string
		wantPort string
		wantErr  bool
	}{
		{name: "test1", args: args{"0.0.0.0:3306"}, wantIp: "0.0.0.0", wantPort: "3306", wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp, gotPort, err := Reipport(tt.args.ipport)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reipport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIp != tt.wantIp {
				t.Errorf("Reipport() gotIp = %v, want %v", gotIp, tt.wantIp)
			}
			if gotPort != tt.wantPort {
				t.Errorf("Reipport() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
		})
	}
}
