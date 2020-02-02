package main

import (
	"time"
)

func main() {
	currentTime := time.Now().Unix()
	time.Sleep(time.Duration((600 - (currentTime)%600)) * time.Second)
	for true {
		initiateNewBLock()
		currentTime = time.Now().Unix()
		time.Sleep(time.Duration((600 - (currentTime)%600)) * time.Second)
	}
}
func initiateNewBLock() {
	node_id := []string{"cusat", "thankappan"}

}
