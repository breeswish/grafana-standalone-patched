package assets

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/macaron.v1"
)

type TplFileSystem struct {
	files []macaron.TemplateFile
}

func NewTemplateFileSystem(filterBaseDirectory string) macaron.TemplateFileSystem {
	tplFs := TplFileSystem{
		files: make([]macaron.TemplateFile, 0, 10),
	}
	exts := []string{".tmpl", ".html"}
	afero.Walk(Fs, filterBaseDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		r, err := filepath.Rel(filterBaseDirectory, path)
		ext := macaron.GetExt(path)
		for _, extension := range exts {
			if ext != extension {
				continue
			}
			data, err := afero.ReadFile(Fs, path)
			if err != nil {
				return nil
			}
			name := filepath.ToSlash((r[0 : len(r)-len(ext)]))
			tplFs.files = append(tplFs.files, macaron.NewTplFile(name, data, ext))
		}
		return nil
	})

	return tplFs
}

func (fs TplFileSystem) ListFiles() []macaron.TemplateFile {
	return fs.files
}

func (fs TplFileSystem) Get(name string) (io.Reader, error) {
	for i := range fs.files {
		if fs.files[i].Name()+fs.files[i].Ext() == name {
			return bytes.NewReader(fs.files[i].Data()), nil
		}
	}
	return nil, fmt.Errorf("file '%s' not found", name)
}
