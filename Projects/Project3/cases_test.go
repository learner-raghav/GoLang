package main

import "errors"

func getTestCases(dbObject DBObject) []testStruct{

	testCases = []testStruct{
		{
			desc: "Basic Read",
			route: "/getEmployeeById",
			method: "GET",
			dbObj: dbObject,
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
			dbObj: dbObject,
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
			dbObj: dbObject,
			employee: Employee{Name: "Raghav",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
			output: createOutput{
				err: nil,
			},
		},
		{
			desc: "Basic Update",
			route: "/update",
			method: "PUT",
			dbObj: dbObject,
			employee: Employee{EmployeeId: 17,Name: "Akash SAMS",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
			output: createOutput{
				err: nil,
			},
		},
		{
			desc: "Basic Update",
			route: "/update",
			method: "PUT",
			dbObj: dbObject,
			employee: Employee{EmployeeId: 200,Name: "Akashsasss",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
			output: createOutput{
				err: errors.New("No modifications done"),
			},
		},
		{
			desc: "Basic Delete",
			route: "/delete",
			method: "DELETE",
			dbObj: dbObject,
			employee: Employee{EmployeeId: 50},
			output: createOutput{
				err: nil,
			},
		},
		{
			desc: "Basic Delete",
			route: "/delete",
			method: "DELETE",
			dbObj: dbObject,
			employee: Employee{EmployeeId: 2900},
			output: createOutput{
				err: errors.New("The id does not exist"),
			},
		},
		{
			desc: "Basic Read with DB Nil",
			route: "/getEmployeeById",
			method: "GET",
			dbObj: DBObject{db: nil},
			employee: Employee{EmployeeId: 100},
			output: createOutput{
				employee: Employee{},
				err: errors.New("DB not configured properly"),
			},
		},
		{
			desc: "Basic Create with DB nil",
			route: "/create",
			method: "POST",
			dbObj: DBObject{db: nil},
			employee: Employee{Name: "Raghav",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
			output: createOutput{
				err: errors.New("DB not configured properly"),
			},
		},
		{
			desc: "Basic Update with DB nil",
			route: "/update",
			method: "PUT",
			dbObj: DBObject{db: nil},
			employee: Employee{EmployeeId: 200,Name: "Akashsasss",Email: "raghav@ZopSmart.com",Phone: "8384852943",RoleId: 4},
			output: createOutput{
				err: errors.New("DB not configured properly"),
			},
		},
		{
			desc: "Basic Delete with DB nil",
			route: "/delete",
			method: "DELETE",
			dbObj: DBObject{db: nil},
			employee: Employee{EmployeeId: 2900},
			output: createOutput{
				err: errors.New("DB not configured properly"),
			},
		},
	}
	return testCases
}

