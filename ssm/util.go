package ssm

import (
	"os"
	"strings"

	"github.com/buildkite/interpolate"
)

var env Env

// difference returns the elements in a that aren't in b
// the second argument is slice of string pointers to suit AWS SDK
func stringSliceDifference(a, b []string) []string {
	mb := map[string]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []string{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func ExpandArgs(args []string) []string {
	var expanded []string
	for _, arg := range args {
		arg = ExpandValue(arg)
		expanded = append(expanded, arg)
	}
	return expanded
}
func ExpandValue(val string) string {
	e, err := interpolate.Interpolate(env, val)
	if err == nil {
		return strings.TrimSpace(string(e))
	}
	return val

}

// just adapt os.LookupEnv to this interface
type Env struct{}

func (e Env) Get(key string) (string, bool) {
	return os.LookupEnv(key)
}
