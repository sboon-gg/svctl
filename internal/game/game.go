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
	Chatlogs() []Artifact
	PRDemos() []Artifact
	BF2Demos() []Artifact
}
