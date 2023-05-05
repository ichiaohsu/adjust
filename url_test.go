package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name     string
		url      []string
		expected []string
		wantErr  bool
	}{
		{
			name:     "should return error when get invalid address",
			url:      []string{"http://:foo"},
			expected: []string{},
			wantErr:  true,
		},
		{
			name:     "should let url untouched when all in needed",
			url:      []string{"https://www.adjust.com", "http://www.google.com"},
			expected: []string{"https://www.adjust.com", "http://www.google.com"},
			wantErr:  false,
		},
		{
			name:     "should add default http scheme",
			url:      []string{"www.google.com"},
			expected: []string{"http://www.google.com"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := parseURL(tt.url)
			if (err != nil) && !tt.wantErr {
				t.Fatalf("test should return error but err is nil")
			}

			count := 0
			for _, p := range parsed {
				for _, te := range tt.expected {
					if p == te {
						count++
						break
					}
				}
			}
			if count != len(tt.expected) {
				t.Errorf("corrected urls doesn't match expected. Want: %v but got %v", parsed, tt.expected)
			}
		})
	}
}
