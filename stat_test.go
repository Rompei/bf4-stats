package stat

import (
	"testing"
)

func TestGetUserStats(t *testing.T) {
	res, err := GetUserStats("zousanvl", "974643959", "pc")
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n", res.Context.ProfileCommon.User.Presence.PresenceStates)
}
