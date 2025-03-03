package types

import (
	"fmt"
	"math"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		VotingRightsExpirationTime:      86000,
		CollectionSize:                  24,
		CollectionPrice:                 sdk.NewInt64Coin("ucredits", 10000000),
		ActiveCollectionsAmount:         3,
		CollectionCreationFee:           sdk.NewInt64Coin("ucredits", int64(5000*math.Pow(10, 6))),
		CollateralDeposit:               sdk.NewInt64Coin("ucredits", int64(50*math.Pow(10, 6))),
        TrialVoteReward:               	 sdk.NewInt64Coin("ucredits", int64(math.Pow(10, 6))),
		WinnerReward:                    int64(math.Pow(10, 6)),
		VotePoolFraction:                int64(math.Pow(10, 6)),
        VotingRewardCap:                 int64(math.Pow(10, 6)),
		HourlyFaucet:                    sdk.NewInt64Coin("ucredits", int64(50*math.Pow(10, 6))),
		InflationRate:                   "1.1", // TODO: Also make this a fixed point number
		RaresPerPack:                    1,
		CommonsPerPack:                  9,
		UnCommonsPerPack:                3,
		TrialPeriod:                     14 * 24 * 500,
		GameVoteRatio:                   20, // This is a fixed point number and will be devided by 100 when used
		CardAuctionPriceReductionPeriod: 20,
		AirDropValue:                    sdk.NewInt64Coin("ubpf", int64(5*math.Pow(10, 6))),
		AirDropMaxBlockHeight:           5000000,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte("VotingRightsExpirationTime"), &p.VotingRightsExpirationTime, validateVotingRightsExpirationTime),
		paramtypes.NewParamSetPair([]byte("CollectionSize"), &p.CollectionSize, validateCollectionSize),
		paramtypes.NewParamSetPair([]byte("CollectionPrice"), &p.CollectionPrice, validateCollectionPrice),
        paramtypes.NewParamSetPair([]byte("TrialVoteReward"), &p.TrialVoteReward, validateCollectionPrice),
		paramtypes.NewParamSetPair([]byte("ActiveCollectionsAmount"), &p.ActiveCollectionsAmount, validateActiveCollectionsAmount),
		paramtypes.NewParamSetPair([]byte("CollectionCreationFee"), &p.CollectionCreationFee, validateCollectionCreationFee),
		paramtypes.NewParamSetPair([]byte("CollateralDeposit"), &p.CollateralDeposit, validateCollateralDeposit),
		paramtypes.NewParamSetPair([]byte("WinnerReward"), &p.WinnerReward, validateWinnerReward),
        paramtypes.NewParamSetPair([]byte("VotePoolFraction"), &p.VotePoolFraction, validateVoterReward),
        paramtypes.NewParamSetPair([]byte("VotingRewardCap"), &p.VotingRewardCap, validateVoterReward),
        paramtypes.NewParamSetPair([]byte("HourlyFaucet"), &p.HourlyFaucet, validateHourlyFaucet),
		paramtypes.NewParamSetPair([]byte("InflationRate"), &p.InflationRate, validateInflationRate),
		paramtypes.NewParamSetPair([]byte("RaresPerPack"), &p.RaresPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("CommonsPerPack"), &p.CommonsPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("UnCommonsPerPack"), &p.UnCommonsPerPack, validatePerPack),
		paramtypes.NewParamSetPair([]byte("TrialPeriod"), &p.TrialPeriod, validateTrialPeriod),
		paramtypes.NewParamSetPair([]byte("GameVoteRatio"), &p.GameVoteRatio, validateGameVoteRatio),
		paramtypes.NewParamSetPair([]byte("CardAuctionPriceReductionPeriod"), &p.CardAuctionPriceReductionPeriod, validateCardAuctionPriceReductionPeriod),
		paramtypes.NewParamSetPair([]byte("AirDropValue"), &p.AirDropValue, validateAirDropValue),
		paramtypes.NewParamSetPair([]byte("AirDropMaxBlockHeight"), &p.AirDropMaxBlockHeight, validateAirDropMaxBlockHeight),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateVotingRightsExpirationTime(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid VotingRightsExpirationTime: %d", v)
	}

	return nil
}

func validateCollectionSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid CollectionSize: %d", v)
	}

	return nil
}

func validateCollectionPrice(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollectionPrice: %v", v)
	}

	return nil
}

func validateActiveCollectionsAmount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid ActiveCollectionsAmount: %d", v)
	}

	return nil
}

func validateTrialPeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid TrialPeriod: %d", v)
	}

	return nil
}

func validatePerPack(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid number per pack: %d", v)
	}

	return nil
}

func validateGameVoteRatio(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid GameVoteRatio: %d", v)
	}

	return nil
}

func validateCardAuctionPriceReductionPeriod(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid CardAuctionPriceReductionPeriod: %d", v)
	}

	return nil
}

func validateCollectionCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollectionCreationFee: %v", v)
	}

	return nil
}

func validateAirDropValue(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid AirDropValue: %v", v)
	}

	return nil
}

func validateCollateralDeposit(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid CollateralDeposit: %v", v)
	}

	return nil
}

func validateWinnerReward(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid WinnerReward: %d", v)
	}

	return nil
}

func validateVoterReward(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid VoterReward: %d", v)
	}

	return nil
}

func validateAirDropMaxBlockHeight(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("invalid AirDropMaxBlockHeight: %d", v)
	}

	return nil
}

func validateHourlyFaucet(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == sdk.NewInt64Coin("ucredits", 0) {
		return fmt.Errorf("invalid HourlyFaucet: %v", v)
	}

	return nil
}

func validateInflationRate(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := strconv.ParseFloat(v, 8)
	if err != nil {
		return fmt.Errorf("invalid parameter: %d", err)
	}
	return nil
}
