package cmd

import (
    "fmt"
    "os"
    "sync"

    "yyax13/gommit/src/core"
    "yyax13/gommit/src/utils"
    "yyax13/gommit/src/config"

    "github.com/spf13/cobra"
)

var (
    flagCommit  bool
    flagPush    bool
    flagForce   bool
    flagBranch  string
)

var genCmd = &cobra.Command{
    Use:   "gen",
    Short: "Generate a commit message",
    Run: func(cmd *cobra.Command, args []string) {
	    var confPath string = config.GetConfigPath()
        config.EnsureConfig(confPath)
    
        cfg, err := config.LoadConfig(confPath)
        if err != nil {
            fmt.Fprintln(os.Stderr, utils.Red("Can't get config"))
            return
        }
    
        diff, err := core.GetDiff()
        if err != nil {
            fmt.Fprintln(os.Stderr, utils.Red("Can't get diff, aborting..."))
            return
        }
    
        var hist string
        if cfg.UseHist {
            hist, err = core.GetHist()
            if err != nil {
                fmt.Fprintln(os.Stderr, utils.Red("Can't get history, aborting"))
                return
            }
        }
    
        var wg sync.WaitGroup
        var finalMessage string
    
        wg.Add(1)
        go core.GetCommitMessage(&finalMessage, diff, hist, &wg, cfg)
        wg.Wait()
    
        fmt.Println(utils.Green("Generated commit message:"))
        fmt.Println(utils.Green(finalMessage))
    
        // SE NÃO PASSOU -c / --commit → só imprime
        if !flagCommit {
            return
        }
    
        // commit usando mensagem gerada
        core.Commit(finalMessage)
    
        if flagPush {
            branch := flagBranch
            if branch == "" {
                branch = "origin"
            }
    
            core.Push(branch, flagForce)
        }
    },
}

func init() {
    rootCmd.AddCommand(genCmd)

    // Flags do comando gen
    genCmd.Flags().BoolVarP(&flagCommit, "commit", "c", false, "Commit the generated message")
    genCmd.Flags().BoolVarP(&flagPush, "push", "p", false, "Push after committing")
    genCmd.Flags().BoolVar(&flagForce, "force", false, "Add --force when pushing")
    genCmd.Flags().StringVar(&flagBranch, "branch", "", "Branch to push to (default: origin)")
}