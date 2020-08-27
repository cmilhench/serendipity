package pkg

import (
	"fmt"
	"log"
)

func (r *Serendipity) Domain() string {
	return fmt.Sprintf("%s.%s", r.Word(), r.TLD())
}

func (r *Serendipity) TLD() string {
	obj, err := r.loadStrings("/tld.txt")
	if err != nil {
		log.Println(err)
		return ""
	}
	count := len(*obj)
	if count == 0 {
		return ""
	}
	i := r.N(0, count-1)
	return (*obj)[i]
}

func (r *Serendipity) URL(suffix ...string) string {
	s := "html"
	if len(suffix) > 0 {
		s = suffix[0]
	}
	return fmt.Sprintf("https://%s/%s/%s.%s", r.Domain(), r.Word(), r.Word(), s)
}
