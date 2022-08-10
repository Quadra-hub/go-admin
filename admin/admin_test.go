package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Photo struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type Supply struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Photo Photo  `json:"photo"`
}

func TestGelAllSubscribedTables(t *testing.T) {
	admin := New()
	admin.Subscribe(Photo{})
	admin.Subscribe(Supply{})
	assert.Equal(t, []string{"Photo", "Supply"}, admin.GetTableNames())
}

func TestGetHeadersOfTables(t *testing.T) {
	admin := New()
	admin.Subscribe(Photo{})
	assert.Equal(t, []string{"Id", "Url"}, admin.GetHeaders(Photo{}))
}

func TestGetValueByFieldName(t *testing.T) {
	admin := New()
	admin.Subscribe(Photo{})

	assert.Equal(t, 1, admin.GetValueByFieldName(Photo{Id: 1}, "Id"))
	assert.Equal(t, "http://www.google.com", admin.GetValueByFieldName(Photo{Url: "http://www.google.com"}, "Url"))
}
