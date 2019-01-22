package main

import (
	"flag"
	"fmt"
	"os"
	cf "rf/commonFunc"
	"strconv"
)

var folder, ext, prefix string
var lenMask int

func init() {
	flag.StringVar(&folder, "path", "", "type path")
	flag.StringVar(&ext, "ext", "jpg", "file extention")
	flag.StringVar(&prefix, "prefix", "F", "prefix for new file name")
	flag.IntVar(&lenMask, "lenMask", 5, "count of digitl in new file name")
}

func main() {
	//folder := ""
	//ext := "jpg"
	//prefix := "a"
	//lenMask := 5
	//folder = "TMP"
	fmt.Println("Program rename files in directory for mask (for example) F00001.jpg, F00002.jpg and so on")
	fmt.Println("file sorted by date")
	fmt.Println("use -? for see otions")
	flag.Parse()
	fmt.Println("Folder:", folder)
	dir, _ := cf.FDirExt(folder, ext)
	fmt.Println("File count:", len(dir))
	dir = cf.FSortFileListByDate(dir)
	for i, f := range dir {
		newFn := createfn(prefix, ext, lenMask, i+1)
		fmt.Println(cf.DdateTiDDMMYYYhhmmss(f.ModTime), "|", f.Name, " -> ", newFn)
		os.Rename(folder+"\\"+f.Name, folder+"\\"+newFn)
	}

}
func createfn(prefix, ext string, lenMask, n int) string {
	rez := strconv.Itoa(n)
	for i := len(rez); i < lenMask; i++ {
		rez = "0" + rez
	}
	rez = prefix + rez + "." + ext
	return rez
}
