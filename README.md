# urlify-go
Transliterates non-ascii characters for use in URLs.  
Inspired from https://github.com/jbroadway/urlify  

### Install  
```
$ go get https://github.com/twisted1919/evs-go.git  
```

### Usage
```
package main  

import (  
	"fmt"  
	urlify "github.com/twisted1919/urlify-go"  
)  

func main() {  
	u := urlify.NewParser()  
	fmt.Println(u.SetText("Lo siento, no hablo español.").Parse())  
	fmt.Println(u.SetText(" J'étudie le français ").Parse())  
	fmt.Println(u.SetText("Some special chars: Çćč ßśš ûüùúū į").Parse())  
	fmt.Println(u.AddToRemoveList("super").SetText("Like a super man").Parse())  
	fmt.Println(u.RemoveFromRemoveList("super").SetText("Like a super man").Parse())  
}

```

The above will output:  
```
lo-siento-no-hablo-espanol  
jetudie-le-francais  
some-special-chars-ccc-ssss-uueuuu-i  
like-man  
like-super-man  
```
