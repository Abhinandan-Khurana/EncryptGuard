package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Abhinandan-Khurana/EncryptGuard/filecrypt"
	"github.com/fatih/color"
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
	magenta := color.New(color.FgMagenta).Add(color.Bold)
	blue := color.New(color.FgCyan).Add(color.Bold)
	green := color.New(color.FgGreen).Add(color.Bold)
	yellow := color.New(color.FgYellow).Add(color.Bold)
	fmt.Println("")
	magenta.Println(`
____ _  _ ____ ____ _   _ ___  ___ ____ _  _ ____ ____ ___
|___ |\ | |    |__/  \_/  |__]  |  | __ |  | |__| |__/ |  \
|___ | \| |___ |  \   |   |     |  |__] |__| |  | |  \ |__/
`)
	fmt.Println("")
	blue.Println("Usage:")
	fmt.Println("")
	yellow.Println("\t EncryptGuard encrypt /path/to/your.file")
	fmt.Println("")
	green.Println("Commands:")
	fmt.Println("")
	fmt.Println("\tencrypt:\tEncrypt a file given a password")
	fmt.Println("\tdecrypt:\tDecrypt a file using a password")
	fmt.Println("\thelp:\t\tDisplays help description")
	fmt.Println("")
}

func encryptHandle() {
	magenta := color.New(color.FgMagenta).Add(color.Bold)
	green := color.New(color.FgGreen).Add(color.Bold)

	if len(os.Args) < 3 {
		println("missing the path to the file. For more info, run --> go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found!")
	}
	password := getPassword()
	magenta.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	green.Println("\n File successfully protected!")
}

func decryptHandle() {
	green := color.New(color.FgGreen).Add(color.Bold)
	magenta := color.New(color.FgMagenta).Add(color.Bold)
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
	magenta.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	green.Println("\n File successfully decrypted!")
}

func getPassword() []byte {
	yellow := color.New(color.FgYellow).Add(color.Bold)
	red := color.New(color.FgRed).Add(color.Bold)

	fmt.Print("Set Password:")
	password, err := readPassword()
	if err != nil {
		red.Println("\nFailed to read password. Error:", err)
		os.Exit(1)
	}
	fmt.Print("\nConfirm Password: ")
	password2, err := readPassword()
	if err != nil {
		red.Println("\nFailed to read password. Error:", err)
		os.Exit(1)
	}
	if !validatePassword(password, password2) {
		yellow.Print("\nPasswords do not match, try again!\n")
		return getPassword()
	}
	return password
}

// func readPassword() ([]byte, error) {
// 	if runtime.GOOS == "windows" {
// 		return term.ReadPassword(int(windows.Handle(os.Stdin.Fd())))
// 	}
// 	return term.ReadPassword(int(syscall.Stdin))
// }

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
