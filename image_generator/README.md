
# Terminal image generator

`generate.py` runs all the `main.go` files under examples and converts terminal outputs to `html` and then `png`

## Prerequisites

* Go
* Python3.7 (preferred)
* Pip

### terminal-to-html

```
$ cd ~/
$ go get github.com/buildkite/terminal-to-html/cmd/terminal-to-html
```

### wkhtmltoimage

#### Arch

```
$ yaourt wkhtmltopdf
```

#### Ubuntu

```
$ sudo apt-get update
$ sudo apt-get install xvfb libfontconfig wkhtmltopdf
```

## Usage

```
$ python3.7 generate.py
```

After running the script, image files should be created under `assets/png`