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
	x := &indexObjectWithSecondaryKey{
		objectPrefix:      []byte("some-proto-object"),
		indexPrefix:       []byte("some-index-name"),
		secondaryKeyValue: []byte("some-secondary-key-value"),
		primaryKey:        []byte("some-primary-key"),
	}
	y := &indexObjectWithSecondaryKey{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts := x.marshal()
		err := y.unmarshal(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
