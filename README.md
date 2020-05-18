# iio
IIO is a simple machine learning tool, written in Go. Its main purpose is OCR problem solving.
Under the hood there's a primitive feed forward architecture, designed to recognize numeric
images (mainly, MNIST database). All contributions and support are welcome! 

## Usage
This software is powered by [lets](https://github.com/lets-cli/lets) - we strongly recommend you
to install and use this CLI. If you have no opportunity to install this command runner, look the
desired commands up in [lets.yaml](https://github.com/xXxRisingTidexXx/iio/blob/master/lets.yaml).
After the repository clonning run these instructions in the project root:
```bash
$ lets workspace
$ lets compile
$ lets run
```
If you need more info about available terminal commands, run this one:
```bash
$ lets help
```
