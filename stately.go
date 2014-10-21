// this is the libarary for stately

package stately

//import "fmt"

const branch = "v0.1.0"

func Version() string {
	return branch
}

type source struct {
	Target     string
	Url        string
	SourceType *vcs
}

type Config struct {
	Name    string
	Sources []source
}

func (c *Config) Get() {
	for _, source := range c.Sources {

	}
}

func (s *source) set_type() {
	s.SourceType = vcsGit
}
