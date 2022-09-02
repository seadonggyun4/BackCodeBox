package structPackage // 패키지 명은 파일명과 동일하게 해도 가능

import "fmt"

type StructPackage struct {
	Name    string 
	Version string
}

func ShowPackageInfo(sp *StructPackage, name string, ver string) {
	//sp := StructPackage{Name: "goBolock", Version: "v1.0.0"}
	//sp = new(StructPackage)
	d := name
	sp.Version = ver
	fmt.Printf("Package Name : %s, Version : %s \n", d, sp.Version)
}
