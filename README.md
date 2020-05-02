# OpenSlides Keyformat Benchmark

This is a test script to check, which Output format for the
openslides-autoupdate-service has the best performance.

Run: `go test -bench=.` to get the result.

The script creates a map with 1.000 keys. Each map is then converted in a possible format.


## Formats

### Key Value Format

The KeyValue format is one map from string to the value. The key is a fqfield.
For example user/5/name. The value is the encoded value of the field.

```
{"a/1/f": "some value", "a/1/b/": 5, "b/1/f": "other value"}
```

The Benchmark is run twice. Once with the go default json package and once with
a manually build function.


### FQID Format

The FQID format is the format, that the datastore returns with its getMany
method. It is a map from fqid to a map of field to value.

```
{"a/1": {"f": "some value", "b": 5}, "b/1": {"f": "other value"}}
```

The Benchmark converts the key-value format into the FQID format and uses the
default json package to decode it.


### 3Parts Fromat

The 3Part format is three nested maps. The first is the collection, the second
the id and the third the field.

```
{"a": {"1": {"f": "some value", "b": 5}}, "b": {"1": {"f": "other value"}}}
```


## Benchmark

On my laptop the results are:

```
BenchmarkKeyValueBuildin-8         25570             47725 ns/op
BenchmarkKeyValueManully-8        146134              8360 ns/op
BenchmarkFQID-8                     7435            159182 ns/op
Benchmark3Parts-8                   7054            172860 ns/op
```

The important number is the last one. It tells, how long the function needed to
convert the given key-values into the format.