package projectpath_test

import (
	"strings"
	"testing"

	"github.com/qadre/qdoc/api/pkg/projectpath"
)

func TestBase(t *testing.T) {
	project := "qdoc/api"
	if got := projectpath.Base(); strings.Index(got, project) < 1 {
		t.Errorf("wrong project path %v", got)
	}
}
