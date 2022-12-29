package app

import (
	"health-checker/internal/services/notification"
	"log"
)

func StartNotification(n notification.Nitifier, msgq <-chan *notification.Message) {

	for {
		msg := <-msgq
		go func() {
			ne := n.Notify(*msg)
			if ne != nil {
				log.Fatal(ne.Error())
			}
		}()
	}
}