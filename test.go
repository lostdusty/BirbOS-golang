package main

import (
	"fmt"
	//"strings"
	. "github.com/klauspost/cpuid"
	"github.com/mackerelio/go-osstat/memory"
	"os"
	"os/exec"
	"runtime"
	"time"
	//"log"
	//"bufio"
	"io/ioutil"
	"github.com/sqweek/dialog"
)

var clear map[string]func() //create a map for storing clear funcs

func check (e error) {
	if e != nil {
	panic(e)
	}
}
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
func post() {
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

func bootloader() {
	fmt.Println("======================== BirbOS Loader ========================")
	fmt.Print("\n\nBooting BirbOS")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(3 * time.Second)
	if _, err := os.Stat("birbtool.CMD"); os.IsNotExist(err) {
		fmt.Println("\nbirbtool.CMD is missing, system cannot boot.\nRestarting in 10 seconds....")
		time.Sleep(10 * time.Second)
		CallClear()
		main()
	} else {
		boot()
	}
}
func boot() {
	fmt.Println("\nDetecting PCI-E devices.....")
	CallClear()
	if _, err := os.Stat("fsetup.e"); os.IsNotExist(err) {
		fmt.Println("Completing first time setup...")
		writefirstsetup()
		time.Sleep(1 * time.Second)
		usrtool()
	} else {
		bos()
	}
}

func bos() {
fmt.Println("Skiping first time setup")
}

func usrtool() {
fmt.Println("Press any key to continue to user creation.")
fmt.Print("\nType your username: ")
var usrname string
fmt.Scan(&usrname)
pwd := dialog.Message("Do you want your account to have a password?").Title("Protect with password?").YesNo()
if pwd {
var passwd string
fmt.Print("\nType your password: ")
fmt.Scan(&passwd)
fmt.Println("\nUsername:", usrname, "password:", passwd)
} else {
fmt.Println("\nLoading your personal settings....")
}
}

func main() {
	post()
	time.Sleep(3 * time.Second)
	CallClear()
	bootloader()
}

func writefirstsetup() {
d1 := []byte("first-setup-complete")
    err := ioutil.WriteFile("fsetup.e", d1, 0644)
    check(err)
}