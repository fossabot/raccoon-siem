package install

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
	"os/user"
	"path"
	"text/template"
)

const (
	darwinDefaultBinaryDir        = "/usr/local/bin"
	darwinDefaultDataDir          = "/usr/local/var"
	darwinDefaultServiceDirSuffix = "Library/LaunchAgents"
)

type darwinServiceTemplateInput struct {
	Name string
	Args []string
}

var darwinCoreServiceTemplate = template.Must(template.New(componentCore).Parse(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>{{.Name}}</string>

  <key>ProgramArguments</key>
  <array>
	{{range .Args}}
		<string>{{.}}</string>	
	{{end}}
  </array>

  <key>RunAtLoad</key>
  <true/>
</dict>
</plist>`))

type installerDarwin struct {
	baseInstaller
}

func (r *installerDarwin) Install(component string) error {
	if err := r.takeDefaultActions(); err != nil {
		return err
	}

	switch component {
	case componentCore:
		return r.installCore()
	case componentCollector:
		return r.installCollector()
	case componentCorrelator:
		return r.installCorrelator()
	case componentBus:
		return r.installBus()
	case componentAL:
		return r.installAL()
	case componentStorage:
		return r.installStorage()
	}

	return nil
}

func (r *installerDarwin) installCore() error {
	input := &darwinServiceTemplateInput{
		Name: "raccoon_core",
		Args: []string{
			path.Join(r.binaryDir, raccoonBinaryName),
			componentCore,
			fmt.Sprintf("--db %s", path.Join(r.dataDir, "raccoon.db")),
			fmt.Sprintf("--listen %s", flagCoreListenAddress),
		},
	}

	serviceFileContent := bytes.NewBuffer(make([]byte, 0))
	if err := darwinCoreServiceTemplate.Execute(serviceFileContent, input); err != nil {
		return err
	}

	servicePath := path.Join(r.serviceDir, fmt.Sprintf("%s.plist", raccoonCoreServiceName))
	if err := ioutil.WriteFile(
		servicePath,
		serviceFileContent.Bytes(),
		serviceFileMode,
	); err != nil {
		return err
	}

	return r.loadService(servicePath)
}

func (r *installerDarwin) installCollector() error {
	return errors.New("not implemented")
}

func (r *installerDarwin) installCorrelator() error {
	return errors.New("not implemented")
}

func (r *installerDarwin) installBus() error {
	return errors.New("not implemented")
}

func (r *installerDarwin) installAL() error {
	return errors.New("not implemented")
}

func (r *installerDarwin) installStorage() error {
	return errors.New("not implemented")
}

func (r *installerDarwin) loadService(path string) error {
	return exec.Command("launchctl", "load", path).Run()
}

func newDarwinInstaller() (Installer, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &installerDarwin{
		baseInstaller: baseInstaller{
			binaryDir:  darwinDefaultBinaryDir,
			dataDir:    darwinDefaultDataDir,
			serviceDir: path.Join(u.HomeDir, darwinDefaultServiceDirSuffix),
		},
	}, nil
}
