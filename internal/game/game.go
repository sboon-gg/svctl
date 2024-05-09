package game

import (
	"context"
	"io"
	"os"
	"time"
)

type Game interface {
	Configs
	Process
	Updater
	Artifacts
}

type Configs interface {
	// Writes configuration file, storing original file.
	// Stores origing only first time, then it will be ignored.
	// Registers file if not already registered.
	WriteConfig(path string, data []byte) error
}

type Process interface {
	// Start game server process.
	Start() (*os.Process, error)
}

type Updater interface {
	// Update game server.
	Update(ctx context.Context, output io.Writer) error
}

type ArtifactType int

const (
	ArtifactUnknown ArtifactType = iota
	ArtifactChatlog
	ArtifactPRDemo
	ArtifactBF2Demo
	ArtifactJSONSummary
)

type Artifact interface {
	// File path
	Location() string

	// Time infered from file name or content metadata.
	Start() time.Time

	// Artifact type
	Type() ArtifactType
}

type Artifacts interface {
	// Returns all chatlogs.
	// Gets chatlogs directory and file pattern from configuration.
	Chatlogs() []Artifact

	// Returns all PR demos.
	// Gets PR demos directory and file pattern from configuration.
	PRDemos() []Artifact

	// Returns all BF2 demos.
	// Gets BF2 demos directory and file pattern from configuration.
	BF2Demos() []Artifact

	// Returns all JSON summaries.
	// Gets JSON summaries directory and file pattern from configuration.
	JSONSummaries() []Artifact
}
