package config

import (
	"os"
	"testing"
)

// TestSanitizeInput tests the SanitizeInput function with various valid and invalid inputs.
func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		// Valid inputs
		{"valid-input", "valid-input", false},
		{"another_valid.input", "another_valid.input", false},
		{"CAPS_and_numbers123", "CAPS_and_numbers123", false},
		{"mixed-123.UP_case", "mixed-123.UP_case", false},

		// Invalid inputs
		{"invalid/input", "", true},     // contains /
		{"input with spaces", "", true}, // contains spaces
		{"invalid@input", "", true},     // contains @
		{"<invalid>", "", true},         // contains < >
		{";ls", "", true},               // contains ;
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := SanitizeInput(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, err)
			}
			if result != tt.expected {
				t.Errorf("Expected result: %s, got: %s", tt.expected, result)
			}
		})
	}
}

// TestEnsureDirectory tests that EnsureDirectory creates a directory if it doesn't exist and handles errors correctly.
func TestEnsureDirectory(t *testing.T) {
	tempDir := os.TempDir() + "/test_ensure_directory"

	if _, err := os.Stat(tempDir); err == nil {
		os.RemoveAll(tempDir)
	}

	err := EnsureDirectory(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if _, err := os.Stat(tempDir); err != nil {
		t.Errorf("Expected directory to exist, but got error: %v", err)
	}

	defer os.RemoveAll(tempDir)
}
