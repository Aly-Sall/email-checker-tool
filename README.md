Mail Checker

Mail Checker is a command-line tool written in Go that checks DNS records for a given domain. It identifies whether the domain has MX, SPF, and DMARC records and outputs the results in a CSV format.

Features

.Checks for MX (mail server) records.
.Checks for and retrieves SPF records.
.Checks for and retrieves DMARC records.
.Outputs results in a clear, structured CSV format.

Prerequisites

Go version 1.18 or newer.

Installation

1.Clone the GitHub repository
2.Build the program: go build -o mail-checker
This will generate an executable binary named mail-checker.

Usage

Input via stdin
You can provide a list of domains directly via stdin: echo -e "example.com\ngoogle.com" | ./mail-checker

Input from a file
To check domains from a file domains.txt: ./mail-checker < domains.txt

Example Output
The output is displayed in CSV format: domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
example.com, true, false, , true, v=DMARC1; p=none
google.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=reject; rua=mailto:dmarc-reports@google.com
