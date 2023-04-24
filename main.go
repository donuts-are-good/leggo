package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	devname     = "donuts-are-good"
	projectpath = filepath.Join(os.Getenv("HOME"), "Projects")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <project name>")
		os.Exit(1)
	}
	name := os.Args[1]
	dir := filepath.Join(projectpath, name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create project directory %s: %v\n", dir, err)
		os.Exit(1)
	}
	if err := setup(name, dir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup project %s: %v\n", name, err)
		os.Exit(1)
	}
}

func setup(n, d string) error {
	if err := createFile(filepath.Join(d, "main.go"), "package main;\n\n func main() { \n\tprintln(\"alive!\")\n}"); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	if err := createFile(filepath.Join(d, "README.md"), getReadmeTxt(n)); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	if err := createFile(filepath.Join(d, "LICENSE.md"), licensetext); err != nil {
		return fmt.Errorf("failed to create LICENSE.md: %w", err)
	}
	if err := createFile(filepath.Join(d, ".gitignore"), fmt.Sprintf(".DS_Store\n.Trash-1000\n%s\nBUILDS", n)); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	if err := execCmd("go", "mod", "init", "github.com/"+devname+"/"+n, d); err != nil {
		return fmt.Errorf("failed to initialize go modules: %w", err)
	}
	if err := execCmd("go", "mod", "tidy", d); err != nil {
		return fmt.Errorf("failed to tidy go modules: %w", err)
	}
	if err := execCmd("git", "init", d); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}
	if err := execCmd("git", "remote", "add", "origin", "https://github.com/"+devname+"/"+n+".git", d); err != nil {
		return fmt.Errorf("failed to add git remote: %w", err)
	}
	if err := execCmd("git", "add", "-A", ".", ".gitignore", d); err != nil {
		return fmt.Errorf("failed to add files to git repository: %w", err)
	}
	if err := execCmd("code", ".", d); err != nil {
		return fmt.Errorf("failed to open project in Visual Studio Code: %w", err)
	}
	return nil
}

func execCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = arg[0]
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command %q: %w", cmd.String(), err)
	}
	return nil
}

func createFile(p, c string) error {
	f, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", p, err)
	}
	defer f.Close()
	if _, err := f.WriteString(c); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", p, err)
	}
	return nil
}
