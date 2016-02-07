# emptyport

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
