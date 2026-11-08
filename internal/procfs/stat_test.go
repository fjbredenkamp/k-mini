package procfs

import (
	"errors"
	"strings"
	"testing"
)

// A trimmed real stat lin: fields 1-24, enough for Stat
const bashStat = "1234 (bash) S 1000 1234 1234 34816 1300 4194304 2400 15000 0 3 12 5 30 10 20 0 1 0 98765 9375744 1350"

func wantBash(comm string) Stat {
	return Stat{
		PID: 1234, Comm: comm, State: "S", PPID: 1000,
		UTime: 12, STime: 5, NumThreads: 1, StartTime: 98765,
		VSize: 9375744, RSS: 1350,
	}
}

func TestParseStat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Stat
		wantErr bool
	}{
		{name: "simple", input: bashStat, want: wantBash("bash")},
		{name: "trailing newline", input: bashStat + "\n", want: wantBash("bash")},
		{name: "extra fields ignored", input: bashStat + " 18446744073709551615 0 0 0", want: wantBash("bash")},
		{name: "comm with spaces", input: strings.Replace(bashStat, "(bash)", "(tmux: server)", 1), want: wantBash("tmux: server")},
		{name: "comm with parens", input: strings.Replace(bashStat, "(bash)", "(my (odd) proc)", 1), want: wantBash("my (odd) proc")},

		{name: "empty", input: "", wantErr: true},
		{name: "no parens", input: strings.Replace(bashStat, "(bash)", "bash", 1), wantErr: true},
		{name: "too few fields", input: "1234 (bash) S 1000", wantErr: true},
		{name: "bad pid", input: strings.Replace(bashStat, "1234 (", "abc (", 1), wantErr: true},
		{name: "bad number", input: strings.Replace(bashStat, "9375744", "lots", 1), wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseStat([]byte(tc.input))
			if tc.wantErr {
				if !errors.Is(err, ErrMalformed) {
					t.Fatalf("ParseStat() error = %v, want ErrMalformed", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseStat() unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("ParseStat() =\n  %+v\nwant\n  %+v", got, tc.want)
			}
		})
	}
}
