package main

import (
	"log"
	"os"
	"os/exec"
)

func runCommand(cmdStr string){
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("command execution error: ", err)
	}
}

func main() {
	log.Println("entering the op-batcher directory... ")
	err := os.Chdir("optimism/op-batcher")
	if err != nil {
		log.Fatal("error entering the op-batcher directory: ", err)
	}

	log.Println("running op-batcher")
	runCmd := `./bin/op-batcher \
    --l2-eth-rpc=http://localhost:8545 \
    --rollup-rpc=http://localhost:8547 \
    --poll-interval=1s \
    --sub-safety-margin=6 \
    --num-confirmations=1 \
    --safe-abort-nonce-too-low-count=3 \
    --resubmission-timeout=30s \
    --rpc.addr=0.0.0.0 \
    --rpc.port=8548 \
    --rpc.enable-admin \
    --max-channel-duration=1 \
    --l1-eth-rpc=$L1_RPC \
    --private-key=$BATCHER_KEY`

	runCommand(runCmd)

	log.Println("op-batcher running successfully! ")

}