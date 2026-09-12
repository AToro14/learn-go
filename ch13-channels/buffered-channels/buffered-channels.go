package main

func addEmailsToQueue(emails []string) chan string {
	emailBatch := make(chan string, len(emails))
	for _, email := range emails {
		emailBatch <- email
	}
	return emailBatch
}
