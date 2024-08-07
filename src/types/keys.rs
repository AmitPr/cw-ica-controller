//! # Keys
//!
//! Contains key constants definitions for the contract such as version info for migrations.

/// `CONTRACT_NAME` is the name of the contract recorded with [`cw2`]
pub const CONTRACT_NAME: &str = "crates.io:cw-ica-controller";
/// `CONTRACT_VERSION` is the version of the cargo package.
/// This is also the version of the contract recorded in [`cw2`]
pub const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

/// This module contains [`cosmwasm_std::SubMsg`] reply ids.
pub mod reply_ids {
    /// `SEND_PACKET` is the reply id for a packet sent using
    /// [`crate::types::msg::ExecuteMsg::SendQueryMsgs`]
    pub const SEND_PACKET: u64 = 0;
    /// `SEND_QUERY_PACKET` is the reply id for a packet sent using
    /// [`crate::types::msg::ExecuteMsg::SendQueryMsgs`] with the `query` feature enabled.
    #[cfg(feature = "query")]
    pub const SEND_QUERY_PACKET: u64 = 1;
}
