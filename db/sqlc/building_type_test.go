package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomBuildingType(t *testing.T) RegistrationBuildingType {

	description := util.RandomString(10)

	item, err := testQueries.CreateBuildingType(context.Background(), description)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListBuildingTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBuildingType(t)
	}

	arg := ListBuildingTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListBuildingTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateBuildingType(t *testing.T) {
	createRandomBuildingType(t)
}

func TestGetBuildingType(t *testing.T) {
	item1 := createRandomBuildingType(t)
	item2, err := testQueries.GetBuildingType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateBuildingTypes(t *testing.T) {
	arg := UpdateBuildingTypeParams{
		ID:          createRandomBuildingType(t).ID,
		Description: util.RandomString(10),
	}

	item2, err := testQueries.UpdateBuildingType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Description, item2.Description)
}

func TestDeleteBuildingType(t *testing.T) {
	item1 := createRandomBuildingType(t)
	err := testQueries.DeleteBuildingType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetBuildingType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
