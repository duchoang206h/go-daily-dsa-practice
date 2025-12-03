.PHONY: new test clean help

# Create new problem folder with templates
new:
ifndef path
	@echo "Usage: make new path=category/problem-name"
	@echo "Example: make new path=strings/two-sum"
	@exit 1
endif
	@echo "Creating new problem at $(path)..."
	@mkdir -p $(path)
	@printf 'package main\n\nimport "fmt"\n\n// TODO: Implement solution here\nfunc solution() {\n\t// Your code here\n}\n\nfunc main() {\n\tfmt.Println("Solution not implemented yet")\n}\n' > $(path)/main.go
	@printf 'package main\n\nimport (\n\t"testing"\n)\n\n// TODO: Add test cases\nfunc TestSolution(t *testing.T) {\n\ttests := []struct {\n\t\tname     string\n\t\tinput    interface{}\n\t\texpected interface{}\n\t}{\n\t\t{\n\t\t\tname:     "Test case 1",\n\t\t\tinput:    nil,\n\t\t\texpected: nil,\n\t\t},\n\t}\n\n\tfor _, test := range tests {\n\t\tt.Run(test.name, func(t *testing.T) {\n\t\t\t// result := solution(test.input)\n\t\t\t// if result != test.expected {\n\t\t\t//     t.Errorf("got %%v, want %%v", result, test.expected)\n\t\t\t// }\n\t\t\tt.Skip("Not implemented yet")\n\t\t})\n\t}\n}\n' > $(path)/main_test.go
	@echo "✅ Created $(path)/"
	@echo "   - main.go"
	@echo "   - main_test.go"

# Run tests for a specific problem
test:
ifndef path
	@echo "Usage: make test path=category/problem-name"
	@echo "Example: make test path=strings/two-sum"
	@echo ""
	@echo "Or run: make test-all to test everything"
	@exit 1
endif
	@echo "Running tests for $(path)..."
	@cd $(path) && go test -v

# Run all tests
test-all:
	@echo "Running all tests..."
	@go test ./...

# Run tests with coverage for a specific problem
coverage:
ifndef path
	@echo "Usage: make coverage path=category/problem-name"
	@echo "Example: make coverage path=strings/two-sum"
	@exit 1
endif
	@echo "Running tests with coverage for $(path)..."
	@cd $(path) && go test -cover -coverprofile=coverage.out
	@cd $(path) && go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: $(path)/coverage.html"

# Clean coverage files
clean:
	@echo "Cleaning coverage files..."
	@find . -name "coverage.out" -delete
	@find . -name "coverage.html" -delete
	@echo "✅ Cleaned"

# Show help
help:
	@echo "Available commands:"
	@echo ""
	@echo "  make new path=category/problem-name"
	@echo "    Create a new problem with template files"
	@echo "    Example: make new path=strings/two-sum"
	@echo ""
	@echo "  make test path=category/problem-name"
	@echo "    Run tests for a specific problem"
	@echo "    Example: make test path=strings/longest-palindromic-substring"
	@echo ""
	@echo "  make test-all"
	@echo "    Run all tests in the repository"
	@echo ""
	@echo "  make coverage path=category/problem-name"
	@echo "    Generate coverage report for a specific problem"
	@echo "    Example: make coverage path=strings/two-sum"
	@echo ""
	@echo "  make clean"
	@echo "    Remove all coverage files"
	@echo ""
	@echo "  make help"
	@echo "    Show this help message"

