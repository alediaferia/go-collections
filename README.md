# go-collections
This project ports **map**-**filter**-**reduce** feature to [Go](http://golang.org).

Supported features
------------------

### Map

Given a collection it's trivial creating a new one editing each element on the fly.

```go
planets := []interface{}{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
    "Pluto",
}

coll := collections.NewFromSlice(planets)
coll = coll.Map(func(v interface{}) interface{} {
    return strings.Join([]string{ "Hello ", v.(string) })
})

fmt.Println(coll)
```

Output:
```go
    "Hello Mercury"
    "Hello Venus"
    "Hello Earth"
    "Hello Mars"
    "Hello Jupiter"
    "Hello Saturn"
    "Hello Uranus"
    "Hello Neptune"
    "Hello Pluto"
```

### Filter
Filtering a for even numbers only:

```go
numbers := []interface{}{
    1,
    3,
    5,
    6,
    7,
    8,
    9,
    10,
}

coll := collections.NewFromSlice(numbers)
coll = coll.Filter(func(v interface{}) bool {
    return v.(int) % 2 == 0
})
```

Output:
```go
    6,
    8,
    10,
```

### Reduce
Reducing for minimum number:

```go
numbers := []interface{}{
    1,
    3,
    5,
    6,
    7,
    8,
    9,
    10,
}

coll := collections.NewFromSlice(numbers)
min  := collections.Reduce(0, func(a, b interface{}) interface{} {
    if a < b { return a } else { return b }
})
```

Output:
```go
1
```

A complete example
------------------
The following example assumes a list of ages and returns the average of the ages greater or equal to 21.

```go
package main

import (
  "github.com/alediaferia/go-collections"
  "fmt"
)

func main() {
  ages := []interface{}{
    3,
    4,
    8,
    17,
    22,
    28,
    34,
    65,
    32,
    24,
  }

  coll := collections.NewFromSlice(ages)
  count := 0
  fmt.Printf("The average age of all the adult people at the party is: %d\n", coll.Filter(func(v interface{}) bool {
    return v.(int) >= 21
  }).Map(func(v interface{}) interface{} {
    count += 1
    return v
  }).Reduce(0, func(a, b interface{}) interface{} {
    return a.(int) + b.(int)
  }).(int) / count) 
}
```

License
-------
The code in this codebase is provided as-is as per the MIT License included [here](LICENSE).

Copyright (c) Alessandro Diaferia 2015