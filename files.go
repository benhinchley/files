package files

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
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
			log.Printf("error: %s", err)
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

// GetHomeDir returns the path of the users home directory.
func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

// Move moves a file from src to dst, making any directories needed
func Move(src, dst string) error {
	dstDir := filepath.Dir(dst)
	if !Exists(dstDir) {
		if err := os.MkdirAll(dstDir, 0777); err != nil {
			return err
		}
	}
	if err := os.Rename(src, dst); err != nil {
		return err
	}
	return nil
}
