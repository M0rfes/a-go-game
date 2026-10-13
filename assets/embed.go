package assets

import (
	"embed"
	"io/fs"

	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed sprites
var files embed.FS

//go:embed default.map
var DefaultMapData []byte

var Textures = make(map[string]*ebiten.Image)

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

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return err
		}
		texture := ebiten.NewImageFromImage(img)

		Textures[path] = texture
		return nil
	})
	if err != nil {
		panic(err)
	}
}
