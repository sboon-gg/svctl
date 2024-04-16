package prbf2update_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sboon-gg/svctl/pkg/prbf2update"
	"github.com/stretchr/testify/require"
)

func TestCache_FetchFor(t *testing.T) {
	dir := t.TempDir()

	currentVersion := "1.7.4.4"
	requiredVersion := "1.7.4.5"

	cache := prbf2update.NewCache(dir)

	err := cache.FetchFor(currentVersion, requiredVersion)
	require.NoError(t, err)

	patchFileName := fmt.Sprintf("prbf2_%s_to_%s_server_patch.prpatch", currentVersion, requiredVersion)

	require.FileExists(t, filepath.Join(dir, patchFileName))
	require.FileExists(t, filepath.Join(os.TempDir(), patchFileName))

	err = cache.FetchFor(currentVersion, requiredVersion)
	require.NoError(t, err)
}
