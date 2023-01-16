package profile

import (
	"encoding/xml"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	gc "github.com/patrickmn/go-cache"
	"strings"
	"sync"
	"time"
)

func Pprof() {
	app := gin.Default()

	pprof.Register(app) // 性能
}

const (
	expiration = time.Minute * 15
	cleanup    = time.Hour
)

var cache = gc.New(expiration, cleanup)

var fetch = struct {
	sync.Mutex
	m map[string]*sync.Mutex
}{
	m: make(map[string]*sync.Mutex),
}

type Searcher interface {
	Search(uid string, term string, found chan<- []Result)
}

// Result represents a search result that was found.
type Result struct {
	Engine  string
	Title   string
	Link    string
	Content string
}

var d Document

type Item struct {
	XMLName     xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

func RssSearch(*gin.Context) {
	results := []Result{}
	//var d Document

	// 在循环中，做小写
	for i := 0; i < 10; i++ {
		description := "item.Descriptionitem.Descriptionitem.Descriptionitem.Descriptionitem.Descriptionitem.Descriptionitem.Descriptionitem.Descriptionitem.Description"
		//Description := strings.ToLower(description)
		termToHight := "termtermtermtermtermtermtermtermtermtermtermtermtermtermterm"
		//term := strings.ToLower(termToHight)
		//fmt.Printf("Description %v\n", Description)
		//fmt.Printf("term %v\n", term)
		if strings.Contains(description, termToHight) {
			results = append(results, Result{
				Engine:  "engine",
				Title:   "item.Title",
				Link:    "item.Link",
				Content: "item.Description",
			})
		}
		//time.Sleep(time.Second * 10)
	}

	return
}
