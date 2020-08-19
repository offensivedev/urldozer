# urldozer
Perform operations on URLs like extracting path, parameter names and/or values, domain name, host name (without HTTP[s]). Useful in penetration testing engagements and bug bounty. 

## Installation
If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

`go get -u github.com/offensivedev/urldozer`

Otherwise download a [release](https://github.com/offensivedev/urldozer/releases/). To make it easier to execute you can put the binary in your $PATH.

## Usage
```
$ urldozer -h
$ cat urls.txt | urldozer
$ echo "https://example.com/admin/?search=test&?file=index.php" | urldozer -param-value
```

## Example use cases
- Create a parameter wordlist from URLs fetched gau/waybackurls
- Extract URL paths to check for vulnerabilities across different environment
- Find all domains used in list of URLs by piping the output like `| sort -u`

## TODO
- Add payload to parameters 
  - Option to inject payloads from file
- Output to file
  - Allow formats for output file (raw, json)


## Author
Amit Kumar <br />
Twitter     - [@offensivedev](https://twitter.com/offensivedev) <br />
LinkedIn    - [Amit Kumar](https://www.linkedin.com/in/offensivedev)