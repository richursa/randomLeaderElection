package main

import (
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/leaderElection/")
	currentTime := time.Now().Unix()
	time.Sleep(time.Duration((20 - (currentTime)%20)) * time.Second)
	for true {
		initiateNewBLock()
		currentTime = time.Now().Unix()
		time.Sleep(time.Duration((20 - (currentTime)%20)) * time.Second)
	}
}
func initiateNewBLock() {

}
