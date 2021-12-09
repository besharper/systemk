package ospkg

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"

	"github.com/containerd/containerd/remotes/docker"
	corev1 "k8s.io/api/core/v1"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// ImageManager manages unitfiles based on images.
type ImageManager struct{}

var _ Manager = (*ImageManager)(nil)

const (
	imageBaseRootPath = "/var/lib/fleet"
)

func (p *ImageManager) Install(container corev1.Container, version string) (bool, error) {
	// rootPath := fmt.Sprintf("%s/%s", imageBaseRootPath, container.Name)
	store := content.NewMemoryStore()
	resolver := docker.NewResolver(docker.ResolverOptions{
		Hosts: func(s string) ([]docker.RegistryHost, error) {
			return []docker.RegistryHost{
				{
					Host:   "hub.docker.com",
					Scheme: "https",
				},
			}, nil
		},
	})

	// pull image
	_, layers, err := oras.Pull(context.TODO(), resolver, container.Image, store)
	if err != nil {
		return false, err
	}
	// find the correct layer of the image
	for _, layer := range layers {
		if _, ok := layer.Annotations["org.opencontainers.image.title"]; ok {
			_, res, ok := store.Get(layer)
			if !ok {
				return false, errors.New("unable to get layer from store")
			}
			b := bytes.NewBuffer(res)
			tr := tar.NewReader(b)
			if _, err := tr.Next(); err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (p *ImageManager) Unitfile(pkg string) (string, error) {
	return "", nil
}
