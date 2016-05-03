package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/benhinchley/log"
)

// Exists checks whether the given path exists.
func Exists(p string) bool {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// StripRoot strips the root of the path off.
func StripRoot(root, p string) string {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	return strings.TrimPrefix(p, root)
}

// ListPath lists the contents of a given path.
func ListPath(p string) <-chan string {
	w := make(fileWalk)
	go func() {
		if err := filepath.Walk(p, w.Walk); err != nil {
			log.Fatal(err)
		}
		close(w)
	}()
	return w
}

type fileWalk chan string

func (f fileWalk) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}
	return nil
}
