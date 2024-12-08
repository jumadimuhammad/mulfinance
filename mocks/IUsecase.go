// Code generated by mockery v2.30.1. DO NOT EDIT.

package ports

import (
	context "context"
	domain "mulfinance/pkg/limit/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

type IUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *IUsecase) EXPECT() *IUsecase_Expecter {
	return &IUsecase_Expecter{mock: &_m.Mock}
}

// ListLimit provides a mock function with given fields: ctx
func (_m *IUsecase) ListLimit(ctx context.Context) ([]domain.Limit, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Limit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Limit, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Limit); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Limit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUsecase_ListLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListLimit'
type IUsecase_ListLimit_Call struct {
	*mock.Call
}

// ListLimit is a helper method to define mock.On call
//   - ctx context.Context
func (_e *IUsecase_Expecter) ListLimit(ctx interface{}) *IUsecase_ListLimit_Call {
	return &IUsecase_ListLimit_Call{Call: _e.mock.On("ListLimit", ctx)}
}

func (_c *IUsecase_ListLimit_Call) Run(run func(ctx context.Context)) *IUsecase_ListLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *IUsecase_ListLimit_Call) Return(limits []domain.Limit, err error) *IUsecase_ListLimit_Call {
	_c.Call.Return(limits, err)
	return _c
}

func (_c *IUsecase_ListLimit_Call) RunAndReturn(run func(context.Context) ([]domain.Limit, error)) *IUsecase_ListLimit_Call {
	_c.Call.Return(run)
	return _c
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
