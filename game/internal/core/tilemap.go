package core

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	defaultTileScale = 4.0
	viewportPadTiles = 1
	tileWorldLayer   = 0
)

type cellKey struct {
	col int
	row int
}

type TileMap struct {
	*Node2D
	grid     [][]int
	textures []*ebiten.Image
	camera   *Camera
	scale    float64
	live     map[cellKey]*Sprite
}

func NewTileMap(grid [][]int, textures []*ebiten.Image, camera *Camera) *TileMap {
	return &TileMap{
		Node2D:   NewNode2D("tile_map"),
		grid:     grid,
		textures: textures,
		camera:   camera,
		scale:    defaultTileScale,
		live:     make(map[cellKey]*Sprite),
	}
}

func (t *TileMap) Update() {
	if t.camera == nil {
		return
	}

	tileSize := float64(Settings.TileSize) * t.scale
	cam := t.camera.GetPosition()
	startCol, startRow, cols, rows := visibleTileRange(
		cam.X(),
		cam.Y(),
		float64(t.camera.width),
		float64(t.camera.height),
		tileSize,
		viewportPadTiles,
	)

	needed := make(map[cellKey]struct{}, cols*rows)
	for row := startRow; row < startRow+rows; row++ {
		for col := startCol; col < startCol+cols; col++ {
			key := cellKey{col: col, row: row}
			needed[key] = struct{}{}
			if _, ok := t.live[key]; !ok {
				t.spawn(col, row, tileSize)
			}
		}
	}

	for key, sprite := range t.live {
		if _, ok := needed[key]; ok {
			continue
		}
		t.DetachChild(sprite)
		delete(t.live, key)
	}
}

func (t *TileMap) spawn(col, row int, tileSize float64) {
	tex := t.textureAt(col, row)
	if tex == nil {
		return
	}

	sprite := NewSprite(fmt.Sprintf("tile_%d_%d", row, col), tex, tileWorldLayer, false)
	sprite.SetPosition(float64(col)*tileSize, float64(row)*tileSize)
	sprite.SetScale(t.scale, t.scale)
	t.AddChild(sprite)
	t.live[cellKey{col: col, row: row}] = sprite
}

func (t *TileMap) textureAt(col, row int) *ebiten.Image {
	if len(t.textures) == 0 {
		return nil
	}

	id := wrapTileID(t.grid, col, row)
	if id < 0 || id >= len(t.textures) || t.textures[id] == nil {
		return t.textures[0]
	}
	return t.textures[id]
}

func wrapTileID(grid [][]int, col, row int) int {
	height := len(grid)
	if height == 0 || len(grid[0]) == 0 {
		return 0
	}

	width := len(grid[0])
	patternRow := ((row % height) + height) % height
	patternCol := ((col % width) + width) % width
	return grid[patternRow][patternCol]
}

func visibleTileRange(camX, camY, camW, camH, tileSize float64, pad int) (startCol, startRow, cols, rows int) {
	startCol = int(math.Floor(camX/tileSize)) - pad
	startRow = int(math.Floor(camY/tileSize)) - pad
	cols = int(math.Ceil(camW/tileSize)) + pad*2
	rows = int(math.Ceil(camH/tileSize)) + pad*2
	return
}
