package ospkg

import (
	"testing"
)

func TestImage(t *testing.T) {
	d := new(ImageManager)
	ok, err := d.Install("docker.io/library/busybox:latest", "")
	if err != nil || !ok {
		// not installed
		t.Error(err)
	}
	path, err := d.Unitfile("docker.io/library/busybox:latest")
	if err != nil || !ok {
		// not installed
		t.Error(err)
	}
	if path == "" {
		t.Errorf("invalid path: %s", path)
	}
}
