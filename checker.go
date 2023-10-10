package main

import (
	"fmt"
	"github.com/simonmartyr/toogoodtogogo"
	"log"
	"time"
)

func (t *TgtgNotifier) Check() error {
	log.Println("Check Starting")
	t.Refresh()
	var itemsToNotify []*toogoodtogo.Item
	for i, x := range t.Config.Items {
		if !x.Notify {
			continue
		}
		item, getErr := t.Client.GetItem(x.ItemId)
		if getErr != nil {
			log.Println(getErr)
			continue
		}
		if item.ItemsAvailable > 0 && ShouldNotify(x.LastNotified) {
			t.Config.Items[i].LastNotified = time.Now().Format("2006-01-02")
			itemsToNotify = append(itemsToNotify, item)
		}
	}
	log.Println("Check Complete")
	if len(itemsToNotify) > 0 {
		log.Println("Sending Notification")
		return t.SendNotification(itemsToNotify)
	}
	return nil
}

func ShouldNotify(toCompare string) bool {
	if toCompare == "" {
		return true
	}
	currentTime := time.Now().Format("2006-01-02")
	current, err1 := time.Parse("2006-01-02", currentTime)
	compare, err2 := time.Parse("2006-01-02", toCompare)
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing date strings:", err1, err2)
		return false
	}
	return current.Sub(compare) >= 24*time.Hour
}
