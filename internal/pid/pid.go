package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type PID struct {
	PID int `json:"pid"`
}

// Path zwraca sciezke do pliku pid
func (pid *PID) Path() string {
	procName := filepath.Base(os.Args[0])
	filename := fmt.Sprintf("%s.pid", procName)
	return filepath.Join(string("/tmp/"), filename)
}

func (pid *PID) Save() error {
	filename := pid.Path()
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return fmt.Errorf("PID file already exists: %s", filename)
	}
	pid.PID = os.Getpid()
	err := ioutil.WriteFile(filename, []byte(strconv.Itoa(pid.PID)), 0666)
	return err
}

func (pid *PID) Close() error {
	filename := pid.Path()
	return os.Remove(filename)
}
