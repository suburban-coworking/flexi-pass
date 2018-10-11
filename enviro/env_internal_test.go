package enviro

import (
	"testing"
)

func Test_env_setupLog(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		wantErr bool
	}{
		{"Valid log level", "info", false},
		{"Invalid log level", "invalid", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e env

			if err := e.setupLog(tt.level); (err != nil) != tt.wantErr {
				t.Errorf("env.setupLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_env_setupDir(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		wantErr bool
	}{
		{"Valid dir", "/tmp/su/flexi-1/", false},
		{"Invalid dir", "/root/su/flexi-2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e env

			if err := e.setupLog("info"); err != nil {
				t.Errorf("env.setupDir() - log setup failed: error = %v", err)
			}

			if err := e.setupDir(tt.dir); (err != nil) != tt.wantErr {
				t.Errorf("env.setupDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
