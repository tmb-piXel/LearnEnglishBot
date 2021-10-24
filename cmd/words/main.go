package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

var path = "dictionaries/"

var urls = map[string]string{
	"🇬🇧English1": "https://reallanguage.club/anglijskie-slova-po-temam/",
	"🇫🇷France":   "https://reallanguage.club/francuzskie-slova-po-temam/",
	"🇮🇹Italy":    "https://reallanguage.club/italyanskie-slova-po-temam/",
	"🇪🇸Spain":    "https://reallanguage.club/ispanskie-slova-po-temam/",
	"🇩🇪German":   "https://reallanguage.club/nemeckie-slova-po-temam/",
}

//4308

func main() {
	tn := time.Now()
	var wg sync.WaitGroup
	wg.Add(5)
	for titleDir, url := range urls {
		go func(titleDir, url string) {
			topics := tableElements(url, "a")
			os.Mkdir(path+titleDir, 0777)
			var wg1 sync.WaitGroup
			wg1.Add(len(topics))
			for _, t := range topics {
				go writeWordsFile(t, titleDir+"/", &wg1)
			}
			wg1.Wait()
			defer wg.Done()
		}(titleDir, url)
	}

	wg.Wait()
	fmt.Println(time.Since(tn))
	fmt.Println("stop")
}

func writeWordsFile(t *html.Node, dir string, wg *sync.WaitGroup) {
	//Создаем файлы со словами
	topicTitle := t.FirstChild.Data
	if t.FirstChild.Data == "span" {
		topicTitle = t.FirstChild.LastChild.Data
	}

	f, err := os.Create(path + dir + topicTitle)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := words(t)
	for _, w := range words {
		f.WriteString(w)
	}

	defer wg.Done()
}

func words(t *html.Node) (words []string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			pairsWords := tableElements(a.Val, "tr")

			for _, at := range pairsWords {
				original := ""
				russian := ""
				spans := cascadia.MustCompile("span").MatchAll(at)
				sizeSpans := len(spans)
				for i, sp := range spans {
					strong := cascadia.MustCompile("strong").MatchAll(sp)
					if strong != nil {
						st := strong[0]
						original = st.FirstChild.Data
					} else if sp.FirstChild.Data != "span" && sizeSpans == 4 {
						if i == 1 {
							original = sp.FirstChild.Data
						} else if i == 3 {
							russian = sp.FirstChild.Data
						}
					} else {
						russian = sp.FirstChild.Data
					}
				}
				words = append(words, original+"----"+russian+"\n")
			}
		}
	}
	return
}

func tableElements(url, tag string) (elements []*html.Node) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatalln(err)
	}

	tables := cascadia.MustCompile("table").MatchFirst(doc)
	elements = cascadia.MustCompile(tag).MatchAll(tables)
	return
}
