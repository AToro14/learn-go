package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	usersMap := make(map[string]user)
	users := []user{}
	if len(names) != len(phoneNumbers) {
		return usersMap, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		users = append(users, user{name: names[i], phoneNumber: phoneNumbers[i]})
		usersMap[users[i].name] = users[i]
	}
	return usersMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
