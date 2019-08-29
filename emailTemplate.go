package main

var (
	UnformattedTextBody string
)

func init() {
	UnformattedTextBody = "%s,\n\n" +
		"Thank you for booking %s. Below are the details of your appointment.\n\n" +
		"Appointment Details:\n\n" +
		"Service: %s\n" +
		"Date: %s\n" +
		"Time: %s\n" +
		"With: %s\n" +
		"Cost: $0\n\n" +
		"Best Regards!\n" +
		"Yeshiva University\n" +
		"212-960-5200\n" +
		"https://yushuttles.com/\n\n" +
		"Please, use the following link in case you'd like to cancel your appointment.\n" +
		"https://yushuttles.com/?bupcancelappointment=n8jvsj1b3t47o8ph2f3csprm7b_1567027197&bupid=288"
}
