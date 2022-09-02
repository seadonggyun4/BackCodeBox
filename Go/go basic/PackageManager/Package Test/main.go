package main

import "kerygma.com/structPackage"

func main() {
	original_sp := structPackage.StructPackage{}
	structPackage.ShowPackageInfo(&original_sp, "stopBlock", "v1.0.1")
}
