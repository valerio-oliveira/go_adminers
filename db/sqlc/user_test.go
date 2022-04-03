package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomUser(t *testing.T) AccessUser {
	login := util.RandomString(8)
	arg := CreateUserParams{
		Login:  login,
		Email:  util.RandomEmail(login),
		Name:   util.RandomString(10) + " " + util.RandomString(10),
		Phone:  util.RandomPhone(),
		Hash:   util.RandomHashMD5(login),
		Active: util.RandomBool(),
	}

	item, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Login, item.Login)
	require.Equal(t, arg.Email, item.Email)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Phone, item.Phone)
	require.Equal(t, arg.Hash, item.Hash)
	require.Equal(t, arg.Active, item.Active)
	require.NotZero(t, item.ID)

	return item
}

func TestListUses(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	item1 := createRandomUser(t)
	item2, err := testQueries.GetUser(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Login, item2.Login)
	require.Equal(t, item1.Email, item2.Email)
	require.Equal(t, item1.Name, item2.Name)
	require.Equal(t, item1.Phone, item2.Phone)
	require.Equal(t, item1.Hash, item2.Hash)
	require.Equal(t, item1.Active, item2.Active)
}

func TestUpdateUser(t *testing.T) {
	item1 := createRandomUser(t)
	login := item1.Login
	arg := UpdateUserParams{
		ID:    item1.ID,
		Email: util.RandomEmail(login),
		Name:  util.RandomString(10) + " " + util.RandomString(10),
		Phone: util.RandomPhone(),
		Hash:  util.RandomHashMD5(login),
	}

	item2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, login, item2.Login)
	require.Equal(t, arg.Email, item2.Email)
	require.Equal(t, arg.Name, item2.Name)
	require.Equal(t, arg.Phone, item2.Phone)
	require.Equal(t, arg.Hash, item2.Hash)
	require.Equal(t, arg.Active, item2.Active)
}

func TestDeleteUser(t *testing.T) {
	item1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetUser(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
