package system

import (
	"os"
	"io"
		"bufio"
	"github.com/pkg/errors"
	"fmt"
	"log"
)

type Detector interface {
	InContainer() (answer bool, err error)
}

type DockerDetector struct {
}

func (dd *DockerDetector) InContainer() (answer bool, err error) {
	var f io.ReadCloser
	f, err = os.Open("/proc/1/cgroup")
	if err != nil {
		err = errors.Wrap(err, "Can't open cgroup")
		return
	}
	buf := bufio.NewScanner(f)

	//
	for buf.Scan() {
		fmt.Println(buf.Text())
	}

	if err := buf.Err() ; err != nil {
		log.Fatal(err)
	}

	return
}
