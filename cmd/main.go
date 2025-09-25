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

	arial, err := grender.NewFont("../testdata/arial.ttf", 24)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := grender.NewRenderer(atlas, []*grender.Font{arial})
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Cleanup()

	grender.SetTargetFPS(60)

	for grender.WindowShouldNotClose() {
		if grender.IsKeyPressed(grender.X) {
			log.Printf("pressed x\n")
		}
		if grender.IsKeyDown(grender.J) {
			log.Printf("down j\n")
			x, y := grender.GetMousePos()
			log.Printf("mouse pos: %.3f, %.3f", x, y)
		}
		if grender.IsKeyReleased(grender.N) {
			log.Printf("release n\n")
		}

		if grender.IsMouseButtonPressed(grender.MouseLeft) {
			log.Printf("pressing mouse button left")
		}

		if grender.IsMouseButtonDown(grender.MouseRight) {
			log.Printf("down mouse button right")
		}

		if grender.IsMouseButtonReleased(grender.MouseMiddle) {
			log.Printf("released mouse button middle")
		}

		// commented out as is spammy
		//if grender.IsMouseButtonUp(grender.Mouse4) {
		//	log.Printf("up mouse button 4")
		//}

		r.Begin()

		r.DrawTexture(tex1, 0, 0)
		r.DrawTexture(tex2, uint32(tex1.Size.X), 0)
		r.DrawTexture(tex3, uint32(tex1.Size.X+tex2.Size.X), 0)
		r.DrawColorTexture(colorTex, uint32(tex1.Size.X+tex2.Size.X+tex3.Size.X), 0, 50, 50)

		r.DrawText(arial, "hello, world!", 0, uint32(tex2.Size.Y))

		r.End()
	}
}
