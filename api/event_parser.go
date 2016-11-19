package api

import (
	"dd-server/types"
	"encoding/json"
	"fmt"
)

// Parse event
//"events": {
//"System": [
//{
//"timestamp": 1476195870.719676,
//"host": "bogon",
//"api_key": "8ea42d1939e248c59333070d5fd2f4c3",
//"msg_text": "Version 5.8.5",
//"event_type": "Agent Startup"
//}
//],
//}
func ParseEvents(req *types.RequestBody) []*types.Event {

	events := make([]*types.Event, 0)

	for k, v := range req.Events {
		fmt.Println("Key: ", k)
		var subEvents []*types.Event

		tmp, err := json.Marshal(v)
		if err != nil {
			fmt.Println("Marshal events error", v, ", error:", err)
			continue
		}

		err = json.Unmarshal(tmp, &subEvents)
		if err != nil {
			fmt.Println("Unmarshal events string error", tmp, ", error:", err)
		}

		events = append(events, subEvents...)

	}

	return events
}
