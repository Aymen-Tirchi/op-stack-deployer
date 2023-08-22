package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func generateKey(role string) (address, privateKey string, err error) {
	cmd := exec.Command("cast", "wallet", "new")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Address:") {
			address = strings.TrimSpace(strings.TrimPrefix(line, "Address:"))
		} else if strings.HasPrefix(line, "Private key:") {
			privateKey = strings.TrimSpace(strings.TrimPrefix(line, "Private key:"))
		}
	}
	return address, privateKey, nil
}

func main() {
	if _, err := os.Stat("keys.txt"); os.IsNotExist(err) {
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
		log.Printf("keys have been generated and saved to %s\n", outputFilePath)
	} else {
		log.Println("keys.txt already exists")
	}

	os.Chdir("optimism/packages/contracts-bedrock")
	if _, err := os.Stat(".envrc"); os.IsNotExist(err) {
        log.Println("Copying the .envrc file...")
        var cpOutput bytes.Buffer
        cpCmd := exec.Command("cp", ".envrc.example", ".envrc")
        cpCmd.Stderr = &cpOutput // Capture the standard error output

        err := cpCmd.Run()
        if err != nil {
            log.Fatalf("Error copying the .envrc file: %s", cpOutput.String()) // Print captured error output
        }
    } else {
		log.Println(".envrc already exists")
	}

	fmt.Println(".envrc copied successfully!")
}
