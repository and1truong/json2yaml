# json2yaml

A command-line utility to convert JSON files or stdin input to YAML format.

## Features

- 🚀 Convert JSON to YAML with proper formatting
- 📁 Support for both file input and stdin piping
- 🛠 Built using Cobra for professional CLI structure
- ✅ Comprehensive error handling
- 🧪 Full test coverage for all use cases
- 📦 Single binary distribution
- 🖥 Cross-platform support (Windows/Linux/macOS)

## Installation

### Prerequisites

- Go 1.16 or later

### Build from source

```bash
git clone https://github.com/and1truong/json2yaml.git
cd json2yaml
go build -o json2yaml
```

### Add to PATH (optional)

```bash
sudo mv json2yaml /usr/local/bin/
```

## Usage

### Convert a JSON file
```bash
./json2yaml input.json
```

### Convert from stdin
```bash
cat input.json | ./json2yaml
```

### Help command
```bash
./json2yaml --help
```

Output:
```
Convert JSON files or stdin input to YAML format.

Usage:
  json2yaml [file]

Examples:
  json2yaml input.json
  cat input.json | json2yaml

Flags:
  -h, --help   help for json2yaml
```

## Examples

**Input JSON file (example.json):**
```json
{
  "name": "John Doe",
  "age": 30,
  "hobbies": ["reading", "cycling"]
}
```

**Conversion command:**
```bash
./json2yaml example.json
```

**Output YAML:**
```yaml
age: 30
hobbies:
  - reading
  - cycling
name: John Doe
```

## Testing

Run the test suite with:
```bash
go test -v
```

Test coverage includes:
- Valid JSON file conversion
- Valid stdin input
- Invalid JSON handling
- Missing file errors
- Empty input handling
- Error message validation

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [YAML v3](https://github.com/go-yaml/yaml) - YAML support

## License

MIT License

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

Please report any issues or feature requests in the GitHub issue tracker.
