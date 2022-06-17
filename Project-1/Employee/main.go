package main

func main() {
	jim := employee{
		person: person{
			firstName: "Jim",
			lastName:  "McMartin",
		},
		contactInfo: contactInfo{
			email:   "jim.mcmartin@gmail.com",
			zipCode: 123456,
		},
		startDate:  "20June2022",
		employeeId: 1,
	}

	ashly := employee{
		person: person{
			firstName: "Ashly",
			lastName:  "Walter",
		},
		contactInfo: contactInfo{
			email:   "ashly.walter@gmail.com",
			zipCode: 653234,
		},
		startDate:  "21June2022",
		employeeId: 2,
	}
	jim.printEmployee()
	jim.updateFistName("Jimmy")
	jim.printEmployee()

	ashly.printEmployee()

}
