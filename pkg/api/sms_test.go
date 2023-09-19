package api

import "testing"

func TestSendBulkSMSViaAfricasTalking(t *testing.T) {
	type args struct {
		message      string
		phoneNumbers []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendBulkSMSViaAfricasTalking(tt.args.message, tt.args.phoneNumbers); (err != nil) != tt.wantErr {
				t.Errorf("SendBulkSMSViaAfricasTalking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
