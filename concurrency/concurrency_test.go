package concurrency

import (
	"testing"
	"time"
)

func slowStubVerificationWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerifyWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "one url"
	}

	for i := 0; i < b.N; i++ {
		VerifyWebsites(slowStubVerificationWebsite, urls)
	}
}
