package services

import (
	"github.com/meilisearch/meilisearch-go"
)

func CreateMeiliClient(host string) (err error, meili *meilisearch.Client) {
	meili = meilisearch.NewClient(meilisearch.ClientConfig{
		Host: host,
	})

	if _, err := meili.GetIndex("users"); err != nil {
		_, err := meili.CreateIndex(&meilisearch.IndexConfig{
			PrimaryKey: "id",
			Uid:        "users",
		})
		if err != nil {
			return err, nil
		}
	}

	return nil, meili
}
