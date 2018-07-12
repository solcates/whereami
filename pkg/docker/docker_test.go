package docker

import (
	"reflect"
	"testing"
)

func TestNewDockerDetector(t *testing.T) {
	tests := []struct {
		name   string
		wantDd *DockerDetector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDd := NewDockerDetector(); !reflect.DeepEqual(gotDd, tt.wantDd) {
				t.Errorf("NewDockerDetector() = %v, want %v", gotDd, tt.wantDd)
			}
		})
	}
}

func TestDockerDetector_InContainer(t *testing.T) {
	tests := []struct {
		name       string
		dd         *DockerDetector
		wantAnswer bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnswer, err := tt.dd.InContainer()
			if (err != nil) != tt.wantErr {
				t.Errorf("DockerDetector.InContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("DockerDetector.InContainer() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
