package serendipity

import (
	"log"
)

func (r *Serendipity) Currency() *Data {
	obj, err := r.loadData("/currency.json")
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
