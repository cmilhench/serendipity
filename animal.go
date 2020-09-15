package serendipity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
)

type AnimalType int

const (
	AnimalTypeUnknown AnimalType = iota
	AnimalTypeOcean
	AnimalTypeDesert
	AnimalTypeGrassland
	AnimalTypeForest
	AnimalTypeFarm
	AnimalTypePet
	AnimalTypeZoo
)

var animalTypeText = map[AnimalType]string{
	AnimalTypeUnknown:   "Unknown",
	AnimalTypeOcean:     "ocean",
	AnimalTypeDesert:    "desert",
	AnimalTypeGrassland: "grassland",
	AnimalTypeForest:    "forest",
	AnimalTypeFarm:      "farm",
	AnimalTypePet:       "pet",
	AnimalTypeZoo:       "zoo",
}

func AnimalTypeText(code AnimalType) string {
	return animalTypeText[code]
}

func (enum AnimalType) String() string {
	if val, ok := animalTypeText[enum]; ok {
		return val
	}
	return animalTypeText[AnimalTypeUnknown]
}

func (s *AnimalType) Scan(value interface{}) error {
	*s = AnimalTypeUnknown
	bytes, ok := value.([]byte)
	if !ok {
		val, ok := value.(string)
		if !ok {
			return nil //errors.New("Scan source is not []byte")
		}
		bytes = []byte(val)
	}
	for k, v := range animalTypeText {
		if string(bytes) == v {
			*s = k
			return nil
		}
	}
	return nil
}

func (s AnimalType) Value() (driver.Value, error) {
	return AnimalTypeText(s), nil
}

func (s AnimalType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(AnimalTypeText(s))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *AnimalType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	return s.Scan([]byte(j))
}

func (r *Serendipity) AnimalType() AnimalType {
	return AnimalType(r.N(int(AnimalTypeOcean), int(AnimalTypeZoo)))
}

func (r *Serendipity) Animal(animalType ...AnimalType) string {
	a := AnimalTypeUnknown
	if len(animalType) > 0 {
		a = animalType[0]
	} else {
		a = r.AnimalType()
	}
	obj, err := r.loadStrings(fmt.Sprintf("/animal_%s.txt", a.String()))
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
