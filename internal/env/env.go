package env

import (
	"github.com/miltian/homie/internal/util/executor"
	"github.com/miltian/homie/pkg/static"
	"log"
)

func MiniCondaInstall() error {

	content, err := static.FS.ReadFile("assets/shell/python/test.sh")
	if err != nil {
		log.Printf("read test sh err: %v", err)
		return err
	}
	err = executor.ExecuteShellScript(string(content))
	if err != nil {
		log.Println("ExecuteShellScript error")
		return err
	}
	return nil
}
