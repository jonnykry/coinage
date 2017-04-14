package mocks

import coinagemodels "github.com/ghmeier/coinage/models"
import gateways "github.com/ghmeier/coinage/gateways"
import mock "github.com/stretchr/testify/mock"
import models "github.com/jakelong95/TownCenter/models"
import stripe "github.com/stripe/stripe-go"
import warehousemodels "github.com/lcollin/warehouse/models"

// Stripe is an autogenerated mock type for the Stripe type
type Stripe struct {
	mock.Mock
}

// AddSource provides a mock function with given fields: id, token
func (_m *Stripe) AddSource(id string, token string) (*stripe.Customer, error) {
	ret := _m.Called(id, token)

	var r0 *stripe.Customer
	if rf, ok := ret.Get(0).(func(string, string) *stripe.Customer); ok {
		r0 = rf(id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stripe.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomer provides a mock function with given fields: id
func (_m *Stripe) DeleteCustomer(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCustomer provides a mock function with given fields: id
func (_m *Stripe) GetCustomer(id string) (*stripe.Customer, error) {
	ret := _m.Called(id)

	var r0 *stripe.Customer
	if rf, ok := ret.Get(0).(func(string) *stripe.Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stripe.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlan provides a mock function with given fields: secret, pid
func (_m *Stripe) GetPlan(secret string, pid string) (*stripe.Plan, error) {
	ret := _m.Called(secret, pid)

	var r0 *stripe.Plan
	if rf, ok := ret.Get(0).(func(string, string) *stripe.Plan); ok {
		r0 = rf(secret, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stripe.Plan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(secret, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccount provides a mock function with given fields: user, roaster
func (_m *Stripe) NewAccount(user *models.User, roaster *models.Roaster) (*stripe.Account, error) {
	ret := _m.Called(user, roaster)

	var r0 *stripe.Account
	if rf, ok := ret.Get(0).(func(*models.User, *models.Roaster) *stripe.Account); ok {
		r0 = rf(user, roaster)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stripe.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User, *models.Roaster) error); ok {
		r1 = rf(user, roaster)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCustomer provides a mock function with given fields: token, userID
func (_m *Stripe) NewCustomer(token string, userID string) (string, error) {
	ret := _m.Called(token, userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(token, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPlan provides a mock function with given fields: secret, item, freq
func (_m *Stripe) NewPlan(secret string, item *warehousemodels.Item, freq coinagemodels.Frequency) (*stripe.Plan, error) {
	ret := _m.Called(secret, item, freq)

	var r0 *stripe.Plan
	if rf, ok := ret.Get(0).(func(string, *warehousemodels.Item, coinagemodels.Frequency) *stripe.Plan); ok {
		r0 = rf(secret, item, freq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stripe.Plan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *warehousemodels.Item, coinagemodels.Frequency) error); ok {
		r1 = rf(secret, item, freq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: roaster, id, planID, quantity
func (_m *Stripe) Subscribe(roaster *coinagemodels.Roaster, id string, planID string, quantity uint64) (string, error) {
	ret := _m.Called(roaster, id, planID, quantity)

	var r0 string
	if rf, ok := ret.Get(0).(func(*coinagemodels.Roaster, string, string, uint64) string); ok {
		r0 = rf(roaster, id, planID, quantity)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*coinagemodels.Roaster, string, string, uint64) error); ok {
		r1 = rf(roaster, id, planID, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ gateways.Stripe = (*Stripe)(nil)
