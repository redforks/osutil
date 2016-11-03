// Package osutil document
package osutil

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// Copy src file content to dstFile, overwrite dstFile if exist,
// dstFile can be directory.
func Copy(dstFile, srcFile string) (err error) {
	var (
		sf, df *os.File
		isdir  bool
	)

	isdir, err = isDir(dstFile)
	if err != nil {
		return
	}

	if isdir {
		dstFile = filepath.Join(dstFile, filepath.Base(srcFile))
	}

	if sf, err = os.Open(srcFile); err != nil {
		return err
	}
	defer func() {
		if err := sf.Close(); err != nil {
			log.Printf("[%s] Close src file failed: %s", "osutil/copy", err)
		}
	}()

	if df, err = os.Create(dstFile); err != nil {
		return err
	}

	_, err = io.Copy(df, sf)
	if e := df.Close(); e != nil {
		if err == nil {
			err = e
		}
	}
	return
}

// isDir returns true if path p is a directory
func isDir(p string) (bool, error) {
	info, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}
