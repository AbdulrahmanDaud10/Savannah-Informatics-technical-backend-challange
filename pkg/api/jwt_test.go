package api

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		userid uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateToken(tt.args.userid); got != tt.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
