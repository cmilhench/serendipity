package serendipity

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"
)

func (r *Serendipity) Name() string {
	return strings.Trim(fmt.Sprintf("%s %s", r.GivenName(), r.FamilyName()), " ")
}

func (r *Serendipity) GivenName(gender ...GenderType) string {
	g := GenderMale
	if len(gender) > 0 {
		g = gender[0]
	}
	obj, err := r.loadStrings(fmt.Sprintf("/given_name_%s.txt", strings.ToLower(g.String())))
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

func (r *Serendipity) FamilyName() string {
	obj, err := r.loadStrings("/family_name.txt")
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

func (r *Serendipity) Username() string {
	var result string
	if r.Bool(.9) {
		result = fmt.Sprintf("%s%s%d", r.Adjective(), r.Noun(), r.N(1, 99))
	} else {
		a := strings.Split(r.Animal(), " ")
		result = fmt.Sprintf("%s%s%d", r.Adjective(), a[len(a)-1], r.N(1, 99))
	}
	return strings.ToLower(strings.Trim(result, " "))
}

func (r *Serendipity) Password() string {
	text := fmt.Sprintf("%s%s%s%d", r.Adjective(), r.Char("-:+"), r.Noun(), r.N(1, 9))
	text = string(bytes.Join([][]byte{bytes.ToUpper([]byte{text[0]}), []byte(text)[1:]}, nil))
	return text
}

func (r *Serendipity) Email() string {
	return strings.ToLower(fmt.Sprintf("%s%s%d@%s", r.GivenName(), r.Letter(), r.N(0, 10), r.Domain()))
}

func (r *Serendipity) Phone() string {
	if r.Bool(.7) {
		return fmt.Sprintf("07%d %d%d", r.N(100, 999), r.N(100, 999), r.N(100, 999))
	}
	if r.Bool(.7) {
		return fmt.Sprintf("01%d %d%d", r.N(100, 999), r.N(100, 999), r.N(100, 999))
	}
	return fmt.Sprintf("02%d %d%d", r.N(100, 999), r.N(100, 999), r.N(100, 999))
}

func (r *Serendipity) Locale() *Data {
	obj, err := r.loadData("/locale.json")
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

func (r *Serendipity) Profession() string {
	obj, err := r.loadStrings("/profession.txt")
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

type PersonInfo struct {
	Sub                 string       `json:"sub"`
	Name                string       `json:"name"`
	GivenName           string       `json:"given_name"`
	FamilyName          string       `json:"family_name"`
	MiddleName          string       `json:"middle_name"`
	Nickname            string       `json:"nickname"`
	Username            string       `json:"preferred_username"`
	Profile             string       `json:"profile"`
	Picture             string       `json:"picture"`
	Website             string       `json:"website"`
	Email               string       `json:"email"`
	EmailVerified       bool         `json:"email_verified"`
	Gender              GenderType   `json:"gender"`
	Birthday            time.Time    `json:"birthday"`
	ZoneInfo            string       `json:"zoneinfo"`
	Locale              string       `json:"locale"`
	PhoneNumber         string       `json:"phone_number"`
	PhoneNumberVerified bool         `json:"phone_number_verified"`
	Address             *AddressInfo `json:"address"`
	Updated             string       `json:"updated_at"`
}

func (r *Serendipity) Person() *PersonInfo {
	gender := r.Gender()
	givenName := r.GivenName(gender)
	middleName := ""
	if r.Bool(.5) {
		middleName = r.GivenName(gender)
	}
	familyName := r.FamilyName()
	username := r.Username()
	email := strings.ToLower(fmt.Sprintf("%s%s%d@%s", givenName, string(familyName[0]), r.N(0, 10), r.Domain()))
	var parts []string
	for _, s := range []string{givenName, middleName, familyName} {
		if strings.TrimSpace(s) != "" {
			parts = append(parts, s)
		}
	}
	name := strings.Join(parts, " ")
	person := PersonInfo{
		Sub:                 r.UUID(),
		Name:                name,
		GivenName:           givenName,
		FamilyName:          familyName,
		MiddleName:          middleName,
		Nickname:            "",
		Username:            username,
		Profile:             fmt.Sprintf("https://%s/%s/%s", r.Domain(), "profile", username),
		Picture:             fmt.Sprintf("https://%s/%s/%s.jpg", r.Domain(), "pictures", r.FakeWord()),
		Website:             fmt.Sprintf("https://%s", r.Domain()),
		Email:               email,
		EmailVerified:       r.Bool(.9),
		Gender:              gender,
		Birthday:            r.Birthday(),
		ZoneInfo:            "Europe/London",
		Locale:              "en-GB",
		PhoneNumber:         r.Phone(),
		PhoneNumberVerified: r.Bool(.7),
		Address:             r.Address(),
		Updated:             time.Now().UTC().Format(time.RFC3339),
	}
	person.Address.Country = "United Kingdom"

	return &person
}
