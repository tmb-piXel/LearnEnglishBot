package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://reallanguage.club/nemeckie-slova-po-temam/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	doc, _ := html.Parse(strings.NewReader(string(body)))

	bd := cascadia.MustCompile("table").MatchFirst(doc)
	c := cascadia.MustCompile("a").MatchAll(bd)

	for _, n := range c {
		f, err := os.Create(n.FirstChild.Data)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		for _, a := range n.Attr {
			if a.Key == "href" {
				resp, err := http.Get(a.Val)
				if err != nil {
					log.Fatalln(err)
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}

				doc, _ := html.Parse(strings.NewReader(string(body)))

				bd := cascadia.MustCompile("table").MatchFirst(doc)
				atr := cascadia.MustCompile("tr").MatchAll(bd)
				for _, at := range atr {
					p := cascadia.MustCompile("span").MatchAll(at)
					origWorld := ""
					russianWorld := ""
					for _, original := range p {
						origW := cascadia.MustCompile("strong").MatchAll(original)
						if origW != nil {
							for _, o := range origW {
								origWorld = o.FirstChild.Data
							}
						} else {
							russianWorld = original.FirstChild.Data
						}
					}
					fmt.Println(origWorld, " - ", russianWorld)
				}
			}
		}
	}
}
