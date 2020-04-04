# iio
IIO is a simple machine learning tool, written in Go. Its main purpose is OCR problem solving.
Under the hood there's a primitive 3-layer perceptron architecture, designed to recognize numeric
images (mainly, MNIST database). All contributions and support are welcome! 

## Usage
> This tool is powered by [lets](https://github.com/lets-cli/lets) - we strongly recommend you to
> install and use this CLI. 

Clone this repository and in the project root run this instruction
```bash
$ lets workspace
``` 
This will install all required dependencies and libraries. After that compile the sources
```bash
$ lets compile
```
And run the script
```bash
$ lets run
```

## Code quality
Run code formatter
```bash
$ lets format
```
or linter ([golangci-lint](https://github.com/golangci/golangci-lint) must be installed)
```bash
$ lets lint
```

## Tests
Execute all test suites with a single operation
```bash
$ lets test
```

## Help
Info about available terminal commands
```bash
$ lets help
```
