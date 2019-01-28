package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


type logins struct {
	Api_key string `yaml:"nytimes_api_key"`
}


func (c *logins) get_nytimes_api_key() *logins {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

type NYTlayer struct {
	Copyright string `json:"copyright"`
	Response  struct {
		Meta struct {
			Hits int `json:"hits"`
		} `json:"meta"`
		Docs []struct {
			WebURL        string        `json:"web_url"`
			Snippet       interface{}   `json:"snippet"`
			LeadParagraph interface{}   `json:"lead_paragraph"`
			Abstract      interface{}   `json:"abstract"`
			PrintPage     string        `json:"print_page"`
			Blog          []interface{} `json:"blog"`
			Source        string        `json:"source"`
			Multimedia    []interface{} `json:"multimedia"`
			Headline      struct {
				Main   string `json:"main"`
				Kicker string `json:"kicker"`
			} `json:"headline"`
			Keywords         []interface{} `json:"keywords"`
			PubDate          time.Time     `json:"pub_date"`
			DocumentType     string        `json:"document_type"`
			NewsDesk         interface{}   `json:"news_desk"`
			SectionName      interface{}   `json:"section_name"`
			SubsectionName   interface{}   `json:"subsection_name"`
			Byline           interface{}   `json:"byline"`
			TypeOfMaterial   string        `json:"type_of_material"`
			ID               string        `json:"_id"`
			WordCount        int           `json:"word_count"`
			SlideshowCredits interface{}   `json:"slideshow_credits"`
		} `json:"docs"`
	} `json:"response"`
}

func main() {
	var c logins
	api_key := c.get_nytimes_api_key().Api_key
	url := fmt.Sprintf("https://api.nytimes.com/svc/archive/v1/2019/1.json?api-key=%s", api_key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record NYTlayer

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println(record.Copyright)
	fmt.Println("da")
}