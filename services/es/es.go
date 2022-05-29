package es

import (
	"Atlantis/config"
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/olivere/elastic/v7"
	"go.elastic.co/apm/module/apmelasticsearch"
)

var client *elastic.Client
var esOnce sync.Once

func Init() *elastic.Client {
	esOnce.Do(func() {
		esURL := config.Get().EsURL
		var err error
		client, err = elastic.NewClient(elastic.SetURL(esURL), elastic.SetHttpClient(&http.Client{
			Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
		}), elastic.SetSniff(false), elastic.SetBasicAuth("elastic", "XLVSwKIpBLLwvD8qmqzsqXaD"), elastic.SetScheme("https"))
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(config.Get().EsURL)
		_, _, err = client.Ping(config.Get().EsURL).Do(context.Background())
		if err != nil {
			panic(err)
		}

	})
	return client
}

func Client() *elastic.Client {
	return client
}
