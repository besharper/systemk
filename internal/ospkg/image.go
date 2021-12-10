package ospkg

import (
	"fmt"

	"github.com/ForAllSecure/rootfs_builder/rootfs"
	corev1 "k8s.io/api/core/v1"
)

// ImageManager manages unitfiles based on images.
type ImageManager struct{}

var _ Manager = (*ImageManager)(nil)

const (
	imageBaseRootPath = "/tmp/fleet"
)

func (p *ImageManager) Install(container corev1.Container, version string) (bool, error) {
	rootPath := fmt.Sprintf("%s/%s", imageBaseRootPath, container.Name)
	image := rootfs.PullableImage{
		Name:    container.Image,
		Retries: 1,
		Spec: rootfs.Spec{
			Dest: rootPath,
			User: "root",
		},
	}

	pulledManifest, err := image.Pull()
	if err != nil {
		return false, fmt.Errorf("failed to pull image manifest: %+v", err)
	}

	// Extract rootfs
	err = pulledManifest.Extract()
	if err != nil {
		return false, fmt.Errorf("failed to extract rootfs: %+v", err)
	}
	return true, nil
}

func (p *ImageManager) Unitfile(pkg string) (string, error) {
	return "", nil
}
