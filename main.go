package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Main takes a list of domain names on stdin and outputs a csv with the results of the
// DNS lookups. The output is in the following format:
// domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

// checkDomain checks MX, SPF, and DMARC records for a domain and prints the results
func checkDomain(domain string) {
	domain = strings.TrimSpace(domain)
	if domain == "" || !strings.Contains(domain, ".") {
		log.Printf("Invalid domain: %s\n", domain)
		return
	}

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Lookup MX records
	if mxRecords, err := net.LookupMX(domain); err == nil && len(mxRecords) > 0 {
		hasMX = true
	} else if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
	}

	// Lookup TXT records
	if txtRecords, err := net.LookupTXT(domain); err == nil {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	} else {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	// Lookup DMARC records
	if dmarcRecords, err := net.LookupTXT("_dmarc." + domain); err == nil {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	} else {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	}

	// Output results
	fmt.Printf("%s, %v, %v, %s, %v, %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
