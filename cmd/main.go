package main

import (
	"grender"
	"log"
)

func main() {
	err := grender.CreateWindow(640, 480, "GRender Test")
	if err != nil {
		log.Fatalln(err)
	}
	defer grender.CloseWindow()

	tex1, err := grender.NewTexture("../testdata/tex1.jxl")
	if err != nil {
		log.Fatalln(err)
	}
	tex2, err := grender.NewTexture("../testdata/tex2.jxl")
	if err != nil {
		log.Fatalln(err)
	}
	tex3, err := grender.NewTexture("../testdata/tex4.jxl")
	if err != nil {
		log.Fatalln(err)
	}
	colorTex := grender.NewTextureColor(255, 107, 192) // some sort of pink5

	atlas := grender.NewAtlas(4096)

	atlas.AddTexture(tex1)
	atlas.AddTexture(tex2)
	atlas.AddTexture(tex3)
	atlas.AddTexture(colorTex)

	atlas.Upload()

	r, err := grender.NewRenderer(atlas)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Cleanup()

	grender.SetTargetFPS(60)

	for grender.WindowShouldNotClose() {
		log.Println(1 / r.GetDeltaTime())
		// loop logic here
		r.Begin()

		r.DrawTexture(tex1, 0, 0)
		r.DrawTexture(tex2, uint32(tex1.Size.X), 0)
		r.DrawTexture(tex3, uint32(tex1.Size.X+tex2.Size.X), 0)
		r.DrawColorTexture(colorTex, uint32(tex1.Size.X+tex2.Size.X+tex3.Size.X), 0, 50, 50)

		r.End()
	}
}
