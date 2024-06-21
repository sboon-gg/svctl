//go:build windows

package game

import (
	"context"
	"io"
)

func (s *Server) update(ctx context.Context, outW, inW, errW io.Writer) error {
	return nil
}
