package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func verifyEmail(email string) bool {
	sp := strings.Split(email, "@")
	//local_part := sp[0]
	domain_part := sp[1]

	mxrecords, _ := net.LookupMX(domain_part)
	mxserver := mxrecords[0].Host[:len(mxrecords[0].Host)-1]

	conn, _ := smtp.Dial(mxserver + ":25")
	conn.Mail("tmp@gmail.com")
	err := conn.Rcpt(email)
	conn.Quit()

	return err == nil
}

func main() {
	sleep_ptr := flag.Int("sleep", 1, "delay for each connection")
	flag.Parse()
	sleep := *sleep_ptr

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		email := scanner.Text()
		result := verifyEmail(email)

		if result == true {
			fmt.Println("\x1b[32m[+] " + email + " is valid.\x1b[0m")
		} else {
			fmt.Println("\x1b[31m[-] " + email + " is not valid.\x1b[0m")
		}

		time.Sleep(time.Second * time.Duration(sleep))
	}
}
