package core

type settings struct {
	ScreenWidth  int
	ScreenHeight int
	TileSize     int
}

var Settings settings

func init() {
	Settings.ScreenWidth = 640
	Settings.ScreenHeight = 480
	Settings.TileSize = 16
}
