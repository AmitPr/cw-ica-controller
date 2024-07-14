/* Code generated by github.com/srdtrk/go-codegen, DO NOT EDIT. */
package cwicacontroller

// The message to instantiate the ICA controller contract.
type InstantiateMsg struct {
	// The options to initialize the IBC channel upon contract instantiation.
	ChannelOpenInitOptions ChannelOpenInitOptions `json:"channel_open_init_options"`
	// The address of the owner of the ICA application. If not specified, the sender is the owner.
	Owner *string `json:"owner,omitempty"`
	// The contract address that the channel and packet lifecycle callbacks are sent to. If not specified, then no callbacks are sent.
	SendCallbacksTo *string `json:"send_callbacks_to,omitempty"`
}

// The messages to execute the ICA controller contract.
type ExecuteMsg struct {
	// `CreateChannel` makes the contract submit a stargate MsgChannelOpenInit to the chain. This is a wrapper around [`options::ChannelOpenInitOptions`] and thus requires the same fields. If not specified, then the options specified in the contract instantiation are used.
	CreateChannel *ExecuteMsg_CreateChannel `json:"create_channel,omitempty"`
	// `CloseChannel` closes the IBC channel.
	CloseChannel *ExecuteMsg_CloseChannel `json:"close_channel,omitempty"`
	/*
	   `SendCosmosMsgs` converts the provided array of [`CosmosMsg`] to an ICA tx and sends them to the ICA host. [`CosmosMsg::Stargate`] and [`CosmosMsg::Wasm`] are only supported if the [`TxEncoding`](crate::ibc::types::metadata::TxEncoding) is [`TxEncoding::Protobuf`](crate::ibc::types::metadata::TxEncoding).

	   **This is the recommended way to send messages to the ICA host.**
	*/
	SendCosmosMsgs *ExecuteMsg_SendCosmosMsgs `json:"send_cosmos_msgs,omitempty"`
	// `UpdateCallbackAddress` updates the contract callback address.
	UpdateCallbackAddress *ExecuteMsg_UpdateCallbackAddress `json:"update_callback_address,omitempty"`
	// Update the contract's ownership. The `action` to be provided can be either to propose transferring ownership to an account, accept a pending ownership transfer, or renounce the ownership permanently.
	UpdateOwnership *ExecuteMsg_UpdateOwnership `json:"update_ownership,omitempty"`
}

// The messages to query the ICA controller contract.
type QueryMsg struct {
	// GetChannel returns the IBC channel info.
	GetChannel *QueryMsg_GetChannel `json:"get_channel,omitempty"`
	// GetContractState returns the contact's state.
	GetContractState *QueryMsg_GetContractState `json:"get_contract_state,omitempty"`
	// Query the contract's ownership information
	Ownership *QueryMsg_Ownership `json:"ownership,omitempty"`
}

// IbcChannel defines all information on a channel. This is generally used in the hand-shake process, but can be queried directly.
type IbcChannel struct {
	// The connection upon which this channel was created. If this is a multi-hop channel, we only expose the first hop.
	ConnectionId string `json:"connection_id"`
	CounterpartyEndpoint IbcEndpoint `json:"counterparty_endpoint"`
	Endpoint IbcEndpoint `json:"endpoint"`
	Order IbcOrder `json:"order"`
	// Note: in ibcv3 this may be "", in the IbcOpenChannel handshake messages
	Version string `json:"version"`
}

// `TxEncoding` is the encoding of the transactions sent to the ICA host.
type TxEncoding string

const (
	// `Protobuf` is the protobuf serialization of the CosmosSDK's Any.
	TxEncoding_Proto3 TxEncoding = "proto3"
	// `Proto3Json` is the json serialization of the CosmosSDK's Any.
	TxEncoding_Proto3Json TxEncoding = "proto3json"
)

type Coin struct {
	Amount Uint128 `json:"amount"`
	Denom string `json:"denom"`
}

// Expiration represents a point in time when some event happens. It can compare with a BlockInfo and will return is_expired() == true once the condition is hit (and for every block in the future)
type Expiration struct {
	// AtHeight will expire when `env.block.height` >= height
	AtHeight *Expiration_AtHeight `json:"at_height,omitempty"`
	// AtTime will expire when `env.block.time` >= time
	AtTime *Expiration_AtTime `json:"at_time,omitempty"`
	// Never will never expire. Used to express the empty variant
	Never *Expiration_Never `json:"never,omitempty"`
}

/*
The message types of the wasm module.

See https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto
*/
type WasmMsg struct {
	/*
	   Dispatches a call to another contract at a known address (with known ABI).

	   This is translated to a [MsgExecuteContract](https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto#L68-L78). `sender` is automatically filled with the current contract's address.
	*/
	Execute *WasmMsg_Execute `json:"execute,omitempty"`
	/*
	   Instantiates a new contracts from previously uploaded Wasm code.

	   The contract address is non-predictable. But it is guaranteed that when emitting the same Instantiate message multiple times, multiple instances on different addresses will be generated. See also Instantiate2.

	   This is translated to a [MsgInstantiateContract](https://github.com/CosmWasm/wasmd/blob/v0.29.2/proto/cosmwasm/wasm/v1/tx.proto#L53-L71). `sender` is automatically filled with the current contract's address.
	*/
	Instantiate *WasmMsg_Instantiate `json:"instantiate,omitempty"`
	/*
	   Instantiates a new contracts from previously uploaded Wasm code using a predictable address derivation algorithm implemented in [`cosmwasm_std::instantiate2_address`].

	   This is translated to a [MsgInstantiateContract2](https://github.com/CosmWasm/wasmd/blob/v0.29.2/proto/cosmwasm/wasm/v1/tx.proto#L73-L96). `sender` is automatically filled with the current contract's address. `fix_msg` is automatically set to false.
	*/
	Instantiate2 *WasmMsg_Instantiate2 `json:"instantiate2,omitempty"`
	/*
	   Migrates a given contracts to use new wasm code. Passes a MigrateMsg to allow us to customize behavior.

	   Only the contract admin (as defined in wasmd), if any, is able to make this call.

	   This is translated to a [MsgMigrateContract](https://github.com/CosmWasm/wasmd/blob/v0.14.0/x/wasm/internal/types/tx.proto#L86-L96). `sender` is automatically filled with the current contract's address.
	*/
	Migrate *WasmMsg_Migrate `json:"migrate,omitempty"`
	// Sets a new admin (for migrate) on the given contract. Fails if this contract is not currently admin of the target contract.
	UpdateAdmin *WasmMsg_UpdateAdmin `json:"update_admin,omitempty"`
	// Clears the admin on the given contract, so no more migration possible. Fails if this contract is not currently admin of the target contract.
	ClearAdmin *WasmMsg_ClearAdmin `json:"clear_admin,omitempty"`
}

type WasmQuery struct {
	// this queries the public API of another contract at a known address (with known ABI) Return value is whatever the contract returns (caller should know), wrapped in a ContractResult that is JSON encoded.
	Smart *WasmQuery_Smart `json:"smart,omitempty"`
	// this queries the raw kv-store of the contract. returns the raw, unparsed data stored at that key, which may be an empty vector if not present
	Raw *WasmQuery_Raw `json:"raw,omitempty"`
	// Returns a [`ContractInfoResponse`] with metadata on the contract from the runtime
	ContractInfo *WasmQuery_ContractInfo `json:"contract_info,omitempty"`
	// Returns a [`CodeInfoResponse`] with metadata of the code
	CodeInfo *WasmQuery_CodeInfo `json:"code_info,omitempty"`
}

type QueryMsg_Ownership struct{}

type ExecuteMsg_SendCosmosMsgs struct {
	// The stargate messages to convert and send to the ICA host.
	Messages []CosmosMsg_for_Empty `json:"messages"`
	// Optional memo to include in the ibc packet.
	PacketMemo *string `json:"packet_memo,omitempty"`
	// The stargate queries to convert and send to the ICA host. The queries are executed after the messages.
	Queries []QueryRequest_for_Empty `json:"queries"`
	// Optional timeout in seconds to include with the ibc packet. If not specified, the [default timeout](crate::ibc::types::packet::DEFAULT_TIMEOUT_SECONDS) is used.
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

type VoteOption string

const (
	VoteOption_Yes        VoteOption = "yes"
	VoteOption_No         VoteOption = "no"
	VoteOption_Abstain    VoteOption = "abstain"
	VoteOption_NoWithVeto VoteOption = "no_with_veto"
)

/*
A fixed-point decimal value with 18 fractional digits, i.e. Decimal(1_000_000_000_000_000_000) == 1.0

The greatest possible value that can be represented is 340282366920938463463.374607431768211455 (which is (2^128 - 1) / 10^18)
*/
type Decimal string

type CosmosMsg_for_Empty struct {
	Bank *CosmosMsg_for_Empty_Bank `json:"bank,omitempty"`
	Custom *CosmosMsg_for_Empty_Custom `json:"custom,omitempty"`
	Staking *CosmosMsg_for_Empty_Staking `json:"staking,omitempty"`
	Distribution *CosmosMsg_for_Empty_Distribution `json:"distribution,omitempty"`
	// A Stargate message encoded the same way as a protobuf [Any](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/any.proto). This is the same structure as messages in `TxBody` from [ADR-020](https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-020-protobuf-transaction-encoding.md)
	Stargate *CosmosMsg_for_Empty_Stargate `json:"stargate,omitempty"`
	Ibc *CosmosMsg_for_Empty_Ibc `json:"ibc,omitempty"`
	Wasm *CosmosMsg_for_Empty_Wasm `json:"wasm,omitempty"`
	Gov *CosmosMsg_for_Empty_Gov `json:"gov,omitempty"`
}

/*
The message types of the staking module.

See https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/staking/v1beta1/tx.proto
*/
type StakingMsg struct {
	// This is translated to a [MsgDelegate](https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/staking/v1beta1/tx.proto#L81-L90). `delegator_address` is automatically filled with the current contract's address.
	Delegate *StakingMsg_Delegate `json:"delegate,omitempty"`
	// This is translated to a [MsgUndelegate](https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/staking/v1beta1/tx.proto#L112-L121). `delegator_address` is automatically filled with the current contract's address.
	Undelegate *StakingMsg_Undelegate `json:"undelegate,omitempty"`
	// This is translated to a [MsgBeginRedelegate](https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/staking/v1beta1/tx.proto#L95-L105). `delegator_address` is automatically filled with the current contract's address.
	Redelegate *StakingMsg_Redelegate `json:"redelegate,omitempty"`
}

// The options needed to initialize the IBC channel.
type ChannelOpenInitOptions struct {
	// The counterparty connection id on the counterparty chain.
	CounterpartyConnectionId string `json:"counterparty_connection_id"`
	// The counterparty port id. If not specified, [`crate::ibc::types::keys::HOST_PORT_ID`] is used. Currently, this contract only supports the host port.
	CounterpartyPortId *string `json:"counterparty_port_id,omitempty"`
	// The order of the channel. If not specified, [`IbcOrder::Ordered`] is used. [`IbcOrder::Unordered`] is only supported if the counterparty chain is using `ibc-go` v8.1.0 or later.
	ChannelOrdering *IbcOrder `json:"channel_ordering,omitempty"`
	// The connection id on this chain.
	ConnectionId string `json:"connection_id"`
}
type ExecuteMsg_UpdateOwnership Action

/*
The message types of the distribution module.

See https://github.com/cosmos/cosmos-sdk/blob/v0.42.4/proto/cosmos/distribution/v1beta1/tx.proto
*/
type DistributionMsg struct {
	// This is translated to a [MsgSetWithdrawAddress](https://github.com/cosmos/cosmos-sdk/blob/v0.42.4/proto/cosmos/distribution/v1beta1/tx.proto#L29-L37). `delegator_address` is automatically filled with the current contract's address.
	SetWithdrawAddress *DistributionMsg_SetWithdrawAddress `json:"set_withdraw_address,omitempty"`
	// This is translated to a [[MsgWithdrawDelegatorReward](https://github.com/cosmos/cosmos-sdk/blob/v0.42.4/proto/cosmos/distribution/v1beta1/tx.proto#L42-L50). `delegator_address` is automatically filled with the current contract's address.
	WithdrawDelegatorReward *DistributionMsg_WithdrawDelegatorReward `json:"withdraw_delegator_reward,omitempty"`
	// This is translated to a [[MsgFundCommunityPool](https://github.com/cosmos/cosmos-sdk/blob/v0.42.4/proto/cosmos/distribution/v1beta1/tx.proto#LL69C1-L76C2). `depositor` is automatically filled with the current contract's address.
	FundCommunityPool *DistributionMsg_FundCommunityPool `json:"fund_community_pool,omitempty"`
}

/*
A human readable address.

In Cosmos, this is typically bech32 encoded. But for multi-chain smart contracts no assumptions should be made other than being UTF-8 encoded and of reasonable length.

This type represents a validated address. It can be created in the following ways 1. Use `Addr::unchecked(input)` 2. Use `let checked: Addr = deps.api.addr_validate(input)?` 3. Use `let checked: Addr = deps.api.addr_humanize(canonical_addr)?` 4. Deserialize from JSON. This must only be done from JSON that was validated before such as a contract's state. `Addr` must not be used in messages sent by the user because this would result in unvalidated instances.

This type is immutable. If you really need to mutate it (Really? Are you sure?), create a mutable copy using `let mut mutable = Addr::to_string()` and operate on that `String` instance.
*/
type Addr string

// The contract's ownership info
type Ownership_for_String struct {
	// The contract's current owner. `None` if the ownership has been renounced.
	Owner *string `json:"owner,omitempty"`
	// The deadline for the pending owner to accept the ownership. `None` if there isn't a pending ownership transfer, or if a transfer exists and it doesn't have a deadline.
	PendingExpiry *Expiration `json:"pending_expiry,omitempty"`
	// The account who has been proposed to take over the ownership. `None` if there isn't a pending ownership transfer.
	PendingOwner *string `json:"pending_owner,omitempty"`
}

// IbcOrder defines if a channel is ORDERED or UNORDERED Values come from https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/ibc/core/channel/v1/channel.proto#L69-L80 Naming comes from the protobuf files and go translations.
type IbcOrder string

const (
	IbcOrder_OrderUnordered IbcOrder = "ORDER_UNORDERED"
	IbcOrder_OrderOrdered   IbcOrder = "ORDER_ORDERED"
)

/*
Binary is a wrapper around Vec<u8> to add base64 de/serialization with serde. It also adds some helper methods to help encode inline.

This is only needed as serde-json-{core,wasm} has a horrible encoding for Vec<u8>. See also <https://github.com/CosmWasm/cosmwasm/blob/main/docs/MESSAGE_TYPES.md>.
*/
type Binary string

type DistributionQuery struct {
	// See <https://github.com/cosmos/cosmos-sdk/blob/c74e2887b0b73e81d48c2f33e6b1020090089ee0/proto/cosmos/distribution/v1beta1/query.proto#L222-L230>
	DelegatorWithdrawAddress *DistributionQuery_DelegatorWithdrawAddress `json:"delegator_withdraw_address,omitempty"`
	// See <https://github.com/cosmos/cosmos-sdk/blob/c74e2887b0b73e81d48c2f33e6b1020090089ee0/proto/cosmos/distribution/v1beta1/query.proto#L157-L167>
	DelegationRewards *DistributionQuery_DelegationRewards `json:"delegation_rewards,omitempty"`
	// See <https://github.com/cosmos/cosmos-sdk/blob/c74e2887b0b73e81d48c2f33e6b1020090089ee0/proto/cosmos/distribution/v1beta1/query.proto#L180-L187>
	DelegationTotalRewards *DistributionQuery_DelegationTotalRewards `json:"delegation_total_rewards,omitempty"`
	// See <https://github.com/cosmos/cosmos-sdk/blob/b0acf60e6c39f7ab023841841fc0b751a12c13ff/proto/cosmos/distribution/v1beta1/query.proto#L202-L210>
	DelegatorValidators *DistributionQuery_DelegatorValidators `json:"delegator_validators,omitempty"`
}

type IbcEndpoint struct {
	ChannelId string `json:"channel_id"`
	PortId string `json:"port_id"`
}

// IcaInfo is the ICA address and channel ID.
type IcaInfo struct {
	ChannelId string `json:"channel_id"`
	Encoding TxEncoding `json:"encoding"`
	IcaAddress string `json:"ica_address"`
}

/*
A thin wrapper around u128 that is using strings for JSON encoding/decoding, such that the full u128 range can be used for clients that convert JSON numbers to floats, like JavaScript and jq.

# Examples

Use `from` to create instances of this and `u128` to get the value out:

``` # use cosmwasm_std::Uint128; let a = Uint128::from(123u128); assert_eq!(a.u128(), 123);

let b = Uint128::from(42u64); assert_eq!(b.u128(), 42);

let c = Uint128::from(70u32); assert_eq!(c.u128(), 70); ```
*/
type Uint128 string

/*
An empty struct that serves as a placeholder in different places, such as contracts that don't set a custom message.

It is designed to be expressable in correct JSON and JSON Schema but contains no meaningful data. Previously we used enums without cases, but those cannot represented as valid JSON Schema (https://github.com/CosmWasm/cosmwasm/issues/451)
*/
type Empty struct{}

// Status is the status of an IBC channel.
type ChannelStatus string

const (
	// Uninitialized is the default state of the channel.
	ChannelStatus_StateUninitializedUnspecified ChannelStatus = "STATE_UNINITIALIZED_UNSPECIFIED"
	// Init is the state of the channel when it is created.
	ChannelStatus_StateInit ChannelStatus = "STATE_INIT"
	// TryOpen is the state of the channel when it is trying to open.
	ChannelStatus_StateTryopen ChannelStatus = "STATE_TRYOPEN"
	// Open is the state of the channel when it is open.
	ChannelStatus_StateOpen ChannelStatus = "STATE_OPEN"
	// Closed is the state of the channel when it is closed.
	ChannelStatus_StateClosed ChannelStatus = "STATE_CLOSED"
	// The channel has just accepted the upgrade handshake attempt and is flushing in-flight packets. Added in `ibc-go` v8.1.0.
	ChannelStatus_StateFlushing ChannelStatus = "STATE_FLUSHING"
	// The channel has just completed flushing any in-flight packets. Added in `ibc-go` v8.1.0.
	ChannelStatus_StateFlushcomplete ChannelStatus = "STATE_FLUSHCOMPLETE"
)

type ExecuteMsg_CreateChannel struct {
	// The options to initialize the IBC channel. If not specified, the options specified in the last channel creation are used. Must be `None` if the sender is not the owner.
	ChannelOpenInitOptions *ChannelOpenInitOptions `json:"channel_open_init_options,omitempty"`
}

type ExecuteMsg_CloseChannel struct{}

type BankQuery struct {
	// This calls into the native bank module for querying the total supply of one denomination. It does the same as the SupplyOf call in Cosmos SDK's RPC API. Return value is of type SupplyResponse.
	Supply *BankQuery_Supply `json:"supply,omitempty"`
	// This calls into the native bank module for one denomination Return value is BalanceResponse
	Balance *BankQuery_Balance `json:"balance,omitempty"`
	// This calls into the native bank module for all denominations. Note that this may be much more expensive than Balance and should be avoided if possible. Return value is AllBalanceResponse.
	AllBalances *BankQuery_AllBalances `json:"all_balances,omitempty"`
	// This calls into the native bank module for querying metadata for a specific bank token. Return value is DenomMetadataResponse
	DenomMetadata *BankQuery_DenomMetadata `json:"denom_metadata,omitempty"`
	// This calls into the native bank module for querying metadata for all bank tokens that have a metadata entry. Return value is AllDenomMetadataResponse
	AllDenomMetadata *BankQuery_AllDenomMetadata `json:"all_denom_metadata,omitempty"`
}

type StakingQuery struct {
	// Returns the denomination that can be bonded (if there are multiple native tokens on the chain)
	BondedDenom *StakingQuery_BondedDenom `json:"bonded_denom,omitempty"`
	// AllDelegations will return all delegations by the delegator
	AllDelegations *StakingQuery_AllDelegations `json:"all_delegations,omitempty"`
	// Delegation will return more detailed info on a particular delegation, defined by delegator/validator pair
	Delegation *StakingQuery_Delegation `json:"delegation,omitempty"`
	/*
	   Returns all validators in the currently active validator set.

	   The query response type is `AllValidatorsResponse`.
	*/
	AllValidators *StakingQuery_AllValidators `json:"all_validators,omitempty"`
	/*
	   Returns the validator at the given address. Returns None if the validator is not part of the currently active validator set.

	   The query response type is `ValidatorResponse`.
	*/
	Validator *StakingQuery_Validator `json:"validator,omitempty"`
}

/*
A thin wrapper around u64 that is using strings for JSON encoding/decoding, such that the full u64 range can be used for clients that convert JSON numbers to floats, like JavaScript and jq.

# Examples

Use `from` to create instances of this and `u64` to get the value out:

``` # use cosmwasm_std::Uint64; let a = Uint64::from(42u64); assert_eq!(a.u64(), 42);

let b = Uint64::from(70u32); assert_eq!(b.u64(), 70); ```
*/
type Uint64 string

// State is the state of the contract.
type State struct {
	// The address of the callback contract.
	CallbackAddress *Addr `json:"callback_address,omitempty"`
	// The Interchain Account (ICA) info needed to send packets. This is set during the handshake.
	IcaInfo *IcaInfo `json:"ica_info,omitempty"`
}

type ExecuteMsg_UpdateCallbackAddress struct {
	// The new callback address. If not specified, then no callbacks are sent.
	CallbackAddress *string `json:"callback_address,omitempty"`
}

// IBCTimeoutHeight Height is a monotonically increasing data type that can be compared against another Height for the purposes of updating and freezing clients. Ordering is (revision_number, timeout_height)
type IbcTimeoutBlock struct {
	// block height after which the packet times out. the height within the given revision
	Height int `json:"height"`
	// the version that the client is currently on (e.g. after resetting the chain this could increment 1 as height drops to 0)
	Revision int `json:"revision"`
}

// In IBC each package must set at least one type of timeout: the timestamp or the block height. Using this rather complex enum instead of two timeout fields we ensure that at least one timeout is set.
type IbcTimeout struct {
	Block *IbcTimeoutBlock `json:"block,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
}

/*
The message types of the bank module.

See https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/bank/v1beta1/tx.proto
*/
type BankMsg struct {
	/*
	   Sends native tokens from the contract to the given address.

	   This is translated to a [MsgSend](https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/bank/v1beta1/tx.proto#L19-L28). `from_address` is automatically filled with the current contract's address.
	*/
	Send *BankMsg_Send `json:"send,omitempty"`
	// This will burn the given coins from the contract's account. There is no Cosmos SDK message that performs this, but it can be done by calling the bank keeper. Important if a contract controls significant token supply that must be retired.
	Burn *BankMsg_Burn `json:"burn,omitempty"`
}

type WeightedVoteOption struct {
	Option VoteOption `json:"option"`
	Weight Decimal `json:"weight"`
}

/*
This message type allows the contract interact with the [x/gov] module in order to cast votes.

## Examples

Cast a simple vote:

``` # use cosmwasm_std::{ #     HexBinary, #     Storage, Api, Querier, DepsMut, Deps, entry_point, Env, StdError, MessageInfo, #     Response, QueryResponse, # }; # type ExecuteMsg = (); use cosmwasm_std::{GovMsg, VoteOption};

#[entry_point] pub fn execute( deps: DepsMut, env: Env, info: MessageInfo, msg: ExecuteMsg, ) -> Result<Response, StdError> { // ... Ok(Response::new().add_message(GovMsg::Vote { proposal_id: 4, vote: VoteOption::Yes, })) } ```

Cast a weighted vote:

``` # use cosmwasm_std::{ #     HexBinary, #     Storage, Api, Querier, DepsMut, Deps, entry_point, Env, StdError, MessageInfo, #     Response, QueryResponse, # }; # type ExecuteMsg = (); # #[cfg(feature = "cosmwasm_1_2")] use cosmwasm_std::{Decimal, GovMsg, VoteOption, WeightedVoteOption};

# #[cfg(feature = "cosmwasm_1_2")] #[entry_point] pub fn execute( deps: DepsMut, env: Env, info: MessageInfo, msg: ExecuteMsg, ) -> Result<Response, StdError> { // ... Ok(Response::new().add_message(GovMsg::VoteWeighted { proposal_id: 4, options: vec![ WeightedVoteOption { option: VoteOption::Yes, weight: Decimal::percent(65), }, WeightedVoteOption { option: VoteOption::Abstain, weight: Decimal::percent(35), }, ], })) } ```

[x/gov]: https://github.com/cosmos/cosmos-sdk/tree/v0.45.12/x/gov
*/
type GovMsg struct {
	// This maps directly to [MsgVote](https://github.com/cosmos/cosmos-sdk/blob/v0.42.5/proto/cosmos/gov/v1beta1/tx.proto#L46-L56) in the Cosmos SDK with voter set to the contract address.
	Vote *GovMsg_Vote `json:"vote,omitempty"`
	// This maps directly to [MsgVoteWeighted](https://github.com/cosmos/cosmos-sdk/blob/v0.45.8/proto/cosmos/gov/v1beta1/tx.proto#L66-L78) in the Cosmos SDK with voter set to the contract address.
	VoteWeighted *GovMsg_VoteWeighted `json:"vote_weighted,omitempty"`
}

/*
A point in time in nanosecond precision.

This type can represent times from 1970-01-01T00:00:00Z to 2554-07-21T23:34:33Z.

## Examples

``` # use cosmwasm_std::Timestamp; let ts = Timestamp::from_nanos(1_000_000_202); assert_eq!(ts.nanos(), 1_000_000_202); assert_eq!(ts.seconds(), 1); assert_eq!(ts.subsec_nanos(), 202);

let ts = ts.plus_seconds(2); assert_eq!(ts.nanos(), 3_000_000_202); assert_eq!(ts.seconds(), 3); assert_eq!(ts.subsec_nanos(), 202); ```
*/
type Timestamp Uint64

type QueryRequest_for_Empty struct {
	Bank *QueryRequest_for_Empty_Bank `json:"bank,omitempty"`
	Custom *QueryRequest_for_Empty_Custom `json:"custom,omitempty"`
	Staking *QueryRequest_for_Empty_Staking `json:"staking,omitempty"`
	Distribution *QueryRequest_for_Empty_Distribution `json:"distribution,omitempty"`
	// A Stargate query is encoded the same way as abci_query, with path and protobuf encoded request data. The format is defined in [ADR-21](https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-021-protobuf-query-encoding.md). The response is protobuf encoded data directly without a JSON response wrapper. The caller is responsible for compiling the proper protobuf definitions for both requests and responses.
	Stargate *QueryRequest_for_Empty_Stargate `json:"stargate,omitempty"`
	Ibc *QueryRequest_for_Empty_Ibc `json:"ibc,omitempty"`
	Wasm *QueryRequest_for_Empty_Wasm `json:"wasm,omitempty"`
}

// State is the state of the IBC application's channel. This application only supports one channel.
type ChannelState struct {
	// The IBC channel, as defined by cosmwasm.
	Channel IbcChannel `json:"channel"`
	// The status of the channel.
	ChannelStatus ChannelStatus `json:"channel_status"`
}

type QueryMsg_GetContractState struct{}

// These are queries to the various IBC modules to see the state of the contract's IBC connection. These will return errors if the contract is not "ibc enabled"
type IbcQuery struct {
	/*
	   Gets the Port ID the current contract is bound to.

	   Returns a `PortIdResponse`.
	*/
	PortId *IbcQuery_PortId `json:"port_id,omitempty"`
	/*
	   Lists all channels that are bound to a given port. If `port_id` is omitted, this list all channels bound to the contract's port.

	   Returns a `ListChannelsResponse`.
	*/
	ListChannels *IbcQuery_ListChannels `json:"list_channels,omitempty"`
	/*
	   Lists all information for a (portID, channelID) pair. If port_id is omitted, it will default to the contract's own channel. (To save a PortId{} call)

	   Returns a `ChannelResponse`.
	*/
	Channel *IbcQuery_Channel `json:"channel,omitempty"`
}

// Actions that can be taken to alter the contract's ownership
type Action interface {
	Implements_Action()
}

var _ Action = (*Action_TransferOwnership)(nil)

type Action_TransferOwnership struct {
	Expiry *Expiration `json:"expiry,omitempty"`
	NewOwner string `json:"new_owner"`
}

func (*Action_TransferOwnership) Implements_Action() {}

var _ Action = (*Action_AcceptOwnership)(nil)

type Action_AcceptOwnership string

/*
Accept the pending ownership transfer.

Can only be called by the pending owner.
*/
const Action_AcceptOwnership_Value Action_AcceptOwnership = "accept_ownership"

func (*Action_AcceptOwnership) Implements_Action() {}

var _ Action = (*Action_RenounceOwnership)(nil)

type Action_RenounceOwnership string

/*
Give up the contract's ownership and the possibility of appointing a new owner.

Can only be invoked by the contract's current owner.

Any existing pending ownership transfer is canceled.
*/
const Action_RenounceOwnership_Value Action_RenounceOwnership = "renounce_ownership"

func (*Action_RenounceOwnership) Implements_Action() {}

// Simplified version of the PageRequest type for pagination from the cosmos-sdk
type PageRequest struct {
	Limit int `json:"limit"`
	Reverse bool `json:"reverse"`
	Key *Binary `json:"key,omitempty"`
}

// These are messages in the IBC lifecycle. Only usable by IBC-enabled contracts (contracts that directly speak the IBC protocol via 6 entry points)
type IbcMsg struct {
	// Sends bank tokens owned by the contract to the given address on another chain. The channel must already be established between the ibctransfer module on this chain and a matching module on the remote chain. We cannot select the port_id, this is whatever the local chain has bound the ibctransfer module to.
	Transfer *IbcMsg_Transfer `json:"transfer,omitempty"`
	// Sends an IBC packet with given data over the existing channel. Data should be encoded in a format defined by the channel version, and the module on the other side should know how to parse this.
	SendPacket *IbcMsg_SendPacket `json:"send_packet,omitempty"`
	// This will close an existing channel that is owned by this contract. Port is auto-assigned to the contract's IBC port
	CloseChannel *IbcMsg_CloseChannel `json:"close_channel,omitempty"`
}

type QueryMsg_GetChannel struct{}

type WasmMsg_ClearAdmin struct {
	ContractAddr string `json:"contract_addr"`
}
type QueryRequest_for_Empty_Distribution DistributionQuery

type DistributionQuery_DelegatorValidators struct {
	DelegatorAddress string `json:"delegator_address"`
}
type QueryRequest_for_Empty_Wasm WasmQuery

type DistributionQuery_DelegationTotalRewards struct {
	DelegatorAddress string `json:"delegator_address"`
}

type Expiration_AtHeight int

type IbcQuery_PortId struct{}
type CosmosMsg_for_Empty_Bank BankMsg

type GovMsg_Vote struct {
	ProposalId int `json:"proposal_id"`
	/*
	   The vote option.

	   This should be called "option" for consistency with Cosmos SDK. Sorry for that. See <https://github.com/CosmWasm/cosmwasm/issues/1571>.
	*/
	Vote VoteOption `json:"vote"`
}

type WasmMsg_Migrate struct {
	ContractAddr string `json:"contract_addr"`
	// msg is the json-encoded MigrateMsg struct that will be passed to the new code
	Msg Binary `json:"msg"`
	// the code_id of the new logic to place in the given contract
	NewCodeId int `json:"new_code_id"`
}
type QueryRequest_for_Empty_Custom Empty

type BankQuery_DenomMetadata struct {
	Denom string `json:"denom"`
}

type WasmMsg_Execute struct {
	ContractAddr string `json:"contract_addr"`
	Funds []Coin `json:"funds"`
	// msg is the json-encoded ExecuteMsg struct (as raw Binary)
	Msg Binary `json:"msg"`
}
type CosmosMsg_for_Empty_Gov GovMsg

type StakingQuery_Validator struct {
	// The validator's address (e.g. (e.g. cosmosvaloper1...))
	Address string `json:"address"`
}
type CosmosMsg_for_Empty_Ibc IbcMsg

type IbcMsg_Transfer struct {
	// packet data only supports one coin https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/ibc/applications/transfer/v1/transfer.proto#L11-L20
	Amount Coin `json:"amount"`
	// existing channel to send the tokens over
	ChannelId string `json:"channel_id"`
	// when packet times out, measured on remote chain
	Timeout IbcTimeout `json:"timeout"`
	// address on the remote chain to receive these tokens
	ToAddress string `json:"to_address"`
}

type IbcQuery_Channel struct {
	ChannelId string `json:"channel_id"`
	PortId *string `json:"port_id,omitempty"`
}
type CosmosMsg_for_Empty_Distribution DistributionMsg
type CosmosMsg_for_Empty_Staking StakingMsg

type QueryRequest_for_Empty_Stargate struct {
	// this is the expected protobuf message type (not any), binary encoded
	Data Binary `json:"data"`
	// this is the fully qualified service path used for routing, eg. custom/cosmos_sdk.x.bank.v1.Query/QueryBalance
	Path string `json:"path"`
}

type WasmQuery_Smart struct {
	ContractAddr string `json:"contract_addr"`
	// msg is the json-encoded QueryMsg struct
	Msg Binary `json:"msg"`
}

type GovMsg_VoteWeighted struct {
	Options []WeightedVoteOption `json:"options"`
	ProposalId int `json:"proposal_id"`
}

type BankMsg_Burn struct {
	Amount []Coin `json:"amount"`
}
type QueryRequest_for_Empty_Staking StakingQuery

type CosmosMsg_for_Empty_Stargate struct {
	TypeUrl string `json:"type_url"`
	Value Binary `json:"value"`
}

type IbcQuery_ListChannels struct {
	PortId *string `json:"port_id,omitempty"`
}

type WasmMsg_Instantiate2 struct {
	/*
	   A human-readable label for the contract.

	   Valid values should: - not be empty - not be bigger than 128 bytes (or some chain-specific limit) - not start / end with whitespace
	*/
	Label string `json:"label"`
	// msg is the JSON-encoded InstantiateMsg struct (as raw Binary)
	Msg Binary `json:"msg"`
	Salt Binary `json:"salt"`
	Admin *string `json:"admin,omitempty"`
	CodeId int `json:"code_id"`
	Funds []Coin `json:"funds"`
}

type BankQuery_AllDenomMetadata struct {
	Pagination *PageRequest `json:"pagination,omitempty"`
}

type BankQuery_Supply struct {
	Denom string `json:"denom"`
}

type Expiration_Never struct{}

type BankQuery_Balance struct {
	Address string `json:"address"`
	Denom string `json:"denom"`
}

type WasmMsg_Instantiate struct {
	CodeId int `json:"code_id"`
	Funds []Coin `json:"funds"`
	/*
	   A human-readable label for the contract.

	   Valid values should: - not be empty - not be bigger than 128 bytes (or some chain-specific limit) - not start / end with whitespace
	*/
	Label string `json:"label"`
	// msg is the JSON-encoded InstantiateMsg struct (as raw Binary)
	Msg Binary `json:"msg"`
	Admin *string `json:"admin,omitempty"`
}
type Expiration_AtTime Timestamp
type CosmosMsg_for_Empty_Wasm WasmMsg

type BankQuery_AllBalances struct {
	Address string `json:"address"`
}

type StakingMsg_Delegate struct {
	Validator string `json:"validator"`
	Amount Coin `json:"amount"`
}

type IbcMsg_SendPacket struct {
	ChannelId string `json:"channel_id"`
	Data Binary `json:"data"`
	// when packet times out, measured on remote chain
	Timeout IbcTimeout `json:"timeout"`
}

type StakingQuery_Delegation struct {
	Delegator string `json:"delegator"`
	Validator string `json:"validator"`
}

type IbcMsg_CloseChannel struct {
	ChannelId string `json:"channel_id"`
}

type DistributionQuery_DelegationRewards struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
}

type DistributionQuery_DelegatorWithdrawAddress struct {
	DelegatorAddress string `json:"delegator_address"`
}

type StakingQuery_AllValidators struct{}

type StakingMsg_Undelegate struct {
	Amount Coin `json:"amount"`
	Validator string `json:"validator"`
}
type CosmosMsg_for_Empty_Custom Empty

type DistributionMsg_SetWithdrawAddress struct {
	// The `withdraw_address`
	Address string `json:"address"`
}

type WasmQuery_CodeInfo struct {
	CodeId int `json:"code_id"`
}

type DistributionMsg_WithdrawDelegatorReward struct {
	// The `validator_address`
	Validator string `json:"validator"`
}

type WasmQuery_Raw struct {
	ContractAddr string `json:"contract_addr"`
	// Key is the raw key used in the contracts Storage
	Key Binary `json:"key"`
}

type DistributionMsg_FundCommunityPool struct {
	// The amount to spend
	Amount []Coin `json:"amount"`
}
type QueryRequest_for_Empty_Bank BankQuery

type BankMsg_Send struct {
	Amount []Coin `json:"amount"`
	ToAddress string `json:"to_address"`
}

type WasmQuery_ContractInfo struct {
	ContractAddr string `json:"contract_addr"`
}

type StakingMsg_Redelegate struct {
	Amount Coin `json:"amount"`
	DstValidator string `json:"dst_validator"`
	SrcValidator string `json:"src_validator"`
}

type StakingQuery_AllDelegations struct {
	Delegator string `json:"delegator"`
}

type StakingQuery_BondedDenom struct{}
type QueryRequest_for_Empty_Ibc IbcQuery

type WasmMsg_UpdateAdmin struct {
	Admin string `json:"admin"`
	ContractAddr string `json:"contract_addr"`
}