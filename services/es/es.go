package es

import (
	"Atlantis/config"
	"Atlantis/services/logger"
	"context"
	"net/http"
	"sync"

	"github.com/olivere/elastic/v7"
	"go.elastic.co/apm/module/apmelasticsearch"
	"go.uber.org/zap"
)

var client *elastic.Client
var esOnce sync.Once

func Init() *elastic.Client {
	esOnce.Do(func() {
		esURL := config.Get().EsURL
		var err error
		client, err = elastic.NewClient(elastic.SetURL(esURL), elastic.SetHttpClient(&http.Client{
			Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
		}), elastic.SetSniff(false))
		if err != nil {
			panic(err.Error())
		}
		_, _, err = client.Ping(config.Get().EsURL).Do(context.Background())
		if err != nil {
			panic(err)
		}
		logger.Client().Info("es connected", zap.String("host", config.Get().EsURL))
	})
	return client
}

func Client() *elastic.Client {
	return client
}
