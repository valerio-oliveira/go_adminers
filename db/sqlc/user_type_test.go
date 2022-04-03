package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomUserType(t *testing.T) AccessUserType {

	description := util.RandomString(10)

	item, err := testQueries.CreateUserType(context.Background(), description)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListUserTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUserType(t)
	}

	arg := ListUserTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListUserTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateUserType(t *testing.T) {
	createRandomUserType(t)
}

func TestGetUserType(t *testing.T) {
	item1 := createRandomUserType(t)
	item2, err := testQueries.GetUserType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateUserTypes(t *testing.T) {
	arg := UpdateUserTypeParams{
		ID:          createRandomUserType(t).ID,
		Description: util.RandomString(10),
	}

	item2, err := testQueries.UpdateUserType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Description, item2.Description)
}

func TestDeleteUserType(t *testing.T) {
	item1 := createRandomUserType(t)
	err := testQueries.DeleteUserType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetUserType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
