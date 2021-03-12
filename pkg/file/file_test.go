package file_test

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/Ryoh827/sorting-file-script/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestDirwalk(t *testing.T) {
	paths := file.Dirwalk("./test")
	sort.Strings(paths)
	actualPaths := []string{"test/a", "test/b", "test/c.txt", "test/d.log", "test/e.txt"}
	sort.Strings(actualPaths)

	assert.Equal(t, actualPaths, paths, "pass")
}

func TestGetExtList(t *testing.T) {
	actualExt := []string{"txt", "log", "other"}
	sort.Strings(actualExt)
	paths := []string{"test/a", "test/b", "test/c.txt", "test/d.log", "test/e.txt"}
	ext := file.GetExtList(paths)
	sort.Strings(ext)
	assert.Equal(t, actualExt, ext, "pass")
}

func TestSort(t *testing.T) {
	actualResult := []string{"test/output/other/a", "test/output/other/b", "test/output/txt/c.txt", "test/output/log/d.log", "test/output/txt/e.txt"}

	paths := []string{"test/a", "test/b", "test/c.txt", "test/d.log", "test/e.txt"}
	rootDir := "./test"
	extList := []string{"txt", "log", "other"}
	file.Sort(paths, rootDir, extList)

	for _, file := range actualResult {
		_, err := os.Stat(file)
		assert.Nil(t, err)
		filename := filepath.Base(file)
		if err := os.Rename(file, filepath.Join("test", filename)); err != nil {
			fmt.Println(err)
		}
	}
}
