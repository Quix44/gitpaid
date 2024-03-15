// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/access/Ownable.sol";

interface IERC20 {
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);
    function transfer(
        address recipient,
        uint256 amount
    ) external returns (bool);
}

contract GitPaid is Ownable(msg.sender) {
    error Unauthorized();
    error InsufficientFunds(uint requested, uint available);
    error TransferFailed();
    error LockTimeNotMet(uint requestedTime, uint currentTime);

    struct Deposit {
        uint256 amount;
        uint256 time;
        address depositor;
    }

    address public updaterAddress;
    uint256 public lockDuration = 1 days;
    uint256 public totalPaid;
    uint256 public totalFunded;

    mapping(address => mapping(string => Deposit)) public repositoryMap;

    event Funded(
        address token,
        string repository,
        address user,
        uint256 amount
    );
    event Withdrawn(
        address token,
        string repository,
        address user,
        uint256 amount
    );
    event PaymentMade(
        address token,
        string repository,
        address to,
        uint256 amount
    );

    constructor(address _updaterAddress) {
        updaterAddress = _updaterAddress;
    }

    modifier onlyUpdater() {
        if (msg.sender != updaterAddress) revert Unauthorized();
        _;
    }

    function fund(
        address tokenAddress,
        string memory repository,
        uint256 amount
    ) external {
        IERC20(tokenAddress).transferFrom(msg.sender, address(this), amount);
        Deposit storage deposit = repositoryMap[tokenAddress][repository];
        deposit.amount += amount;
        deposit.time = block.timestamp;
        deposit.depositor = msg.sender;
        totalFunded += amount;

        emit Funded(tokenAddress, repository, msg.sender, amount);
    }

    function withdraw(
        address tokenAddress,
        string memory repository,
        uint256 amount
    ) external {
        Deposit storage deposit = repositoryMap[tokenAddress][repository];

        if (msg.sender != deposit.depositor) revert Unauthorized();
        if (block.timestamp < deposit.time + lockDuration)
            revert LockTimeNotMet(deposit.time + lockDuration, block.timestamp);
        if (amount > deposit.amount)
            revert InsufficientFunds(amount, deposit.amount);

        bool success = IERC20(tokenAddress).transfer(msg.sender, amount);
        if (!success) revert TransferFailed();

        deposit.amount -= amount;
        emit Withdrawn(tokenAddress, repository, msg.sender, amount);
    }

    function pay(
        address tokenAddress,
        string memory repository,
        address to,
        uint256 amount
    ) external onlyUpdater {
        Deposit storage deposit = repositoryMap[tokenAddress][repository];

        if (amount > deposit.amount)
            revert InsufficientFunds(amount, deposit.amount);

        bool success = IERC20(tokenAddress).transfer(to, amount);
        if (!success) revert TransferFailed();

        deposit.amount -= amount;
        totalPaid += amount;
        emit PaymentMade(tokenAddress, repository, to, amount);
    }

    function setUpdaterAddress(address newUpdater) external onlyOwner {
        updaterAddress = newUpdater;
    }
}
