package search

import (
	"github.com/blevesearch/bleve"
)

var Engine bleve.Index

func InitStorage() {
	var err error
	Engine, err = bleve.Open("blogsdata")
	if err != nil {
		imp := bleve.NewIndexMapping()
		blogMapping := bleve.NewDocumentMapping()
		imp.AddDocumentMapping("blog", blogMapping)
		Engine, err = bleve.New("blogsdata", imp)
		if err != nil {
			panic(err)
		}
	}
}

func AddDoc(id string, data string) error {
	return Engine.Index(id, data)
}

func DelDoc(id string) error {
	return Engine.Delete(id)
}

func Find(key string) []string {
	var result = make([]string, 0)
	query := bleve.NewMatchQuery(key)
	search := bleve.NewSearchRequest(query)
	searchResults, err := Engine.Search(search)
	if err != nil {
		return result
	}
	re := searchResults.Hits
	for _, v := range re {
		result = append(result, v.ID)
	}
	return result
}
