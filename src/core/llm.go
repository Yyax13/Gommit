package core

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"
    "yyax13/gommit/src/utils"
    "yyax13/gommit/src/config"

    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

var (
    ctx    = context.Background()
    client *genai.Client
    model  *genai.GenerativeModel
)

func init() {
    var err error

    var confPath string = config.GetConfigPath()
    config.EnsureConfig(confPath)

    cfg, err := config.LoadConfig(confPath)
    if cfg.GeminiApiKey == "" {
	    log.Fatal(utils.Red("GeminiApiKey not set"))
    
    }
    client, err = genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiApiKey))
    if err != nil {
        log.Fatal(utils.Red(err.Error()))
    }

    model = client.GenerativeModel("models/gemini-2.5-flash")
}

var prompt string = `
You are generating a git commit message.

Follow STRICTLY this commit convention:

Format:
<type>(<scope>): <past-tense message>

Rules:
- The message MUST start with a past-tense verb.
- The scope MUST NOT be empty.
- The message must describe the change clearly and concisely.
- Do NOT add explanations, only output the commit message.
- Use the git diff as the source of truth.

Allowed types:
feat  → New feature
fix   → Bug fix
ref   → Refactor without behavior change
docs  → Documentation changes
chore → Build / tooling / maintenance
test  → Tests
perf  → Performance improvements

Examples:

feat(encoders/xor): Added settings option to use different hash algorithms
feat(app/main): Added tests for the new hashing algorithms
feat(utils/hash): Created custom hashing
fix(utils/random): Fixed modulo bias in randr using limit and rejecting trash values
ref(cli/loading): Deleted useless cli/loading modules
chore(make): Added new source files to SRC_ENTRYPOINT
docs(cli/logs): Updated documentation for bytesf
`

var promptDelimDesc string = `
If a valid history is provided, you can get context based in it

The diff is between === START OF DIFF === and ==== END OF DIFF =====
The history is between === START OF HIST === and ==== END OF HIST ====
`

func GetCommitMessage(out *string, diff string, hist string, wg *sync.WaitGroup, cfg *utils.Config) {
    defer wg.Done()
    var basePrompt string
    if cfg.OverWriteDefaultCommitPatternPrompt {
	   	basePrompt = fmt.Sprintf(`
%s
%s
		`, cfg.CommitPatternPrompt, promptDelimDesc)
					
    } else {
	   	basePrompt = fmt.Sprintf(`
%s
%s
		`, prompt, promptDelimDesc)
					
    }

    var finalPrompt string
    if cfg.UseHist {
    finalPrompt = fmt.Sprintf(`
%s

=== START OF DIFF ===
%s
==== END OF DIFF ====

=== START OF HIST === 
%s
==== END OF HIST ====
	`, basePrompt, diff, hist)
					
    } else {
   		finalPrompt = fmt.Sprintf(`
%s

=== START OF DIFF ===
%s
==== END OF DIFF ====
		`, basePrompt, diff)
    
    }
    

    msg, err := runLLM(finalPrompt)

    if err != nil {
		fmt.Printf("Can't get: %s\n", err.Error())
        log.Fatal(utils.Red("Can't generate the commit message, aborting..."))
        os.Exit(1)
    }

    *out = msg
}

func runLLM(prompt string) (string, error) {
    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return "", err
    }

    var output strings.Builder

    for _, cand := range resp.Candidates {
        for _, part := range cand.Content.Parts {
            if txt, ok := part.(genai.Text); ok {
                output.WriteString(string(txt))
            }
        }
    }

    return strings.TrimSpace(output.String()), nil
}
