package patreon

import (
	"encoding/json"
	"errors"
)

type Includes struct {
	Items []interface{}
}

func (i *Includes) UnmarshalJSON(b []byte) error {
	var items []*json.RawMessage
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	count := len(items)
	i.Items = make([]interface{}, count)

	s := struct {
		Type string `json:"type"`
	}{}

	for idx, raw := range items {
		if err := json.Unmarshal(*raw, &s); err != nil {
			return err
		}

		var obj interface{}

		// Depending on the type, we can run json.Unmarshal again on the same byte slice
		// But this time, we'll pass in the appropriate struct instead of a map
		if s.Type == "user" {
			obj = &User{}
		} else if s.Type == "reward" {
			obj = &Reward{}
		} else if s.Type == "goal" {
			obj = &Goal{}
		} else {
			return errors.New("unsupported type found")
		}

		if err := json.Unmarshal(*raw, obj); err != nil {
			return err
		}

		i.Items[idx] = obj
	}

	return nil
}
