package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
)

func newBalanceFromCoins(coins []*Coin) (*Balance, error) {
	b := new(strings.Builder)
	for i, c := range coins {
		// if it's not the first element we write the comma to separate denoms
		if i != 0 {
			_, _ = b.WriteString(",")
		}
		_, _ = b.WriteString(fmt.Sprintf("%s%s", c.Amount, c.Denom))
	}
	sdkCoins, err := types.ParseCoinsNormalized(b.String())
	if err != nil {
		return nil, err
	}
	return &Balance{Coins: sdkCoins}, nil
}

type Balance struct {
	types.Coins
}

func (b *Balance) ToCoins() []*Coin {
	c := make([]*Coin, len(b.Coins))
	for i, x := range b.Coins {
		c[i] = &Coin{Amount: x.Amount.String(), Denom: x.Denom}
	}
	return c
}

// SafeSub subtracts from initial the amount the provided coins
// NOTE: if anything is negative an error is returned.
func SafeSub(initial []*Coin, toSub []*Coin) (result []*Coin, err error) {
	i, err := newBalanceFromCoins(initial)
	if err != nil {
		return nil, err
	}
	s, err := newBalanceFromCoins(toSub)
	if err != nil {
		return nil, err
	}
	if i.IsAnyNegative() || s.IsAnyNegative() {
		return nil, fmt.Errorf("negative coins")
	}
	sdkRes, ok := i.SafeSub(s.Coins)
	if !ok {
		return nil, fmt.Errorf("negative balance")
	}
	return (&Balance{Coins: sdkRes}).ToCoins(), nil
}

// SafeAdd adds to initial the provided amount
// NOTE: if anything is negative an error is returned
func SafeAdd(initial []*Coin, toAdd []*Coin) (result []*Coin, err error) {
	i, err := newBalanceFromCoins(initial)
	if err != nil {
		return nil, err
	}
	s, err := newBalanceFromCoins(toAdd)
	if err != nil {
		return nil, err
	}
	if i.IsAnyNegative() || s.IsAnyNegative() {
		return nil, fmt.Errorf("negative coins")
	}
	sdkRes := i.Add(s.Coins...)
	return (&Balance{Coins: sdkRes}).ToCoins(), nil
}
