package main

import (
	"net/http"
	"os"
	// "errors"
	// "bufio"
	"fmt"
	"regexp"
	"bytes"
	"strings"
)
var url string
func main() {
	fmt.Print("Please input a url:  ")
	fmt.Scanln(&url)
	getIds()
}

func getIds() {
	response, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        if err != nil {
            fmt.Printf("%s", err)
        }
	}
	buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    resp := buf.String()
	re := regexp.MustCompile(`/sound/[0-9]*/`)
	results := re.FindAllString(resp, -1)
	str := strings.Join(results,"")
	ids := strings.Split(str, "/sound")
	fmt.Printf("%q\n", ids)
}

func download(){
	
}