package myflag

import (
	"reflect"
	"testing"
)

func TestMakeFlag(t *testing.T) {
	ip := "0.0.0.0"
	port := "3306"
	tests := []struct {
		name     string
		wantIp   *string
		wantPort *string
		wantErr  bool
	}{
		{name: "test1", wantIp: &ip, wantPort: &port, wantErr: false}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp, gotPort, err := MakeFlag()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIp, tt.wantIp) {
				t.Errorf("MakeFlag() gotIp = %v, want %v", gotIp, tt.wantIp)
			}
			if !reflect.DeepEqual(gotPort, tt.wantPort) {
				t.Errorf("MakeFlag() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
		})
	}
}
