package graphics

import (
	"flat_game/utils"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	ProgramId uint32
	Name      string
}

func NewShaderFromFiles(name string, vertexShaderFileName string, fragmentShaderFileName string) (*Shader, error) {
	vertexShader, err := utils.ReadTextFile(vertexShaderFileName)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := utils.ReadTextFile(fragmentShaderFileName)
	if err != nil {
		panic(err)
	}

	shader, err := NewShader(name, vertexShader+"\x00", fragmentShader+"\x00")
	if err != nil {
		panic(err)
	}

	return shader, nil
}

func NewShader(name string, vertexShaderSource string, fragmentShaderSource string) (*Shader, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	shaderProgram := gl.CreateProgram()

	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{shaderProgram, name}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.ProgramId)
}

func (shader *Shader) SetInteger(name string, value int32, useShader bool) {
	if useShader {
		shader.Use()
	}

	gl.Uniform1i(gl.GetUniformLocation(shader.ProgramId, gl.Str(name+"\x00")), value)
}

func (shader *Shader) SetVector3f(name string, value mgl32.Vec3, useShader bool) {
	if useShader {
		shader.Use()
	}

	gl.Uniform3f(gl.GetUniformLocation(shader.ProgramId, gl.Str(name+"\x00")), value.X(), value.Y(), value.Z())
}

func (shader *Shader) SetMatrix4(name string, matrix *mgl32.Mat4, useShader bool) {
	if useShader {
		shader.Use()
	}

	gl.UniformMatrix4fv(
		gl.GetUniformLocation(shader.ProgramId, gl.Str(name+"\x00")),
		1,
		false,
		&matrix[0],
	)
}
