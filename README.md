# bcashe
Golang small cache library

## Use

```golang
package main

import (
    "fmt"
    "github.com/mirrr/bcashe"
    "time"
)

func main() {
    cache := bcache.Create()
    cache.Updater(func(key string) interface{} {
        if key == "mykey" {
            return 777
        }
        return nil
    })
    cache.Set("mykey2", 888)

    fmt.Println(cache.Get("mykey"))  // 777
    fmt.Println(cache.Get("mykey2")) // 888

    time.Sleep(time.Second * 2)

    fmt.Println(cache.Get("mykey"))  //  777
    fmt.Println(cache.Get("mykey2")) //  <nil>
}
```

\* Any key kept in cache no longer than a second. Byt upon receipt of the value life time is extended for a second. 

## Documentation
[Docs on godoc.org](https://godoc.org/github.com/mirrr/bcashe)


## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2015 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
