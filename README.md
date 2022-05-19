
map+RWMuter vs sync.Map

```BenchmarkRwMapStore
BenchmarkRwMapStore-4                    3363078               344.8 ns/op            55 B/op          0 allocs/op
BenchmarkSyncMapStore
BenchmarkSyncMapStore-4                  1000000              1252 ns/op             179 B/op          5 allocs/op
BenchmarkRwMapLoad
BenchmarkRwMapLoad-4                     5072104               238.1 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapLoad
BenchmarkSyncMapLoad-4                   2779568               404.8 ns/op             0 B/op          0 allocs/op
BenchmarkRwMapDelete
BenchmarkRwMapDelete-4                   8844811               134.4 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapDelete
BenchmarkSyncMapDelete-4                 8576566               131.9 ns/op             0 B/op          0 allocs/op
BenchmarkRwMapStoreAndLoad
BenchmarkRwMapStoreAndLoad-4              105675             13457 ns/op            5287 B/op         15 allocs/op
BenchmarkSyncMapStoreAndLoad
BenchmarkSyncMapStoreAndLoad-4             69193             35049 ns/op            5545 B/op          6 allocs/op```