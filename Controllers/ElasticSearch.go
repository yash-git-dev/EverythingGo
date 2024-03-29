package Controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch"
	//"github.com/elastic/go-elasticsearch/v8"
)

func GetDataFromElastic(word string) interface{} {
	client1 := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	clcll, _ := elasticsearch.NewClient(client1)
	fmt.Println(clcll)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}
	fmt.Println(es)
	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"field": word,
			},
		},
	}
	err = json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return err
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return res.IsError()
	}
	fmt.Println(res.Body)
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return response
}
