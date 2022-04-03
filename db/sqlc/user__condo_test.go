package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserCondo(t *testing.T) AccessUserCondo {
	arg := CreateUserCondoParams{
		IDUser:     createRandomUser(t).ID,
		IDCondo:    createRandomCondo(t).ID,
		IDUserType: createRandomUserType(t).ID,
	}

	item, err := testQueries.CreateUserCondo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDUser, item.IDUser)
	require.Equal(t, arg.IDCondo, item.IDCondo)
	require.Equal(t, arg.IDUserType, item.IDUserType)

	return item
}

func TestListUserCondos(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUserCondo(t)
	}

	arg := ListUserCondosParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListUserCondos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.IDUser)
		require.NotZero(t, item.IDCondo)
		require.NotZero(t, item.IDUserType)
		require.NotEmpty(t, item)
	}
}

func TestCreateUserCondo(t *testing.T) {
	createRandomUserCondo(t)
}

func TestGetUserCondo(t *testing.T) {
	item1 := createRandomUserCondo(t)
	arg := GetUserCondoParams{
		IDUser:  item1.IDUser,
		IDCondo: item1.IDCondo,
	}
	item2, err := testQueries.GetUserCondo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.IDUser, item2.IDUser)
	require.Equal(t, item1.IDCondo, item2.IDCondo)
	require.Equal(t, item1.IDUserType, item2.IDUserType)
}

func TestUpdateUserCondo(t *testing.T) {
	item1 := createRandomUserCondo(t)
	arg := UpdateUserCondoParams{
		IDUser:     item1.IDUser,
		IDCondo:    item1.IDCondo,
		IDUserType: item1.IDUserType,
	}

	item, err := testQueries.UpdateUserCondo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDUser, item.IDUser)
	require.Equal(t, arg.IDCondo, item.IDCondo)
	require.Equal(t, arg.IDUserType, item.IDUserType)
}

func TestDeleteUserCondo(t *testing.T) {
	item1 := createRandomUserCondo(t)
	arg1 := DeleteUserCondoParams{
		IDUser:  item1.IDUser,
		IDCondo: item1.IDCondo,
	}
	err := testQueries.DeleteUserCondo(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetUserCondoParams{
		IDUser:  item1.IDUser,
		IDCondo: item1.IDCondo,
	}
	item2, err := testQueries.GetUserCondo(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
