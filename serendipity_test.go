package serendipity

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestAgeRange(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.AgeRange()
	if a < AgeRangeUnknown || AgeRangeSenior < a {
		t.Errorf("expected valid age range, got %s", a)
	}
}

func TestBirthday(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	var tests = []struct {
		arg AgeRangeType
	}{
		{AgeRangeUnknown},
		{AgeRangeChild},
		{AgeRangeTeen},
		{AgeRangeAdult},
		{AgeRangeSenior},
	}
	for _, test := range tests {
		a := r.Birthday(test.arg)
		now := time.Now().UTC()
		max := now.Unix()
		min := now.AddDate(-(&test.arg).Max(), 0, 0).Unix()
		if a.Unix() < min || max < a.Unix() {
			t.Errorf("expected valid birthday, got %s", a)
		}
	}
	g := AgeRangeUnknown
	a := r.Birthday()
	now := time.Now().UTC()
	max := now.Unix()
	min := now.AddDate(-(&g).Max(), 0, 0).Unix()
	if a.Unix() < min || max < a.Unix() {
		t.Errorf("expected valid birthday, got %s", a)
	}
}

func TestAnimalType(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.AnimalType()
	if a < AnimalTypeUnknown || AnimalTypeZoo < a {
		t.Errorf("expected valid animal type, got %s", a)
	}
}

func TestAnimal(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	var tests = []struct {
		arg AnimalType
	}{
		//{AnimalTypeUnknown},
		{AnimalTypeOcean},
		{AnimalTypeDesert},
		{AnimalTypeGrassland},
		{AnimalTypeForest},
		{AnimalTypeFarm},
		{AnimalTypePet},
		{AnimalTypeZoo},
	}
	for _, test := range tests {
		a := r.Animal(test.arg)
		if len(a) == 0 {
			t.Errorf("expected animal, got %s", a)
		}
	}
}

func TestBool(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	var tests = []struct {
		f    func() bool
		want float64
	}{
		{func() bool { return r.Bool() }, 50},
		{func() bool { return r.Bool(.9) }, 90},
		{func() bool { return r.Bool(.1) }, 10},
	}

	for _, test := range tests {
		var tot float64 = 0
		var num float64 = 1000
		for n := float64(0); n < num; n++ {
			got := test.f()
			if got {
				tot += (float64(100) / num)
			}
		}
		if math.Abs(test.want-tot) > 5 {
			t.Errorf("expected %.1f, got %.1f", test.want, tot)
		}
	}
}

func BenchmarkBool(b *testing.B) {
	r := New()
	r.Seed(time.Now().Unix())
	var tests = []struct {
		name string
		f    func() string
	}{
		{"AgeRange", func() string { return fmt.Sprintf("%+v", r.AgeRange()) }},
		{"Birthday", func() string { return fmt.Sprintf("%s", r.Birthday()) }},
		{"AnimalType", func() string { return fmt.Sprintf("%+v", r.AnimalType()) }},
		{"Animal", func() string { return fmt.Sprint(r.Animal()) }},
		{"AnimalDesert", func() string { return fmt.Sprint(r.Animal(AnimalTypeDesert)) }},
		{"AnimalFarm", func() string { return fmt.Sprint(r.Animal(AnimalTypeFarm)) }},
		{"AnimalForest", func() string { return fmt.Sprint(r.Animal(AnimalTypeForest)) }},
		{"AnimalGrassland", func() string { return fmt.Sprint(r.Animal(AnimalTypeGrassland)) }},
		{"AnimalOcean", func() string { return fmt.Sprint(r.Animal(AnimalTypeOcean)) }},
		{"AnimalPet", func() string { return fmt.Sprint(r.Animal(AnimalTypePet)) }},
		{"AnimalZoo", func() string { return fmt.Sprint(r.Animal(AnimalTypeZoo)) }},
		{"Bool", func() string { return fmt.Sprint(r.Bool()) }},
		{"Colour", func() string { return fmt.Sprintf("%+v", r.Country()) }},
		{"Currency", func() string { return fmt.Sprintf("%+v", r.Currency()) }},
		{"Domain", func() string { return fmt.Sprint(r.Domain()) }},
		{"TLD", func() string { return fmt.Sprint(r.TLD()) }},
		{"URL", func() string { return fmt.Sprint(r.URL()) }},
		{"Gender", func() string { return fmt.Sprintf("%+v", r.Gender()) }},
		{"Street", func() string { return fmt.Sprint(r.Street()) }},
		{"StreetSuffix", func() string { return fmt.Sprint(r.StreetSuffix()) }},
		{"Locality", func() string { return fmt.Sprint(r.Locality()) }},
		{"Region", func() string { return fmt.Sprint(r.Region()) }},
		{"PostalCode", func() string { return fmt.Sprint(r.PostalCode()) }},
		{"Country", func() string { return fmt.Sprintf("%+v", r.Country()) }},
		{"Latitude", func() string { return fmt.Sprintf("%.4f", r.Latitude()) }},
		{"Longitude", func() string { return fmt.Sprintf("%.4f", r.Longitude()) }},
		{"Timezone", func() string { return fmt.Sprint(r.Timezone()) }},
		{"Address", func() string { return fmt.Sprintf("%+v", r.Address()) }},
		{"Name", func() string { return fmt.Sprint(r.Name()) }},
		{"GivenName", func() string { return fmt.Sprint(r.GivenName()) }},
		{"GivenFemale", func() string { return fmt.Sprint(r.GivenName(GenderFemale)) }},
		{"GivenMale", func() string { return fmt.Sprint(r.GivenName(GenderMale)) }},
		{"FamilyName", func() string { return fmt.Sprint(r.FamilyName()) }},
		{"Username", func() string { return fmt.Sprint(r.Username()) }},
		{"Password", func() string { return fmt.Sprint(r.Password()) }},
		{"Email", func() string { return fmt.Sprint(r.Email()) }},
		{"Phone", func() string { return fmt.Sprint(r.Phone()) }},
		{"Locale", func() string { return fmt.Sprintf("%+v", r.Locale()) }},
		{"Profession", func() string { return fmt.Sprint(r.Profession()) }},
		{"Person", func() string { return fmt.Sprintf("%+v", r.Person()) }},
		{"N", func() string { return fmt.Sprintf("%d", r.N(1, 9)) }},
		{"N64", func() string { return fmt.Sprintf("%d", r.N64(1, 9)) }},
		{"UUID", func() string { return fmt.Sprint(r.UUID()) }},
		{"Char", func() string { return fmt.Sprint(r.Char("aeiou")) }},
		{"Letter", func() string { return fmt.Sprint(r.Letter()) }},
		{"Vowel", func() string { return fmt.Sprint(r.Vowel()) }},
		{"Consonant", func() string { return fmt.Sprint(r.Consonant()) }},
		{"Syllable", func() string { return fmt.Sprint(r.Syllable()) }},
		{"FakeWord", func() string { return fmt.Sprint(r.FakeWord()) }},
		{"Word", func() string { return fmt.Sprint(r.Word()) }},
		{"Adjective", func() string { return fmt.Sprint(r.Adjective()) }},
		{"Noun", func() string { return fmt.Sprint(r.Noun()) }},
		{"Sentence", func() string { return fmt.Sprint(r.Sentence()) }},
		{"Sentence!", func() string { return fmt.Sprint(r.Sentence(true)) }},
		{"Paragraph", func() string { return fmt.Sprint(r.Paragraph()) }},
		{"Punctuation", func() string { return fmt.Sprint(r.Punctuation()) }},
	}
	for _, test := range tests {
		b.Run(fmt.Sprintf("method%v", test.name), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = test.f()
			}
		})
	}
}

func TestColour(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Colour()
	if a == nil {
		t.Errorf("expected valid colour, got %s", a)
	}
}

func TestCurrency(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Currency()
	if a == nil {
		t.Errorf("expected valid currency, got %s", a)
	}
}
func TestDomain(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Domain()
	if len(a) == 0 {
		t.Errorf("expected valid domain, got %s", a)
	}
}
func TestTLD(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.TLD()
	if len(a) == 0 {
		t.Errorf("expected valid tld, got %s", a)
	}
}
func TestURL(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.URL()
	if len(a) == 0 {
		t.Errorf("expected valid url, got %s", a)
	}
}

func TestGender(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Gender()
	if a < GenderUnknown || GenderFemale < a {
		t.Errorf("expected valid gender, got %s", a)
	}
}

func TestStreet(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Street()
	if len(a) == 0 {
		t.Errorf("expected valid street, got %s", a)
	}
}

func TestStreetSuffix(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.StreetSuffix()
	if len(a) == 0 {
		t.Errorf("expected valid street suffix, got %s", a)
	}
}

func TestLocality(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Locality()
	if len(a) == 0 {
		t.Errorf("expected valid locality, got %s", a)
	}
}

func TestRegion(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Region()
	if len(a) == 0 {
		t.Errorf("expected valid region, got %s", a)
	}
}

func TestPostalCode(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.PostalCode()
	if len(a) == 0 {
		t.Errorf("expected valid postal code, got %s", a)
	}
}

func TestCountry(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Country()
	if a == nil {
		t.Errorf("expected valid country, got %s", a)
	}
}

func TestLatitude(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Latitude()
	if a == 0 {
		t.Errorf("expected valid latitude, got %.4f", a)
	}
}

func TestLongitude(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Longitude()
	if a == 0 {
		t.Errorf("expected valid longitude, got %.4f", a)
	}
}

func TestTimezone(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Timezone()
	if len(a) == 0 {
		t.Errorf("expected valid timezone, got %s", a)
	}
}

func TestAddress(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Address()
	if a == nil {
		t.Errorf("expected valid address got %s", a)
	}
}

func TestName(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Name()
	if len(a) == 0 {
		t.Errorf("expected valid name, got %s", a)
	}
}

func TestGivenName(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.GivenName()
	if len(a) == 0 {
		t.Errorf("expected valid given name, got %s", a)
	}
}

func TestFamilyName(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.FamilyName()
	if len(a) == 0 {
		t.Errorf("expected valid family name, got %s", a)
	}
}

func TestUsername(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Username()
	if len(a) == 0 {
		t.Errorf("expected valid username, got %s", a)
	}
}

func TestPassword(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Password()
	if len(a) == 0 {
		t.Errorf("expected valid password, got %s", a)
	}
}

func TestEmail(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Email()
	if len(a) == 0 {
		t.Errorf("expected valid email, got %s", a)
	}
}

func TestPhone(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Phone()
	if len(a) == 0 {
		t.Errorf("expected valid phone, got %s", a)
	}
}

func TestLocale(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Locale()
	if a == nil {
		t.Errorf("expected valid locale, got %s", a)
	}
}

func TestProfession(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Profession()
	if len(a) == 0 {
		t.Errorf("expected valid profession, got %s", a)
	}
}

func TestPerson(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Person()
	if a == nil {
		t.Errorf("expected valid person, got %+v", a)
	}
}

func TestN(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.N(1, 9)
	if a == 0 {
		t.Errorf("expected valid number, got %d", a)
	}
}

func TestN64(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.N64(1, 9)
	if a == 0 {
		t.Errorf("expected valid number, got %d", a)
	}
}

func TestUUID(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.UUID()
	if len(a) == 0 {
		t.Errorf("expected valid uuid, got %s", a)
	}
}

func TestChar(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Char("1+=-*")
	if len(a) != 1 {
		t.Errorf("expected valid char, got %s", a)
	}
}

func TestLetter(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Letter()
	if len(a) != 1 {
		t.Errorf("expected valid letter, got %s", a)
	}
}

func TestVowel(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Vowel()
	if len(a) != 1 {
		t.Errorf("expected valid vowel, got %s", a)
	}
}

func TestConsonant(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Consonant()
	if len(a) != 1 {
		t.Errorf("expected valid consonant, got %s", a)
	}
}

func TestSyllable(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Syllable()
	if len(a) == 0 {
		t.Errorf("expected valid syllable, got %s", a)
	}
}

func TestFakeWord(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.FakeWord()
	if len(a) == 0 {
		t.Errorf("expected valid word, got %s", a)
	}
}

func TestWord(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Word()
	if len(a) == 0 {
		t.Errorf("expected valid word, got %s", a)
	}
}

func TestAdjective(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Adjective()
	if len(a) == 0 {
		t.Errorf("expected valid adjective, got %s", a)
	}
}

func TestNoun(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Noun()
	if len(a) == 0 {
		t.Errorf("expected valid noun, got %s", a)
	}
}

func TestSentence(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Sentence()
	if len(a) == 0 {
		t.Errorf("expected valid sentence, got %s", a)
	}
}

func TestParagraph(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Paragraph()
	if len(a) == 0 {
		t.Errorf("expected valid paragraph, got %s", a)
	}
}

func TestPunctuation(t *testing.T) {
	r := New()
	r.Seed(time.Now().Unix())
	a := r.Punctuation()
	if len(a) != 1 {
		t.Errorf("expected valid punctuation, got %s", a)
	}
}
