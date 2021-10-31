package space

import (
	"fmt"
	"os"
	"path"
)

type Space struct {
	path string
	name string
	f    *os.File
}

func NewSpace(spacePath, name string) (s *Space, err error) {
	s = new(Space)
	if spacePath == "" {
		spacePath, err = os.Getwd()
		if err != nil {
			return
		}
	}
	dirPath := path.Join(spacePath, name)
	dirInfo, err := os.Stat(dirPath)
	if ok := os.IsExist(err); !ok {
		err = os.MkdirAll(dirPath, 0777)
		if err != nil {
			return
		}
	} else {
		if !dirInfo.IsDir() {
			err = fmt.Errorf("path: %v not directory", dirPath)
			return
		}
	}
	s.f, err = os.OpenFile(path.Join(dirPath, name), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	s.path = dirPath
	s.name = name
	return
}

func (s *Space) Close() error {
	return s.f.Close()
}

func (s *Space) GetDir() string {
	return s.path
}
