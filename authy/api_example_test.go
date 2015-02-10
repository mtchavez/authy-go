package authy

import "fmt"

func ExampleClient_NewUser() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &NewUserReq{
		User: &User{
			Email:       "user@domain.com",
			Cellphone:   "839-338-9302",
			CountryCode: 1,
		},
	}
	resp := client.NewUser(req)
	fmt.Println("Message:", resp.Message)
	fmt.Println("User created?", resp.Success)
	fmt.Println("User id:", resp.User.ID)
	// Output:
	// Message: User created successfully.
	// User created? true
	// User id: 3
}

func ExampleClient_NewUser_WithErrors() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &NewUserReq{
		User: &User{},
	}
	resp := client.NewUser(req)
	fmt.Println("Message:", resp.Message)
	fmt.Println("User created?", resp.Success)
	fmt.Println("User:", resp.User)
	fmt.Printf("%+v Errors:\n", len(resp.Errors)-1)
	fmt.Println("cellphone", resp.Errors["cellphone"])
	fmt.Println("country_code", resp.Errors["country_code"])
	fmt.Println("email", resp.Errors["email"])
	fmt.Println("message", resp.Errors["message"])
	// Output:
	// Message: User was not valid.
	// User created? false
	// User: <nil>
	// 3 Errors:
	// cellphone is invalid
	// country_code is invalid
	// email is invalid and can't be blank
	// message User was not valid.
}

func ExampleClient_Verify() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &VerifyReq{
		Token:   "0000000",
		AuthyID: 1,
	}
	resp := client.Verify(req)
	fmt.Println("token:", resp.Token)
	// API is currently returning a string for success here for some reason
	// fmt.Println("success:", resp.Success)
	// success: true

	// Output:
	// token: is valid
}

func ExampleClient_Verify_WithErrors() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &VerifyReq{
		Token:   "1234567",
		AuthyID: 1,
	}
	resp := client.Verify(req)
	fmt.Printf("token:%+v\n", resp.Token)
	fmt.Println("success:", resp.Success)
	fmt.Println("error:", resp.Errors["message"])
	// Output:
	// token:
	// success: false
	// error: User doesn't exist.
}

func ExampleClient_Verify_Force() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &VerifyReq{
		Token:   "939393",
		AuthyID: 3,
		Force:   true,
	}
	resp := client.Verify(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("error:", resp.Errors["message"])
	// Output:
	// success: false
	// error: User doesn't exist.
}

func ExampleClient_Verify_CustomAction() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &VerifyReq{
		Token:        "0000000",
		AuthyID:      1,
		Force:        true,
		CustomAction: "change_preferences",
	}
	resp := client.Verify(req)
	// API is currently returning a string for success here for some reason
	// fmt.Println("success:", resp.Success)
	// success: true
	fmt.Println("message:", resp.Message)
	fmt.Println("token:", resp.Token)
	// Output:
	// message: Token is valid.
	// token: is valid
}
