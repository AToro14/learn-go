package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var msgCharLimit = 100
	if membershipType == "premium" {
		msgCharLimit = 1000
	}
	nUser := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: msgCharLimit,
		},
	}
	return nUser
}
