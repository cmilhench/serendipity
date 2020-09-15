package serendipity

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type GenderType int

const (
	GenderUnknown GenderType = iota
	GenderMale
	GenderFemale
)

var genderText = map[GenderType]string{
	GenderUnknown: "Unknown",
	GenderMale:    "Male",
	GenderFemale:  "Female",
}

func GenderText(code GenderType) string {
	return genderText[code]
}

func (enum GenderType) String() string {
	if val, ok := genderText[enum]; ok {
		return val
	}
	return genderText[GenderUnknown]
}

func (s *GenderType) Scan(value interface{}) error {
	*s = GenderUnknown
	bytes, ok := value.([]byte)
	if !ok {
		val, ok := value.(string)
		if !ok {
			return nil //errors.New("Scan source is not []byte")
		}
		bytes = []byte(val)
	}
	for k, v := range genderText {
		if string(bytes) == v {
			*s = k
			return nil
		}
	}
	return nil
}
func (s GenderType) Value() (driver.Value, error) {
	return GenderText(s), nil
}

func (s GenderType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(GenderText(s))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *GenderType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	return s.Scan([]byte(j))
}

func (r *Serendipity) Gender() GenderType {
	if r.Bool() {
		return GenderMale
	} else {
		return GenderFemale
	}
}
