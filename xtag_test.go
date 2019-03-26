package xtag

import "testing"

func TestFindLatest(t *testing.T) {
	tests := []struct {
		xtag string
		tags []string
		want string
	}{
		{"v1.x", []string{"v1.0", "v1.5", "v1.4", "v2.0"}, "v1.5"},
		{"1.x", []string{"1.0", "1.5", "1.4", "2.0"}, "1.5"},
		{"v2.x", []string{"v1.0", "v2.0.0", "v3.0"}, "v2.0.0"},
		{"v1.2.x", []string{"v1.0", "v2.0.0", "v1.2.1", "v1.2.3"}, "v1.2.3"},
		{"vx", []string{"v1.0", "v2.0.0", "v1.2.1"}, "v2.0.0"},
		{"x", []string{"1.0", "2.0.0", "1.2.1"}, "2.0.0"},
	}
	for _, tt := range tests {
		got, err := FindLatest(tt.xtag, tt.tags)
		if err != nil {
			t.Errorf("%s: got error: %v", tt.xtag, err)
			continue
		}
		if got != tt.want {
			t.Errorf("FindLatest(%q, %s) = %q, want %q", tt.xtag, tt.tags, got, tt.want)
		}
	}
}
