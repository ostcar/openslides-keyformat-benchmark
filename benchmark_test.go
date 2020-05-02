package main

import (
	"encoding/json"
	"testing"
)

func TestKeyValueManuelly(t *testing.T) {
	kv := buildKeyValue(10, 10, 10, `"some \"value\""`)
	encoded := encodeKeyValueManuelly(kv)

	var result map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &result); err != nil {
		t.Fatal(err)
	}

	if len(kv) != len(result) {
		t.Errorf("map not the same len")
	}

	for k := range kv {
		if kv[k] != string(result[k]) {
			t.Errorf("key %s not the same", k)
		}
	}

}

func BenchmarkKeyValueBuildin(b *testing.B) {
	kv := buildKeyValue(10, 10, 10, `"some value"`)

	for n := 0; n < b.N; n++ {
		encodeKeyValueBuildin(kv)
	}
}

func BenchmarkKeyValueManully(b *testing.B) {
	kv := buildKeyValue(10, 10, 10, `"some value"`)

	for n := 0; n < b.N; n++ {
		encodeKeyValueManuelly(kv)
	}

}

func BenchmarkFQID(b *testing.B) {
	kv := buildKeyValue(10, 10, 10, `"some value"`)

	for n := 0; n < b.N; n++ {
		encodeFQID(kv)
	}
}

func Benchmark3Parts(b *testing.B) {
	kv := buildKeyValue(10, 10, 10, `"some value"`)

	for n := 0; n < b.N; n++ {
		encode3Parts(kv)
	}

}
