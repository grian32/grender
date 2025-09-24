package grender

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var glfwWindow *glfw.Window
var windowWidth, windowHeight uint32
var targetDelta float64
var inputManager *Input

func init() {
	runtime.LockOSThread()
}

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
	inputManager = &Input{
		prevKeys:    make(map[glfw.Key]bool),
		currKeys:    make(map[glfw.Key]bool),
		prevButtons: make(map[glfw.MouseButton]bool),
		currButtons: make(map[glfw.MouseButton]bool),
	}

	return nil
}

func WindowShouldNotClose() bool {
	return !glfwWindow.ShouldClose()
}

func SetTargetFPS(fps uint32) {
	targetDelta = 1.0 / float64(fps)
}

func CloseWindow() {
	glfw.Terminate()
}
