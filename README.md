# urldozer
Perform operations on URLs like extracting paths, parameter names and/or values, domain name, host name (without HTTP[s]).

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