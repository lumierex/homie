package executor

import (
	"fmt"
	"os/exec"
)

func ExecuteShellScript(cmd string) error {
    c := exec.Command("bash", "-c", cmd)
    // 此处是windows版本
    // c := exec.Command("cmd", "/C", cmd)
    output, err := c.CombinedOutput()
    fmt.Println(string(output))
    return err
}