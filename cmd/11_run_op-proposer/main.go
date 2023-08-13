package main

import (
	"log"
	"os"
	"os/exec"
)

func runCommand(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("command execution error: ", err)
	}
}

func main() {
	log.Println("entering the op-proposer directory... ")
	err := os.Chdir("optimism/op-proposer")
	if err != nil {
		log.Fatal("error entering the op-proposer directory: ", err)
	}

	runCmd := `./bin/op-proposer \
    --poll-interval=12s \
    --rpc.port=8560 \
    --rollup-rpc=http://localhost:8547 \
    --l2oo-address=$L2OO_ADDR \
    --private-key=$PROPOSER_KEY \
    --l1-eth-rpc=$L1_RPC`

	runCommand(runCmd)
	log.Println("op-proposer running successfully!")
}