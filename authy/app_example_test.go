package authy

import "fmt"

func ExampleClient_AppDetails() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	resp := client.AppDetails()
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Println("app:")
	fmt.Println("name:", resp.App.Name)
	fmt.Println("plan:", resp.App.Plan)
	fmt.Println("sms_enabled:", resp.App.SmsEnabled)
	fmt.Println("white_label:", resp.App.WhiteLabel)
	// Output:
	// success: true
	// message: Application information.
	// app:
	// name: Authy Docs
	// plan: starter
	// sms_enabled: false
	// white_label: false
}
