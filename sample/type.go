package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func typeStructToString() {
	println("Struct to string start:")
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

func typeStringConvert() {
	println("String to int: ", types.StringToInt("1"))
	fmt.Println("String to float: ", types.StringToFloat("1.1"))
	println("String to bool: ", types.StringToBool("true"))
}

func typeCastToString() {
	println("Cast int to string: ", types.IntToString(1))
	fmt.Println("Cast float to string: ", types.FloatToString(1.1))
	println("Cast bool to string: ", types.BoolToString(true))
	println("Cast interface to string: ", types.InterfaceToString([]int{1, 2, 3}))
}

func typeCheck() {
	println("Check type is int: ", types.IsInt(1))
	println("Check type is uint: ", types.IsUnInt(-1))
	println("Check type is float: ", types.IsFloat(1.1))
	println("Check type is complex: ", types.IsComplex(1.1))
	println("Check type is number: ", types.IsNumber(1.23))
	println("Check type is string: ", types.IsString("1"))
	println("Check type is rune: ", types.IsRune('a'))
	println("Check type is byte: ", types.IsByte(byte('a')))
	println("Check type is uintptr: ", types.IsUintptr(uintptr(1)))
	println("Check type is bool: ", types.IsBool(true))
	println("Check type is slice: ", types.IsSlice([]int{1, 2, 3}))
	println("Check type is map: ", types.IsMap(map[string]int{"a": 1}))
	println("Check type is struct: ", types.IsStruct(struct{}{}))
	println("Check type is error: ", types.IsError(new(error)))
	println("Check type is interface: ", types.IsInterface(new(interface{})))
	println("Check type is channel: ", types.IsChannel(make(chan int)))
	println("Check type is array: ", types.IsArray([3]int{1, 2, 3}))
	println("Check type is unsafe pointer: ", types.IsUnsafePointer(new(*int)))
	println("Check type is pointer: ", types.IsPointer(new(int)))
	println("Check type is func: ", types.IsFunc(func() {}))
	println("Check type is nil: ", types.IsNil(nil))
	println("Check type is empty: ", types.IsEmpty(""))
	println("Check type is empty: ", types.IsEmpty(0))
	println("Check type is zero: ", types.IsZero(0))
}

func typeCheckUUID() {
	isValid, err := types.IsUUID("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		println(err.Error())
	}
	println("Check is valid uuid: ", isValid)
}

func typeCheckEmail() {
	emails := []string{
		"okahehe@gmail.com",
		"test@example.com",
		"short@ex.com",
		"toolong@exampledomain.com",
		"missingat.sign.com",
		"invalid@.domain.com",
		"user2@yahoo.com",
		"bean@outlook.com",
		"ititiu19228@student.hcmiu.edu.vn",
	}

	println("--- Check valid gmail ---")
	for _, email := range emails {
		isEmail, err := types.IsGmail(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid gmail\n", email)
		}
	}

	println("\n--- Check valid yahoo ---")
	for _, email := range emails {
		isEmail, err := types.IsYahoo(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid yahoo\n", email)
		}
	}

	println("\n--- Check valid outlook ---")
	for _, email := range emails {
		isEmail, err := types.IsOutlook(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid outlook\n", email)
		}
	}

	println("\n--- Check valid education mail ---")
	for _, email := range emails {
		isEmail, err := types.IsEdu(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid edu\n", email)
		}
	}

	println("\n--- Check valid custom domain ---")
	for _, email := range emails {
		isEmail, err := types.IsEmail(email, "edu.vn")
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid custom domain\n", email)
		}
	}
}

func typeCheckUsername() {
	usernames := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
	}

	println("--- Check valid username ---")
	for _, username := range usernames {
		isUsername, err := types.IsUsername(username)
		if err != nil {
			fmt.Printf("%s - %s\n", username, err)
		}

		if isUsername {
			fmt.Printf("%s - valid username\n", username)
		}
	}
}

// func typeCheckPassword() {
// 	passwords := []string{
// 		"okahehe",
// 		"missingat.sign",
// 		"#$%&*()",
// 	}

// 	println("--- Check valid password ---")
// 	for _, password := range passwords {
// 		isPassword, err := types.IsPassword(password)
// 		if err != nil {
// 			fmt.Printf("%s - %s\n", password, err)
// 		}

// 		if isPassword {
// 			fmt.Printf("%s - valid password\n", password)
// 		}
// 	}
// }

func TypeExample() {
	typeStructToString()
	typeStringConvert()
	typeCastToString()

	typeCheck()
	typeCheckUUID()
	typeCheckEmail()
	typeCheckUsername()
	// typeCheckPassword()
}
