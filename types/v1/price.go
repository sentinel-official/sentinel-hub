package v1

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type QuotePriceFunc func(ctx sdk.Context, basePrice sdk.DecCoin) (sdk.Coin, error)

func NewPriceFromString(s string) (Price, error) {
	parts := strings.Split(s, ";")
	if len(parts) != 3 {
		return Price{}, errors.New("invalid format: expected BaseValue;QuoteValue;Denom")
	}

	baseValue, err := sdkmath.LegacyNewDecFromStr(parts[0])
	if err != nil {
		return Price{}, fmt.Errorf("invalid base value: %w", err)
	}

	quoteValue, ok := sdkmath.NewIntFromString(parts[1])
	if !ok {
		return Price{}, errors.New("invalid quote value")
	}

	denom := parts[2]
	if err := sdk.ValidateDenom(denom); err != nil {
		return Price{}, fmt.Errorf("invalid denomination: %w", err)
	}

	return Price{
		Denom:      denom,
		BaseValue:  baseValue,
		QuoteValue: quoteValue,
	}, nil
}

func ZeroPrice(denom string) Price {
	return Price{
		Denom:      denom,
		BaseValue:  sdkmath.LegacyZeroDec(),
		QuoteValue: sdkmath.ZeroInt(),
	}
}

func (p Price) BasePrice() sdk.DecCoin {
	return sdk.DecCoin{
		Denom:  p.Denom,
		Amount: p.BaseValue,
	}
}

func (p Price) QuotePrice() sdk.Coin {
	return sdk.Coin{
		Denom:  p.Denom,
		Amount: p.QuoteValue,
	}
}

func (p Price) Copy() Price {
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue,
		QuoteValue: p.QuoteValue,
	}
}

func (p Price) String() string {
	return fmt.Sprintf("%s;%s;%s", p.BaseValue.String(), p.QuoteValue.String(), p.Denom)
}

func (p Price) IsEqual(v Price) bool {
	return p.Denom == v.Denom &&
		p.BaseValue.Equal(v.BaseValue) &&
		p.QuoteValue.Equal(v.QuoteValue)
}

func (p Price) IsValid() bool {
	if sdk.ValidateDenom(p.Denom) != nil {
		return false
	}
	if p.BaseValue.IsNegative() || p.QuoteValue.IsNegative() {
		return false
	}

	return true
}

func (p Price) negative() Price {
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Neg(),
		QuoteValue: p.QuoteValue.Neg(),
	}
}

func (p Price) Add(v Price) Price {
	if p.Denom != v.Denom {
		panic(errors.New("denominations do not match"))
	}

	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Add(v.BaseValue),
		QuoteValue: p.QuoteValue.Add(v.QuoteValue),
	}
}

func (p Price) Sub(v Price) Price {
	if p.Denom != v.Denom {
		panic(errors.New("denominations do not match"))
	}

	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Sub(v.BaseValue),
		QuoteValue: p.QuoteValue.Sub(v.QuoteValue),
	}
}

func (p Price) UpdateQuoteValue(ctx sdk.Context, fn QuotePriceFunc) (Price, error) {
	// If BaseValue is zero, return the original Price without modification
	if p.BaseValue.IsZero() {
		return p, nil
	}

	// Get the base price using BasePrice()
	basePrice := p.BasePrice()

	// Compute the new quote value using the provided function
	newQuote, err := fn(ctx, basePrice)
	if err != nil {
		return Price{}, err
	}

	// Return a new Price instance with the updated QuoteValue
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue,
		QuoteValue: newQuote.Amount,
	}, nil
}

type Prices []Price

func NewPricesFromString(s string) (Prices, error) {
	parts := strings.Split(s, ",")

	prices := make(Prices, len(parts))
	for i, part := range parts {
		price, err := NewPriceFromString(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price at index %d: %w", i, err)
		}

		prices[i] = price
	}

	return prices.Sort(), nil
}

func (p Prices) Copy() Prices {
	v := make(Prices, len(p))
	for i := range p {
		v[i] = p[i].Copy()
	}

	return v
}

func (p Prices) String() string {
	var parts []string
	for _, price := range p {
		parts = append(parts, price.String())
	}

	return strings.Join(parts, ",")
}

func (p Prices) IsEqual(v Prices) bool {
	if len(p) != len(v) {
		return false
	}

	for i := range p {
		if !p[i].IsEqual(v[i]) {
			return false
		}
	}

	return true
}

func (p Prices) IsSorted() bool {
	for i := 1; i < len(p); i++ {
		if p[i].Denom < p[i-1].Denom {
			return false
		}
	}

	return true
}

func (p Prices) IsValid() bool {
	for i := 0; i < len(p); i++ {
		if i > 0 && p[i].Denom <= p[i-1].Denom {
			return false
		}
		if !p[i].IsValid() {
			return false
		}
	}

	return true
}

func (p Prices) Len() int {
	return len(p)
}

func (p Prices) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Prices) Less(i, j int) bool {
	return p[i].Denom < p[j].Denom
}

func (p Prices) Sort() Prices {
	sort.Sort(p)
	return p
}

func (p Prices) AmountOf(denom string) (sdkmath.LegacyDec, sdkmath.Int) {
	index := p.IndexOf(denom)
	if index != -1 {
		return p[index].BaseValue, p[index].QuoteValue
	}

	return sdkmath.LegacyZeroDec(), sdkmath.ZeroInt()
}

func (p Prices) IndexOf(denom string) int {
	index := sort.Search(len(p), func(i int) bool {
		return p[i].Denom >= denom
	})
	if index < len(p) && p[index].Denom == denom {
		return index
	}

	return -1
}

func (p Prices) Find(denom string) (Price, bool) {
	index := p.IndexOf(denom)
	if index != -1 {
		return p[index], true
	}

	return Price{}, false
}

func (p Prices) merge(v Prices) Prices {
	m := make(map[string]Price)
	for _, price := range p {
		m[price.Denom] = price
	}

	for _, price := range v {
		curr, ok := m[price.Denom]
		if !ok {
			curr = ZeroPrice(price.Denom)
		}

		m[price.Denom] = curr.Add(price)
	}

	res := make(Prices, 0, len(m))
	for _, price := range m {
		res = append(res, price)
	}

	return res.Sort()
}

func (p Prices) negative() Prices {
	v := make(Prices, len(p))
	for i, price := range p {
		v[i] = price.negative()
	}

	return v
}

func (p Prices) add(v Prices) Prices {
	return p.merge(v)
}

func (p Prices) Add(items ...Price) Prices {
	return p.add(items)
}

func (p Prices) sub(v Prices) Prices {
	return p.merge(v.negative())
}

func (p Prices) Sub(items ...Price) Prices {
	return p.sub(items)
}
