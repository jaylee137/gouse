package dir

import "github.com/thuongtruong1009/gouse/io"

func ioHierarchyDir() {
	data, err := io.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
