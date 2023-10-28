package watcher

import (
	"github.com/choigonyok/goopt/pkg/manifest"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher
}

// New create new fsnotify file watcher
func New() (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	return &Watcher{watcher: watcher}, err
}

// GetEventChannel returns file change event channel
func (w *Watcher) GetEventChannel() chan fsnotify.Event {
	return w.watcher.Events
}

// WatchFile add file path that watcher detects
func (w *Watcher) WatchFile(m manifest.Manifest) error {
	fullFilePath := m.Path + "/" + m.Name + "." + m.Extension
	return w.watcher.Add(fullFilePath)
}

// Close stop watching file
func (w *Watcher) Close() error {
	return w.watcher.Close()
}

// GetWatchList returns watcher's currently watching file list
func (w *Watcher) GetWatchList() []string {
	return w.watcher.WatchList()
}
