package authy

import "fmt"

func ExampleClient_UserStatu() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &UserStatusReq{
		AuthyId: 2,
	}
	resp := client.UserStatus(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("status:")
	fmt.Println("authy_id:", resp.Status.AuthyId)
	fmt.Println("confirmed:", resp.Status.Confirmed)
	fmt.Println("registered:", resp.Status.Registered)
	fmt.Println("country_code:", resp.Status.CountryCode)
	fmt.Println("phone_number:", resp.Status.PhoneNumber)
	fmt.Println("devices:", resp.Status.Devices)
	// Output:
	// success: true
	// message: User status.
	// status:
	// authy_id: 2
	// confirmed: true
	// registered: false
	// country_code: 1
	// phone_number: XXX-XXX-9302
	// devices: []
}
