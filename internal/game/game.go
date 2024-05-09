package game

import "time"

type Game interface {
	Configs
	Process
	Updater
	Artifacts
}

type Configs interface {
	// Regsiter configuration file path
	RegisterConfig(path string)

	// Writes configuration file, storing original file.
	// Stores origing only first time, then it will be ignored.
	// Registers file if not already registered.
	WriteConfig(path string, data []byte) error
}

type Status int

const (
	StatusUnknown Status = iota
	StatusRunning
	StatusStopped
	StatusUpdating
)

type Process interface {
	// Adopt running process.
	// NOTE: not sure if this is needed since we can just
	// store PID and use it to check if process is running.
	// The intent is to be able to recover from daemon crash.
	// Adopt(*os.Process) error

	// Start game server process.
	// Returns error if process is already running.
	Start() error

	// Stop game server process.
	// If process is not running, returns nil.
	Stop() error

	// Checks if game server process is running.
	// Only checks if process is running,
	// to check if server is healthy use Status().
	IsRunning() bool

	// Returns game server status.
	Status() (Status, error)
}

// TODO: consider returning more information about the update:
// new version, changed config files, etc.
type Updater interface {
	// Update game server.
	// Returns error update is in progress.
	Update() error
}

type ArtifactType int

const (
	ArtifactUnknown ArtifactType = iota
	ArtifactChatlog
	ArtifactPRDemo
	ArtifactBF2Demo
)

type Artifact interface {
	// File path
	Location() string

	// Time infered from file name or content metadata.
	Start() time.Time

	// Artifact type
	Type() ArtifactType

	// Whether the round the artifact belongs to is finished.
	Finished() bool
}

// TODO: improve this interface
type Artifacts interface {
	Chatlogs() []Artifact
	PRDemos() []Artifact
	BF2Demos() []Artifact
}
