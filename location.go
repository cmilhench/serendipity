package serendipity

import (
	"fmt"
	"log"
)

func (r *Serendipity) Street() string {
	return fmt.Sprintf("%s %s", r.FamilyName(), r.StreetSuffix())
}

func (r *Serendipity) StreetSuffix() string {
	obj, err := r.loadStrings("/street_suffix.txt")
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

func (r *Serendipity) Locality() string {
	obj, err := r.loadStrings("/locality-GB.txt")
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

func (r *Serendipity) Region() string {
	obj, err := r.loadStrings("/region-GB.txt")
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

func (r *Serendipity) PostalCode() string {
	from := "ABCDEFGHJKLMNPRSTUVWXYZ"
	text := fmt.Sprintf("%s%s%d %d%s%s", r.Char(from), r.Char(from), r.N(1, 59), r.N(1, 9), r.Char(from), r.Char(from))
	return text
}

func (r *Serendipity) Country() *Data {
	obj, err := r.loadData("/country.json")
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

func (r *Serendipity) Latitude() float64 {
	min, max := float64(-90), float64(90)
	return min + r.Float64()*(max-min)
}

func (r *Serendipity) Longitude() float64 {
	min, max := float64(-180), float64(180)
	return min + r.Float64()*(max-min)
}

func (r *Serendipity) Timezone() string {
	obj, err := r.loadStrings("/timezone.txt")
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

type AddressInfo struct {
	Street     string `json:"street_address"`
	Locality   string `json:"locality"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

func (r *Serendipity) Address() *AddressInfo {
	ctry := r.Country()
	addr := AddressInfo{
		Street:     fmt.Sprintf("%d %s", r.N(1, 2999), r.Street()),
		Locality:   r.Locality(),
		Region:     r.Region(),
		PostalCode: r.PostalCode(),
	}
	if ctry != nil {
		addr.Country = ctry.Name
	}
	return &addr
}
