package main

import (
	"net/http"
	"time"

	"./nodelist"
	"github.com/gorilla/mux"
)

var nodeCommitmentMap map[string]string
var nodeRevealMap map[string]string

func main() {

	nodeCommitmentMap = make(map[string]string)

	nodeRevealMap = make(map[string]string)
	blankMap(nodeCommitmentMap, nodeRevealMap)
	router := mux.NewRouter()
	router.HandleFunc("/api/leaderElection/commit/{nodeID}/{hash}", commitmentHandler)
	router.HandleFunc("/api/leaderElection/reveal/{nodeID}/{key}", revealHandler)
	go http.ListenAndServe(":9090", router)
	currentTime := time.Now().Unix()
	time.Sleep(time.Duration((60 - (currentTime)%60)) * time.Second)
	for true {
		initiateNewBLock(nodeCommitmentMap, nodeRevealMap)
		currentTime = time.Now().Unix()
		time.Sleep(time.Duration((60 - (currentTime)%60)) * time.Second)
	}
}
func blankMap(nodeCommitmentMap map[string]string, nodeRevealMap map[string]string) {
	for i := 0; i < len(nodelist.Nodelist); i++ {
		nodeCommitmentMap[nodelist.Nodelist[i]] = ""
		nodeRevealMap[nodelist.Nodelist[i]] = ""
	}
}
func commitmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeCommitmentMap[vars["nodeID"]] = vars["hash"]
}

func revealHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeRevealMap[vars["nodeID"]] = vars["key"]
}
func initiateNewBLock(nodeCommitmentMap map[string]string, nodeRevealMap map[string]string) {
	//update stuff here
}
