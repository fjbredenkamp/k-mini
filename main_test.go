package main

import "testing"

func TestVersionString(t *testing.T) {
	tests := []struct {
		name    string
		version string
		commit  string
		want    string
	}{
		{name: "defaults", version: "dev", commit: "none", want: "k-mini dev (none)"},
		{name: "release", version: "v0.1.0", commit: "abc1234", want: "k-mini v0.1.0 (abc1234)"},
		{name: "no commit", version: "v0.1.0", commit: "", want: "k-mini v0.1.0"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := versionString(tc.version, tc.commit)
			if got != tc.want {
				t.Errorf("versionString(%q, %q) = %q, want %q", tc.version, tc.commit, got, tc.want)
			}
		})
	}
}
