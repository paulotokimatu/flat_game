package graphics

import (
	"flat_game"
	"flat_game/input"
	"flat_game/utils"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type OpenGl struct {
	// game        internal.IGame
	bgColor     utils.Vec3
	keysPressed [glfw.KeyLast]bool
	renderer    *SpriteRenderer
	shaders     map[string]*Shader
	window      *glfw.Window
}

func (openGl *OpenGl) Name() string {
	return "opengl"
}

func (openGl *OpenGl) Setup(config *flat_game.GameConfig) {
	// openGl.game = game
	openGl.bgColor = config.BgColor

	shaders := make(map[string]*Shader)
	openGl.shaders = shaders

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(int(config.Size.X), int(config.Size.Y), config.Name, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	window.SetKeyCallback(openGl.keyCallback)

	openGl.window = window

	// Initialize Glow (I do not know if this block should be here)
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	shaderFolder := basePath() + "/shaders/"
	spriteShader, err := NewShaderFromFiles("sprite", shaderFolder+"sprite.vert", shaderFolder+"sprite.frag")
	if err != nil {
		panic(err)
	}

	spriteShader.Use()

	spriteShader.SetInteger("sprite", 0, false)

	projection := mgl32.Ortho2D(0.0, config.Size.X, config.Size.Y, 0)
	spriteShader.SetMatrix4("projection", &projection, false)

	renderer, err := NewSpriteRenderer(spriteShader)
	if err != nil {
		panic(err)
	}

	openGl.renderer = renderer
}

func (openGl *OpenGl) Terminate() {
	glfw.Terminate()
}

func (openGl *OpenGl) WindowShouldClose() bool {
	return openGl.window.ShouldClose()
}

func (openGl *OpenGl) PreTick() {
	glfw.PollEvents()

	gl.ClearColor(openGl.bgColor.X, openGl.bgColor.Y, openGl.bgColor.Z, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (openGl *OpenGl) Tick() {
}

func (openGl *OpenGl) PostTick() {
	openGl.window.SwapBuffers()
}

func (openGl *OpenGl) IsKeyPressed(key input.Key) bool {
	return openGl.window.GetKey(glfw.Key(key)) == glfw.Press
}

func (openGl *OpenGl) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key < 0 || key >= 1024 {
		return
	}

	// We do not need to do much to map glfw values since our map was based on glfw itself
	// transformedKey := input.Key(key)

	if action == glfw.Press && !openGl.keysPressed[key] {
		openGl.keysPressed[key] = true
		// openGl.game.OnKeyEvent(transformedKey, input.EventKeyPressed)
	} else if action == glfw.Release && openGl.keysPressed[key] {
		openGl.keysPressed[key] = false
		// openGl.game.OnKeyEvent(transformedKey, input.EventKeyReleased)
	}
}

func (openGl *OpenGl) DrawSprite(texture flat_game.ITexture, position *utils.Vec2, size *utils.Vec2, color utils.Vec3) {
	openGl.renderer.DrawSprite(texture, position, size, color)
}

func basePath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}
