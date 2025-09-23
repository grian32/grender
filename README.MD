# GRender

A high-performance 2D Rendering library made in Go, designed for game development.
Provides a simple API that allows the rendering of textures and fonts, handling input, etc., while achieving high performance through automatic texture atlasing and instanced rendering. Built on OpenGL 3.3 Core and GLFW

Currently, JPEG XL is the only format supported for texture loading, JPEG XL is a modern format that is superior to traditional formats. Support for traditional formats may be added in the future.

GRender was created as a replacement for Raylib for [GRPG](https://github.com/grian32/grpg). You may notice some similarities.

## Current Features

- Texture management & automatic atlasing.
- Rendering API.

## Planned Features

- Text Rendering
- API for rendering solid color primitives.
- Input Handling
- API for providing lighting along with the necessary shaders.

## Examples

```go
// initialization of textures, renderer, window, etc.
for grender.WindowShouldNotClose() {
	// loop logic here
	r.Begin()

	r.DrawTexture(tex1, 0, 0)
	r.DrawTexture(tex2, uint32(tex1.Size.X), 0)
	r.DrawTexture(tex3, uint32(tex1.Size.X+tex2.Size.X), 0)

	r.End()
}
```

The full example code can be found in `cmd/main.go`.