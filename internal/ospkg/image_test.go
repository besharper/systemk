package ospkg

import (
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestImage(t *testing.T) {
	d := new(ImageManager)
	c := v1.Container{
		Name:  "startup",
		Image: "docker.io/library/busybox:latest",
	}
	ok, err := d.Install(c, "")
	if err != nil || !ok {
		// not installed
		t.Error(err)
	}
}
