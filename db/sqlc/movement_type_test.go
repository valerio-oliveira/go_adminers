package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valerio-oliveira/go_adminers/util"
)

func createRandomMovementType(t *testing.T) FinanceMovementType {

	arg := CreateMovementTypeParams{
		Description:       util.RandomString(10),
		Direction:         util.RandomDirection(),
		IDDefaultSupplier: sql.NullInt32{Int32: createRandomSupplier(t).ID, Valid: true},
	}

	item, err := testQueries.CreateMovementType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, arg.Direction, item.Direction)
	require.Equal(t, arg.IDDefaultSupplier, item.IDDefaultSupplier)
	require.NotZero(t, item.ID)

	return item
}

func TestListMovementTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMovementType(t)
	}

	arg := ListMovementTypesParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListMovementTypes(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)

	for _, item := range items {
		require.NotZero(t, item.ID)
		require.NotEmpty(t, item)
	}
}

func TestCreateMovementType(t *testing.T) {
	createRandomMovementType(t)
}

func TestGetMovementType(t *testing.T) {
	item1 := createRandomMovementType(t)
	item2, err := testQueries.GetMovementType(context.Background(), item1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Description, item2.Description)
	require.Equal(t, item1.Direction, item2.Direction)
	require.Equal(t, item1.IDDefaultSupplier, item2.IDDefaultSupplier)
}

func TestUpdateMovementTypes(t *testing.T) {
	arg := UpdateMovementTypeParams{
		ID:                createRandomMovementType(t).ID,
		Description:       util.RandomString(10),
		Direction:         util.RandomDirection(),
		IDDefaultSupplier: sql.NullInt32{Int32: createRandomSupplier(t).ID, Valid: true},
	}

	item, err := testQueries.UpdateMovementType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.Equal(t, arg.ID, item.ID)
	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, arg.Direction, item.Direction)
	require.Equal(t, arg.IDDefaultSupplier, item.IDDefaultSupplier)
}

func TestDeleteMovementType(t *testing.T) {
	item1 := createRandomMovementType(t)
	err := testQueries.DeleteMovementType(context.Background(), item1.ID)
	require.NoError(t, err)

	item2, err := testQueries.GetMovementType(context.Background(), item1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, item2)
}
