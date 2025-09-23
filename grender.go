package grender

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var glfwWindow *glfw.Window
var windowWidth, windowHeight uint32

func CreateWindow(width, height uint32, title string) error {
	if err := glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return err
	}
	glfwWindow = window
	windowWidth = width
	windowHeight = height

	return nil
}

func WindowShouldNotClose() bool {
	return !glfwWindow.ShouldClose()
}

func CloseWindow() {
	glfw.Terminate()
}
