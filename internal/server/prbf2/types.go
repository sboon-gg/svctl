package prbf2

//go:generate go run golang.org/x/tools/cmd/stringer -type=State,Action -linecomment -output=types_string.go

type State int

const (
	StateStopped       State = iota // Stopped
	StateStopping                   // Stopping
	StateStarting                   // Starting
	StateAdopting                   // Adopting
	StateRunning                    // Running
	StateRestarting                 // Restarting
	StateExited                     // Exited
	StateErrored                    // Errored
	StateCleaningError              // CleaningError
)

type Action int

const (
	ActionStop    Action = iota // Stop
	ActionStart                 // Start
	ActionAdopt                 // Adopt
	ActionRestart               // Restart
)
