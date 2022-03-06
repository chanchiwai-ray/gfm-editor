# GitHub Flavor Markdown Editor

A simple web-based editor that allows you to preview the content of a markdown file just as what you will see in GitHub.

## Installation

### Local Build

```go
go build && go install
```

## Usage

```
Usage of github-markdown:

A simple web-based editor that allows you to preview the
content of a markdown file just as what you will see in GitHub.

Options:
  -f string
    	The path to the markdown file (Required).
  -port string
    	Start the webserver at this port. (default "8080")
```

## Example

```
github-markdown -f READMD.md
```
