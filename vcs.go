package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type vcs struct {
	name string
	cmd  string // name of binary to invoke command

	createCmd      string // command to download a fresh copy of a repository
	downloadCmd    string // command to download updates into an existing repository
	BaseBranchName string // the name of the default branch/revision to get when non is given
}

func (v *vcs) create(dir, repo, branch string) error {
	return v.run(".", v.createCmd, "dir", dir, "repo", repo, "branch", branch)
}

// download downloads any new changes for the repo in dir.
func (v *vcs) download(dir string) error {
	return v.run(dir, v.downloadCmd)
}

var vcsGit = &vcs{
	name: "Git",
	cmd:  "git",

	createCmd:      "clone {repo} {dir} -b {branch}",
	downloadCmd:    "pull --ff-only",
	BaseBranchName: "master",
}

var vcsSvn = &vcs{
	name: "Subversion",
	cmd:  "svn",

	createCmd:      "checkout {repo} {dir} -r {branch}",
	downloadCmd:    "update",
	BaseBranchName: "HEAD",
}

func (v *vcs) run(dir string, cmd string, keyval ...string) error {
	_, err := v.run1(dir, cmd, keyval, true)
	return err
}

// runVerboseOnly is like run but only generates error output to standard error in verbose mode.
func (v *vcs) runVerboseOnly(dir string, cmd string, keyval ...string) error {
	_, err := v.run1(dir, cmd, keyval, false)
	return err
}

// runOutput is like run but returns the output of the command.
func (v *vcs) runOutput(dir string, cmd string, keyval ...string) ([]byte, error) {
	return v.run1(dir, cmd, keyval, true)
}

// run1 is the generalized implementation of run and runOutput.
func (v *vcs) run1(dir string, cmdline string, keyval []string, verbose bool) ([]byte, error) {
	m := make(map[string]string)
	for i := 0; i < len(keyval); i += 2 {
		m[keyval[i]] = keyval[i+1]
	}
	args := strings.Fields(cmdline)
	for i, arg := range args {
		args[i] = expand(m, arg)
	}

	_, err := exec.LookPath(v.cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"go: missing %s command. See http://golang.org/s/gogetcmd\n",
			v.name)
		return nil, err
	}

	cmd := exec.Command(v.cmd, args...)
	cmd.Dir = dir
	//cmd.Env = envForDir(cmd.Dir)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err = cmd.Run()
	out := buf.Bytes()
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "# cd %s; %s %s\n", dir, v.cmd, strings.Join(args, " "))
			os.Stderr.Write(out)
		}
		return nil, err
	}
	return out, nil
}

func expand(match map[string]string, s string) string {
	for k, v := range match {
		s = strings.Replace(s, "{"+k+"}", v, -1)
	}
	return s
}
