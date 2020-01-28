package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hhkbp2/go-logging"
	hlogging "github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
	h "github.com/redmeros/htrade/web/helpers"
)

var logger logging.Logger

func init() {
	logger = hlogging.NewLogger("dataCollector.log")
}

func resolveCollectorExec() (string, error) {
	var paths = []string{
		"../dist/data_collector",
		"data_collector",
	}
	for _, path := range paths {
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		if !info.IsDir() {
			return path, nil
		}
	}
	return "", errors.New("Żaden ze standardowych plików kolektora nie istnieje")
}

// CollectorStatus zwraca message
// ze statusem dzialania collectora
func CollectorStatus(c *gin.Context) {
	data, err := ioutil.ReadFile("/tmp/data_collector.pid")
	if err != nil {
		h.Badf(c, "Cannot read pid file: %s", 500, err)
		return
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		h.Badf(c, "PID file has wrong content", 500, err)
		return
	}

	_, err = os.FindProcess(pid)
	if err != nil {
		h.Badf(c, "Cannot find process with PID: %d: %s", 500, pid, err)
		return
	}
	c.JSON(models.Msg200f("Pid file found with pid %d, process still running", pid))

}

// CollectorStart startuje nowy data collector
// i tworzy dla niego pid
func CollectorStart(c *gin.Context) {
	path, err := resolveCollectorExec()
	if err != nil {
		h.Bad(c, "cannot find collector exec", 500)
		c.Abort()
		return
	}

	pidpath := "/tmp/data_collector.pid"
	if info, err := os.Stat(pidpath); err == nil {
		h.Bad(c, fmt.Sprintf("pidfile for datacollector already exists: %s", info.Name()), 400)
		c.Abort()
		return
	}

	cmd := exec.Command(path)
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		h.Bad(c, fmt.Sprintf("error during starting data collector: %s", err.Error()), 500)
		c.Abort()
		return
	}
	logger.Infof("Started process with pid %d", cmd.Process.Pid)
	c.JSON(models.Msg200f("Collector started with pid %d", cmd.Process.Pid))
}

// CollectorStop zamyka datacollectora
// na podstawiwe pliku pid
func CollectorStop(c *gin.Context) {
	path := "/tmp/data_collector.pid"
	var err error
	var pidcontent []byte
	if pidcontent, err = ioutil.ReadFile(path); err != nil {
		h.Badf(c, "pidfile does not exists: %s", 400, err.Error())
		return
	}
	pid, err := strconv.Atoi(string(pidcontent))
	proc, err := os.FindProcess(pid)
	if err != nil {
		h.Badf(c, "Cannot find process", 400, err.Error())
		return
	}

	if err := proc.Signal(syscall.SIGINT); err != nil {
		h.Badf(c, "Cannot send SIGINT signall process %s", 500, err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "dataCollector to be stopped"})
	return
}
