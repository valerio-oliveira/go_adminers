package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomSupplierType(t *testing.T) RegistrationSupplierType {

	description := util.RandomString(10)

	item, err := testQueries.CreateSupplierType(context.Background(), description)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListSupplierTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSupplierType(t)
	}

	arg := ListSupplierTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListSupplierTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateSupplierType(t *testing.T) {
	createRandomSupplierType(t)
}

func TestGetSupplierType(t *testing.T) {
	item1 := createRandomSupplierType(t)
	item2, err := testQueries.GetSupplierType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateSupplierTypes(t *testing.T) {
	arg := UpdateSupplierTypeParams{
		ID:          createRandomSupplierType(t).ID,
		Description: util.RandomString(10),
	}

	item2, err := testQueries.UpdateSupplierType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Description, item2.Description)
}

func TestDeleteSupplierType(t *testing.T) {
	item1 := createRandomSupplierType(t)
	err := testQueries.DeleteSupplierType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetSupplierType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
