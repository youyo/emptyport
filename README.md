# emptyport

[![wercker status](https://app.wercker.com/status/36d0825cc499a52cf476c8379342c8ea/s/master "wercker status")](https://app.wercker.com/project/byKey/36d0825cc499a52cf476c8379342c8ea)

This Library find the tcp port that is not used.

## Installation

```
$ go get gopkg.in/youyo/emptyport.v1
```

## Usage

Just call `emptyport.Get()`.

```
import (
	"gopkg.in/youyo/emptyport.v1"
	"fmt"
)

p, _ := emptyport.Get()
fmt.Println(p)
```

## License

MIT License.
