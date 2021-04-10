package machine

import (
	"log"
	"os/exec"
	"regexp"
)

func GetCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Println(err)
	}
	str := string(out)
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
