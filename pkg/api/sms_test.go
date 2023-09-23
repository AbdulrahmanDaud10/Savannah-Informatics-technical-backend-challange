package api

import "testing"

func TestSendAfricastalkingBulkSMS(t *testing.T) {
	type args struct {
		africasTalkingSettings AfricasTalkingSettings
		message                string
		recipients             []string
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
			if err := SendAfricastalkingBulkSMS(tt.args.africasTalkingSettings, tt.args.message, tt.args.recipients); (err != nil) != tt.wantErr {
				t.Errorf("SendAfricastalkingBulkSMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
