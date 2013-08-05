package gci

import (
	"testing"
	"fmt"
)

func TestCubridVersion(t *testing.T) {
	var version string
	gci_get_version_string(&version)
	fmt.Println(version)
}

