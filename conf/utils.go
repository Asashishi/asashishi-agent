package conf

import "os/exec"

func EnvDetect() string {
	var err error
	if err = exec.Command("uname").Run(); err != nil {
		return Windows
	}
	return Linux
}
