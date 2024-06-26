package fs_watcher

func (w *Watcher) GetEventsChan() <-chan string {
	return w.events
}
