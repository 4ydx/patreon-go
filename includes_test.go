package patreon

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseIncludes(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(includesJson), &includes)
	require.NoError(t, err)
	require.Len(t, includes.Items, 5)

	user, ok := includes.Items[0].(*User)
	require.True(t, ok)
	require.Equal(t, "2822191", user.Id)
	require.Equal(t, "user", user.Type)
	require.Equal(t, "podsync", user.Attributes.Vanity)

	reward, ok := includes.Items[1].(*Reward)
	require.True(t, ok)
	require.Equal(t, "12312312", reward.Id)
	require.Equal(t, "reward", reward.Type)
	require.Equal(t, 100, reward.Attributes.Amount)

	goal, ok := includes.Items[2].(*Goal)
	require.True(t, ok)
	require.Equal(t, "2131231", goal.Id)
	require.Equal(t, "goal", goal.Type)
	require.Equal(t, 1000, goal.Attributes.Amount)

	campaign, ok := includes.Items[3].(*Campaign)
	require.True(t, ok)
	require.Equal(t, "12312321", campaign.Id)
	require.Equal(t, "campaign", campaign.Type)

	pledge, ok := includes.Items[4].(*Pledge)
	require.True(t, ok)
	require.Equal(t, 100, pledge.Attributes.AmountCents)
	require.Equal(t, time.Date(2017, 6, 20, 23, 21, 34, 514822000, time.UTC).Unix(), pledge.Attributes.CreatedAt.Unix())
	require.True(t, pledge.Attributes.PatronPaysFees)
	require.Equal(t, 100, pledge.Attributes.PledgeCapCents)
}

func TestParseUnsupportedInclude(t *testing.T) {
	includes := Includes{}
	err := json.Unmarshal([]byte(unknownIncludeJson), &includes)
	require.Error(t, err)
	require.Equal(t, "unsupported type 'unknown'", err.Error())
}

const includesJson = `
[
	{
		"attributes": {
			"vanity": "podsync"
		},
		"id": "2822191",
		"relationships": {},
		"type": "user"
	},
	{
		"attributes": {
			"amount": 100
		},
		"id": "12312312",
		"relationships": {},
		"type": "reward"
	},
	{
		"attributes": {
			"amount": 1000
		},
		"id": "2131231",
		"type": "goal"
	},
	{
		"attributes": {},
		"id": "12312321",
		"type": "campaign"
	},
	{
		"attributes": {
			"amount_cents": 100,
			"created_at": "2017-06-20T23:21:34.514822+00:00",
			"declined_since": null,
			"patron_pays_fees": true,
			"pledge_cap_cents": 100
		},
		"id": "2321312",
		"relationships": {
			"card": {
				"data": {
					"id": "bt_1231232132",
					"type": "card"
				},
				"links": {
					"related": "https://www.patreon.com/api/cards/bt_123123213"
				}
			},
			"creator": {
				"data": {
					"id": "12312321321312",
					"type": "user"
				},
				"links": {
					"related": "https://www.patreon.com/api/user/12312321321312"
				}
			},
			"patron": {
				"data": {
					"id": "213213213",
					"type": "user"
				},
				"links": {
					"related": "https://www.patreon.com/api/user/213213213"
				}
			},
			"reward": {
				"data": {
					"id": "12312321321",
					"type": "reward"
				},
				"links": {
					"related": "https://www.patreon.com/api/rewards/12312321321"
				}
			}
		},
		"type": "pledge"
	}
]
`

const unknownIncludeJson = `
[
	{
		"attributes": {
			"vanity": "podsync"
		},
		"id": "2822191",
		"relationships": {},
		"type": "user"
	},
	{
		"attributes": {},
		"id": "12312312",
		"relationships": {},
		"type": "unknown"
	}
]
`
