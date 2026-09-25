package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	usERR := em.recipient.updateStatus(em.status)
	if usERR != nil {
		return fmt.Errorf("error updating user status: %w", usERR)
	}
	tERR := a.track(em.status)
	if tERR != nil {
		return fmt.Errorf("error tracking user bounce: %w", tERR)
	}
	return nil
}
