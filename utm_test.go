package utm

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	tables := []struct {
		url      string
		config   Config
		expected string
		err      error
	}{
		{"https://test.com", Config{Source: "google", Medium: "email", Campaign: "Awesome Test"}, "https://test.com?utm_campaign=Awesome+Test&utm_medium=email&utm_source=google", nil},
		{"https://test.com?hello=world", Config{Source: "google", Medium: "email", Campaign: "Awesome Test"}, "https://test.com?hello=world&utm_campaign=Awesome+Test&utm_medium=email&utm_source=google", nil},
		{"https://test.com", Config{Source: "bing", Medium: "search", Campaign: "Some Awesome &^%%^ Campaign", Term: "winner winner chicken dinner"}, "https://test.com?utm_campaign=Some+Awesome+%26%5E%25%25%5E+Campaign&utm_medium=search&utm_source=bing&utm_term=winner+winner+chicken+dinner", nil},
		{"", Config{}, "", ErrURLRequired},
		{"ht//test.com", Config{Source: "google", Medium: "email", Campaign: "Awesome Test"}, "", ErrInvalidURL},
		{"//test", Config{Source: "google", Medium: "email", Campaign: "Awesome Test"}, "", ErrInvalidURL},
		{"https://test.com", Config{}, "", ErrSourceRequired},
		{"https://test.com", Config{Source: "google"}, "", ErrMediumRequired},
		{"https://test.com", Config{Source: "google", Medium: "email"}, "", ErrCampaignRequired},
		{"https://test.com?hello=world", Config{Source: "google", Medium: "email", Campaign: "Awesome", Content: "Hello World"}, "https://test.com?hello=world&utm_campaign=Awesome&utm_content=Hello+World&utm_medium=email&utm_source=google", nil},
		{"https://test.com?hello=world#foobar", Config{Source: "google", Medium: "email", Campaign: "Awesome", Content: "Hello World"}, "https://test.com?hello=world&utm_campaign=Awesome&utm_content=Hello+World&utm_medium=email&utm_source=google#foobar", nil},
	}

	for _, table := range tables {
		result, err := BuildURL(table.url, table.config)

		if err != table.err {
			if err != nil {
				t.Errorf("BuildURL threw an unexpected error, got: %s.", err.Error())
			} else {
				t.Errorf("BuildURL expected to throw error: %s.", table.err)
			}
		}

		if result != table.expected {
			t.Errorf("BuildURL output not expected, got: %s, want: %s.", result, table.expected)
		}
	}
}
