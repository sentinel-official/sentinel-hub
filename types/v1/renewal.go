package v1

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// String converts a RenewalPricePolicy to its string representation.
func (r RenewalPricePolicy) String() string {
	switch r {
	case RenewalPricePolicyIfLesser:
		return "if_lesser"
	case RenewalPricePolicyIfLesserOrEqual:
		return "if_lesser_or_equal"
	case RenewalPricePolicyIfEqual:
		return "if_equal"
	case RenewalPricePolicyIfNotEqual:
		return "if_not_equal"
	case RenewalPricePolicyIfGreater:
		return "if_greater"
	case RenewalPricePolicyIfGreaterOrEqual:
		return "if_greater_or_equal"
	case RenewalPricePolicyAlways:
		return "always"
	default:
		return "unspecified"
	}
}

// Equal checks if two RenewalPricePolicy values are equal.
func (r RenewalPricePolicy) Equal(v RenewalPricePolicy) bool {
	return r == v
}

// IsValid checks whether the RenewalPricePolicy is a valid value.
func (r RenewalPricePolicy) IsValid() bool {
	switch r {
	case RenewalPricePolicyUnspecified,
		RenewalPricePolicyIfLesser,
		RenewalPricePolicyIfLesserOrEqual,
		RenewalPricePolicyIfEqual,
		RenewalPricePolicyIfNotEqual,
		RenewalPricePolicyIfGreater,
		RenewalPricePolicyIfGreaterOrEqual,
		RenewalPricePolicyAlways:
		return true
	default:
		return false
	}
}

// RenewalPricePolicyFromString converts a string to a RenewalPricePolicy.
func RenewalPricePolicyFromString(s string) RenewalPricePolicy {
	s = strings.ToLower(s)
	switch s {
	case "if_lesser":
		return RenewalPricePolicyIfLesser
	case "if_lesser_or_equal":
		return RenewalPricePolicyIfLesserOrEqual
	case "if_equal":
		return RenewalPricePolicyIfEqual
	case "if_not_equal":
		return RenewalPricePolicyIfNotEqual
	case "if_greater":
		return RenewalPricePolicyIfGreater
	case "if_greater_or_equal":
		return RenewalPricePolicyIfGreaterOrEqual
	case "always":
		return RenewalPricePolicyAlways
	default:
		return RenewalPricePolicyUnspecified
	}
}

// Validate validates whether a subscription can be renewed based on the policy and given DecCoin conditions.
// Returns an error if the renewal is not allowed or invalid.
func (r RenewalPricePolicy) Validate(currentPrice, previousPrice sdk.DecCoin) error {
	// Skip denomination check for RenewalPricePolicyAlways
	if r != RenewalPricePolicyAlways && currentPrice.Denom != previousPrice.Denom {
		return fmt.Errorf("current price denom %s does not match previous price denom %s", currentPrice.Denom, previousPrice.Denom)
	}

	// Compare prices based on the policy
	switch r {
	case RenewalPricePolicyUnspecified:
		return fmt.Errorf("renewal policy unspecified")
	case RenewalPricePolicyIfLesser:
		if !currentPrice.Amount.LT(previousPrice.Amount) {
			return fmt.Errorf("current price %s is not less than previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyIfLesserOrEqual:
		if !currentPrice.Amount.LTE(previousPrice.Amount) {
			return fmt.Errorf("current price %s is not less than or equal to previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyIfEqual:
		if !currentPrice.Amount.Equal(previousPrice.Amount) {
			return fmt.Errorf("current price %s is not equal to previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyIfNotEqual:
		if currentPrice.Amount.Equal(previousPrice.Amount) {
			return fmt.Errorf("current price %s is equal to previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyIfGreater:
		if !currentPrice.Amount.GT(previousPrice.Amount) {
			return fmt.Errorf("current price %s is not greater than previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyIfGreaterOrEqual:
		if !currentPrice.Amount.GTE(previousPrice.Amount) {
			return fmt.Errorf("current price %s is not greater than or equal to previous price %s", currentPrice, previousPrice)
		}
	case RenewalPricePolicyAlways:
		return nil
	default:
		return errors.New("invalid renewal policy")
	}

	return nil
}
