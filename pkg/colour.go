package pkg

import (
	"log"
)

func (r *Serendipity) Colour() *Data {
	obj, err := r.loadData("/colour.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	count := len(*obj)
	if count == 0 {
		return nil
	}
	i := r.N(0, count-1)
	return &(*obj)[i]
}
