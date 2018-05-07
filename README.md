# Bump

Bump handles semantic version updates.

## Installation

```
$ go install github.com/olivere/bump
```

## In a nutshell

```
$ echo v1.0.0 | bump
v1.0.1
$ echo 1.0.0 | bump
1.0.1
$ echo 1.0.1 | bump -kind=patch
1.0.2
$ echo 1.0.0 | bump -kind=minor
1.1.0
$ echo 1.0.0 | bump -kind=major
2.0.0
$ cat VERSION
v1.0.0
$ bump -kind=minor -i VERSION -o VERSION.out
$ cat VERSION.out
v1.0.1
$ bump -h
Usage of bump:

	bump [flags]

Flags:
  -i string
    	Input file (default: stdin)
  -kind string
    	Kind of update: major, minor, or patch (default: patch) (default "patch")
  -o string
    	Output file (default: stdout)
```

## License

MIT. See [LICENSE](https://github.com/olivere/bump//blob/master/LICENSE).
