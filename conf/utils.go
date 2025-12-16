package conf

import "os/exec"

func EnvDetect() string {
	if exec.Command("bash", "-c", "uname -s").Run().Error() != "" {
		return Windows
	}
	return Linux
}
