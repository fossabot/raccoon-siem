package install

import (
	"os"
	"runtime"
)

const (
	osNameLinux     = "linux"
	osNameDarwin    = "darwin"
	serviceFileMode = 0644
)

type Installer interface {
	Install(component string) error
}

type baseInstaller struct {
	dataDir    string
	binaryDir  string
	serviceDir string
}

func (r *baseInstaller) takeDefaultActions() error {
	dirs := []string{r.dataDir, r.binaryDir, r.serviceDir}
	return r.createDirs(dirs)
}

func (r *baseInstaller) createDirs(dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func newInstaller() (inst Installer, err error) {
	switch runtime.GOOS {
	case osNameDarwin:
		inst, err = newDarwinInstaller()
	default:
		err = errUnsupportedOS
	}
	return
}
