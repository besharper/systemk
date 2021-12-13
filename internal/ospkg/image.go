package ospkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ForAllSecure/rootfs_builder/rootfs"
	"github.com/virtual-kubelet/systemk/internal/system"
	"github.com/virtual-kubelet/systemk/internal/unit"
)

// ImageManager manages unitfiles based on images.
type ImageManager struct{}

var _ Manager = (*ImageManager)(nil)

const (
	imageBaseRootPath = "/tmp/fleet"
)

const baseUnit = `[Unit]

[Service]

[Install]
`

func (p *ImageManager) Install(pkg, version string) (bool, error) {
	rootPath := getRootPath(pkg)
	if _, err := os.Stat(rootPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(rootPath, os.ModePerm)
			if err != nil {
				return false, err
			}
		}
	}
	image := rootfs.PullableImage{
		Name:    pkg,
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
	basicPath := ""
	serviceName := prepareServiceName(pkg)

	// Determine OS
	systemID := system.ID()
	switch systemID {
	case "debian", "ubuntu":
		basicPath = debianSystemdUnitfilesPathPrefix + serviceName + unit.ServiceSuffix
	case "arch":
		basicPath = archlinuxSystemdUnitfilesPathPrefix + serviceName + unit.ServiceSuffix
	}

	// Get/Create unitfile
	uf := &unit.File{}
	if _, err := os.Stat(basicPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			uf, err = unit.NewFile(baseUnit)
			if err != nil {
				return basicPath, err
			}
		}
	} else {
		buf, err := ioutil.ReadFile(basicPath)
		if err != nil {
			return basicPath, err
		}
		uf, err = unit.NewFile(string(buf))
		if err != nil {
			return basicPath, err
		}
	}

	// Required fields
	uf = uf.Overwrite("Service", "RootDirectory", fmt.Sprintf("%s/rootfs", getRootPath(pkg)))

	// Write contents
	bContents := []byte(uf.String())
	log.Infof("writing systemd unit %q (%dB written)", basicPath, len(bContents))
	err := ioutil.WriteFile(basicPath, bContents, os.FileMode(0644))
	if err != nil {
		return basicPath, err
	}

	return basicPath, nil
}

func getRootPath(pkg string) string {
	return fmt.Sprintf("%s/%s", imageBaseRootPath, pkg)
}

func prepareServiceName(pkg string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pkg, "/", "-"), ":", "-"), ".", "-")
}
