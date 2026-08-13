package main

// ?
func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if len(message) > u.MessageCharLimit {
		return "", false
	} else {
		return message, true
	}
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
