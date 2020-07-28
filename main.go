package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"net/url"
	"strings"
)

func main() {
	
	// get domains only
	var getDomainOnly bool
	flag.BoolVar(&getDomainOnly, "domain", false, "Return domain without their path from list of URL")

	// get hosts only
	var getHostOnly bool
	flag.BoolVar(&getHostOnly, "host", false, "Return host from list of URL")
	
	// get path only
	var getPathOnly bool
	flag.BoolVar(&getPathOnly, "path", false, "Returns path collected from URL")

	// get params only
	var getParamOnly bool
	flag.BoolVar(&getParamOnly, "param", false, "Returns parameter collected from URL")

	// get params with values
	var getParamsWithValues bool
	flag.BoolVar(&getParamsWithValues, "param-value", false, "Returns parameter with their value collected from URL")

	flag.Parse()

	if (getHostOnly == false && getDomainOnly == false && getParamOnly == false && getParamsWithValues == false && getPathOnly == false) {
		fmt.Printf("Please select atleast one mode\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		u, err := url.Parse(sc.Text())
		if err != nil {
			fmt.Printf("Error in parsing the URL :(\n")
			panic(err)
		}
		
		if getDomainOnly {
			fmt.Printf("%s\n", u.Scheme + "://" + u.Host)
		}
		if getHostOnly {
			fmt.Printf("%s\n", u.Host)
		}
		if getPathOnly {
			fmt.Printf("%s\n", u.Path)
		}
		if getParamOnly {
			param, _ := url.ParseQuery(u.RawQuery)

			for name := range param {
				name = strings.ReplaceAll(string(name), "?", "")
				fmt.Println(name)
			}
		}
		if getParamsWithValues {
			param, _ := url.ParseQuery(u.RawQuery)

			for name, value := range param {
				name = strings.ReplaceAll(string(name), "?", "")
				fmt.Println(name)
				fmt.Println(value[0])
			}
		}
	}
}