package docker

import (
	"io"
	"os"
	"github.com/pkg/errors"
	"bufio"
	"strings"
	"log"
	"runtime"
)

type ContainerDetector interface {
	InContainer() (answer bool, err error)
}

type DockerDetector struct {
}

func NewDockerDetector() (dd *DockerDetector) {
	dd = &DockerDetector{
	}

	return
}

func (dd *DockerDetector) InContainer() (answer bool, err error) {

	switch runtime.GOOS {
	case "linux":
		var f io.ReadCloser
		f, err = os.Open("/proc/1/cgroup")
		if err != nil {
			err = errors.Wrap(err, "Can't open cgroup")
			return
		}
		buf := bufio.NewScanner(f)

		//
		for buf.Scan() {
			sp := strings.Split(buf.Text(), ":")
			if len(sp) != 3 {
				return false, errors.New("Unknown Cgroup format... you sure this is linux?")
			}
			if sp[1] == "pids" {
				if sp[2] == "/" {
					answer = false
				} else {
					answer = true
				}
			}
		}

		if err := buf.Err(); err != nil {
			log.Fatal(err)
		}
	default:
		answer = false
	}

	return
}
