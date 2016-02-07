# emptyport

[![Circle CI](https://circleci.com/gh/youyo/emptyport.svg?style=svg)](https://circleci.com/gh/youyo/emptyport)

This Library find the tcp port that is not used.

## Usage

Just call `emptyport.Get()`.

```
import (
	"github.com/youyo/emptyport"
	"fmt"
)

p := emptyport.Get()
fmt.Println(p)
```

## License

MIT License.
