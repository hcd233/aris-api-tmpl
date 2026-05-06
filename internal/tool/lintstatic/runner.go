// Package lintstatic runs static analysis commands for this repository.
package lintstatic

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	allPackagesPattern = "./..."
	goCommand          = "go"
	goVetCommand       = "vet"
	golangciCommand    = "golangci-lint"
)

// Result describes static analysis execution output.
type Result struct {
	Output string
	Err    error
}

// Run executes go vet and golangci-lint when golangci-lint is available.
func Run(args []string) Result {
	targets := normalizeTargets(args)
	var out strings.Builder
	var hasErr bool

	vetOut, vetErr := exec.Command(goCommand, append([]string{goVetCommand}, targets...)...).CombinedOutput()
	appendOutput(&out, vetOut)
	if vetErr != nil {
		hasErr = true
	}

	lintOut, lintErr := runGolangCILint(targets)
	appendOutput(&out, lintOut)
	if lintErr != nil {
		hasErr = true
	}

	result := Result{Output: out.String()}
	if hasErr {
		result.Err = fmt.Errorf("static checks failed")
	}
	return result
}

// Write writes lint output to the provided writer.
func (r Result) Write(w io.Writer) {
	if strings.TrimSpace(r.Output) == "" {
		_, _ = io.WriteString(w, "[lintstatic] all static checks passed\n")
		return
	}
	_, _ = io.WriteString(w, r.Output)
	if !strings.HasSuffix(r.Output, "\n") {
		_, _ = io.WriteString(w, "\n")
	}
}

func normalizeTargets(args []string) []string {
	if len(args) == 0 {
		return []string{allPackagesPattern}
	}
	return args
}

func appendOutput(out *strings.Builder, data []byte) {
	if len(data) == 0 {
		return
	}
	out.Write(data)
	if !strings.HasSuffix(out.String(), "\n") {
		out.WriteByte('\n')
	}
}

func runGolangCILint(targets []string) ([]byte, error) {
	command := os.Getenv("GOLANGCI_LINT")
	if command == "" {
		commandPath, lookErr := exec.LookPath(golangciCommand)
		if lookErr != nil {
			return []byte("[lintstatic] golangci-lint not found, skipping. Install with: go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest\n"), nil
		}
		command = commandPath
	}
	return exec.Command(command, append([]string{"run"}, targets...)...).CombinedOutput()
}
