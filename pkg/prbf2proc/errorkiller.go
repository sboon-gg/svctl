package prbf2proc

import (
	"context"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var killer = newErrorKiller()

type errorKiller struct {
	watchedProcesses []*PRBF2Process
	ticker           *time.Ticker
	cancel           context.CancelFunc
}

func newErrorKiller() *errorKiller {
	return &errorKiller{
		watchedProcesses: make([]*PRBF2Process, 0),
		ticker:           &time.Ticker{},
	}
}

func (ek *errorKiller) Watch(proc *PRBF2Process) {
	if len(ek.watchedProcesses) == 0 && runtime.GOOS == "windows" {
		ek.start()
	}

	ek.watchedProcesses = append(ek.watchedProcesses, proc)
}

func (ek *errorKiller) Unwatch(proc *PRBF2Process) {
	for i, p := range ek.watchedProcesses {
		if p == proc {
			ek.watchedProcesses = append(ek.watchedProcesses[:i], ek.watchedProcesses[i+1:]...)
			break
		}
	}

	if len(ek.watchedProcesses) == 0 {
		ek.stop()
	}
}

func (ek *errorKiller) start() {
	ek.ticker.Reset(time.Millisecond * 500)
	ctx, cancel := context.WithCancel(context.Background())
	ek.cancel = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				ek.ticker.Stop()
				return
			case <-ek.ticker.C:
				ek.checkErrors()
			}
		}
	}()
}

func (ek *errorKiller) stop() {
	if ek.cancel != nil {
		ek.cancel()
	}
}

func (ek *errorKiller) checkErrors() {
	erroredPIDs, err := findErroredPIDs()
	if err != nil {
		return
	}

	for _, proc := range ek.watchedProcesses {
		pid := proc.Pid()
		if _, ok := erroredPIDs[pid]; ok {
			_ = proc.Stop()
		}
	}
}

func findErroredPIDs() (map[int]struct{}, error) {
	var filters = [4]string{
		"WINDOWTITLE eq BF2 Memory Error",
		"WINDOWTITLE eq BF2 Error",
		"WINDOWTITLE eq Microsoft Visual C++ Runtime Library",
		"STATUS eq NOT RESPONDING",
	}

	erroredPIDs := make(map[int]struct{})

	for _, f := range filters {
		procs, err := runTaskList(&taskListOpts{
			Filter: f,
		})
		if err != nil {
			return nil, err
		}

		for _, p := range procs {
			erroredPIDs[p.PID] = struct{}{}
		}
	}

	return erroredPIDs, nil
}

type taskListProc struct {
	Name        string `csv:"Image Name"`
	PID         int    `csv:"PID"`
	SessionName string `csv:"Session Name"`
	SessionID   int    `csv:"Session Name"`
	Memory      string `csv:"Mem Usage"`
}

type taskListOpts struct {
	Filter string
}

func runTaskList(opts *taskListOpts) ([]taskListProc, error) {
	var procs []taskListProc
	args := []string{
		"/FO", "csv",
	}

	if opts == nil {
		opts = &taskListOpts{}
	}

	if opts.Filter != "" {
		args = append(args, "/FI", opts.Filter)
	}

	cmd := exec.Command("tasklist", args...)

	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return procs, err
	}

	if strings.Contains(out.String(), "PID") {
		err := gocsv.Unmarshal(strings.NewReader(out.String()), &procs)
		if err != nil {
			return procs, err
		}
	}

	return procs, nil
}
