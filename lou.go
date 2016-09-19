package lougo

import (
	"strings"
	"sync"

	"github.com/ikawaha/kagome/tokenizer"
)

var (
	loukenizer tokenizer.Tokenizer
	once       sync.Once
)

// Lou returns filtered input text like as Lou Ohshiba. Let's together!!
func Lou(input string) (output string, err error) {
	once.Do(func() {
		loukenizer, err = loader()
	})
	if err != nil {
		return
	}

	tokens := loukenizer.Tokenize(input)

	var out []string
	for _, token := range tokens {
		switch token.Class.String() {
		case "DUMMY":
			continue

		case "USER":
			lou := token.Features()[2]
			out = append(out, lou)

		default:
			out = append(out, token.Surface)
		}
	}

	output = strings.Join(out, "")
	return
}

func loader() (t tokenizer.Tokenizer, err error) {
	f, err := Assets.Open("/data/loudic.txt")
	if err != nil {
		return
	}

	r, err := tokenizer.NewUserDicRecords(f)
	if err != nil {
		return
	}

	udic, err := r.NewUserDic()
	if err != nil {
		return
	}

	t = tokenizer.New()
	t.SetUserDic(udic)
	return
}
