//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info
package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStats(servers map[string]int) {
	fmt.Println("\n There are", len(servers), "servers")
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("server error bro!!!")

		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are Offline")
	fmt.Println(stats[Maintenance], "servers are undergoing Maintenance")
	fmt.Println(stats[Retired], "servers are Retired")

}
func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w2359", "baseline"}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	printServerStats(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	printServerStats(serverStatus)
	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	printServerStats(serverStatus)
}
