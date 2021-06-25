package main

import (
	"Dhash/Imagehash"
	_ "github.com/mattn/go-sqlite3"
)


func main()  {
	Imagehash.Save("sample1")
	Imagehash.Save("sample2")
	Imagehash.Save("sample3")
	Imagehash.Save("sample4")
	//Imagehash.Delete()
}


