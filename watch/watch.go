package watch

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	watch    *fsnotify.Watcher
	timeChan chan uint
}

func InitWatch(path string, timeChan chan uint) {
	ww, _ := fsnotify.NewWatcher()
	w := Watch{
		watch:    ww,
		timeChan: timeChan,
	}
	w.watchDir(path)
}

//监控目录
func (w *Watch) addTimeChan() {
	w.timeChan <- uint(time.Now().UnixNano() % 1000000000000 / 1000)
}

func (w *Watch) watchDir(dir string) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						w.addTimeChan()
						fmt.Println("创建 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							_ = w.watch.Add(ev.Name)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						w.addTimeChan()
						fmt.Println("写入 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						w.addTimeChan()
						fmt.Println("删除 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							_ = w.watch.Remove(ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						w.addTimeChan()
						fmt.Println("重命名 : ", ev.Name)
						_ = w.watch.Remove(ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
