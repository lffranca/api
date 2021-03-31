package environment

import (
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	os.Setenv("API_VERSION", "/conductor/v1/api")
	os.Setenv("API_PORT", "8000")
	os.Setenv("SQLITE_PATH", "./app.db")

	tests := []struct {
		name    string
		want    *Environment
		wantErr bool
	}{
		{
			name: "environment TestGet 01",
			want: &Environment{
				APIVersion: "/conductor/v1/api",
				APIPort:    "8000",
				SQLitePath: "./app.db",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
