package v1

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default period for deposits & voting
const (
	DefaultPeriod          time.Duration = time.Hour * 24 * 2 // 2 days
	DefaultExpeditedPeriod time.Duration = time.Hour * 24     // 1 day
)

// Default governance params
var (
	DefaultMinDepositTokens          = sdk.NewInt(10000000)
	DefaultMinExpeditedDepositTokens = sdk.NewInt(20000000)
	DefaultQuorum                    = sdk.NewDecWithPrec(334, 3)
	DefaultExpeditedQuorum           = sdk.NewDecWithPrec(667, 3)
	DefaultThreshold                 = sdk.NewDecWithPrec(5, 1)
	DefaultExpeditedThreshold        = sdk.NewDecWithPrec(667, 3)
	DefaultVetoThreshold             = sdk.NewDecWithPrec(334, 3)
)

// Parameter store key
var (
	ParamStoreKeyDepositParams = []byte("depositparams")
	ParamStoreKeyVotingParams  = []byte("votingparams")
	ParamStoreKeyTallyParams   = []byte("tallyparams")
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyDepositParams, DepositParams{}, validateDepositParams),
		paramtypes.NewParamSetPair(ParamStoreKeyVotingParams, VotingParams{}, validateVotingParams),
		paramtypes.NewParamSetPair(ParamStoreKeyTallyParams, TallyParams{}, validateTallyParams),
	)
}

// NewDepositParams creates a new DepositParams object
func NewDepositParams(minDeposit sdk.Coins, minExpeditedDeposit sdk.Coins, maxDepositPeriod time.Duration) DepositParams {
	return DepositParams{
		MinDeposit:          minDeposit,
		MaxDepositPeriod:    &maxDepositPeriod,
		MinExpeditedDeposit: minExpeditedDeposit,
	}
}

// DefaultDepositParams default parameters for deposits
func DefaultDepositParams() DepositParams {
	return NewDepositParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultMinDepositTokens)),
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultMinExpeditedDepositTokens)),
		DefaultPeriod,
	)
}

// Equal checks equality of DepositParams
func (dp DepositParams) Equal(dp2 DepositParams) bool {
	return sdk.Coins(dp.MinDeposit).IsEqual(dp2.MinDeposit) &&
		sdk.Coins(dp.MinExpeditedDeposit).IsEqual(dp2.MinExpeditedDeposit) &&
		dp.MaxDepositPeriod == dp2.MaxDepositPeriod
}

func validateDepositParams(i interface{}) error {
	v, ok := i.(DepositParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !sdk.Coins(v.MinDeposit).IsValid() {
		return fmt.Errorf("invalid minimum deposit: %s", v.MinDeposit)
	}
	if !sdk.Coins(v.MinExpeditedDeposit).IsValid() {
		return fmt.Errorf("invalid minimum expedited deposit: %s", v.MinExpeditedDeposit)
	}
	if sdk.Coins(v.MinExpeditedDeposit).IsAllLTE(v.MinDeposit) {
		return fmt.Errorf("minimum expedited deposit: %s should be larger than minimum deposit: %s", v.MinExpeditedDeposit, v.MinDeposit)
	}
	if v.MaxDepositPeriod == nil || v.MaxDepositPeriod.Seconds() <= 0 {
		return fmt.Errorf("maximum deposit period must be positive: %d", v.MaxDepositPeriod)
	}

	return nil
}

// GetMinimumDeposit returns minimum deposit based on the value isExpedited
func (dp DepositParams) GetMinimumDeposit(isExpedited bool) sdk.Coins {
	if isExpedited {
		return dp.MinExpeditedDeposit
	}
	return dp.MinDeposit
}

// NewTallyParams creates a new TallyParams object
func NewTallyParams(quorum, expeditedQuorum, threshold, expeditedThreshold, vetoThreshold sdk.Dec) TallyParams {
	return TallyParams{
		Quorum:             quorum.String(),
		ExpeditedQuorum:    expeditedQuorum.String(),
		Threshold:          threshold.String(),
		VetoThreshold:      vetoThreshold.String(),
		ExpeditedThreshold: expeditedThreshold.String(),
	}
}

// DefaultTallyParams default parameters for tallying
func DefaultTallyParams() TallyParams {
	return NewTallyParams(DefaultQuorum, DefaultExpeditedQuorum, DefaultThreshold, DefaultExpeditedThreshold, DefaultVetoThreshold)
}

// Equal checks equality of TallyParams
func (tp TallyParams) Equal(other TallyParams) bool {
	return tp.Quorum == other.Quorum &&
		tp.ExpeditedQuorum == other.ExpeditedQuorum &&
		tp.Threshold == other.Threshold &&
		tp.ExpeditedThreshold == other.ExpeditedThreshold &&
		tp.VetoThreshold == other.VetoThreshold
}

// GetTallyThreshold returns threshold based on the value isExpedited
func (tp TallyParams) GetTallyThreshold(isExpedited bool) sdk.Dec {
	if isExpedited {
		return sdk.MustNewDecFromStr(tp.ExpeditedThreshold)
	}
	return sdk.MustNewDecFromStr(tp.Threshold)
}

// GetTallyQuorum returns quorum based on the value isExpedited
func (tp TallyParams) GetTallyQuorum(isExpedited bool) sdk.Dec {
	if isExpedited {
		return sdk.MustNewDecFromStr(tp.ExpeditedQuorum)
	}
	return sdk.MustNewDecFromStr(tp.Quorum)
}

func validateTallyParams(i interface{}) error {
	v, ok := i.(TallyParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	quorum, err := sdk.NewDecFromStr(v.Quorum)
	if err != nil {
		return fmt.Errorf("invalid quorum string: %w", err)
	}
	if quorum.IsNegative() {
		return fmt.Errorf("quorom cannot be negative: %s", quorum)
	}
	if quorum.GT(sdk.OneDec()) {
		return fmt.Errorf("quorom too large: %s", v)
	}

	expeditedQuorum, err := sdk.NewDecFromStr(v.ExpeditedQuorum)
	if err != nil {
		return fmt.Errorf("invalid expedited quorum string: %w", err)
	}
	if expeditedQuorum.IsNegative() {
		return fmt.Errorf("expedited quorom cannot be negative: %s", v.ExpeditedQuorum)
	}
	if expeditedQuorum.GT(sdk.OneDec()) {
		return fmt.Errorf("expedited quorom too large: %s", v.ExpeditedQuorum)
	}
	if expeditedQuorum.LTE(quorum) {
		return fmt.Errorf("expedited quorum %s, must be greater than the regular quorum %s", v.ExpeditedQuorum, v.Quorum)
	}

	threshold, err := sdk.NewDecFromStr(v.Threshold)
	if err != nil {
		return fmt.Errorf("invalid threshold string: %w", err)
	}
	if !threshold.IsPositive() {
		return fmt.Errorf("vote threshold must be positive: %s", threshold)
	}
	if threshold.GT(sdk.OneDec()) {
		return fmt.Errorf("vote threshold too large: %s", v)
	}

	expeditedThreshold, err := sdk.NewDecFromStr(v.ExpeditedThreshold)
	if err != nil {
		return fmt.Errorf("invalid expedited threshold string: %w", err)
	}
	if !expeditedThreshold.IsPositive() {
		return fmt.Errorf("expedited vote threshold must be positive: %s", expeditedThreshold)
	}
	if expeditedThreshold.GT(sdk.OneDec()) {
		return fmt.Errorf("expedited vote threshold too large: %s", v)
	}

	vetoThreshold, err := sdk.NewDecFromStr(v.VetoThreshold)
	if err != nil {
		return fmt.Errorf("invalid vetoThreshold string: %w", err)
	}
	if !vetoThreshold.IsPositive() {
		return fmt.Errorf("veto threshold must be positive: %s", vetoThreshold)
	}
	if vetoThreshold.GT(sdk.OneDec()) {
		return fmt.Errorf("veto threshold too large: %s", v)
	}

	return nil
}

// NewVotingParams creates a new VotingParams object
func NewVotingParams(votingPeriod, expeditedVotingPeriod time.Duration) VotingParams {
	return VotingParams{
		VotingPeriod:          &votingPeriod,
		ExpeditedVotingPeriod: &expeditedVotingPeriod,
	}
}

// DefaultVotingParams default parameters for voting
func DefaultVotingParams() VotingParams {
	return NewVotingParams(DefaultPeriod, DefaultExpeditedPeriod)
}

// Equal checks equality of TallyParams
func (vp VotingParams) Equal(other VotingParams) bool {
	return vp.VotingPeriod == other.VotingPeriod
}

func (vp VotingParams) GetVotingPeriodWithExpedited(isExpedited bool) *time.Duration {
	if isExpedited {
		return vp.ExpeditedVotingPeriod
	}

	return vp.VotingPeriod
}

func validateVotingParams(i interface{}) error {
	v, ok := i.(VotingParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.VotingPeriod == nil {
		return errors.New("voting period must not be nil")
	}

	if v.ExpeditedVotingPeriod == nil {
		return errors.New("expedited voting period must not be nil")
	}

	if v.VotingPeriod.Seconds() <= 0 {
		return fmt.Errorf("voting period must be positive: %s", v.VotingPeriod)
	}

	if v.ExpeditedVotingPeriod.Seconds() <= 0 {
		return fmt.Errorf("expedited voting period must be positive: %s", v.ExpeditedVotingPeriod)
	}

	return nil
}

// Params returns all of the governance params
type Params struct {
	VotingParams  VotingParams  `json:"voting_params" yaml:"voting_params"`
	TallyParams   TallyParams   `json:"tally_params" yaml:"tally_params"`
	DepositParams DepositParams `json:"deposit_params" yaml:"deposit_params"`
}

func (gp Params) String() string {
	return gp.VotingParams.String() + "\n" +
		gp.TallyParams.String() + "\n" + gp.DepositParams.String()
}

// NewParams creates a new gov Params instance
func NewParams(vp VotingParams, tp TallyParams, dp DepositParams) Params {
	return Params{
		VotingParams:  vp,
		DepositParams: dp,
		TallyParams:   tp,
	}
}

// DefaultParams default governance params
func DefaultParams() Params {
	return NewParams(DefaultVotingParams(), DefaultTallyParams(), DefaultDepositParams())
}
