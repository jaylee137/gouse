package cast

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func SampleTypeStructConvert() {
	type CompanyInfo struct {
		Company string
		Address string
		Contact string
	}

	companyInfo := CompanyInfo{
		Company: "GeeksforGeeks",
		Address: "Noida",
		Contact: "+91 9876543210",
	}

	println("Struct to string: ", types.StructToString(companyInfo))
	fmt.Println("Struct to map: ", types.StructToMap(companyInfo))
}
