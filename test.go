package main

import (
	"fmt"
	//"strings"
	"os"
	"os/exec"
	"runtime"
	"time"
	"github.com/mackerelio/go-osstat/memory"
	. "github.com/klauspost/cpuid"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
func boot() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Println("AMIBIOS(C) 2021 American Megatrends, Inc.")
	fmt.Println("Birb-OS Bios Date: 14/04/21 Rev: 1.0")
	time.Sleep(2 * time.Second)
	fmt.Println("CPU  :  ", CPU.BrandName)
	fmt.Println("Count:  ", CPU.LogicalCores)
	time.Sleep(2 * time.Second)
	fmt.Print("\nTotal Memory: ")
	time.Sleep(1 * time.Second)
	fmt.Print(memory.Total, "K  OK")
	time.Sleep(1 * time.Second)
	fmt.Print("\nInitializing USB Controllers ..")
	time.Sleep(3 * time.Second)
	fmt.Print(" Done.")
	time.Sleep(2 * time.Second)
	fmt.Print("\n\n\nAuto detecting AHCI PORT 0 ..")
	time.Sleep(3 * time.Second)
	fmt.Print(" ATAPI Hard Disk")
	time.Sleep(1 * time.Second)
	fmt.Println("\n\n\nPress F1 to Run SETUP\nPress F8 to select boot device")
	
}
func main() {
	boot()
	time.Sleep(30 * time.Second)
	CallClear()
}