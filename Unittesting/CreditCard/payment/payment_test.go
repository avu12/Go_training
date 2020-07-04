package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../entity"
	"github.com/stretchr/testify/mock"
)

type AttemptHistory struct {
	mock.Mock
}

func (a *AttemptHistory) IncrementFailure(user entity.User) error {
	args := a.Called(user)

	return args.Error(0)
}

func (a *AttemptHistory) CountFailures(user entity.User) (int, error) {
	args := a.Called(user)

	return args.Int(0), args.Error(1)
}

type GatewayMock struct {
	mock.Mock
}

func (gm *GatewayMock) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	args := gm.Called(user, creditCard)

	return args.Bool(0), args.Error(1)
}
func (gm *GatewayMock) Pay(creditCard entity.CreditCard, amount int) error {
	args := gm.Called(creditCard, amount)

	return args.Error(0)
}

func TestShouldHaveASuccessfullAuthorisation(t *testing.T) {

	attemptHistory := &AttemptHistory{}
	user := entity.User{}
	creditCard := entity.CreditCard{}

	attemptHistory.On("CountFailures", user).Return(2, nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(true, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	IsAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	attemptHistory.AssertNotCalled(t, "IncrementFailure", user)
	assert.True(t, IsAuthorized)

}
func TestShouldHaveFailedAuthorisation(t *testing.T) {
	attemptHistory := &AttemptHistory{}
	user := entity.User{}
	creditCard := entity.CreditCard{}

	attemptHistory.On("CountFailures", user).Return(3, nil)
	attemptHistory.On("IncrementFailure", user).Return(nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(false, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	IsAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	attemptHistory.AssertCalled(t, "IncrementFailure", user)
	assert.False(t, IsAuthorized)
}

func TestShouldHaveAForcedFailedAuthorisation(t *testing.T) {
	attemptHistory := &AttemptHistory{}
	user := entity.User{}
	creditCard := entity.CreditCard{}

	attemptHistory.On("CountFailures", user).Return(8, nil)
	gateway := &GatewayMock{}
	paymentService := NewPaymentService(attemptHistory, gateway)

	IsAuthorized, err := paymentService.IsAuthorized(user, creditCard)
	if err != nil {
		t.Fatal(err.Error())
	}
	attemptHistory.AssertNotCalled(t, "IncrementFailure", user)
	attemptHistory.AssertNotCalled(t, "IsAuthorized", user, creditCard)
	assert.False(t, IsAuthorized)
}
