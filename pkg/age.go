package pkg

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type AgeRange int

const (
	AgeRangeUnknown AgeRange = iota
	AgeRangeChild
	AgeRangeTeen
	AgeRangeAdult
	AgeRangeSenior
)

var ageRangeText = map[AgeRange]string{
	AgeRangeUnknown: "Unknown",
	AgeRangeChild:   "Child",
	AgeRangeTeen:    "Teen",
	AgeRangeAdult:   "Adult",
	AgeRangeSenior:  "Senior",
}

var ageRangeAge = map[AgeRange][]int{
	AgeRangeUnknown: []int{0, 115},
	AgeRangeChild:   []int{0, 12},
	AgeRangeTeen:    []int{13, 18},
	AgeRangeAdult:   []int{19, 65},
	AgeRangeSenior:  []int{65, 115},
}

func AgeRangeText(code AgeRange) string {
	return ageRangeText[code]
}

func (enum AgeRange) String() string {
	if val, ok := ageRangeText[enum]; ok {
		return val
	}
	return ageRangeText[AgeRangeUnknown]
}

func (s *AgeRange) Scan(value interface{}) error {
	*s = AgeRangeUnknown
	bytes, ok := value.([]byte)
	if !ok {
		val, ok := value.(string)
		if !ok {
			return nil //errors.New("Scan source is not []byte")
		}
		bytes = []byte(val)
	}
	for k, v := range ageRangeText {
		if string(bytes) == v {
			*s = k
			return nil
		}
	}
	return nil
}

func (s AgeRange) Value() (driver.Value, error) {
	return AgeRangeText(s), nil
}

func (s AgeRange) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(AgeRangeText(s))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *AgeRange) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	return s.Scan([]byte(j))
}

func (s *AgeRange) Min() int {
	return ageRangeAge[*s][0]
}

func (s *AgeRange) Max() int {
	return ageRangeAge[*s][1]
}

func (r *Serendipity) AgeRange() AgeRange {
	if r.Bool(.2) {
		if r.Bool(.3) {
			return AgeRangeChild // 6%
		}
		return AgeRangeTeen // 14%
	}
	if r.Bool(.8) {
		return AgeRangeAdult // 64%
	}
	return AgeRangeSenior // 16%
}

func (r *Serendipity) Birthday(ageRange ...AgeRange) time.Time {
	a := AgeRangeUnknown
	if len(ageRange) > 0 {
		a = ageRange[0]
	} else {
		a = r.AgeRange()
	}
	now := time.Now().UTC()
	max := now.Unix()
	min := now.AddDate(-a.Max(), 0, 0).Unix()
	sec := r.N64(min, max)
	return time.Unix(sec, 0)
}
