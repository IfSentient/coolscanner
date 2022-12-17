package linux

import (
	"bytes"
	"coolscanner/pkg/models"
	"os/exec"
	"strings"
)

type RPMCollector struct {
}

func (r *RPMCollector) Collect() ([]models.Package, error) {
	c := exec.Command("rpm", "-qa", "--queryformat", "%{NAME} %{EPOCH} %{VERSION} %{RELEASE} %{ARCH} %{VENDOR}\\n")
	resp := bytes.Buffer{}

	c.Stdout = &resp
	err := c.Run()
	if err != nil {
		return nil, err
	}

	pkgs := make([]models.Package, 0)
	for _, pkg := range strings.Split(resp.String(), "\n") {
		if pkg == "" {
			continue
		}
		parts := strings.Split(pkg, " ")
		pkgs = append(pkgs, models.Package{
			Type:    models.PackageType("rpm"),
			Name:    parts[0],
			Epoch:   parts[1],
			Version: parts[2],
			Release: parts[3],
			Arch:    parts[4],
		})
	}

	return pkgs, nil
}
