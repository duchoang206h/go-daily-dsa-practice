# DSA Practice

Daily Golang DSA practice.

## Structure

Problems are organized by category:
- `strings/` - String manipulation problems
- _(more categories updating...)_

Each problem contains:
- `main.go` - Solution implementation
- `main_test.go` - Test cases

## Running Tests

Navigate to a problem directory and run:

```bash
go test -v
```

To generate coverage:

```bash
go test -cover -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

