package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("[?] What is your name?")
	user.Age = readInt("[?] How old are you?")
	user.FavoriteNumber = readFloat("[?] What is your favorite number?")
	user.OwnsADog = haveADog("[?] Do you own a dog? (y/n)")
	fmt.Printf("[+] Your name is %s. You are %d years old. Your favorite number is %.2f. Owns a dog: %t.\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog,
	)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a whole value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("[!] Please enter a number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("[!] Please enter a number")
		} else {
			return num
		}
	}
}

func haveADog(s string) bool {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "y" || userInput == "Y" {
			return true
		} else if userInput == "n" || userInput == "N" {
			return false
		} else {
			fmt.Println("[!] Only y or n values are accepted")
		}
	}
}
