package fsm

//go:generate go run golang.org/x/tools/cmd/stringer -type=StateT,Action -linecomment -output=types_string.go

type StateT int

const (
	StateTStopped    StateT = iota // Stopped
	StateTRunning                  // Running
	StateTRestarting               // Restarting
	StateTErrored                  // Errored
)

type Action int

const (
	ActionStop    Action = iota // Stop
	ActionStart                 // Start
	ActionAdopt                 // Adopt
	ActionRestart               // Restart
)
