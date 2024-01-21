package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"syscall"

	"github.com/Abhinandan-Khurana/EncryptGuard/filecrypt"
	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt tp decrypt a file.")
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("")
	fmt.Println("File Encryption")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t go run . encrypt /path/tp/your.file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\tEncrypt:\tEncrypt a file given a password")
	fmt.Println("\tDecrypt:\tDecrypt a file using a password")
	fmt.Println("\tHelp:\t\tDisplays help description")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run --> go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found!")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n File successfully protected!")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run --> go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found!")
	}
	fmt.Println("Enter Password:")
	password, _ := readPassword()
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n File successfully decrypted!")
}

func getPassword() []byte {
	fmt.Print("Enter Password:")
	password, err := readPassword()
	if err != nil {
		fmt.Println("\nFailed to read password. Error:", err)
		os.Exit(1)
	}
	fmt.Print("\nConfirm Password: ")
	password2, err := readPassword()
	if err != nil {
		fmt.Println("\nFailed to read password. Error:", err)
		os.Exit(1)
	}
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match, try again!\n")
		return getPassword()
	}
	return password
}

func readPassword() ([]byte, error) {
	if runtime.GOOS == "windows" {
		return term.ReadPassword(int(windows.Handle(os.Stdin.Fd())))
	}
	return term.ReadPassword(int(syscall.Stdin))
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
