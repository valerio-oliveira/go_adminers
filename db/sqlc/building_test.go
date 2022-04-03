package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomBuilding(t *testing.T) RegistrationBuilding {
	arg := CreateBuildingParams{
		IDCondo:        createRandomCondo(t).ID,
		IDBuildingType: createRandomBuildingType(t).ID,
		Description:    util.RandomString(10),
	}

	item, err := testQueries.CreateBuilding(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDCondo, item.IDCondo)
	require.Equal(t, arg.IDBuildingType, item.IDBuildingType)
	require.Equal(t, arg.Description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListBuildings(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBuilding(t)
	}

	arg := ListBuildingsParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListBuildings(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateBuilding(t *testing.T) {
	createRandomBuilding(t)
}

func TestGetBuilding(t *testing.T) {
	item1 := createRandomBuilding(t)
	item2, err := testQueries.GetBuilding(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.IDCondo, item2.IDCondo)
	require.Equal(t, item1.IDBuildingType, item2.IDBuildingType)
	require.Equal(t, item1.Description, item2.Description)
	// require.WithinDuration(t, item1.CreatedAt, item2.CreatedAt, time.Second)
}

func TestUpdateBuildings(t *testing.T) {
	arg := UpdateBuildingParams{
		ID:             createRandomBuilding(t).ID,
		IDBuildingType: createRandomBuildingType(t).ID,
		Description:    util.RandomString(10),
	}

	item, err := testQueries.UpdateBuilding(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.ID, item.ID)
	require.Equal(t, arg.IDBuildingType, item.IDBuildingType)
	require.Equal(t, arg.Description, item.Description)
}

func TestDeleteBuilding(t *testing.T) {
	item1 := createRandomBuilding(t)
	err := testQueries.DeleteBuilding(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetBuilding(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
