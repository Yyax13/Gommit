package config

import (
	"path/filepath"
	"log"
	"github.com/spf13/viper"
	"fmt"
	"yyax13/gommit/src/utils"
	"os"

)

func LoadConfig(path string) (*utils.Config, error) {
    viper.SetConfigFile(path)

    viper.SetDefault("UseHist", false)
    viper.SetDefault("OverWriteDefaultCommitPatternPrompt", false)

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("can't read config: %w", err)
    
    }

    var cfg utils.Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("can't unmarshal config: %w", err)
        
    }

    return &cfg, nil
    
}

func EnsureConfig(path string) {
    dir := filepath.Dir(path)

    if _, err := os.Stat(dir); os.IsNotExist(err) {
        if err := os.MkdirAll(dir, 0755); err != nil {
            log.Fatal("error creating dir:", err)
        }
    }

    if _, err := os.Stat(path); os.IsNotExist(err) {

        defaultConfig := `GeminiApiKey: ""
UseHist: false
CommitPatternPrompt: ""
OverwriteDefaultCommitPatternPrompt: false

`

        err := os.WriteFile(path, []byte(defaultConfig), 0644)
        if err != nil {
            log.Fatal("error creating config:", err)
        }
    }
}

func GetConfigPath() string {
    home, err := os.UserHomeDir()
    if err != nil {
        panic(err)
    }

    return filepath.Join(home, ".config", "gommit", "settings.yaml")
}