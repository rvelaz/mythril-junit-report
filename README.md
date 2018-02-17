# mythril-junit-report

It parses the output of mythril analysis on solidity files and creates an ouput compatible with JUnit
 

## Installation
Go version 1. or higher is required. Install or update using the go get command:

```bash
go get -u github.com/rvelaz/mythril-junit-report
```


## Usage
mythril-junit-report reads the mythril test output in json format from standard in and writes junit compatible XML to standard out.


```bash
myth -o -x origin.sol 2>&1 | mythril-junit-report > report.xml
```

