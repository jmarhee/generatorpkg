package generatorpkg

import (
	generatorpkg "gitlab.com/jmarhee/generatorpkg"
	"log"
	"testing"
)

func TestMac(t *testing.T) {
	mac := generatorpkg.NewMacAddr()
	log.Println(mac)
}
