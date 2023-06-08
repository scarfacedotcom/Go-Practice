package main

import "time"

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}
type Member struct {
	name  Name
	books map[Title]LendAudit
}

func main() {

}
