package cron_test

import (
	"testing"

	keepertest "github.com/neutron-org/neutron/testutil/cron/keeper"
	"github.com/neutron-org/neutron/testutil/cron/nullify"

	"github.com/neutron-org/neutron/x/cron"
	"github.com/neutron-org/neutron/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ScheduleList: []types.Schedule{
		{
			Index: "0",
},
		{
			Index: "1",
},
	},
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CronKeeper(t)
	cron.InitGenesis(ctx, *k, genesisState)
	got := cron.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ScheduleList, got.ScheduleList)
// this line is used by starport scaffolding # genesis/test/assert
}