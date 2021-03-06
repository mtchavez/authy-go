package authy

import "fmt"

func ExampleClient_SMS() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &SMSReq{
		AuthyID: 1,
	}
	resp := client.SMS(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("cellphone:", resp.Cellphone)
	// Output:
	// success: true
	// message: SMS token was sent
	// cellphone: +1-XXX-XXX-XX86
}

func ExampleClient_SMS_Force() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &SMSReq{
		AuthyID: 1,
		Force:   true,
	}
	resp := client.SMS(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("cellphone:", resp.Cellphone)
	// Output:
	// success: true
	// message: SMS token was sent
	// cellphone: +1-XXX-XXX-XX86
}

func ExampleClient_SMS_Shortcode() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &SMSReq{
		AuthyID:   1,
		Force:     true,
		Shortcode: true,
	}
	resp := client.SMS(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("cellphone:", resp.Cellphone)
	// Output:
	// success: true
	// message: SMS token was sent
	// cellphone: +1-XXX-XXX-XX86
}

func ExampleClient_SMS_CustomAction() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &SMSReq{
		AuthyID:      1,
		CustomAction: "change_preferences",
	}
	resp := client.SMS(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("cellphone:", resp.Cellphone)
	// Output:
	// success: true
	// message: SMS token was sent
	// cellphone: +1-XXX-XXX-XX86
}
