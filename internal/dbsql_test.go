package internal

import (
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	ctx, w := workspaceTest(t)

	srcs, err := w.DataSources.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := w.Queries.Create(ctx, sql.QueryPostContent{
		Name:         RandomName("go-sdk/test/"),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SHOW TABLES",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Queries.DeleteByQueryId(ctx, query.Id)
		require.NoError(t, err)
	})

	byId, err := w.Queries.GetByQueryId(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.Query, byId.Query)

	byName, err := w.Queries.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	updated, err := w.Queries.Update(ctx, sql.QueryPostContent{
		QueryId:      query.Id,
		Name:         RandomName("go-sdk-updated"),
		DataSourceId: srcs[0].Id,
		Description:  "UPDATED: test query from Go SDK",
		Query:        "SELECT 2+2",
	})
	require.NoError(t, err)
	assert.NotEqual(t, updated.Query, byId.Query)
}

func TestAccAlerts(t *testing.T) {
	ctx, w := workspaceTest(t)

	srcs, err := w.DataSources.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := w.Queries.Create(ctx, sql.QueryPostContent{
		Name:         RandomName("go-sdk/test/"),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SHOW TABLES",
	})
	require.NoError(t, err)
	defer w.Queries.DeleteByQueryId(ctx, query.Id)

	alert, err := w.Alerts.Create(ctx, sql.EditAlert{
		Name:    RandomName("go-sdk-"),
		QueryId: query.Id,
	})
	require.NoError(t, err)
	defer w.Alerts.DeleteByAlertId(ctx, alert.Id)

	err = w.Alerts.Update(ctx, sql.EditAlert{
		AlertId: alert.Id,
		Name:    RandomName("go-sdk-updated-"),
		QueryId: query.Id,
	})
	require.NoError(t, err)

	byId, err := w.Alerts.GetByAlertId(ctx, alert.Id)
	require.NoError(t, err)

	byName, err := w.Alerts.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Alerts.List(ctx)
	require.NoError(t, err)

	names, err := w.Alerts.AlertNameToIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, alert.Id, names[byId.Name])

	schedule, err := w.Alerts.CreateSchedule(ctx, sql.CreateRefreshSchedule{
		AlertId:      alert.Id,
		Cron:         "5 4 * * *",
		DataSourceId: srcs[0].Id,
	})
	require.NoError(t, err)
	defer w.Alerts.DeleteScheduleByAlertIdAndScheduleId(ctx, alert.Id, schedule.Id)

	schedules, err := w.Alerts.ListSchedulesByAlertId(ctx, alert.Id)
	require.NoError(t, err)
	assert.True(t, len(schedules) >= 1)

	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)

	userId, err := strconv.ParseInt(me.Id, 10, 64)
	require.NoError(t, err)

	sub, err := w.Alerts.Subscribe(ctx, sql.CreateSubscription{
		AlertId: alert.Id,
		UserId:  userId,
	})
	require.NoError(t, err)

	allSubs, err := w.Alerts.GetSubscriptionsByAlertId(ctx, alert.Id)
	require.NoError(t, err)
	assert.True(t, len(allSubs) >= 1)

	err = w.Alerts.UnsubscribeByAlertIdAndSubscriptionId(ctx, alert.Id, sub.Id)
	require.NoError(t, err)
}

func TestAccDashboards(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.Dashboards.Create(ctx, sql.CreateDashboardRequest{
		Name:                    RandomName("go-sdk-"),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	require.NoError(t, err)

	defer w.Dashboards.DeleteByDashboardId(ctx, created.Id)

	byId, err := w.Dashboards.GetByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	byName, err := w.Dashboards.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Dashboards.ListAll(ctx, sql.ListDashboardsRequest{})
	require.NoError(t, err)

	names, err := w.Dashboards.DashboardNameToIdMap(ctx, sql.ListDashboardsRequest{})
	require.NoError(t, err)
	assert.Equal(t, created.Id, names[byId.Name])
	assert.Equal(t, len(all), len(names))

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	err = w.Dashboards.Restore(ctx, sql.RestoreDashboardRequest{
		DashboardId: created.Id,
	})
	require.NoError(t, err)
}
