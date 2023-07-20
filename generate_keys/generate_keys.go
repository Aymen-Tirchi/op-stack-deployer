package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func generateKey(role string) (address, privateKey string, err error){
	cmd := exec.Command("cast", "wallet", "new")
	output, err := cmd.Output()
	if err != nil{
		return "", "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Address:"){
			address = strings.TrimSpace(strings.TrimPrefix(line, "Address:"))
		} else if strings.HasPrefix(line, "Private key:"){
			privateKey = strings.TrimSpace(strings.TrimPrefix(line, "Private key:"))
		}
	}
	return address, privateKey, nil
}

func main() {
	log.Println("entering the contracts-bedrock directory... ")
	os.Chdir("/optimism/packages/contracts-bedrock")

	log.Println("generating keys... ")
	accounts := []string{"Admin", "Proposer", "Batcher", "Sequencer"}
	var keyData strings.Builder

	for _, role := range accounts {
		address, privateKey, err := generateKey(role)
		if err != nil {
			log.Fatalf("error generating %s keys: %v", role, err)
		}
		fmt.Fprintf(&keyData, "%s: \nAddress: %s\nPrivate Key: %s\n\n", role, address, privateKey)
	}

	outputFilePath := "keys.txt"
	err := os.WriteFile(outputFilePath, []byte(keyData.String()), 0644)
	if err != nil {
		log.Fatal("error writing keys to file: ", err)
	}
	fmt.Printf("keys have been generated and saved to %s\n", outputFilePath)
}