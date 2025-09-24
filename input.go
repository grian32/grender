package grender

import "github.com/go-gl/glfw/v3.3/glfw"

type Input struct {
	prevKeys map[glfw.Key]bool
	currKeys map[glfw.Key]bool
}

func (i *Input) Update() {
	for k, v := range i.currKeys {
		i.prevKeys[k] = v
	}

	for k := glfw.KeySpace; k <= glfw.KeyLast; k++ {
		i.currKeys[k] = glfwWindow.GetKey(k) == glfw.Press
	}
}

func GetMousePos() (x, y float64) {
	return glfwWindow.GetCursorPos()
}

// TODO: decouple from glfw here
func IsKeyDown(k glfw.Key) bool {
	return inputManager.currKeys[k]
}

func IsKeyPressed(k glfw.Key) bool {
	return inputManager.currKeys[k] && !inputManager.prevKeys[k]
}

func IsKeyReleased(k glfw.Key) bool {
	return !inputManager.currKeys[k] && inputManager.prevKeys[k]
}
