package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runCommand(cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

var GREETER = "575E9B4f2c3945d7CF07cb76628d29DF471692B8"

func main() {
	if _, err := os.Stat("optimism-tutorial"); os.IsNotExist(err) {
		log.Println("cloning optimism-tutorial repo... ")
		cloneCmd := `git clone https://github.com/ethereum-optimism/optimism-tutorial.git`
		err := runCommand(cloneCmd)
		if err != nil {
			log.Fatal("error cloning the optimism-tutorial repo: ", err)
		}
	} else {
		log.Println("optimism-tutorial is already cloned")
	}

	log.Println("entering the getting-started/foundry directory... ")
	err := os.Chdir("optimism-tutorial/getting-started/foundry")
	if err != nil {
		log.Fatal("error entering the getting-started/foundry directory: ", err)
	}

	if _, err := os.Stat("mnem.delme"); os.IsNotExist(err) {
		log.Println("creating the mnem.delme file... ")
		fmt.Println("Enter your mnemonic phrase here :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		mnem := scanner.Text()
		filepath := "mnem.delme"
		data := []byte(mnem)
		err = os.WriteFile(filepath, data, 0644)
		if err != nil {
			log.Fatal("error writing the mnemonic to file: ", err)
		}
		log.Printf("Mnemonic phrase has been written to %s\n", filepath)
	} else {
		log.Println("the mnem.delme file already exists")
	}

	// // Run forge config --fix to address the foundry.toml warning
	// forgeConfigFix := "forge config --fix"
	// err = runCommand(forgeConfigFix)
	// if err != nil {
	// 	log.Fatal("error fixing forge config: ", err)
	// }

	// log.Println("compiling greeter...")
	// greetCompile := `forge create --mnemonic-path ./mnem.delme Greeter --constructor-args "hello" --gas-limit 5000000 | tee deployment`
	// err = runCommand(greetCompile)
	// if err != nil {
	// 	log.Fatal("error compiling greeter: ", err)
	// }
}


	// log.Println("calling greet()... ")
	// greetCall := `cast call 0x` + GREETER + ` "greet()"`
	// err = runCommand(greetCall)
	// if err != nil {
	// 	log.Fatal("error calling greet(): ", err)
	// }

	// log.Println("calling greet() translated to ASCII... ")
	// greetCallAscii := `cast call 0x` + GREETER + ` "greet()" | cast --to-ascii`
	// err = runCommand(greetCallAscii)
	// if err != nil {
	// 	log.Fatal("error calling greet() translated to ASCII: ", err)
	// }

	// log.Println("sending transaction... ")
	// sendTx := `cast send --mnemonic-path mnem.delme 0x` + GREETER + ` "setGreeting(string)" "Foundry hello" --legacy `
	// err = runCommand(sendTx)
	// if err != nil {
	// 	log.Fatal("error sending the transaction: ", err)
	// }

	// log.Println("Test that the greeting has changed... ")
	// greetTest := `cast call $GREETER "greet()" | cast --to-ascii`
	// err = runCommand(greetTest)
	// if err != nil {
	// 	log.Fatal("error testing: ", err)
	// }

	// log.Println("deploying the contract... ")
	// greetDeploy := `forge create --mnemonic-path ./mnem.delme Greeter \
	// --constructor-args "Greeter from Foundry" --legacy`
	// err = runCommand(greetDeploy)
	// if err != nil {
	// 	log.Fatal("error deploying the contract: ", err)
	// }