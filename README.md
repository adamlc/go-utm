# UTM Tag Builder
Use UTM Tag Builder to add Google utm tags to a URL.

## Installation
```bash
go get -u github.com/adamlc/go-utm
```

## Usage

```go
url := utm.BuildURL("https://test.com", Config{
	Source: "google",
	Medium: "email",
	Campaign: "Awesome Test"},
)

// https://test.com?utm_campaign=Awesome+Test&utm_medium=email&utm_source=google
```

#### func  BuildURL

```go
func BuildURL(rawurl string, config Config) (string, error)
```
BuildURL adds UTM tags to the passed URL

#### type Config

```go
type Config struct {
	Source   string
	Medium   string
	Campaign string
	Content  string
	Term     string
}
```

Config is used to configure UTM tags for a URL