package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <option>")
		fmt.Println("Options:")
		fmt.Println("  (no option) - Use SHA256")
		fmt.Println("  -b          - Use SHA384")
		fmt.Println("  -c          - Use SHA512")
		return
	}

	arg := os.Args[1]
	var data string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		data += sc.Text() + "\n"
	}

	if sc.Err() != nil {
		fmt.Println("Error reading input:", sc.Err())
		return
	}

	switch arg {
	case "":
		hash := sha256.Sum256([]byte(data))
		fmt.Printf("SHA256 sum: %x\n", hash)
	case "-b":
		hash := sha512.Sum384([]byte(data))
		fmt.Printf("SHA384 sum: %x\n", hash)
	case "-c":
		hash := sha512.Sum512([]byte(data))
		fmt.Printf("SHA512 sum: %x\n", hash)
	default:
		fmt.Println("Invalid option. Use -b for SHA384, -c for SHA512, or no option for SHA256.")
	}
}
