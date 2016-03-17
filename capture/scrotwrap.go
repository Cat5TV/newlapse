package capture

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

const (
	// Quality of images
	Quality = "50"
)

// Folder does everything
func Folder(dest string, interv int) {
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	tick := time.Tick(time.Duration(interv) * time.Second)
	callScrot(dest, time.Now())
	pic1, _ := ioutil.ReadDir(dest)
	var size1 float64
	for _, f := range pic1 {
		if n := f.Size(); n > 4096 {
			size1 = float64(n) / float64(interv)
			break
		}
	}
	if size1 <= 0 {
		fmt.Println("cannot estimate bytes/second")
	} else {
		fmt.Printf("%5.3E byte/sec = %5.3f kb/sec = %5.3f mb/sec\n", size1, size1/1000.0, size1/1000000.0)
		size1 *= 60
		fmt.Printf("%5.3E byte/min = %5.3f kb/min = %5.3f mb/min\n", size1, size1/1000.0, size1/1000000.0)
		fmt.Printf("1GB of storage will be filled in %.2f Minutes\n", 1e9/size1)
	}
	var numpic int64
	numpic = int64(len(pic1))
	fmt.Printf("picture #%010d taken\n", numpic)
	numpic++
scrotloop:
	for {
		select {
		case <-sigs:
			break scrotloop
		case t := <-tick:
			callScrot(dest, t)
			fmt.Printf("picture #%010d taken\n", numpic)
			numpic++
		}
	}
	signal.Reset(os.Interrupt, os.Kill)
	fmt.Println("")
}

func callScrot(dest string, n time.Time) {
	filename := fmt.Sprintf("%s/%d-%02d-%02d_%02d-%02d-%02d_$wx$h.png", dest, n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
	cmd := exec.Command("scrot", "-q", Quality, "-z", filename)
	cmd.Run()
}

func getImageSize(file string) int64 {
	f, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return f.Size()
}
