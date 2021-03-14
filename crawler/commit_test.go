package crawler

import (
	"log"
	"testing"
)

func TestReport(t *testing.T) {
	isSuccess, err := Report("20S051030", "lgj147258369.")
	if isSuccess && err == nil {
		log.Println("上报成功！")
	}
}