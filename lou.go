package lougo

import (
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

// Lou returns filtered input text like as Lou Ohshiba. Let's together!!
func Lou(input string) (output string, err error) {
	t, err := louTokenizer()
	if err != nil {
		return
	}

	tokens := t.Tokenize(input)

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

func louTokenizer() (t tokenizer.Tokenizer, err error) {
	t = tokenizer.New()

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

	t.SetUserDic(udic)
	return
}
