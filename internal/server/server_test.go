package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		setup    func() error
		cleanup  func()
		expected bool
	}{
		{
			name:     "file exists",
			filename: "yarn.lock",
			setup: func() error {
				return os.WriteFile("yarn.lock", []byte("test"), 0644)
			},
			cleanup: func() {
				os.Remove("yarn.lock")
			},
			expected: true,
		},
		{
			name:     "file does not exist",
			filename: "nonexistent_file.txt",
			setup: func() error {
				return nil
			},
			cleanup: func() {
			},
			expected: false,
		},
		{
			name:     "directory exists",
			filename: "public",
			setup: func() error {
				return os.Mkdir("public", 0755)
			},
			cleanup: func() {
				os.Remove("public")
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.setup(); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
			defer tt.cleanup()

			result := fileExists(tt.filename)
			if result != tt.expected {
				t.Errorf("fileExists(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestCacheControlHeaders(t *testing.T) {
	handler := noCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	getRecorder := httptest.NewRecorder()
	getRequest := httptest.NewRequest("GET", "/test.js", nil)
	handler.ServeHTTP(getRecorder, getRequest)

	if getRecorder.Header().Get("Cache-Control") == "" {
		t.Errorf("GET request: Cache-Control header not set")
	}
}
