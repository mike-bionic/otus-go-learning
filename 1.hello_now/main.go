package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetNtpTime(host string) time.Time {
	time, err := ntp.Time(host)
	if err != nil {
		fmt.Fprint(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return time
}

func main() {
	fmt.Print("Host")
	var host string
	fmt.Scanln(&host)

	time := GetNtpTime(host)
	fmt.Println(time)
}
