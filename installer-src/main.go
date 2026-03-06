package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	installCmd := `
cd /tmp/
git clone https://github.com/Yyax13/Gommit
cd Gommit
make build
make install
cd
`
	cmd := exec.Command("/bin/sh", "-c", installCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("Successfully installed gommit!")
	
}