package generatorpkg 

import (
	"fmt"
	"math/rand"
	"log"
)

type uuid struct {
	UUID string
}

func NewUuid() interface{} {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
	    log.Fatal(err)
	}
	newUuid := fmt.Sprintf("%x-%x-%x-%x-%x",
	    b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	id := uuid{UUID: newUuid}
	return id.UUID
}