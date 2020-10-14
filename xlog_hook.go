package xlog

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

// ConnectElasticsearch will add Hook to logrus and push log in async mode to elasticsearch
func ConnectElasticsearch(url, username, password, host, index string, logLevel logrus.Level) {
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(username, password))
	if err != nil {
		logrus.Panic(err)
	}
	hook, err := elogrus.NewAsyncElasticHook(client, host, logLevel, index)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.AddHook(hook)
}
