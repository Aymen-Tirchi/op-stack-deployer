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
	log.Println("entering the op-node directory... ")
	err := os.Chdir("optimism/op-node")
	if err != nil {
		log.Fatal("error entering the op-node directory: ", err)
	}

	log.Println("running op-node... ")
	runCmd := `./bin/op-node \
	--l2=http://localhost:8551 \
	--l2.jwt-secret=./jwt.txt \
	--sequencer.enabled \
	--sequencer.l1-confs=3 \
	--verifier.l1-confs=3 \
	--rollup.config=./rollup.json \
	--rpc.addr=0.0.0.0 \
	--rpc.port=8547 \
	--p2p.disable \
	--rpc.enable-admin \
	--p2p.sequencer.key=$SEQ_KEY \
	--l1=$L1_RPC \
	--l1.rpckind=$RPC_KIND`

	runCommand(runCmd)
	log.Println("op-node running successfully! ")
}