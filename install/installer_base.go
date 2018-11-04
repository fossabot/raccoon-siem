package install

import (
	"errors"
	"os"
	"os/user"
	"path"
	"runtime"
)

const (
	osNameLinux  = "linux"
	osNameDarwin = "darwin"
)



type baseInstaller struct {
	dataDir    string
	binaryDir  string
	serviceDir string
}

func (r *installer) install() error {
	switch r.os {
	case osNameDarwin:
		return r.installDarwin()
	case osNameLinux:
		return r.installLinux()
	default:
		return errUnsupportedOS
	}
}

func (r *installer) installDarwin() error {

}

func (r *installer) installLinux() error {

}

func (r *installer) createDirectories() {
	os.MkdirAll(r.)
}

func newInstaller() (inst *installer, err error) {
	switch runtime.GOOS {
	case osNameDarwin:
		inst, err = newDarwinInstaller()
	case osNameLinux:
		inst, err = newLinuxInstaller()
	default:
		err = errUnsupportedOS
	}
	return
}

func newDarwinInstaller() (*installer, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &installer{
		os:         osNameDarwin,
		dataDir:    "/usr/local/var/raccoon",
		binaryDir:  "/usr/local/bin",
		serviceDir: path.Join(u.HomeDir, "Library/LaunchAgents"),
	}, nil
}

func newLinuxInstaller() (*installer, error) {
	return &installer{
		os:         osNameLinux,
		dataDir:    "/var/lib/raccoon",
		binaryDir:  "/usr/local/bin",
		serviceDir: "/etc/systemd/system",
	}, nil
}
