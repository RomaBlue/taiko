// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/

pragma solidity ^0.8.20;

// AddressDelegater lets someone use their mainnet Galxe account (the
// delegater) to set a delegate (the testnet proposer/prover they will use on L3
// to complete the Galxe tasks). This will let a user safely complete their
// testnet
// tasks, while only using their mainnet wallet to sign a transaction to this
// contract
// to set the delegate.
// The galxe API in the go eventindexer will accept the mainnet Galxe wallet in
// the route,
// and look up this contract to see if a delegate has been set. If so, this
// address
// will be used when comparing the `event.Proposer` or `event.Prover`,
// removing the need for using the mainnet private key to propose or prove.
/// @custom:security-contact hello@taiko.xyz
contract AddressDelegater {
    struct Delegate {
        bool approved;
        address delegate;
    }

    mapping(address delegater => Delegate delegate) public proverToDelegate;

    mapping(address delegater => Delegate delegate) public proposerToDelegate;

    // prevent double claim
    mapping(address delegate => bool used) public usedDelegateProverAddresses;
    mapping(address delegate => bool used) public usedDelegateProposerAddresses;

    function delegateProver(address delegate) public {
        require(!usedDelegateProverAddresses[delegate], "delegate address used");
        proverToDelegate[msg.sender] = Delegate(false, delegate);
        usedDelegateProverAddresses[delegate] = true;
    }

    function acceptProverDelegate(address delegate) public {
        require(
            proverToDelegate[delegate].delegate != address(0),
            "delegate not set"
        );
        proverToDelegate[delegate].approved = true;
    }

    function delegateProposer(address delegate) public {
        require(
            !usedDelegateProposerAddresses[delegate], "delegate address used"
        );
        proposerToDelegate[msg.sender] = Delegate(false, delegate);
        usedDelegateProposerAddresses[delegate] = true;
    }

    function acceptProposerDelegate(address delegate) public {
        require(
            proposerToDelegate[delegate].delegate != address(0),
            "delegate not set"
        );
        proposerToDelegate[delegate].approved = true;
    }
}
