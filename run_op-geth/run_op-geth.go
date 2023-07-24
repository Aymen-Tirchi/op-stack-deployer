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
	log.Println("entering the op-geth directory...")
	err := os.Chdir("op-geth")
	if err != nil {
		log.Fatal("error entering the op-geth directory: ", err)
	}

	log.Println("running op-geth...")
	cmdStr := `./build/bin/geth \
			--datadir ./datadir \
			--http \
			--http.corsdomain="*" \
			--http.vhosts="*" \
			--http.addr=0.0.0.0 \
			--http.api=web3,debug,eth,txpool,net,engine \
			--ws \
			--ws.addr=0.0.0.0 \
			--ws.port=8546 \
			--ws.origins="*" \
			--ws.api=debug,eth,txpool,net,engine \
			--syncmode=full \
			--gcmode=archive \
			--nodiscover \
			--maxpeers=0 \
			--networkid=42069 \
			--authrpc.vhosts="*" \
			--authrpc.addr=0.0.0.0 \
			--authrpc.port=8551 \
			--authrpc.jwtsecret=./jwt.txt \
			--rollup.disabletxpoolgossip=true`

	runCommand(cmdStr)

	log.Println("op-geth running successfully!")
}
