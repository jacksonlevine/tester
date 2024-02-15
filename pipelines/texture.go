package pipelines

import (
	"os"
	"io"
	"log"
	"github.com/go-gl/gl/v3.3-compatibility/gl"
)

type Pipeline struct {
	VertexShaderPath string
	FragmentShaderPath string
	ProgramID uint32
}

func loadAndCompileShader(shaderPath string, shaderType uint32) uint32 {
	file, err := os.Open(shaderPath)
	if err != nil {
        log.Fatal(err)
    }
	defer func() {
		err := file.Close()
		if(err != nil) {
			log.Fatal(err)
		}
	}()
	text, err := io.ReadAll(file)
	if(err != nil) {
		log.Fatal(err)
	}
	cStr, freeCStr := gl.Strs(string(text) + "\x00")
	defer freeCStr();
	var shader uint32 = gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, cStr, nil)
	gl.CompileShader(shader)
	return shader;
}

func (pipe *Pipeline) Setup() {
	vertShader := loadAndCompileShader(pipe.VertexShaderPath, gl.VERTEX_SHADER)
	fragShader := loadAndCompileShader(pipe.FragmentShaderPath, gl.FRAGMENT_SHADER)
	defer func() {
		gl.DeleteShader(vertShader)
		gl.DeleteShader(fragShader)
	}()
	pipe.ProgramID = gl.CreateProgram()
	gl.AttachShader(pipe.ProgramID, vertShader)
	gl.AttachShader(pipe.ProgramID, fragShader)
	gl.LinkProgram(pipe.ProgramID)
}