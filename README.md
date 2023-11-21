# how to?
clone and:
```
make
make test
./pack-sizes
```

other shells session:
```
curl -s "localhost:12345/split?orders=12001" | jq .
[
  {
    "bucketSize": 5000,
    "count": 2
  },
  {
    "bucketSize": 2000,
    "count": 1
  },
  {
    "bucketSize": 250,
    "count": 1
  }
]
```

# implementation
Implements two [change-making algorithms](https://en.wikipedia.org/wiki/Change-making_problem) - greedy and dymamic, uses first by default due to performance 
reasons.
