package elasticsearch

import (
	"time"

	"github.com/go_elasticsearch/elastic/internal/pkg/storage"
)

var _ storage.PostStorer = PostStorage{}

type PostStorage struct {
	elastic ElasticSearch
	timeout time.Duration
}

func NewPostStorage(elastic ElasticSearch) (PostStorage, error) {
	return PostStorage{
		elastic: elastic,
		timeout: time.Second * 10,
	}, nil
}
