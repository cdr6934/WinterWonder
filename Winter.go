package main

import (
	"github.com/go-p5/p5"
	"image/color"
	"math/rand"
)


func main() {
	p5.Run(setup, draw)

}


func setup() {
	p5.Canvas(1000,1000)
	p5.Background(color.Gray{Y: 220})

}

func draw() {
	p5.StrokeWidth(2)
	p5.Fill(color.Gray{Y: 110})
	p5.Ellipse(rand.Float64()*1000,rand.Float64()*1000,80,80)
}