package main

import (
	"image"
	"image/color"
	"image/draw"
	"slices"
)

var pipeColor = color.RGBA{0, 255, 0, 255}
var pathColor = color.RGBA{0, 0, 255, 255}
var outsideColor = color.RGBA{255, 255, 255, 255}
var inclosedColor = color.RGBA{255, 0, 0, 255}
var startColor = color.RGBA{255, 255, 0, 255}

var mazePalette color.Palette = color.Palette{
	image.Transparent,
	image.Black,
	pipeColor,
	outsideColor,
	pathColor,
	inclosedColor,
	startColor,
}

func createSprite(p Pipe) image.Image {
	img := image.NewPaletted(image.Rect(0, 0, 4, 4), mazePalette)
	switch p {
	case 'S', '|', '-', 'L', 'J', '7', 'F':
		c := pipeColor
		if p == 'S' {
			c = startColor
		}
		for y := 1; y <= 2; y++ {
			for x := 1; x <= 2; x++ {
				img.Set(x, y, c)
			}
		}
		if slices.Contains([]Pipe{'S', '|', 'L', 'J'}, p) {
			img.Set(1, 0, c)
			img.Set(2, 0, c)
		}
		if slices.Contains([]Pipe{'S', '|', '7', 'F'}, p) {
			img.Set(1, 3, c)
			img.Set(2, 3, c)
		}
		if slices.Contains([]Pipe{'S', '-', 'J', '7'}, p) {
			img.Set(0, 1, c)
			img.Set(0, 2, c)
		}
		if slices.Contains([]Pipe{'S', '-', 'L', 'F'}, p) {
			img.Set(3, 1, c)
			img.Set(3, 2, c)
		}
	case 'O', 'I':
		c := outsideColor
		if p == 'I' {
			c = inclosedColor
		}
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
	}
	return img
}

func drawMaze(m *PipeMaze) *image.Paletted {
	step, scale := 1, 4
	if m.expanded {
		step, scale = 2, 2
	}
	img := image.NewPaletted(image.Rect(0, 0, m.size.x*scale, m.size.y*scale), mazePalette)
	for y := 0; y < m.size.y; y += step {
		for x := 0; x < m.size.x; x += step {
			p := XY{x: x, y: y}
			sprite := createSprite(m.At(p))
			dp := image.Point{x * scale, y * scale}
			r := image.Rectangle{dp, dp.Add(sprite.Bounds().Size())}
			draw.Draw(img, r, sprite, sprite.Bounds().Min, draw.Src)
		}
	}
	return img
}
