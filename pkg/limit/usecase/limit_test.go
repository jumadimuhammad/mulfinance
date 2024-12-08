package usecase

import (
	"context"
	limitMock "mulfinance/mocks"
	"mulfinance/pkg/limit/repository/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repository = &limitMock.IRepository{}
)

func TestListLimit(t *testing.T) {
	res := []entities.Limit{
		{
			ID:        1,
			Limit:     100000,
			Tenor:     1,
			CreatedBy: 1,
			UpdatedBy: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Limit:     200000,
			Tenor:     2,
			CreatedBy: 1,
			UpdatedBy: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Limit:     500000,
			Tenor:     3,
			CreatedBy: 1,
			UpdatedBy: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        4,
			Limit:     700000,
			Tenor:     4,
			CreatedBy: 1,
			UpdatedBy: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	repository.On("ListLimit", mock.Anything).Return(res, nil).Once()

	us := NewUsecase(repository)

	resp, err := us.ListLimit(context.Background())

	assert.NotEmpty(t, resp)
	assert.NotNil(t, resp)
	assert.Len(t, resp, len(res))
	assert.Nil(t, err)
}
