package serendipity

import (
	"bytes"
	"log"
	"strings"
)

func (r *Serendipity) Char(from string) string {
	return string(from[r.Intn(len(from))])
}

func (r *Serendipity) Letter() string {
	return r.Char("abcdefghijklmnopqrstuvwxyz")
}

func (r *Serendipity) Vowel() string {
	return r.Char("aeoui")
}

func (r *Serendipity) Consonant() string {
	return r.Char("bcdfghjklmnpqrstvwxyz")
}

func (r *Serendipity) Syllable() string {
	length := r.N(2, 3)
	text := make([]string, length)
	b := r.Bool()
	for i := 0; i < length; i++ {
		if (i%2 == 0) == b {
			text[i] = r.Vowel()
		} else {
			text[i] = r.Consonant()
		}
	}
	return strings.Join(text, "")
}

func (r *Serendipity) FakeWord() string {
	length := r.N(1, 3)
	text := make([]string, length)
	for i := 0; i < length; i++ {
		text[i] = r.Syllable()
	}
	return strings.Join(text, "")
}

func (r *Serendipity) Word() string {
	if r.Bool() {
		return r.Adjective()
	}
	return r.Noun()
}

func (r *Serendipity) Adjective() string {
	obj, err := r.loadStrings("/adjective.txt")
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

func (r *Serendipity) Noun() string {
	obj, err := r.loadStrings("/noun.txt")
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

func (r *Serendipity) Sentence(punctuation ...bool) string {
	p := false
	if len(punctuation) > 0 {
		p = punctuation[0]
	} else {
		p = r.Bool()
	}
	count := r.N(12, 18)
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = r.FakeWord()
	}
	text := strings.Join(words, " ")
	text = string(bytes.Join([][]byte{bytes.ToUpper([]byte{text[0]}), []byte(text)[1:]}, nil))
	if p {
		text += r.Punctuation()
	}
	return text
}

func (r *Serendipity) Paragraph() string {
	count := r.N(3, 7)
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = r.Sentence(true)
	}
	text := strings.Join(words, " ")
	return text
}

func (r *Serendipity) Punctuation() string {
	return r.Char(".?;!:")
}
