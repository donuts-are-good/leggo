package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	devname     = "donuts-are-good"
	projectpath = filepath.Join(os.Getenv("HOME"), "Projects")
	licensetext = `
MIT License

Copyright (c) 2023 donuts-are-good <https://github.com/donuts-are-good>
	
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`
)

func main() {

	// listen for input
	scanner := bufio.NewScanner(os.Stdin)

	// get the program name from the user
	fmt.Print("Program name: ")

	// the thing we're listening for is a program or lib name
	name := scanInput(scanner)

	// it's going in ~/Projects/... by deault

	dir := filepath.Join(projectpath, name)

	// see if it is a library or program
	fmt.Print("Library or program (l/p)? ")
	lp := scanInput(scanner)

	// make the directory for the proejct
	os.MkdirAll(dir, 0755)

	// present a summary of the actions to the user
	summary(name, dir, lp)
	if strings.ToLower(scanInput(scanner)) != "y" {
		return
	}
	setup(name, dir, lp)
}

func scanInput(s *bufio.Scanner) string { s.Scan(); return s.Text() }

func setup(n, d, lp string) {
	if lp == "l" {
		createFile(filepath.Join(d, n+".go"), "package "+n)
	} else {
		createFile(filepath.Join(d, "main.go"), "package main;\n\n func main() { \n\tprintln(\"alive!\")\n}")
	}
	createFile(filepath.Join(d, "README.md"), "# "+n+"\n")
	if lp == "p" {
		createLicense(d)
	}
	execCmd("go", "mod", "init", "github.com/"+devname+"/"+n, d)
	execCmd("go", "mod", "tidy", d)
	execCmd("git", "init", d)
	execCmd("git", "remote", "add", "origin", "https://github.com/"+devname+"/"+n+".git", d)
	execCmd("git", "add", "-A", ".", ".gitignore", d)
	createFile(filepath.Join(d, ".gitignore"), fmt.Sprintf(".DS_Store\n.Trash-1000\n%s\nBUILDS", n))
	execCmd("code", ".", d)
}

func execCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg[1:]...)
	cmd.Dir = arg[0]
	cmd.Run()
}

func createFile(p, c string) { f, _ := os.Create(p); f.WriteString(c); f.Close() }

func createLicense(d string) {

	createFile(filepath.Join(d, "LICENSE.md"), licensetext)
}

func summary(n, d, lp string) {
	fmt.Printf("Actions:\n1. Create dir: %s\n2. Init Go mod\n3. Init Git repo\n4. Create %s\n5. Create README.md\n", d, fileType(lp))
	if lp == "p" {
		fmt.Println("6. Create LICENSE.md (MIT License)")
	}
	fmt.Println("7. Open project in Visual Studio Code\n\n Press `y` and then enter if this is correct")
}

func fileType(lp string) string {
	if lp == "l" {
		return "library.go"
	}
	return "main.go"
}
