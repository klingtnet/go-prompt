package main

import (
	"errors"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/klingtnet/termenv"
	"gopkg.in/src-d/go-git.v4"
)

func mustUser() string {
	u, err := user.Current()
	if err != nil {
		return "unknown-user"
	}
	return u.Username
}

func mustHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return hostname
}

func mustWd() string {
	wd, err := os.Getwd()
	if err != nil {
		return "unknown-dir"
	}
	return filepath.Clean(wd)
}

func shortenPath(wd string) string {
	home, err := os.UserHomeDir()
	if err == nil {
		if strings.HasPrefix(wd, home) {
			return strings.Replace(wd, home, "~", 1)
		}
	}
	return wd
}

func findGitRepo(wd string) *git.Repository {
	l := strings.Split(wd, string(os.PathSeparator))
	for idx := len(l); idx > 0; idx-- {
		path := "/" + filepath.Join(l[:idx]...)

		repo, err := git.PlainOpen(path)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				continue
			}
			log.Println(err)
			return nil
		}
		return repo
	}
	return nil
}

func gitInfo(wd string) string {
	repo := findGitRepo(wd)
	if repo == nil {
		return ""
	}
	ref, err := repo.Head()
	if err != nil {
		return ""
	}
	s := strings.Split(ref.Name().Short(), "/")
	return "(" + s[len(s)-1] + ")"
}

const prompt = "$"

type colorist struct {
	profile termenv.Profile
}

func newColorist() *colorist {
	return &colorist{
		profile: termenv.ColorProfile(),
	}
}

func (c *colorist) colored(s, color string) string {
	return termenv.String(s).Foreground(c.profile.Color(color)).String()
}

type field struct {
	value string
	color string
}

func main() {
	var statusCode string
	if os.Args[1] != "0" {
		statusCode = os.Args[1]
	}

	termenv.ForceColor = true
	c := newColorist()

	wd := mustWd()
	fields := []field{
		{mustUser(), "#ff0000"},
		{"@", "#ff5f00"},
		{mustHostname(), "#ff8700"},
		{shortenPath(wd), "#ffaf00"},
		{gitInfo(wd), "#ffd700"},
		{statusCode, "#ff00ff"},
		{prompt, "#ffff00"},
	}
	line := ""
	for _, f := range fields {
		line += c.colored(f.value, f.color)
	}
	line += " "
	io.WriteString(os.Stdout, line)
}
