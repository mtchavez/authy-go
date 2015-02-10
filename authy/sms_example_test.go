package authy

import "fmt"

func ExampleClient_SMS() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &SMSReq{
		AuthyId: 1,
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
