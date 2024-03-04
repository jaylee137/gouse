package check

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func SampleTypeCheckUUID() {
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

func SampleTypeCheckGmail() {
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

func SampleTypeCheckYahoo() {
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

func SampleTypeCheckOutlook() {
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

func SampleTypeCheckEdu() {
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

func SampleTypeCheckEmail() {
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

func SampleTypeCheckUsername() {
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

func SampleTypeCheckPassword() {
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

func SampleTypeCheckPhone() {
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
