# golang-playground

Go fun projects for educational purposes only.

Currently there's not much to see here, but:

- AppStoreXtractor -- a Go re-implementation of [AppStore Extractor](https://github.com/maxschlapfer/MacAdminHelpers/tree/master/AppStoreExtract)
  - uses [fsevents](https://github.com/fsnotify/fsevents) to monitor cache folder
  - uses [go-plist](github.com/DHowett/go-plist) to read the `manifest.plist`
  - does not rely on any external tools

# Build projects

If checked out (or symlinked) to `~/go`, setting GOPATH as outlined below is not required.

... or just run `make` -- avoids need to set `GOPATH` manually.

## fish build example

```bash
set GOPATH (pwd)
export GOPATH
go install AppStoreXtractor
bin/AppStoreXtractor
```

## bash build example

```bash
export GOPATH=`pwd`
go install AppStoreXtractor
bin/AppStoreXtractor
```
