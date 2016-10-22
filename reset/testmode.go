package reset

import (
	"os"
	"strings"
	"sync"
)

var (
	testModeOnce = sync.Once{}
	_testMode    bool
)

// TestMode returns true if run as unit test
func TestMode() bool {
	testModeOnce.Do(func() {
		_testMode = strings.HasSuffix(os.Args[0], ".test")
	})
	return _testMode
}
