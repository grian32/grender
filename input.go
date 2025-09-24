package grender

import "github.com/go-gl/glfw/v3.3/glfw"

type Input struct {
	prevKeys    map[glfw.Key]bool
	currKeys    map[glfw.Key]bool
	prevButtons map[glfw.MouseButton]bool
	currButtons map[glfw.MouseButton]bool
}

func (i *Input) Update() {
	for k, v := range i.currKeys {
		i.prevKeys[k] = v
	}

	for k, v := range i.currButtons {
		i.prevButtons[k] = v
	}

	for k := glfw.KeySpace; k <= glfw.KeyLast; k++ {
		i.currKeys[k] = glfwWindow.GetKey(k) == glfw.Press
	}

	for k := glfw.MouseButton1; k <= glfw.MouseButtonLast; k++ {
		i.currButtons[k] = glfwWindow.GetMouseButton(k) == glfw.Press
	}
}

func GetMousePos() (x, y float64) {
	return glfwWindow.GetCursorPos()
}

type Key uint16

const (
	Unknown Key = iota
	Space
	Apostrophe
	Comma
	Minus
	Period
	Slash
	Zero
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Semicolon
	Equal
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	LeftBracket
	Backslash
	RightBracket
	GraveAccent
	World1
	World2
	Escape
	Enter
	Tab
	Backspace
	Insert
	Delete
	Right
	Left
	Down
	Up
	PageUp
	PageDown
	Home
	End
	CapsLock
	ScrollLock
	NumLock
	PrintScreen
	Pause
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	KP0
	KP1
	KP2
	KP3
	KP4
	KP5
	KP6
	KP7
	KP8
	KP9
	KPDecimal
	KPDivide
	KPMultiply
	KPSubtract
	KPAdd
	KPEnter
	KPEqual
	LeftShift
	LeftControl
	LeftAlt
	LeftSuper
	RightShift
	RightControl
	RightAlt
	RightSuper
	Menu
	Last
)

var grenderKeyToGlfw = map[Key]glfw.Key{
	Unknown:      glfw.KeyUnknown,
	Space:        glfw.KeySpace,
	Apostrophe:   glfw.KeyApostrophe,
	Comma:        glfw.KeyComma,
	Minus:        glfw.KeyMinus,
	Period:       glfw.KeyPeriod,
	Slash:        glfw.KeySlash,
	Zero:         glfw.Key0,
	One:          glfw.Key1,
	Two:          glfw.Key2,
	Three:        glfw.Key3,
	Four:         glfw.Key4,
	Five:         glfw.Key5,
	Six:          glfw.Key6,
	Seven:        glfw.Key7,
	Eight:        glfw.Key8,
	Nine:         glfw.Key9,
	Semicolon:    glfw.KeySemicolon,
	Equal:        glfw.KeyEqual,
	A:            glfw.KeyA,
	B:            glfw.KeyB,
	C:            glfw.KeyC,
	D:            glfw.KeyD,
	E:            glfw.KeyE,
	F:            glfw.KeyF,
	G:            glfw.KeyG,
	H:            glfw.KeyH,
	I:            glfw.KeyI,
	J:            glfw.KeyJ,
	K:            glfw.KeyK,
	L:            glfw.KeyL,
	M:            glfw.KeyM,
	N:            glfw.KeyN,
	O:            glfw.KeyO,
	P:            glfw.KeyP,
	Q:            glfw.KeyQ,
	R:            glfw.KeyR,
	S:            glfw.KeyS,
	T:            glfw.KeyT,
	U:            glfw.KeyU,
	V:            glfw.KeyV,
	W:            glfw.KeyW,
	X:            glfw.KeyX,
	Y:            glfw.KeyY,
	Z:            glfw.KeyZ,
	LeftBracket:  glfw.KeyLeftBracket,
	Backslash:    glfw.KeyBackslash,
	RightBracket: glfw.KeyRightBracket,
	GraveAccent:  glfw.KeyGraveAccent,
	World1:       glfw.KeyWorld1,
	World2:       glfw.KeyWorld2,
	Escape:       glfw.KeyEscape,
	Enter:        glfw.KeyEnter,
	Tab:          glfw.KeyTab,
	Backspace:    glfw.KeyBackspace,
	Insert:       glfw.KeyInsert,
	Delete:       glfw.KeyDelete,
	Right:        glfw.KeyRight,
	Left:         glfw.KeyLeft,
	Down:         glfw.KeyDown,
	Up:           glfw.KeyUp,
	PageUp:       glfw.KeyPageUp,
	PageDown:     glfw.KeyPageDown,
	Home:         glfw.KeyHome,
	End:          glfw.KeyEnd,
	CapsLock:     glfw.KeyCapsLock,
	ScrollLock:   glfw.KeyScrollLock,
	NumLock:      glfw.KeyNumLock,
	PrintScreen:  glfw.KeyPrintScreen,
	Pause:        glfw.KeyPause,
	F1:           glfw.KeyF1,
	F2:           glfw.KeyF2,
	F3:           glfw.KeyF3,
	F4:           glfw.KeyF4,
	F5:           glfw.KeyF5,
	F6:           glfw.KeyF6,
	F7:           glfw.KeyF7,
	F8:           glfw.KeyF8,
	F9:           glfw.KeyF9,
	F10:          glfw.KeyF10,
	F11:          glfw.KeyF11,
	F12:          glfw.KeyF12,
	F13:          glfw.KeyF13,
	F14:          glfw.KeyF14,
	F15:          glfw.KeyF15,
	F16:          glfw.KeyF16,
	F17:          glfw.KeyF17,
	F18:          glfw.KeyF18,
	F19:          glfw.KeyF19,
	F20:          glfw.KeyF20,
	F21:          glfw.KeyF21,
	F22:          glfw.KeyF22,
	F23:          glfw.KeyF23,
	F24:          glfw.KeyF24,
	F25:          glfw.KeyF25,
	KP0:          glfw.KeyKP0,
	KP1:          glfw.KeyKP1,
	KP2:          glfw.KeyKP2,
	KP3:          glfw.KeyKP3,
	KP4:          glfw.KeyKP4,
	KP5:          glfw.KeyKP5,
	KP6:          glfw.KeyKP6,
	KP7:          glfw.KeyKP7,
	KP8:          glfw.KeyKP8,
	KP9:          glfw.KeyKP9,
	KPDecimal:    glfw.KeyKPDecimal,
	KPDivide:     glfw.KeyKPDivide,
	KPMultiply:   glfw.KeyKPMultiply,
	KPSubtract:   glfw.KeyKPSubtract,
	KPAdd:        glfw.KeyKPAdd,
	KPEnter:      glfw.KeyKPEnter,
	KPEqual:      glfw.KeyKPEqual,
	LeftShift:    glfw.KeyLeftShift,
	LeftControl:  glfw.KeyLeftControl,
	LeftAlt:      glfw.KeyLeftAlt,
	LeftSuper:    glfw.KeyLeftSuper,
	RightShift:   glfw.KeyRightShift,
	RightControl: glfw.KeyRightControl,
	RightAlt:     glfw.KeyRightAlt,
	RightSuper:   glfw.KeyRightSuper,
	Menu:         glfw.KeyMenu,
	Last:         glfw.KeyLast,
}

func IsKeyDown(k Key) bool {
	return inputManager.currKeys[grenderKeyToGlfw[k]]
}

func IsKeyPressed(k Key) bool {
	return inputManager.currKeys[grenderKeyToGlfw[k]] && !inputManager.prevKeys[grenderKeyToGlfw[k]]
}

func IsKeyReleased(k Key) bool {
	return !inputManager.currKeys[grenderKeyToGlfw[k]] && inputManager.prevKeys[grenderKeyToGlfw[k]]
}

type MouseButton uint16

const (
	MouseLeft MouseButton = iota
	MouseRight
	MouseMiddle
	Mouse4
	Mouse5
	Mouse6
	Mouse7
	Mouse8
)

var grenderMouseToGlfw = map[MouseButton]glfw.MouseButton{
	MouseLeft:   glfw.MouseButton1,
	MouseRight:  glfw.MouseButton2,
	MouseMiddle: glfw.MouseButton3,
	Mouse4:      glfw.MouseButton4,
	Mouse5:      glfw.MouseButton5,
	Mouse6:      glfw.MouseButton6,
	Mouse7:      glfw.MouseButton7,
	Mouse8:      glfw.MouseButton8,
}

func IsMouseButtonPressed(b MouseButton) bool {
	return inputManager.currButtons[grenderMouseToGlfw[b]] && !inputManager.prevButtons[grenderMouseToGlfw[b]]
}

func IsMouseButtonDown(b MouseButton) bool {
	return inputManager.currButtons[grenderMouseToGlfw[b]]
}

func IsMouseButtonReleased(b MouseButton) bool {
	return !inputManager.currButtons[grenderMouseToGlfw[b]] && inputManager.prevButtons[grenderMouseToGlfw[b]]
}

func IsMouseButtonUp(b MouseButton) bool {
	return !inputManager.currButtons[grenderMouseToGlfw[b]]
}
