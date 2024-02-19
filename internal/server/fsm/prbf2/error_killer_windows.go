//go:build windows

package prbf2

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var filters = []string{
	"WINDOWTITLE eq BF2 Memory Error",
	"WINDOWTITLE eq BF2 Error",
	"WINDOWTITLE eq Microsoft Visual C++ Runtime Library",
	"STATUS eq NOT RESPONDING",
}

func findErrors(pid int) (bool, error) {
	for _, f := range filters {
		procs, err := runTaskList(&taskListOpts{
			Filter: f,
		})
		if err != nil {
			return false, err
		}
		for _, p := range procs {
			if p.PID == pid {
				return true, nil
			}
		}
	}

	return false, nil
}

func startErrorKiller(ctx context.Context, proc *os.Process) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if proc != nil {
					found, err := findErrors(proc.Pid)
					if err != nil {
						log.Printf("Couldn't find errors in ErrorKiller: %s", err)
					}
					if found {
						stopProcess(proc)
						return
					}
				}

				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
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
