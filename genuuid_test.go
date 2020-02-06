package generatorpkg

import (
	generatorpkg "gitlab.com/jmarhee/generatorpkg"
	"log"
	"testing"
)

func TestUuid(t *testing.T) {
	uuid := generatorpkg.NewUuid()
	log.Println(uuid)
}
