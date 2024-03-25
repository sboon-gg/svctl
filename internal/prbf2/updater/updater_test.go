package updater_test

import (
	"testing"

	"github.com/sboon-gg/svctl/internal/prbf2/updater"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindInsertions(t *testing.T) {
	files := map[string][2][]byte{
		"file.txt": {
			[]byte(`
test: 1
test2: 2
test3: 3
			`),
			[]byte(`
test: 1
test2: changed
test3: 3
test4: 4
			`),
		},
	}

	patches, err := updater.PreparePatches(files)
	require.NoError(t, err)
	println(patches[0])

	require.Len(t, patches, 1)
	assert.Equal(t, `diff --git a/file.txt b/file.txt
index 481c3e9dddee4569cc5aef2452e640a101d47222..d52524feb8b2506120b8374840bc305bb48e7827 100755
--- a/file.txt
+++ b/file.txt
@@ -1,5 +1,6 @@
 
 test: 1
-test2: 2
+test2: changed
 test3: 3
+test4: 4
 			
\ No newline at end of file
`, patches[0])
}
