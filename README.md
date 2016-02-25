# glash

Go Light-on-features tries-to-be-Anonymous SHell

### Features

1. Opinionated. Absolutely no external configuration. Customized to my needs.
2. Very light on features.
3. No history stored anywhere.
4. To be used on the go, not installed.
5. Provides a scratch directory in `/tmp` which will be cleaned out upon e`x`it'ing the shell.
6. Cleanup quickly using `Ctrl-\`.

### Motivation

Use different campus computers, to `ssh` back to my own machine in the network. Need to keep my contact on those machines as discrete as possible. This is a step in that direction.

### Compiling

I use Docker in my workflow.
Start a container:
```
docker run -it --rm -v `pwd`:/ws -w /ws golang
```

Build the image as :
```
GOOS=<osname> GOARCH=<arch> go build glash.go
```

Example:
```
GOOS=linux GOARCH=386 go build glash.go
```
builds for linux 32 bit.
