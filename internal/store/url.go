package store

import "github.com/aliazizii/url-shotener/internal/model"

type Url interface {
	FindByHash(Hash string) (string, error)
	Insert(insert model.Url) (string, error)
}
