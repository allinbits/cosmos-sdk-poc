package indexes

import "testing"

func Benchmark_typedPrefixedKey_bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		typePrefixedKey{
			primaryKey: []byte("primary_key"),
			typePrefix: []byte("osu.King.v1alpha1"),
		}.bytes()
	}
}

func Benchmark_IndexerKey(b *testing.B) {
	x := &indexerKey{
		objectPrefix: []byte("some-proto-object"),
		indexName:    []byte("some-index-name"),
		indexValue:   []byte("some-index-value"),
		primaryKey:   []byte("some-primary-key"),
	}
	y := &indexerKey{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts := x.marshal()
		err := y.unmarshal(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
