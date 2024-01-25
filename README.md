<div align="center">
	<h1>RollKeys</h1>
	<p>
		<b>Make your API Credentials Rotational</b>
	</p>
	<br>
	<br>
	<br>
</div>

This is a simple library that allows you to rotate your API credentials. 
Let's say there is a free API key limit 5 request per second. And to avoid this limitation is to create multiple API keys.
and use them in rotation.

# install
```bash
go get github.com/FPNL/RollKeys
```

# usage example
```go
package main

import (
	"context"
	"log"

	rollKeys "github.com/FPNL/RollKeys"
	"golang.org/x/sync/errgroup"
)

func main() {
	// means each key could make 5 request per second
	givenAPIKeys := []string{"key1", "key2"}
	givenRequestTimes := 100
	givenRate := 5

	keys, err := rollKeys.NewRotationalSlice(givenAPIKeys, givenRate)
	if err != nil {
		log.Fatal(err)
	}

	eg := errgroup.Group{}

	for i := 0; i < givenRequestTimes; i++ {
		eg.Go(func() error {
			// if all the keys are fully loaded, it will wait.
			// most expect wait duration is (givenRequestTimes / (len(givenAPIKeys) * givenRate)) - 1
			// if there is 100 request at same time, in this example is 100 / (2 * 5) - 1 = 9 seconds
			key, err := keys.Get(context.TODO())
			if err != nil {
				return err
			}

			// Do what you want to key
			log.Println(key)

			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
```
