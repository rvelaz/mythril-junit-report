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
myth -o json -x origin.sol 2>&1 | mythril-junit-report > report.xml
```

I you have already a report, for example that contains several results:

```bash
cat mythril-results.json | mythril-junit-report > report.xml
```


### Using docker
Let's suppose we have a result file (mythril-results.json) containing multiple test outputs in */path/to/folder_with_test_result*.
The report in JUnit format could be generated using docker with the following command:  

```bash
docker run --rm -it -v /path/to/folder_with_test_result:/usr/src/myapp \
        -w /usr/src/myapp golang:1.9 \
        bash -c "go get -u github.com/rvelaz/mythril-junit-report && cat mythril-results.json | mythril-junit-report > report.xml"
```