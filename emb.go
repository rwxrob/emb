// Copyright 2022 emb Authors
// SPDX-License-Identifier: Apache-2.0

package emb

import (
	"embed"
	"io"
	"io/fs"
	"path/filepath"

	_fs "github.com/rwxrob/fs"
)

// set from init by importers, but also changeable at runtime
var (
	FS  embed.FS // assigned
	Top string   // path to top directory (ex: files)
)

func glob(pattern string) ([]string, error) {
	return fs.Glob(FS, filepath.Join(Top, pattern))
}

func RelPaths() []string {
	return _fs.RelPaths(FS, Top)
}

func Cat(globs ...string) ([]byte, error) {
	var buf []byte
	for _, gl := range globs {
		files, err := glob(gl)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			f, err := FS.Open(file)
			info, err := f.Stat()
			if err != nil {
				return nil, err
			}
			if info.IsDir() {
				continue
			}
			if err != nil {
				return nil, err
			}
			b, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			buf = append(buf, b...)
		}
	}
	return buf, nil
}
