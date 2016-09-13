package lougo

import (
	"os"
	"path"
	"strings"

	mecab "github.com/shogo82148/go-mecab"
)

var args = map[string]string{
	"userdic": path.Join(os.Getenv("GOPATH"), "src", "github.com/tomill/lougo", "lou.dic"),
}

// Lou returns filtered input text like as Lou Ohshiba. Let's together!!
func Lou(input string) (output string, err error) {
	m, err := mecab.New(args)
	if err != nil {
		return
	}
	defer m.Destroy()

	node, err := m.ParseToNode(input)
	if err != nil {
		return
	}

	var out []string
	for ; node != (mecab.Node{}); node = node.Next() {
		out = append(out, translate(node))
	}

	output = strings.Join(out, "")
	return
}

type feature struct {
	surface                           string
	pos, iType, iForm, lemma, reading string
	lou, aux                          string
}

func scan(node mecab.Node) *feature {
	if node == (mecab.Node{}) || node.Stat() != mecab.NormalNode {
		return &feature{}
	}

	v := strings.Split(node.Feature()+",,,,", ",")

	return &feature{
		surface: node.Surface(),
		pos:     v[0],
		iType:   v[4],
		iForm:   v[5],
		lemma:   v[6],
		reading: v[7],
		lou:     v[9],  // extra column from userdic
		aux:     v[10], // extra column from userdic
	}
}

func translate(node mecab.Node) string {
	curr := scan(node)
	next := scan(node.Next())

	switch {
	case curr.pos == "接頭詞" && (curr.reading == "ゴ" || curr.reading == "オ") && next.lou != "":
		return ""
	case curr.lou == "":
		return curr.surface // no change
	case curr.pos == "形容詞" && curr.iForm == "基本形" && (next.pos == "助詞" || next.pos == "記号"):
		return curr.lou // no aux
	default:
		return curr.lou + curr.aux
	}
}
