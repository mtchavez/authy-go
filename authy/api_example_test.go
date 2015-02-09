package authy

import "fmt"

func ExampleClient_NewUser() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &NewUserReq{
		User: &AuthyUser{
			Email:       "user@domain.com",
			Cellphone:   "839-338-9302",
			CountryCode: 1,
		},
	}
	resp := client.NewUser(req)
	fmt.Println("Message:", resp.Message)
	fmt.Println("User created?", resp.Success)
	fmt.Println("User id:", resp.User.Id)
	// Output:
	// Message: User created successfully.
	// User created? true
	// User id: 3
}

func ExampleClient_NewUser_WithErrors() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &NewUserReq{
		User: &AuthyUser{},
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
