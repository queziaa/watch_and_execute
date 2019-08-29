package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// "github.com/queziaa/watch_and_execute/execute"
// "github.com/queziaa/watch_and_execute/watch"

func main() {

	// var uintChan = make(chan uint)
	// watch.InitWatch(`C:\Users\PC\Desktop\ycserver`, uintChan)
	// execute.InitExecute("\n~~~~~~~", uintChan)
	a()
	select {}
}
func a() {
	_ = os.Chdir("C:\\Users\\PC\\Desktop\\ycserver\\cmd")

	ctx, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx, "go", "run", "C:\\Users\\PC\\Desktop\\ycserver\\cmd\\main.go", "-c", "C:\\Users\\PC\\Desktop\\ycserver\\configs\\t.conf.json")
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	// if err := cmd.Run(); err != nil {
	// 	fmt.Printf("Run returns: %s\n", err)
	// }
	// fmt.Printf("Stderr: %s\n", string(w.Bytes()))

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Stdout = os.Stdout
	_ = cmd.Start()

	fmt.Println("3")
	time.Sleep(time.Second * 10)
	fmt.Println("退出程序中...", cmd.Process.Pid)
	cancel()

	_ = cmd.Process.Kill()
	_ = cmd.Wait()
}
