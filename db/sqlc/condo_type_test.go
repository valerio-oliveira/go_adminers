package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomCondoType(t *testing.T) RegistrationCondoType {

	description := util.RandomString(10)

	item, err := testQueries.CreateCondoType(context.Background(), description)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, description, item.Description)
	require.NotZero(t, item.ID)

	return item
}

func TestListCondoTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCondoType(t)
	}

	arg := ListCondoTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListCondoTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateCondoType(t *testing.T) {
	createRandomCondoType(t)
}

func TestGetCondoType(t *testing.T) {
	item1 := createRandomCondoType(t)
	item2, err := testQueries.GetCondoType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
}

func TestUpdateCondoTypes(t *testing.T) {
	arg := UpdateCondoTypeParams{
		ID:          createRandomCondoType(t).ID,
		Description: util.RandomString(10),
	}

	item2, err := testQueries.UpdateCondoType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Description, item2.Description)
}

func TestDeleteCondoType(t *testing.T) {
	item1 := createRandomCondoType(t)
	err := testQueries.DeleteCondoType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetCondoType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
