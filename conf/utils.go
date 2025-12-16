package conf

import "os/exec"

func EnvDetect() string {
	var err error
	if err = exec.Command("bash", "-c", "uname -s").Run(); err != nil {
		return Windows
	}
	return Linux
}
