# golang-playground

Go fun projects for educational purposes only.

Currently there's not much to see here, but...

## AppStoreXtractor -- a Go re-implementation of [AppStore Extractor](https://github.com/maxschlapfer/MacAdminHelpers/tree/master/AppStoreExtract) [for Mac]

- uses [fsevents](https://github.com/fsnotify/fsevents) to monitor cache folder
- uses [go-plist](github.com/DHowett/go-plist) to read the `manifest.plist`
- does not rely on any external tools

Download a [release](../../releases) or build from source:

```bash
go get github.com/schnoddelbotz/golang-playground/AppStoreXtractor
```

## License

[WTFPL](http://www.wtfpl.net/)
