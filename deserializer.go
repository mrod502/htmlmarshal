package htmlmarshal

import (
	"bytes"

	"golang.org/x/net/html"
)

type Deserializer interface {
	Deserialize([]byte, interface{}) error
}

func NewDeserializer() Deserializer {
	var out = new(htmlDeserializer)

	return out
}

type htmlDeserializer struct {
}

func (h *htmlDeserializer) Deserialize(b []byte, v interface{}) error {
	t := html.NewTokenizer(bytes.NewReader(b))
	_, err := tokenizer2Map(t)
	return err
}

func tokenizer2Map(t *html.Tokenizer) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
