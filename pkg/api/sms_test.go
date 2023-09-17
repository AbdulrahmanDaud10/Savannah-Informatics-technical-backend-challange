package api

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestAtClient_SendBulkSMS(t *testing.T) {
	type fields struct {
		apiKey     string
		endpoint   string
		httpClient *http.Client
		username   string
	}
	type args struct {
		ctx   context.Context
		input BulkSMSInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BulkSMSResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at := &AtClient{
				apiKey:     tt.fields.apiKey,
				endpoint:   tt.fields.endpoint,
				httpClient: tt.fields.httpClient,
				username:   tt.fields.username,
			}
			got, err := at.SendBulkSMS(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AtClient.SendBulkSMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtClient.SendBulkSMS() = %v, want %v", got, tt.want)
			}
		})
	}
}
