package lougo

import "testing"

var tests = map[string]string{
	"今年もよろしくお願いいたします。":         "ディスイヤーもよろしくプリーズいたします。",
	"信頼と安心のルー語となるよう邁進してまいります。": "トラストとピースオブマインドのルーランゲージとなるようプッシュフォワードしてまいります。",
}

func TestLou(t *testing.T) {
	for input, expected := range tests {
		output, err := Lou(input)
		if err != nil {
			t.Fatal(err)
		}
		if output != expected {
			t.Errorf(`
                input:    %v
                output:   %v
                expected: %v`, input, output, expected)
		}
	}
}
