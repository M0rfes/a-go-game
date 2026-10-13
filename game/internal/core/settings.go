package core

type settings struct {
	ScreenWidth  int
	ScreenHeight int
	TileSize     int
}

var Settings settings
var FloorTileIndex = [][2]int{{0, 0}, {0, 1}, {0, 2}}

func init() {
	Settings.ScreenWidth = 640
	Settings.ScreenHeight = 480
	Settings.TileSize = 16
}
