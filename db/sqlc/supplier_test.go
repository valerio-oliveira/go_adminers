package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomSupplier(t *testing.T) RegistrationSupplier {
	arg := CreateSupplierParams{
		IDSupplierType: createRandomSupplierType(t).ID,
		Description:    util.RandomString(10),
	}

	item, err := testQueries.CreateSupplier(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.IDSupplierType, item.IDSupplierType)
	require.Equal(t, arg.Description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListSuppliers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSupplier(t)
	}

	arg := ListSuppliersParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListSuppliers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateSupplier(t *testing.T) {
	createRandomSupplier(t)
}

func TestGetSupplier(t *testing.T) {
	item1 := createRandomSupplier(t)
	item2, err := testQueries.GetSupplier(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.IDSupplierType, item2.IDSupplierType)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateSuppliers(t *testing.T) {
	arg := UpdateSupplierParams{
		ID:             createRandomSupplier(t).ID,
		IDSupplierType: createRandomSupplierType(t).ID,
		Description:    util.RandomString(10),
	}

	item, err := testQueries.UpdateSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.ID, item.ID)
	require.Equal(t, arg.IDSupplierType, item.IDSupplierType)
	require.Equal(t, arg.Description, item.Description)
}

func TestDeleteSupplier(t *testing.T) {
	item1 := createRandomSupplier(t)
	err := testQueries.DeleteSupplier(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetSupplier(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
