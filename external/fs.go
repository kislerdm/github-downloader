/* Copyright (c) 2021 Dmitry Kisler <dkisler.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

package external

import (
	"io"
	"os"
	"path/filepath"
)

// fWrite writes file on disk.
func fWrite(b io.ReadCloser, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, b)
	return err
}

// mkDirsIfParentNotExists mkdir -p given the full path to file
func mkDirsIfParentNotExists(path string) error {
	d := filepath.Dir(path)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		return os.MkdirAll(d, 0755)
	}
	return nil
}

// Write creates the directory if not exists and stores the file there.
func Write(data io.ReadCloser, path string) error {
	if err := mkDirsIfParentNotExists(path); err != nil {
		return err
	}
	return fWrite(data, path)
}
