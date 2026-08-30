package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	usersMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return usersMap, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		usersMap[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
		// for i, name := range names {
		// 	usersMap[name] = user{name: name, phoneNumber: phoneNumbers[i]}
	}
	return usersMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
