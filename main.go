package main

import (
	"fmt"
	"os"
	"os/user"
	"os/exec"
	"os/signal"
	"syscall"
	"strings"
	. "github.com/logrusorgru/aurora"
	"github.com/Songmu/prompter"
	"path/filepath"
	"bufio"
)

func main() {
	exitHandle()
	reader := bufio.NewReader(os.Stdin)
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	shorthand := "/home/" + strings.ToLower(user.Name) + "/"

	hostname, _ := os.Hostname()

	for {
		path, _ := filepath.Abs(".")
		path = strings.Replace(path, shorthand, "~/", -1)

		fmt.Printf("%v%s%s%s%s%s", Bold(Green(strings.ToLower(user.Name))), Bold(Green("@")), Bold(Green(hostname)), ":", Bold(Blue(path)), "$ ")

		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)

		if command == "oops" {
			os.Exit(0)
		} else if strings.Contains(command, "sudo") {
			sudoPrompt := "[sudo] password for " + strings.ToLower(user.Name)
			password := prompter.Password(sudoPrompt)
			fmt.Println("Your password is: " + password)
		} else {
			cmd, err := exec.Command("bash", "-c", command).Output()
			if err != nil {
				fmt.Printf("bash: %s: command not found\n", command)
			}
			fmt.Println(string(cmd))
		}

	}

}

//Prevents user from using CTRL+C to exit the program
func exitHandle() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Nice Try...") //Temporary
	}()
}
