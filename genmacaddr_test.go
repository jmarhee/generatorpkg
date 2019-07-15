package generatorpkg

import (
	generatorpkg "gitlab.com/jmarhee/generatorpkg"
	"log"
	"testing"
	"os"
)

func TestMac(t *testing.T) {
	mac, err := generatorpkg.NewMacAddr()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(mac)
}
