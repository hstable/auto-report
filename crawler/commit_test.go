package crawler

import (
	"log"
	"testing"
)

func TestCommit(t *testing.T) {
	err := Commit("20S051030", "lgj147258369.")
	log.Println(err)
}