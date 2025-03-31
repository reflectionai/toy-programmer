package main

import (
	"context"
	"dagger/workspace/internal/dagger"
)

// A workspace that can edit files and run 'go build'
type Workspace struct {
	// The workspace's container state.
	// +internal-use-only
	Container *dagger.Container
}

// New creates a new Workspace with the given container.
// The container should be configured with the appropriate runtime environment
// and workdir set to "/app".
func New(container *dagger.Container) Workspace {
	return Workspace{
		Container: container,
	}
}

// Read a file
func (w *Workspace) Read(ctx context.Context, path string) (string, error) {
	return w.Container.File(path).Contents(ctx)
}

// Write a file
func (w Workspace) Write(path, content string) Workspace {
	w.Container = w.Container.WithNewFile(path, content)
	return w
}

// Build the code at the current directory in the workspace
func (w *Workspace) Build(ctx context.Context) error {
	_, err := w.Container.WithExec([]string{"go", "build", "./..."}).Stderr(ctx)
	return err
}
