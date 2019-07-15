package generatorpkg

import (
	generatorpkg "gitlab.com/jmarhee/generatorpkg"
	"log"
	"testing"
)

func TestUuid(t *testing.T) {
	mac, err := generatorpkg.NewUuid()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(mac)
}
