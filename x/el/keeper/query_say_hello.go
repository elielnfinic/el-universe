package keeper

import (
	"context"
	"fmt"

	"el/x/el/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	fmt.Println("Handling SayHello query")

	return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Jambo %s!", req.Name)}, nil
}
