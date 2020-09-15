package serendipity

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cmilhench/serendipity/internal/asset"
	"github.com/cmilhench/serendipity/internal/cache"
)

// This is used to localy cache file data for [ttl:2] minutes
// A tick occurs every [interval:1] minute to clean stale data entries
var (
	ttl   = 2 * time.Minute
	store = cache.New(1 * time.Minute)
)

type Serendipity struct {
	*rand.Rand
}

func New() *Serendipity {
	return &Serendipity{rand.New(rand.NewSource(1))}
}

func (r *Serendipity) Seed(seed int64) {
	r.Rand.Seed(seed)
}

func (r *Serendipity) N(min, max int) int {
	return r.Intn(max-min+1) + min
}

func (r *Serendipity) N64(min, max int64) int64 {
	return r.Int63n(max-min+1) + min
}

func (r *Serendipity) UUID() string {
	b := make([]byte, 16)
	_, err := r.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] &^ 0xf0) | 0x40
	b[8] = (b[8] &^ 0xc0) | 0x80
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func (r *Serendipity) Coalese(value ...string) (result string) {
	for _, result := range value {
		if result == "" {
			continue
		}
		break
	}
	return
}

func (r *Serendipity) loadStrings(name string) (*[]string, error) {
	value, found := store.Get(name)
	if found {
		return value.(*[]string), nil
	}
	buf := asset.Get(name)
	obj := make([]string, 0)
	f := bytes.NewReader(buf)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		obj = append(obj, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	store.Set(name, &obj, ttl)
	return &obj, nil
}

type Data struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

func (r *Serendipity) loadData(name string) (*[]Data, error) {
	value, found := store.Get(name)
	if found {
		return value.(*[]Data), nil
	}
	buf := asset.Get(name)
	obj := make([]Data, 0)
	if len(buf) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(buf, &obj)
	if err != nil {
		return nil, err
	}
	store.Set(name, &obj, ttl)
	return &obj, nil
}

/*
 * Top-level convenience functions
 */

var globalSerendipity = New()

func Seed(seed int64) { globalSerendipity.Seed(seed) }

func Person() *PersonInfo { return globalSerendipity.Person() }

// More TK
