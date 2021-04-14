package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	var url string
	fmt.Print("Please input a url:  ")
	_, err := fmt.Scanln(&url)
	if err != nil { // How swag
		fmt.Println(err)
	}
	download(url)
	fmt.Println("Completed. Check the root folder of where you installed this software.\nPress enter to end.")
	fmt.Scanln()
}

func getIds(url string) []string {
	response, err := http.Get(url) // Gets URL provided by user.
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp := buf.String()
	re := regexp.MustCompile(`/sound/[0-9]*/`) // Regex to get sound ID's from the page
	results := re.FindAllString(resp, -1)
	str := strings.Join(results, "")
	ids := strings.Split(str, "/sound") // Splitting garbage, how fun
	return ids
}

func download(url string) {
	ids := getIds(url)
	ids[0] = ids[len(ids)-1] // this basically just removes an empty object that the regex grabs
	ids[len(ids)-1] = ""
	ids = ids[:len(ids)-1]
	i := 0
	for i < len(ids) {
		// Get the data
		resp, err := http.Get("https://www.sounds-resource.com/download" + ids[i]) // Download endpoint
		re := regexp.MustCompile(`"([^"]*)"`)
		joined := strings.Join(resp.Header.Values("Content-Disposition"), " ") // Gets file name
		header := re.FindString(joined)
		cleaned := strings.ReplaceAll(header, `"`, "")
		if err != nil {
			fmt.Println(err)
		}
		name := cleaned
		defer resp.Body.Close()
		out, err := os.Create(name)
		if err != nil {
			return
		}
		defer out.Close()
		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(2 * time.Second) // This delay was added by request by the sites manager. Don't remove this.
		i++
	}
}
