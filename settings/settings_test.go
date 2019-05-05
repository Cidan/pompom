package settings

import "testing"

func TestSetup(t *testing.T) {
	Setup("./fixtures")
}

func Test_setDefaults(t *testing.T) {
	setDefaults()
}
