package pid

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/hhkbp2/testify/require"
)

func TestProcessName(t *testing.T) {
	pid := new(PID)
	pid.Path()
	require.Equal(t, "/tmp/pid.test.pid", pid.Path())
}

func TestSave(t *testing.T) {
	pid := new(PID)
	defer pid.Close()
	require.NoError(t, pid.Save())
	pidname := "/tmp/pid.test.pid"
	cont, err := ioutil.ReadFile(pidname)
	require.NoError(t, err)
	filepid, err := strconv.Atoi(string(cont))
	require.NoError(t, err)
	require.Equal(t, os.Getpid(), filepid)
}

func TestClose(t *testing.T) {
	pid := new(PID)
	defer pid.Close()

	require.NoError(t, pid.Save())
	require.NoError(t, pid.Close())
	_, err := os.Stat(pid.Path())
	require.Error(t, err)
}

func TestCreateExistingPID(t *testing.T) {
	pid := new(PID)
	pid.Save()
	defer pid.Close()

	pid2 := new(PID)
	require.Error(t, pid2.Save())
	defer pid2.Close()
}
