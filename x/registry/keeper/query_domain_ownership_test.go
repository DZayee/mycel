package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDomainOwnershipQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDomainOwnership(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDomainOwnershipRequest
		response *types.QueryGetDomainOwnershipResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDomainOwnershipRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryGetDomainOwnershipResponse{DomainOwnership: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDomainOwnershipRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryGetDomainOwnershipResponse{DomainOwnership: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDomainOwnershipRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.DomainOwnership(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDomainOwnershipQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDomainOwnership(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDomainOwnershipRequest {
		return &types.QueryAllDomainOwnershipRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DomainOwnershipAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DomainOwnership), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DomainOwnership),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DomainOwnershipAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DomainOwnership), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DomainOwnership),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DomainOwnershipAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.DomainOwnership),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DomainOwnershipAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
