package generatorpkg

import (
	"fmt"
	"math/rand"
	"log"
)

type mac struct {
	MACAddr string
}

func NewMacAddr() interface{} {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
	    log.Fatal(err)
	}
	newMac := fmt.Sprintf("%x:%x:%x:%x:%x:%x",
	    b[0], b[1], b[2], b[3], b[4], b[5])
	newMacAddr := mac{MACAddr: newMac}
	return _,newMacAddr.MACAddr
}
