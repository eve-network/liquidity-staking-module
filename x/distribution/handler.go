package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/iqlusioninc/liquidity-staking-module/x/distribution/keeper"
	"github.com/iqlusioninc/liquidity-staking-module/x/distribution/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSetWithdrawAddress:
			res, err := msgServer.SetWithdrawAddress(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgWithdrawDelegatorReward:
			res, err := msgServer.WithdrawDelegatorReward(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgWithdrawValidatorCommission:
			res, err := msgServer.WithdrawValidatorCommission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgFundCommunityPool:
			res, err := msgServer.FundCommunityPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawTokenizeShareRecordReward:
			res, err := msgServer.WithdrawTokenizeShareRecordReward(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawAllTokenizeShareRecordReward:
			res, err := msgServer.WithdrawAllTokenizeShareRecordReward(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distribution message type: %T", msg)
		}
	}
}

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}
