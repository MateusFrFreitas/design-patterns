package main

import "fmt"

func main() {
	var o *Otp

	smsOTP := &Sms{}
	o = &Otp{
		smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = &Otp{
		emailOTP,
	}
	o.genAndSendOTP(4)

}
