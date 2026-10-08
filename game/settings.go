package game

type settings struct {
	ScreenWidth  int
	ScreenHeight int
}

var Settings settings

func init() {
	Settings.ScreenWidth, Settings.ScreenHeight = 640, 480
}
