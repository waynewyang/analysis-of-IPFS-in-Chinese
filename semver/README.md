## semver
用途版本管理,暂不深究

## 示例

```
  3 import (
  4     "fmt"
  5     "gx/ipfs/QmcrrEpx3VMUbrbgVroH3YiYyUS5c4YAykzyPJWKspUYLa/go-semver/semver"
  6     "os"
  7 )
  8 
  9 func main() {
 10     vA, err := semver.NewVersion(os.Args[1])
 11     if err != nil {
 12         fmt.Println(err.Error())
 13     }
 14     vB, err := semver.NewVersion(os.Args[2])
 15     if err != nil {
 16         fmt.Println(err.Error())
 17     }
 18 
 19     fmt.Printf("%s < %s == %t\n", vA, vB, vA.LessThan(*vB))
 20 }


wayne@wayne:~/go/src/gx/ipfs/QmcrrEpx3VMUbrbgVroH3YiYyUS5c4YAykzyPJWKspUYLa/go-semver$ ./example 1.1.1 2.2.2
1.1.1 < 2.2.2 == true
wayne@wayne:~/go/src/gx/ipfs/QmcrrEpx3VMUbrbgVroH3YiYyUS5c4YAykzyPJWKspUYLa/go-semver$ 
```