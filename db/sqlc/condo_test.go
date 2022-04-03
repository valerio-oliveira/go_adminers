package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomCondo(t *testing.T) RegistrationCondo {
	arg := CreateCondoParams{
		IDCondoType: createRandomCondoType(t).ID,
		Name:        util.RandomString(10),
		Nickname:    util.RandomString(10),
		Address1:    util.RandomStringNull(10),
		Address2:    util.RandomStringNull(10),
		Phone1:      util.RandomStringNull(10),
		Phone2:      util.RandomStringNull(10),
		Email:       util.RandomString(10),
		Cnpj:        util.RandomInt64(),
	}

	item, err := testQueries.CreateCondo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDCondoType, item.IDCondoType)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Nickname, item.Nickname)
	require.Equal(t, arg.Address1, item.Address1)
	require.Equal(t, arg.Address2, item.Address2)
	require.Equal(t, arg.Phone1, item.Phone1)
	require.Equal(t, arg.Phone2, item.Phone2)
	require.Equal(t, arg.Email, item.Email)
	require.Equal(t, arg.Cnpj, item.Cnpj)
	require.NotZero(t, item.ID)

	return item
}

func TestListCondos(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCondo(t)
	}

	arg := ListCondosParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListCondos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateCondo(t *testing.T) {
	createRandomCondo(t)
}

func TestGetCondo(t *testing.T) {
	item1 := createRandomCondo(t)
	item2, err := testQueries.GetCondo(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.IDCondoType, item2.IDCondoType)
	require.Equal(t, item1.Name, item2.Name)
	require.Equal(t, item1.Nickname, item2.Nickname)
	require.Equal(t, item1.Address1, item2.Address1)
	require.Equal(t, item1.Address2, item2.Address2)
	require.Equal(t, item1.Phone1, item2.Phone1)
	require.Equal(t, item1.Phone2, item2.Phone2)
	require.Equal(t, item1.Email, item2.Email)
	require.Equal(t, item1.Cnpj, item2.Cnpj)
}

func TestUpdateCondos(t *testing.T) {
	arg := UpdateCondoParams{
		ID:          createRandomCondo(t).ID,
		IDCondoType: createRandomCondoType(t).ID,
		Name:        util.RandomString(10),
		Nickname:    util.RandomString(10),
		Address1:    util.RandomStringNull(10),
		Address2:    util.RandomStringNull(10),
		Phone1:      util.RandomStringNull(10),
		Phone2:      util.RandomStringNull(10),
		Email:       util.RandomString(10),
		Cnpj:        util.RandomInt64(),
	}

	item, err := testQueries.UpdateCondo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.ID, item.ID)
	require.Equal(t, arg.IDCondoType, item.IDCondoType)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.Nickname, item.Nickname)
	require.Equal(t, arg.Address1, item.Address1)
	require.Equal(t, arg.Address2, item.Address2)
	require.Equal(t, arg.Phone1, item.Phone1)
	require.Equal(t, arg.Phone2, item.Phone2)
	require.Equal(t, arg.Email, item.Email)
	require.Equal(t, arg.Cnpj, item.Cnpj)
}

func TestDeleteCondo(t *testing.T) {
	item1 := createRandomCondo(t)
	err := testQueries.DeleteCondo(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetCondo(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)

}
