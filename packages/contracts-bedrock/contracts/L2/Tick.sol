// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { Semver } from "../universal/Semver.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

interface ITick {
    function tick() external;
}

/**
 * @custom:proxied
 * @custom:predeploy 0x42000000000000000000000000000000000000A0
 * @title Tick
 * @notice The Tick predeploy ticks the chain.
 */
contract Tick is Semver, Ownable, ITick {
    /**
     * @notice Address of the special depositor account.
     */
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    /**
     * @notice Address of the tick contract to be called.
     */
    address public target;

    /**
     * @param _owner Address that will initially own this contract.
     * @custom:semver 0.0.1
     */
    constructor(address _owner) Ownable() Semver(0, 0, 1) {
        transferOwnership(_owner);
    }

    /**
     * @notice Allows the owner to modify the target address.
     * @param _target New target address.
     */
    function setTarget(address _target) public onlyOwner {
        target = _target;
    }

    /**
     * @notice Calls the tick function in the target contract.
     */
    function tick() external {
        require(msg.sender == DEPOSITOR_ACCOUNT, "Tick: only the depositor account can tick");
        if (target == address(0)) {
            return;
        }
        ITick(target).tick();
    }
}
