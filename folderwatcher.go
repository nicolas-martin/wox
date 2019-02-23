package wox

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

const watchInterval = 100

// FolderWatcher ...
type FolderWatcher struct {
	Watcher *watcher.Watcher
	Folder  string
}

// NewFolderWatcher ...
func NewFolderWatcher(folder string) FolderWatcher {

	w := watcher.New()

	wa := FolderWatcher{
		Watcher: w,
		Folder:  folder,
	}

	wa.Watcher.IgnoreHiddenFiles(true)
	return wa
}

// Start ...
func (w FolderWatcher) Start() error {

	go func() {
		for {
			select {
			case event := <-w.Watcher.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Watcher.Error:
				log.Fatalln(err)
			case <-w.Watcher.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.Watcher.AddRecursive(w.Folder); err != nil {
		return err
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.Watcher.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {
		w.Watcher.Wait()
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Watcher.Start(time.Millisecond * watchInterval); err != nil {
		return err
	}

	return nil
}
