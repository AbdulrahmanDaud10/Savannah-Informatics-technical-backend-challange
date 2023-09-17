package repository

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestSetUpDatabaseConnection(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUpDatabaseConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUpDatabaseConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
