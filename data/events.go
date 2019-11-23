package data

type event struct {
	Id    string
	Title string
}

// This commented function doesn't work.
// 		But another version of the GetAllEvents() function works fine.

// func GetAllEvents() []event {

// 	var allEvents [2]event

// 	eventOne := event{Id: "1", Title: "Event One"}
// 	eventTwo := event{Id: "2", Title: "Event Two"}

// 	allEvents[0] = eventOne
// 	allEvents[1] = eventTwo

// 	return allEvents
// }

func GetAllEvents() []event {

	allEvents := []event{
		{Id: "1", Title: "Event One"},
		{Id: "2", Title: "Event Two"},
	}

	return allEvents
}
