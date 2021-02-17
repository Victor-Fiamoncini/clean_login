package mocks

import (
	"context"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// Database struct
type Database struct {
	mock.Mock
}

// InsertOne database method
func (d *Database) InsertOne(ctx context.Context, user *entities.User) error {
	ret := d.Called(ctx, user)

	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		return rf(ctx, user)
	}

	return ret.Error(0)
}

// FindOne Database method
func (d *Database) FindOne(ctx context.Context, filter interface{}) (*entities.User, error) {
	ret := d.Called(ctx, filter)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *entities.User); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
