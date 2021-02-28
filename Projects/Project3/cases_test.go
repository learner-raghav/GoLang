package main

import "errors"

type createOutput struct {
	employee Employee
	err error
}

var testCases = []struct{
	desc string
	route string
	method string
	employee Employee
	output createOutput
}{
	{
		desc: "Basic Read",
		route: "/getEmployeeById",
		method: "GET",
		employee: Employee{EmployeeId: 5},
		output: createOutput{
			employee: Employee{
				EmployeeId: 5,
				Name: "D",
				Email: "D@gmail.com",
				Phone: "8384852943",
				RoleId: 5,
			},
			err: nil,
		},
	},
	{
		desc: "Basic Read",
		route: "/getEmployeeById",
		method: "GET",
		employee: Employee{EmployeeId: 100},
		output: createOutput{
			employee: Employee{},
			err: errors.New("Employee record does not exist"),
		},
	},
	{
		desc: "Basic Create",
		route: "/create",
		method: "POST",
		employee: Employee{Name: "Raghav",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
		output: createOutput{
			err: nil,
		},
	},
	{
		desc: "Basic Update",
		route: "/update",
		method: "PUT",
		employee: Employee{EmployeeId: 17,Name: "Akashsssssss",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
		output: createOutput{
			err: nil,
		},
	},
	{
		desc: "Basic Update",
		route: "/update",
		method: "PUT",
		employee: Employee{EmployeeId: 200,Name: "Akashsasss",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
		output: createOutput{
			err: errors.New("No modifications done"),
		},
	},
	{
		desc: "Basic Delete",
		route: "/delete",
		method: "DELETE",
		employee: Employee{EmployeeId: 31},
		output: createOutput{
			err: nil,
		},
	},
	{
		desc: "Basic Delete",
		route: "/delete",
		method: "DELETE",
		employee: Employee{EmployeeId: 2900},
		output: createOutput{
			err: errors.New("The id does not exist"),
		},
	},
}
