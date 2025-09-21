# GRender

GRender is a 2D Rendering library made in Go, primarily meant to be used for game development.
It aims to provide a relatively simple API that allows the rendering of textures and fonts, handling input, etc.

GRender achieves high performance by automatically atlasing all textures and batching rendering. To achieve this it uses OpenGL 3.3 Core in conjuction with GLFW.