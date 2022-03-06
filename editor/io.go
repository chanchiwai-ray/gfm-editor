package editor

import (
	"os"
	"path/filepath"
)

type File struct {
	Path    string
	Name    string
	Content []byte
}

func (f *File) Save() error {
	filename := f.Name
	return os.WriteFile(filename, f.Content, 0600)
}

func ReadMD(path string) (*File, error) {
	// Expecting user to throw a markdown file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file := &File{Path: path, Name: filepath.Base(path), Content: data}
	return file, nil
}
