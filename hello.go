// NOTES

// MESSAGE WITH COMMAND VAR USING IF ELSE
/*
	if command == 1 {
		fmt.Println("Monitoring")
	} else if command == 2 {
		fmt.Println("System initialized and are working fine")
	} else if command == 3 {
		fmt.Println("Saindo")
	} else {
		fmt.Println("Esta opção não existe")
	}
*/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	//RETURN TWO OR MORE VAR -> MAIN CODE
	/*
		name, age := returnNameAndAge()
		fmt.Println(name, "is", age, "years old")
	*/

	sayGreeting()

	displayMenu()

	command := readVar()

	switch command {
	case 0:
		fmt.Println("Exiting")
		os.Exit(0)
	case 1:
		startMonitoringSystem()
	case 2:
		fmt.Println("System initialized and are working fine")
	default:
		fmt.Println("Esta opção não existe")
		os.Exit(-1)
	}

}

//RETURN TWO OR MORE VAR -> FUNC
/*
func returnNameAndAge() (string, int) {
	name := "Artur"
	age := 18
	return name, age
}
*/

func sayGreeting() {
	name := "Artur"
	version := 0.1

	fmt.Println("\nHello, sr.", name)
	fmt.Println("\nProgram Version", version)
}

func readVar() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("Command", command)

	return command
}

func displayMenu() {
	fmt.Println("1 - Monitoring System")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Exit")
}

func startMonitoringSystem() {
	fmt.Println("Monitoring")

	site := "https://www.fiap.com.br/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site is operating")
	} else {
		fmt.Println("Site is inoperating")
	}
}
