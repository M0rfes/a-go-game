package assets

import (
	"embed"
	"io/fs"
)

//go:embed hero-sheet.jpg
//go:embed hero
var files embed.FS

var Images = map[string][]byte{}

func init() {
	err := fs.WalkDir(files, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		data, err := files.ReadFile(path)
		if err != nil {
			return err
		}
		Images[path] = data
		return nil
	})
	if err != nil {
		panic(err)
	}
}
