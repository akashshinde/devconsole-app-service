package main

import "github.com/redhat-developer/app-service/watcher"

func main() {
	w := watcher.NewWatcher("myproject")
	w.StartWatcher()
	w.ListenWatcher()
}
