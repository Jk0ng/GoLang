package main

import "fmt"

func main() {
	var a  = "initial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int 
	fmt.Println(e) 

	f := "apple"
	fmt.Println(f)
 
	congrats := "happy birthday"
	fmt.Println(congrats)

	penniesPerText :=2.0 
	fmt.Printf("The type of penniesPerTest is %T\n", penniesPerText)

// 	Within the main function, declare a float called averageOpenRate and string called displayMessage on the same line.

// Initialize them to values of .23 and is the average open rate of your messages respectively before they are printed.
	averageOpenRate, displayMessage := .23, "is the average rate of your messages"
	fmt.Println(averageOpenRate)
	fmt.Println(displayMessage)

// 	Our Textio customers want to know how long they have had accounts with us.

// Follow the instructions in the comment provided. You will create a new variable called accountAgeInt that will be the truncated, integer version of accountAge.

	accountAge := 2.6

	//create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer

	accountAgeInt := int(accountAge)
	fmt.Println("your account has existed for", accountAgeInt, "years!")



// 	Something weird is happening in this code.

// What should be happening is that we create 2 separate constants: premiumPlanName and basicPlanName. Right now it looks like we're trying to overwrite one of them.

	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	//assignment: 
	// Keeping track of time in a message-sending application is critical. Imagine getting an appointment reminder an hour after your doctor's visit.
	// Complete the code using a computed constant to print the number of seconds in an hour.

	const secondsInMinute = 60 
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour 


	fmt.Println("number of seconds in an hour:", secondsInHour)

	fmt.Printf("I am %d years old",10)


	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("hi %s, your open rate is %.2f percent ", name, openRate )
	// don't edit below this line

	fmt.Println(msg)

	messageLen := 29
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("message sent")
	} else {
		fmt.Println("message not sent")
	}

	var smsSendingList int
	var costPerSMS float64
	var hasPermission bool
	var userName string

	fmt.Printf("%v %f %v %q\n", smsSendingList, costPerSMS, hasPermission, userName)

	var numCar int = 10
	var numCar1 float64 = 10.0
	fmt.Println(numCar)
	fmt.Println(numCar1)

	fmt.Printf("I am %d years old", numCar1)

}