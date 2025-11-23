# Contributing to AbsOS

Thank you for your interest in contributing to AbsOS! We welcome contributions from the community.

## How to Contribute

### Reporting Issues

- Use the GitHub issue tracker to report bugs
- Describe the issue in detail, including steps to reproduce
- Include your Go version and operating system
- Provide code samples when applicable

### Suggesting Enhancements

- Open an issue to discuss your proposed enhancement
- Explain the use case and benefits
- Be open to feedback and discussion

### Pull Requests

1. **Fork the repository** and create your branch from the main branch
2. **Write clear commit messages** describing your changes
3. **Add tests** for new functionality
4. **Update documentation** as needed (README, GoDoc comments)
5. **Ensure tests pass** by running `go test ./...`
6. **Run go fmt** to format your code: `go fmt ./...`
7. **Submit a pull request** with a clear description of your changes

## Code Guidelines

### Go Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format code
- Run `go vet` to catch common issues
- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines

### Documentation

- Add GoDoc comments for all exported types, functions, and methods
- Use complete sentences in comments
- Start comments with the name of the element being described
- Include code examples for complex functionality

### Testing

- Write unit tests for new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions

### Interface Design

- Keep interfaces minimal and focused
- Design for extensibility
- Consider backward compatibility
- Document expected behavior clearly

## Implementing New Providers

If you're implementing support for a new object storage provider:

1. **Create a separate repository** for your implementation
2. **Implement the required interfaces**: `ObjectStore`, `Bucket`, `Object`, etc.
3. **Add comprehensive tests** using real or mocked provider APIs
4. **Document setup and usage** in your repository's README
5. **Open an issue or PR** to have your implementation linked from the main README

Your implementation does not need to be merged into this repository. We encourage separate packages for each provider to maintain modularity.

## Code of Conduct

### Our Standards

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Accept criticism gracefully

### Unacceptable Behavior

- Harassment or discriminatory language
- Personal attacks or trolling
- Publishing others' private information
- Other conduct inappropriate in a professional setting

## Questions?

Feel free to open an issue for questions or discussion. We're here to help!

## License

By contributing to AbsOS, you agree that your contributions will be licensed under the MIT License.
