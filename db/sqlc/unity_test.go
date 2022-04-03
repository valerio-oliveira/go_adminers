package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomUnity(t *testing.T) RegistrationUnity {
	arg := CreateUnityParams{
		IDBuilding:  createRandomBuilding(t).ID,
		IDUnityType: createRandomUnityType(t).ID,
		UnityNumber: util.RandomString(4),
	}

	item, err := testQueries.CreateUnity(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDBuilding, item.IDBuilding)
	require.Equal(t, arg.IDUnityType, item.IDUnityType)
	require.Equal(t, arg.UnityNumber, item.UnityNumber)
	require.NotZero(t, item.ID)

	return item
}

func TestListUnities(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUnity(t)
	}

	arg := ListUnitysParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListUnitys(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateUnity(t *testing.T) {
	createRandomUnity(t)
}

func TestGetUnity(t *testing.T) {
	item1 := createRandomUnity(t)
	item2, err := testQueries.GetUnity(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.IDBuilding, item2.IDBuilding)
	require.Equal(t, item1.IDUnityType, item2.IDUnityType)
	require.Equal(t, item1.UnityNumber, item2.UnityNumber)
}

func TestUpdateUnity(t *testing.T) {
	arg := UpdateUnityParams{
		ID:          createRandomUnity(t).ID,
		IDBuilding:  createRandomBuilding(t).ID,
		IDUnityType: createRandomUnityType(t).ID,
		UnityNumber: util.RandomString(10),
	}

	item, err := testQueries.UpdateUnity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.ID, item.ID)
	require.Equal(t, arg.IDUnityType, item.IDUnityType)
	require.Equal(t, arg.UnityNumber, item.UnityNumber)
}

func TestDeleteUnity(t *testing.T) {
	item1 := createRandomUnity(t)
	err := testQueries.DeleteUnity(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetUnity(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
