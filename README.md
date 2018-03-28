# smartcrop-cli

[![GitHub release](http://img.shields.io/github/release/jakewarren/smartcrop-cli.svg?style=flat-square)](https://github.com/jakewarren/smartcrop-cli/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/smartcrop-cli/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/smartcrop-cli)](https://goreportcard.com/report/github.com/jakewarren/smartcrop-cli)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)
> smartcrop finds good image crops for arbitrary crop sizes

A command line client for the wonderful [smartcrop](https://github.com/muesli/smartcrop) library.

## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/smartcrop-cli/releases/latest](https://github.com/jakewarren/smartcrop-cli/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/smartcrop-cli
```

## Usage
```
❯ smartcrop-cli --help
usage: smartcrop-cli [<flags>] <input>

smartcrop finds good image crops for arbitrary crop sizes

Optional flags:
      --help        Show context-sensitive help (also try --help-long and --help-man).
  -w, --width=250   preferred width of image (px)
  -h, --height=250  preferred height of image (px)
  -o, --output-file="./smartcrop.jpg"  
                    name of the output file
  -V, --version     Show application version.

Args:
  <input>  input file
```
## Changes

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this project adheres to [Semantic Versioning](http://semver.org/).

## License

MIT © 2018 Jake Warren

[changelog]: https://github.com/jakewarren/smartcrop-cli/blob/master/CHANGELOG.md
