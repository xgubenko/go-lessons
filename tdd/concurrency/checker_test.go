package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://some.site"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://vk.com",
			"http://some.site",
		}

		want := map[string]bool{
			"http://google.com": true,
			"http://vk.com":     true,
			"http://some.site":  false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v, got %v", want, got)
		}
	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
