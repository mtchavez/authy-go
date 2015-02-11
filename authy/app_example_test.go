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

func ExampleClient_AppStats() {
	client := NewSandboxClient("d57d919d11e6b221c9bf6f7c882028f9")
	resp := client.AppStats()
	fmt.Println("success:", resp.Success)
	fmt.Println("message:", resp.Message)
	fmt.Printf("%+v stats returned", len(resp.Stats))
	// Output:
	// success: true
	// message: Monthly statistics.
	// 12 stats returned
}
