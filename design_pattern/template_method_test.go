package design_pattern

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// https://refactoring.guru/design-patterns/template-method/go/example?_gl=1*kj5xyw*_ga*MjM5NDI2NTM4LjE2OTU3MTgwNTQ.*_ga_SR8Y3GYQYC*MTY5NTcxODA1My4xLjEuMTY5NTcxODA5Ni4xNy4wLjA.
// OTP means one-time-password

type IOneTimePwdDelivery interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	buildMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type OTPDeliveryTemplate struct {
	// NOTE 潜入接口类型是为了潜入一个实现了该接口的具体类型的对象。
	// 这个具体类型实现了此 Template 中未实现的行为
	IOneTimePwdDelivery

	msgBuilder func(msgTmpl string, args ...any) string
}

// genAndSendOTP is the entry point method
func (t *OTPDeliveryTemplate) genAndSendOTP(otpLength int) error {
	otp := t.genRandomOTP(otpLength)
	t.saveOTPCache(otp)
	message := t.buildMessage(otp)
	err := t.sendNotification(message)
	if err != nil {
		return err
	}
	t.publishMetric()
	return nil
}

// func (t *OTPDeliveryTemplate) genRandomOTP(len int) string {
//	randomOTP := "1234"
//	fmt.Printf("OTPDeliveryTemplate: generating random one-time-pwd %s\n", randomOTP)
//	return randomOTP
// }

// --- deliverBySMS start ---

type deliverBySMS struct {
	// NOTE 内嵌 Template 是为了：
	// 1. Template 已经实现了接口中的部分行为，内嵌它，当前类型就不需要实现接口中的这些行为了
	// 2. 当前类型可能会使用 Template 特有的数据成员、行为成员
	OTPDeliveryTemplate
}

func (s *deliverBySMS) genRandomOTP(len int) string {
	randomOTP := "123456"
	fmt.Printf("SMS: generating random one-time-pwd %s\n", randomOTP)
	return randomOTP
}

func (s *deliverBySMS) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving one-time-pwd: %s to cache\n", otp)
}

func (s *deliverBySMS) buildMessage(otp string) string {
	return s.msgBuilder("SMS: one-time-pwd for login is %s", otp)
}

func (s *deliverBySMS) sendNotification(message string) error {
	fmt.Printf("SMS: sending short message: %s\n", message)
	return nil
}

func (s *deliverBySMS) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

// --- deliverBySMS end ---

// --- email start ---

type email struct {
	OTPDeliveryTemplate
}

func (s *email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Email: generating random OTPDeliveryTemplate %s\n", randomOTP)
	return randomOTP
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving OTPDeliveryTemplate: %s to cache\n", otp)
}

func (s *email) buildMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}

func (s *email) publishMetric() {
	fmt.Printf("Email: publishing metrics\n")
}

// --- email end ---

// Tests start

func TestSMS1(t *testing.T) {
	convey.Convey(`Given an OTPDeliveryTemplate object
 and embed an deliverBySMS object into it
 and OTPDeliveryTemplate.msgBuilder is not set
 and deliverBySMS.OTPDeliveryTemplate is not set`, t, func() {
		sms := &deliverBySMS{}
		template := OTPDeliveryTemplate{}
		template.IOneTimePwdDelivery = sms

		convey.Convey("When call genAndSendOTP() on template", func() {
			template.genAndSendOTP(6)
			convey.Convey(
				`Then panic should happen:
 runtime error: invalid memory address or nil pointer dereference,
 because OTPDeliveryTemplate.msgBuilder is nil`)
		})
	})
}

func TestSMS2(t *testing.T) {
	convey.Convey(`Given an OTPDeliveryTemplate object
 and embed an deliverBySMS object into it
 and OTPDeliveryTemplate.msgBuilder is set
 and deliverBySMS.OTPDeliveryTemplate is set to the template object`, t, func() {
		sms := &deliverBySMS{}
		template := OTPDeliveryTemplate{}
		template.msgBuilder = fmt.Sprintf
		template.IOneTimePwdDelivery = sms
		sms.OTPDeliveryTemplate = template

		convey.Convey("When call genAndSendOTP() on template", func() {
			template.genAndSendOTP(6)
			convey.Convey(
				`Then it should be OK`, func() {
					convey.So(true, convey.ShouldBeTrue)
				})
		})
	})
}

func TestTemplate(t *testing.T) {
	// OTPDeliveryTemplate := OTPDeliveryTemplate{}

	// smsOTP := &deliverBySMS{
	//  OTPDeliveryTemplate: OTPDeliveryTemplate,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  OTPDeliveryTemplate: OTPDeliveryTemplate,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &deliverBySMS{}
	o := OTPDeliveryTemplate{
		IOneTimePwdDelivery: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &email{}
	o = OTPDeliveryTemplate{
		IOneTimePwdDelivery: emailOTP,
	}
	o.genAndSendOTP(4)

	var x interface{} = "foo"
	x = emailOTP
	// Invalid type switch guard: v := emailOTP.(type) (non-interface type *email on the left)
	// switch v := emailOTP.(type) {
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil") // here v has type interface{}
	case int:
		fmt.Println("x is", v) // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	case OTPDeliveryTemplate:
		fmt.Println("x is OTPDeliveryTemplate")
	case email:
		fmt.Println("x is email")
	case IOneTimePwdDelivery:
		fmt.Println("x is IOneTimePwdDelivery")
	default:
		fmt.Println("type unknown") // here v has type interface{}
	}
}
