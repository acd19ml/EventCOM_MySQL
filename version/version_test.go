package version_test

import (
	"fmt"
	"testing"

	"github.com/acd19ml/EventCOM_MySQL/version"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
