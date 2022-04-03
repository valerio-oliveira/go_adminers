package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomUnityType(t *testing.T) RegistrationUnityType {

	description := util.RandomString(10)

	item, err := testQueries.CreateUnityType(context.Background(), description)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListUnityTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUnityType(t)
	}

	arg := ListUnityTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListUnityTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateUnityType(t *testing.T) {
	createRandomUnityType(t)
}

func TestGetUnityType(t *testing.T) {
	item1 := createRandomUnityType(t)
	item2, err := testQueries.GetUnityType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateUnityTypes(t *testing.T) {
	arg := UpdateUnityTypeParams{
		ID:          createRandomUnityType(t).ID,
		Description: util.RandomString(10),
	}

	item2, err := testQueries.UpdateUnityType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Description, item2.Description)
}

func TestDeleteUnityType(t *testing.T) {
	item1 := createRandomUnityType(t)
	err := testQueries.DeleteUnityType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetUnityType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
