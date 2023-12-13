package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func typeStructConvert() {
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
	println("Convert string to int: ", types.StringToInt("1"))
	fmt.Println("Convert string to float: ", types.StringToFloat("1.1"))
	println("Convert string to bool: ", types.StringToBool("true"))
	fmt.Println("Convert string to bytes: ", string(types.StringToBytes("1")), "->", types.StringToBytes("1"))
	fmt.Println("Convert strings to bytes: ", string(types.StringsToBytes([]string{"1", "2", "3"})), "->", types.StringsToBytes([]string{"1", "2", "3"}))
}

func typeCastToString() {
	println("Cast int to string: ", types.IntToString(1))
	fmt.Println("Cast float to string: ", types.FloatToString(1.1))
	println("Cast bool to string: ", types.BoolToString(true))
	println("Cast interface to string: ", types.InterfaceToString([]int{1, 2, 3}))
	println("Cast bytes to string: ", types.BytesToString([]byte{49, 50, 51}))
	println("Cast rune to string: ", types.RuneToString('a'))
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

var emails = []string{
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

func typeCheckGmail() {
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
}

func typeCheckYahoo() {
	println("--- Check valid yahoo ---")
	for _, email := range emails {
		isEmail, err := types.IsYahoo(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid yahoo email\n", email)
		}
	}
}

func typeCheckOutlook() {
	println("--- Check valid outlook ---")
	for _, email := range emails {
		isEmail, err := types.IsOutlook(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid outlook email\n", email)
		}
	}
}

func typeCheckEdu() {
	println("--- Check valid education mail ---")
	for _, email := range emails {
		isEmail, err := types.IsEdu(email)
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid education email\n", email)
		}
	}
}

func typeCheckEmail() {
	println("--- Check valid custom domain ---")
	for _, email := range emails {
		isEmail, err := types.IsEmail(email, "edu.vn")
		if err != nil {
			fmt.Printf("%s - %s\n", email, err)
		}

		if isEmail {
			fmt.Printf("%s - valid custom domain email\n", email)
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

func typeCheckPassword() {
	passwords := []string{
		"okahehe",
		"missingat.sign",
		"#$%&*()",
		"Username01$",
	}

	println("--- Check valid password ---")
	for _, password := range passwords {
		isPassword, err := types.IsPassword(password)
		if err != nil {
			fmt.Printf("%s - %s\n", password, err)
		}

		if isPassword {
			fmt.Printf("%s - valid password\n", password)
		}
	}
}

func typeCheckPhone() {
	//  Note: Phone format syntax: +<country_calling_code> (<area_Prefix_mobile_code>) <phone_number>
	// Reference at https://en.wikipedia.org/wiki/List_of_mobile_telephone_prefixes_by_country#:~:text=Property%20Value%20%20Country%20or%20unrecognized%20territory%20,73%20%20%20Etisalat%20%20%20www.etisalat.af%20
	phoneNumbers := []string{
		"+1 (123) 456-7890",
		"+44 (20) 1234-5678",
		"+81 (3) 1234-5678",
		"InvalidPhoneNumber",
		"+1 (123) 45-67890",
		"+84 (3) 668-22796",
	}

	println("--- Check valid phone number ---")
	for _, phoneNumber := range phoneNumbers {
		isPhoneNumber, err := types.IsPhone(phoneNumber)
		if err != nil {
			fmt.Printf("%s - %s\n", phoneNumber, err)
		}

		if isPhoneNumber {
			fmt.Printf("%s - valid phone number\n", phoneNumber)
		}
	}
}

func TypeExample() {
	// typeStructConvert()
	// typeStringConvert()
	// typeCastToString()

	// typeCheck()
	// typeCheckUUID()
	// typeCheckGmail()
	// typeCheckYahoo()
	// typeCheckOutlook()
	// typeCheckEdu()
	// typeCheckEmail()
	// typeCheckUsername()
	// typeCheckPassword()
	// typeCheckPhone()
}
