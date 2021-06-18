package secp256k1

import (
	"io"
	"testing"

	crypto2 "github.com/fdymylja/tmos/x/authn/crypto"

	"github.com/fdymylja/tmos/x/authn/crypto/internal/benchmarking"
)

func BenchmarkKeyGeneration(b *testing.B) {
	b.ReportAllocs()
	benchmarkKeygenWrapper := func(reader io.Reader) crypto2.PrivKey {
		priv := genPrivKey(reader)
		return &PrivKey{Key: priv}
	}
	benchmarking.BenchmarkKeyGeneration(b, benchmarkKeygenWrapper)
}

func BenchmarkSigning(b *testing.B) {
	b.ReportAllocs()
	priv := GenPrivKey()
	benchmarking.BenchmarkSigning(b, priv)
}

func BenchmarkVerification(b *testing.B) {
	b.ReportAllocs()
	priv := GenPrivKey()
	benchmarking.BenchmarkVerification(b, priv)
}
