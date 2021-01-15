package mycnf

import "testing"

func TestGenerateMyCnf(t *testing.T) {
	type args struct {
		Ip      string
		PortNum string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1", args: args{Ip: "192.168.171.1", PortNum: "3306"}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateMyCnf(tt.args.Ip, tt.args.PortNum); (err != nil) != tt.wantErr {
				t.Errorf("GenerateMyCnf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
