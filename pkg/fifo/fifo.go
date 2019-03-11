package fifo

import (
	"bufio"
	"fmt"
	"github.com/go-mesh/openlogging"
	"os"
	"path/filepath"
	"syscall"
)

const (

	//agent read from /paas-apm/collectors/pinpoint/{pid}/{app}/{service}
	PathTemplate = "/paas-apm/collectors/pinpoint/%d/%s"
)

func NewPath(app, service string) (string, error) {
	folder := fmt.Sprintf(PathTemplate, os.Getpid(), app)
	_, err := os.Stat(folder)
	if err != nil {
		if !os.IsNotExist(err) {
			openlogging.Error("apm: " + err.Error())
			return "", err
		}
	}
	if os.IsNotExist(err) {
		openlogging.Info("fifo folder do not exist, creating")
		if err := os.MkdirAll(folder, 0700); err != nil {
			openlogging.Error("apm: " + err.Error())
			return "", err
		}
		openlogging.Info("fifo folder created")
	}

	path := fmt.Sprintf(filepath.Join(folder, service))
	return path, nil
}

//NewWriter create writer for you, you can write bytes to fifo,
//an apm agent will read bytes and send to cloud service
func NewWriter(app, service string) (*bufio.Writer, error) {
	path, err := NewPath(app, service)
	if err != nil {
		openlogging.Error("can not get path: " + err.Error())
		return nil, err
	}
	_, err = os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}
	if os.IsNotExist(err) {
		openlogging.Info("fifo do not exist, creating")
		if err := syscall.Mkfifo(path, 0600); err != nil {
			return nil, err
		}
		openlogging.Info("fifo created")
	}
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		openlogging.Error("can not open fifo: " + err.Error())
		return nil, err
	}
	w := bufio.NewWriter(f)
	return w, nil
}

func NewReader(app, service string) (*bufio.Reader, error) {
	path, err := NewPath(app, service)
	if err != nil {
		openlogging.Error("can not get path: " + err.Error())
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		openlogging.Error("can not open fifo: " + err.Error())
		return nil, err
	}
	openlogging.Info("fifo reader created: " + f.Name())
	return bufio.NewReader(f), nil
}
