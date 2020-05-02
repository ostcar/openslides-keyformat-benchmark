package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	kv := buildKeyValue(10, 10, 10, `"some value"`)
	fmt.Println(kv)
}

func buildKeyValue(collections, ids, fields int, value string) map[string]string {
	kv := make(map[string]string, collections*ids*fields)
	for i := 0; i < collections; i++ {
		for j := 0; j < ids; j++ {
			for k := 0; j < fields; j++ {
				kv[fmt.Sprintf("c%d/%d/f%d", i, j, k)] = value
			}
		}
	}
	return kv
}

func encodeKeyValueBuildin(kv map[string]string) []byte {
	b, err := json.Marshal(kv)
	if err != nil {
		panic(err)
	}
	return b
}

func encodeKeyValueManuelly(kv map[string]string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte('{')
	for key, value := range kv {
		buf.WriteByte('"')
		buf.WriteString(key)
		buf.WriteString(`":`)
		buf.WriteString(value)
		buf.WriteString(`,`)
	}
	bs := buf.Bytes()[:buf.Len()-1]
	bs = append(bs, '}')
	return bs
}

func encodeFQID(kv map[string]string) []byte {
	fqidValues := make(map[string]map[string]json.RawMessage)
	for k, v := range kv {
		keyParts := strings.SplitN(k, "/", 3)
		if len(keyParts) != 3 {
			log.Panicf("invalid key %s, %d", k, len(k))
		}

		fqid := keyParts[0] + "/" + keyParts[1]
		if fqidValues[fqid] == nil {
			fqidValues[fqid] = make(map[string]json.RawMessage)
		}

		fqidValues[fqid][keyParts[2]] = json.RawMessage(v)
	}
	bs, err := json.Marshal(fqidValues)
	if err != nil {
		panic(err)
	}
	return bs
}

func encode3Parts(kv map[string]string) []byte {
	fqidValues := make(map[string]map[string]map[string]json.RawMessage)
	for k, v := range kv {
		keyParts := strings.SplitN(k, "/", 3)
		if len(keyParts) != 3 {
			log.Panicf("invalid key %s, %d", k, len(k))
		}

		if fqidValues[keyParts[0]] == nil {
			fqidValues[keyParts[0]] = make(map[string]map[string]json.RawMessage)
		}
		if fqidValues[keyParts[0]][keyParts[1]] == nil {
			fqidValues[keyParts[0]][keyParts[1]] = make(map[string]json.RawMessage)
		}

		fqidValues[keyParts[0]][keyParts[1]][keyParts[2]] = json.RawMessage(v)
	}
	bs, err := json.Marshal(fqidValues)
	if err != nil {
		panic(err)
	}
	return bs
}
