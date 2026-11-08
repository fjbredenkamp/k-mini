// Package procfs reads process info from Linux /proc filesystem.
package procfs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrMalformed is returned when a stat line can't be parsed.
var ErrMalformed = errors.New("malformed stat")

// Stat holds the fields of the /proc/<pid>/stat that K-mini uses.
// Numbers in the comments are field positions from the proc_pid_stat(5).
type Stat struct {
	PID        int    // (1)
	Comm       string // (2) executable name, without brackets.
	State      string // (3) R running, S sleeping, D disk wait, Z zombie etc.
	PPID       int    // (4) parent PID.
	UTime      uint64 // (14) user-mode CPU time, in clock ticks.
	STime      uint64 // (15) kernel-mode CPU time, in clock ticks.
	NumThreads int    // (20)
	StartTime  uint64 // (22) start time after boot, in clock ticks.
	VSize      uint64 // (23) virtual memory size, in bytes.
	RSS        uint64 // (24) resident set size, in pages.
}

// ParseStat parses the contents of a /proc/<pid>/stat file.
func ParseStat(data []byte) (Stat, error) {
	s := strings.TrimSpace(string(data))

	lparen := strings.IndexByte(s, '(')
	rparen := strings.LastIndexByte(s, ')')
	if lparen < 0 || rparen < lparen {
		return Stat{}, fmt.Errorf("%w: no parentheses", ErrMalformed)
	}
	pid, err := strconv.Atoi(strings.TrimSpace(s[:lparen]))
	if err != nil {
		return Stat{}, fmt.Errorf("%w: pid: %w", ErrMalformed, err)
	}
	st := Stat{PID: pid, Comm: s[lparen+1 : rparen]}

	fields := strings.Fields(s[rparen+1:])
	if len(fields) < 22 {
		return Stat{}, fmt.Errorf("%w: want at least 22 fields after comm, got %d", ErrMalformed, len(fields))
	}

	st.State = fields[0]

	st.PPID, err = strconv.Atoi(fields[1])
	if err != nil {
		return Stat{}, fmt.Errorf("%w: ppid: %w", ErrMalformed, err)
	}

	st.UTime, err = strconv.ParseUint(fields[11], 10, 64)
	if err != nil {
		return Stat{}, fmt.Errorf("%w: utime: %w", ErrMalformed, err)
	}

	st.STime, err = strconv.ParseUint(fields[12], 10, 64)
	if err != nil {
		return Stat{}, fmt.Errorf("%w: stime: %w", ErrMalformed, err)
	}

	st.NumThreads, err = strconv.Atoi(fields[17])
	if err != nil {
		return Stat{}, fmt.Errorf("%w: numthreads: %w", ErrMalformed, err)
	}

	st.StartTime, err = strconv.ParseUint(fields[19], 10, 64)
	if err != nil {
		return Stat{}, fmt.Errorf("%w: starttime: %w", ErrMalformed, err)
	}

	st.VSize, err = strconv.ParseUint(fields[20], 10, 64)
	if err != nil {
		return Stat{}, fmt.Errorf("%w: vsize: %w", ErrMalformed, err)
	}

	st.RSS, err = strconv.ParseUint(fields[21], 10, 64)
	if err != nil {
		return Stat{}, fmt.Errorf("%w: rss: %w", ErrMalformed, err)
	}

	return st, nil
}
