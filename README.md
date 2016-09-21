# ルーGo変換

**Lougo** is a golang version of perl5 ルー語変換 library [Acme::Lou](https://metacpan.org/pod/Acme::Lou).

## Download

<https://github.com/tomill/lougo/releases>

## Usage

```
$ echo "人生には、大切な三つの袋があります。" | lougo
ライフには、インポータントな三つの袋があります。
```

## Library

```
package main

import (
	"fmt"

	"github.com/tomill/lougo"
)

func main() {
	out, _ := lougo.Lou("信頼と安心のルー語となるよう邁進してまいります。")

	fmt.Printf(out)
	// Output: トラストとピースオブマインドのルーランゲージとなるようプッシュフォワードしてまいります。
}
```

## See also

* https://lou5.jp
* https://github.com/ikawaha/kagome/
* https://github.com/tomill/Acme-Lou
