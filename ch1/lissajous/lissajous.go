// Lissajous генерирует анимированный GIF из случайных фигур Лиссажу.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.Black, color.RGBA{0, 255, 0, 0xff}, color.Opaque, color.RGBA{128, 64, 255, 0xff}}

const (
	blackIndex = 0 // Первый цвет палитры
	greenIndex = 1 // Второй цвет палитры
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	freq := r.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size*0.5), size+int(y*size*0.5), uint8(r.Int()%len(pallete)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// 3. Создаем файл для записи
	f, err := os.Create("animation.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 4. Кодируем и записываем GIF
	err = gif.EncodeAll(f, &anim)
	if err != nil {
		panic(err)
	}
}
