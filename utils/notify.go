package utils

import "gopkg.in/toast.v1"

func Notify(title, message string, action toast.Action) {
	notification := toast.Notification{
		AppID:    "EveryNASA",
		Title:    title,
		Message:  message,
		Duration: toast.Long,
		Actions: []toast.Action{
			action,
		},
	}

	err := notification.Push()
	if err != nil {
		panic(err)
	}
}
