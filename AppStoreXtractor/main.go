/*
 * Monitor AppStore download cache directory and save packages to ~/Desktop.
 * Go re-implementation of
 * https://github.com/maxschlapfer/MacAdminHelpers/tree/master/AppStoreExtract
 * Heavily based on
 * https://github.com/fsnotify/fsevents/blob/master/example/main.go
 * https://godoc.org/github.com/DHowett/go-plist#example-Decoder-Decode
 */

package main

import (
  "bufio"
  "bytes"
  "fmt"
  "log"
  "os"
  "io/ioutil"
  "path"
  "runtime"
  "time"
  "strings"
  "github.com/fsnotify/fsevents"
  "github.com/DHowett/go-plist"
)

func main() {
  path := path.Clean( os.Getenv("TMPDIR") + "../C/com.apple.appstore" )
  destinationPath := os.Getenv("HOME") + "/Desktop"
  log.Println("Watching:", path)
  log.Println("Saving Packages to:", destinationPath)

  dev, err := fsevents.DeviceForPath(path)
  if err != nil {
    log.Fatalf("Failed to retrieve device for path:", err)
  }

  es := &fsevents.EventStream{
    Paths:   []string{path},
    Latency: 50 * time.Millisecond,
    Device:  dev,
    Flags:   fsevents.FileEvents | fsevents.WatchRoot}
  es.Start()
  ec := es.Events

  go func() {
    for msg := range ec {
      for _, event := range msg {
        handleFSevent(event,path,destinationPath)
      }
    }
  }()

  in := bufio.NewReader(os.Stdin)

  log.Print("Started, press enter to stop")
  in.ReadString('\n')
  runtime.GC()
  es.Stop()
}

func handleFSevent(event fsevents.Event, srcTopDir string, dest string) {
  if strings.HasSuffix(event.Path, ".pkg") {
    basename := path.Base(event.Path)
    if event.Flags & fsevents.ItemCreated == fsevents.ItemCreated {
      linkTarget := dest + "/" + getOutputFilename(srcTopDir+"/manifest.plist")
      log.Printf("Hardlink: %s to %s", basename, linkTarget)
      os.Link("/"+event.Path, linkTarget)
    }
    if event.Flags & fsevents.ItemRemoved == fsevents.ItemRemoved {
      log.Printf("Download completed: %s", basename)
    }
  }
}

func getOutputFilename(manifestPath string) string {
  mdata, err := ioutil.ReadFile(manifestPath)
  if err != nil {
    log.Fatal("Fatal error reading %s: %s", manifestPath, err)
  }
  type sparseBundleHeader struct {
      PkgData     []interface{} `plist:"representations"`
  }
  buf := bytes.NewReader(mdata)
  var data sparseBundleHeader
  decoder := plist.NewDecoder(buf)
  decoderError := decoder.Decode(&data)
  if decoderError != nil {
      fmt.Println(decoderError)
  }
  // (not only) this might deserve some error handling...?
  myMap := data.PkgData[0]
  me := myMap.(map[string]interface{})
  appVersion := me["bundle-version"]
  appTitle := me["title"]
  appTitleNoWhiteSpace := strings.Replace(fmt.Sprintf("%s", appTitle), " ", "_", -1)

  return fmt.Sprintf("%s-%s.pkg", appTitleNoWhiteSpace, appVersion)
}
