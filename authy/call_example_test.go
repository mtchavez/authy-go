package authy

import "fmt"

func ExampleClient_Call() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &CallReq{
		AuthyID: 2,
	}
	resp := client.Call(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	// Output:
	// success: true
	// message: Call started...
}

func ExampleClient_Call_NotFound() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &CallReq{
		AuthyID: 3,
	}
	resp := client.Call(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	// Output:
	// success: false
	// message: User not found.
}

func ExampleClient_Force() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	req := &CallReq{
		AuthyID: 1,
		Force:   true,
	}
	resp := client.Call(req)
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	// Output:
	// success: true
	// message: Call started...
}
