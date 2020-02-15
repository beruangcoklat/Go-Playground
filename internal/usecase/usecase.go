package usecase

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/beruangcoklat/Go-Playground/config"
	"github.com/beruangcoklat/Go-Playground/internal/entity"
)

func RunCode(code string) (*entity.RunCodeResponse, error) {
	filepath := fmt.Sprintf("%v%v.go", config.COMPILE_PATH, time.Now().UnixNano())

	err := ioutil.WriteFile(filepath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	var outb, errb bytes.Buffer

	cmd := exec.Command("go", "run", filepath)
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	cmd.Start()

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err = <-done:
		//
	case <-time.After(1 * time.Second):
		cmd.Process.Kill()
	}

	var resp string
	if err == nil {
		resp = outb.String()
	} else {
		resp = errb.String()
	}

	return &entity.RunCodeResponse{Data: resp}, nil
}
