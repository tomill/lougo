# ルーGo変換

**Lougo** is a golang version of perl5 ルー語変換 library [Acme::Lou](https://metacpan.org/pod/Acme::Lou).

## Example

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

## Depends

* mecab and utf-8 compiled system dictionary
* https://github.com/shogo82148/go-mecab/

## See also

* [lou5.jp](https://lou5.jp)
