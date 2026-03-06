package core

import (
	"fmt"
	"os/exec"
	"os"
	"yyax13/gommit/src/utils"
)

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
    diffBytes, err := cmd.Output()
    if err != nil {
        return "", err

    }

    diff := string(diffBytes)
    if len(diff) == 0 {
        fmt.Println(utils.Green("No staged changes to commit"))
        os.Exit(1);

    };

	return diff, nil

};

func GetHist() (string, error) {
	cmd := exec.Command("git", "log", "--oneline", "-n", "1")
	histBytes, err := cmd.Output()
	if err != nil {
		return "", err
		
	}
	
	hist := string(histBytes)
	if len(hist) == 0 {
		fmt.Println(utils.Yellow("Current repository don't have hist entries"))
		os.Exit(1)
		
	}
	
	return hist, nil
	
}

func Commit(commitMessage string) {
	cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	cmdCommit.Run()
	fmt.Println(utils.Green("Committed!"))

}

func Push(branch string, force bool) {
	var cmdPush *exec.Cmd
	if force {
		cmdPush = exec.Command("git", "push", branch, "--force")
		
	} else {
		cmdPush = exec.Command("git", "push", branch, "--force")
		
	}
	
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	cmdPush.Run()
	fmt.Println(utils.Green("Pushed!"))
	
}