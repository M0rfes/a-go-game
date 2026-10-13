package entites

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/M0rfes/a-go-game/assets"
	"github.com/M0rfes/a-go-game/game/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

const tileWorldLayer = 0

func tileTextures() []*ebiten.Image {
	return []*ebiten.Image{
		assets.Textures["sprites/tiny-dungeon/tile_0048.png"],
		assets.Textures["sprites/tiny-dungeon/tile_0049.png"],
		assets.Textures["sprites/tiny-dungeon/tile_0014.png"],
		assets.Textures["sprites/tiny-dungeon/tile_0036.png"],
		assets.Textures["sprites/tiny-dungeon/tile_0037.png"],
		assets.Textures["sprites/tiny-dungeon/tile_0038.png"],
	}
}

func ParseMap(data []byte) ([][]int, int, int) {
	var grid [][]int
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		row := make([]int, 0, len(fields))
		for _, f := range fields {
			val, err := strconv.Atoi(f)
			if err != nil {
				continue
			}
			row = append(row, val)
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}
	h := len(grid)
	w := 0
	if h > 0 {
		w = len(grid[0])
	}
	return grid, w, h
}

func NewTileMap(world *core.World) *core.TileMap {
	grid, _, _ := ParseMap(assets.DefaultMapData)
	tilemap := core.NewTileMap(grid, tileTextures(), world.Camera())
	world.AddNodeToLayer(tilemap, tileWorldLayer)
	return tilemap
}
