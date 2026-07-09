package main

import "fmt"

type Role struct {
	RoleName string
	HasRole  bool
}

type ClientWithRoles struct {
	ClientId string
	Roles    []Role
}

func main() {
	roles := []Role{}
	templateRole := Role{
		RoleName: "a",
		HasRole:  false,
	}

	names := []string{
		"a", "b", "c", "d",
	}

	for i := range names {
		r := templateRole
		r.RoleName = names[i]
		r.HasRole = true
		roles = append(roles, r)

		fmt.Println(templateRole)
	}

	fmt.Println(roles)
	fmt.Println(templateRole)
}
