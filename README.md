# urldozer
Perform operations on URLs like extracting paths, parameter names and/or values, domain name, host name (without HTTP[s]).

## Usage
```
$ urldozer -h
$ cat urls.txt | urldozer
$ echo "https://example.com/admin/?search=test&?file=index.php" | urldozer -param-value
```