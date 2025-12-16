package conf

import "os/exec"

func EnvDetect() string {
	if exec.Command("uname", "-s").Run().Error() != "" {
		return Windows
	}
	return Linux
}
