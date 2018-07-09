package utm

import (
	"errors"
	"net/url"
)

// Config is used to configure UTM tags for a URL
type Config struct {
	Source   string
	Medium   string
	Campaign string
	Content  string
	Term     string
}

var (
	// ErrURLRequired is thrown when the URL passed is empty
	ErrURLRequired = errors.New("A URL is required")

	// ErrInvalidURL is thrown when the passed URL doesn't validate
	ErrInvalidURL = errors.New("The URL is invalid")

	// ErrSourceRequired is thrown when no source is passed
	ErrSourceRequired = errors.New("A Source is required")

	// ErrMediumRequired is thrown when no source is passed
	ErrMediumRequired = errors.New("A Medium is required")

	// ErrCampaignRequired is thrown when no campaign is passed
	ErrCampaignRequired = errors.New("A Campaign is required")
)

// BuildURL adds UTM tags to the passed URL
func BuildURL(rawurl string, config Config) (string, error) {
	if rawurl == "" {
		return "", ErrURLRequired
	}
	if config.Source == "" {
		return "", ErrSourceRequired
	}
	if config.Medium == "" {
		return "", ErrMediumRequired
	}
	if config.Campaign == "" {
		return "", ErrCampaignRequired
	}

	// First parse the URL
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return "", ErrInvalidURL
	}

	// Now add any new params to the query string
	query := url.Query()
	query.Add("utm_source", config.Source)
	query.Add("utm_medium", config.Medium)
	query.Add("utm_campaign", config.Campaign)

	if config.Content != "" {
		query.Add("utm_content", config.Content)
	}

	if config.Term != "" {
		query.Add("utm_term", config.Term)
	}

	//Reassign the query back
	url.RawQuery = query.Encode()

	return url.String(), nil
}
