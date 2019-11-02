package coinswap

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	keeper "github.com/PirosB3/coinswap/incubator/coinswap/internal/keeper"
	asd "github.com/PirosB3/coinswap/incubator/coinswap/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic app module basics object
type AppModuleBasic struct{}

// Name defines module name
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec registers module codec
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	asd.RegisterCodec(cdc)
}

// DefaultGenesis default genesis state
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return asd.ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// ValidateGenesis module validate genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := asd.ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// RegisterRESTRoutes registers rest routes
func (AppModuleBasic) RegisterRESTRoutes(_ context.CLIContext, _ *mux.Router) {

}

// GetTxCmd returns the root tx command of this module
func (AppModuleBasic) GetTxCmd(_ *codec.Codec) *cobra.Command { return nil }

// GetQueryCmd returns the root query command of this module
func (AppModuleBasic) GetQueryCmd(_ *codec.Codec) *cobra.Command { return nil }

// AppModule coinswap app module
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper) module.AppModule {
	return module.NewGenesisOnlyAppModule(AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	})
}

// Route module message route name
func (AppModule) Route() string {
	return asd.RouterKey
}

// NewHandler module handler
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// QuerierRoute module querier route name
func (AppModule) QuerierRoute() string {
	return asd.QuerierRoute
}

// NewQuerierHandler module querier
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

// InitGenesis coinswap module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	asd.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return asd.ModuleCdc.MustMarshalJSON(gs)
}

// BeginBlock module begin-block
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock module end-block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
