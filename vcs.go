package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type vcs struct {
	name string
	cmd  string // name of binary to invoke command

	createCmd           string // command to download a fresh copy of a repository
	downloadCmd         string // command to download updates into an existing repository
	BaseBranchName      string // the name of the default branch/revision to get when non is given
	matchRegexUrl       string // a regex to match a url for the vcs type
	vcsTypeDirMatchFunc func(path string) (b bool, err error)
	urlGetFunc          func(path string, self *vcs) string
	branchGetFunc       func(path string, self *vcs) string
}

func (v *vcs) create(dir, repo, branch string) error {
	return v.run(".", v.createCmd, "dir", dir, "repo", repo, "branch", branch)
}

// download downloads any new changes for the repo in dir.
func (v *vcs) download(dir string) error {
	return v.run(dir, v.downloadCmd)
}

func git_dir_type(path string) (b bool, err error) {
	path = fmt.Sprintf("%s%s.git", path, os.PathSeparator)
	b, err = exists(path)
	return b, err
}

var vcsGit = &vcs{
	name: "Git",
	cmd:  "git",

	createCmd:      "clone {repo} {dir} -b {branch}",
	downloadCmd:    "pull --ff-only",
	BaseBranchName: "master",
	vcsTypeDirMatchFunc: func(path string) (b bool, err error) {
		path = filepath.Join(path, ".git")
		b, err = exists(path)
		return b, err
	},
	matchRegexUrl: "(\\w+://){0,1}(\\w+@)([\\w\\d\\.]+)(:[\\d]+){0,1}/*(.*)",
	urlGetFunc: func(path string, self *vcs) string {
		r, _ := regexp.Compile("(\\w+://){0,1}(\\w+@)([\\w\\d\\.]+)(:[\\d]+){0,1}/*(.*)")
		output, _ := self.runOutput(path, "remote -v")
		s := string(output)
		url := strings.Split(r.FindString(s), " ")
		//fmt.Println(res[0])
		//fmt.Println(s)
		return url[0]
	},
	branchGetFunc: func(path string, self *vcs) string {
		output, _ := self.runOutput(path, "status")
		lines := strings.Split(string(output), "\n")
		branch := strings.Split(lines[0], " ")
		//fmt.Println(res[0])
		//fmt.Println(s)
		return branch[2]
	},
}

var vcsSvn = &vcs{
	name: "Subversion",
	cmd:  "svn",

	createCmd:      "checkout {repo} {dir} -r {branch}",
	downloadCmd:    "update",
	BaseBranchName: "HEAD",
	matchRegexUrl:  "svn(\\+.*)?://",
	vcsTypeDirMatchFunc: func(path string) (b bool, err error) {
		path = filepath.Join(path, ".svn")
		//fmt.Println("svn:", path)
		b, err = exists(path)
		return b, err
	},
	urlGetFunc: func(path string, self *vcs) string {
		r, _ := regexp.Compile("URL:\\s(.*)")
		output, _ := self.runOutput(path, "info")
		s := string(output)
		url := strings.Replace(r.FindString(s), "URL: ", "", -1)
		//fmt.Println(res[0])
		fmt.Println(url)
		return url
	},
	branchGetFunc: func(path string, self *vcs) string {
		// here should be a function to return the revision number but i'm not sure of its usfullness
		return ""
	},
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
			"stately: missing %s command.\n",
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

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func expand(match map[string]string, s string) string {
	for k, v := range match {
		s = strings.Replace(s, "{"+k+"}", v, -1)
	}
	return s
}

var known_types = []*vcs{
	vcsSvn,
	vcsGit,
}

func (s *source) get_url() string {
	if s.SourceType == nil {
		return ""
	}
	// here get the url of the source
	return ""
}
func (s *source) set_type() error {
	err := errors.New("Error:Error cannot set type")
	if s.Url != "" {
		//fmt.Println("Using url")
		s.SourceType = get_type_by_url(s.Url)
		if s.SourceType == nil {
			fmt.Println("Cannot choose source type 1")
			return err
		}
	} else if s.Target != "" {
		//fmt.Println("Using Target")
		s.SourceType = get_type_by_dir(s.Target)
		if s.SourceType == nil {
			fmt.Println("Cannot choose source type 2")
			return err
		}
	} else {
		fmt.Println("# cannot check source type defaults to git")
		s.SourceType = vcsGit
	}
	return nil
}
func get_type_by_url(url string) *vcs {
	for _, v := range known_types {
		res, _ := regexp.MatchString(v.matchRegexUrl, url)
		if res {
			return v
		}
	}
	return nil
}
func get_type_by_dir(path string) *vcs {
	// will use the specific directory type that should be in the target directory already
	for _, v := range known_types {
		res, _ := v.vcsTypeDirMatchFunc(path)
		if res {
			return v
		}
	}
	return nil
}
