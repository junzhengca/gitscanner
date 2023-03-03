# gitscanner

An utility to scan all public GitHub repositories in an organization for leaks.

## Requirements

* [gitleaks](https://github.com/zricethezav/gitleaks) (for kali you can just apt install)
* A GitHub personal access token

## Installation

```
$ go install
```

## Usage

```
NAME:
   gitscanner - Scan public git repositories for vulnerabilities

USAGE:
   gitscanner [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value, -t value  GitHub token, you can generate one on GitHub settings page
   --org value, -o value    GitHub organization name
   --help, -h               show help
```

After running the program, output will be written to `findings` directory.