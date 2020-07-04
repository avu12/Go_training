package calculator

import (
	"testing"
)

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountCalculator(t *testing.T) {

	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}
	testCases := []testCase{
		{name: "Should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, expectedAmount: 130},
		{name: "Should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 200, expectedAmount: 160},
		{name: "Should apply 60", minimumPurchaseAmount: 100, purchaseAmount: 350, expectedAmount: 290},
		{name: "Should not apply", minimumPurchaseAmount: 100, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				//FailNow + logf
				t.Fatalf("Could not instantiate the calculator: %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)
			if tc.expectedAmount != amount {
				t.Errorf("expected %v, got %v. Fail caused by unexpected amount number ", tc.expectedAmount, amount)
			}
		})

	}
}

func TestDiscountCalculatorShouldFailWithZeroMinimumAmount(t *testing.T) {

	discountRepositoryMock := DiscountRepositoryMock{}
	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		//FailNow + logf
		t.Fatalf("Should not create the calculator with zero purchase amount")
	}

}
