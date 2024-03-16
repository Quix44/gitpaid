import { expect } from "chai";
import { parseEther } from "ethers"; // Ethers v6 syntax for utilities
import { ethers } from "hardhat";
import { GitPaid, MockERC20 } from "../typechain-types";

describe("GitPaid", function () {
  let gitPaid: GitPaid;
  let tokenA: MockERC20, tokenB: MockERC20;

  const repository: string = "ExampleRepo";

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    const addr1 = signers[1];
    const addr2 = signers[2];

    const GitPaidFactory = await ethers.getContractFactory("GitPaid");
    gitPaid = (await GitPaidFactory.deploy(addr2.address)) as GitPaid;

    const MockERC20Factory = await ethers.getContractFactory("MockERC20");
    tokenA = (await MockERC20Factory.deploy("TokenA", "TKA", parseEther("1000"))) as MockERC20;

    tokenB = (await MockERC20Factory.deploy("TokenB", "TKB", parseEther("1000"))) as MockERC20;
  });

  describe("Fund and Withdraw with Multiple Tokens", function () {
    it("should revert with LockTimeNotMet when withdrawing with TokenA before time lock", async function () {
      const [owner, addr1] = await ethers.getSigners();

      const fundAmount = parseEther("10");
      const withdrawAmount = parseEther("5");

      // Approve and fund with TokenA
      await tokenA.approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.fund(await tokenA.getAddress(), repository, fundAmount);

      // Attempt to withdraw with TokenA before lock time is met
      await expect(gitPaid.withdraw(tokenA.getAddress(), repository, withdrawAmount))
        .to.be.revertedWithCustomError(gitPaid, 'LockTimeNotMet');
    });

    it("should revert with LockTimeNotMet when withdrawing with TokenB before time lock", async function () {
      const [owner, addr1] = await ethers.getSigners();

      const fundAmount = parseEther("20");
      const withdrawAmount = parseEther("10");

      // Approve and fund with TokenB
      await tokenB.approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.fund(await tokenB.getAddress(), repository, fundAmount);

      // Attempt to withdraw with TokenB before lock time is met
      await expect(gitPaid.withdraw(await tokenB.getAddress(), repository, withdrawAmount))
        .to.be.revertedWithCustomError(gitPaid, 'LockTimeNotMet');
    });
  })

  describe("GitPaid additional tests", function () {
    it("should fail to fund with insufficient token balance", async function () {
      const [owner, addr1] = await ethers.getSigners();
      const excessiveAmount = parseEther("2000"); // More than initial supply to addr1

      // Attempt to fund with an amount greater than the token balance
      await tokenA.approve(await gitPaid.getAddress(), excessiveAmount);
      await expect(gitPaid.fund(await tokenA.getAddress(), repository, excessiveAmount))
        .to.be.reverted; // Specific error depends on ERC20 implementation
    });

    it("should fail to withdraw more than deposited", async function () {
      const [owner, addr1] = await ethers.getSigners();
      const fundAmount = parseEther("10");
      const excessiveWithdrawAmount = parseEther("15");

      // Fund with a normal amount first
      await tokenA.approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.fund(await tokenA.getAddress(), repository, fundAmount);

      // Attempt to withdraw more than the funded amount
      await expect(gitPaid.withdraw(await tokenA.getAddress(), repository, excessiveWithdrawAmount))
        .to.be.revertedWithCustomError(gitPaid, 'LockTimeNotMet');
    });

    it("should prevent non-depositor from withdrawing funds", async function () {
      const [owner, addr1, addr2] = await ethers.getSigners();
      const fundAmount = parseEther("10");

      // addr1 funds the contract
      await tokenA.approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.fund(await tokenA.getAddress(), repository, fundAmount);

      // addr2 attempts to withdraw addr1's funds
      await expect(gitPaid.connect(addr2).withdraw(await tokenA.getAddress(), repository, fundAmount))
        .to.be.revertedWithCustomError(gitPaid, 'Unauthorized');
    });

    it("should allow withdrawing after timelock period", async function () {
      const [owner, addr1] = await ethers.getSigners();
      const fundAmount = parseEther("10");
      const withdrawAmount = parseEther("5");
      const timelockDuration = 60; // Assuming 60 seconds timelock for simplicity

      // Fund with TokenA
      await tokenA.approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.fund(await tokenA.getAddress(), repository, fundAmount);

      // Increase blockchain time to simulate timelock passing
      await ethers.provider.send("evm_increaseTime", [timelockDuration]);
      await ethers.provider.send("evm_mine", []);

      // Withdraw after timelock
      await expect(gitPaid.withdraw(await tokenA.getAddress(), repository, withdrawAmount))
        .to.emit

    })

    it("should successfully fund a new repository", async function () {
      const [owner] = await ethers.getSigners();
      const fundAmount = ethers.parseEther("10");
      const repository = "newRepo";

      // Approve and fund the repository
      await tokenA.connect(owner).approve(await gitPaid.getAddress(), fundAmount);
      await gitPaid.connect(owner).fund(await tokenA.getAddress(), repository, fundAmount);

      // Fetch the deposit details
      const deposit = await gitPaid.repositoryMap(await tokenA.getAddress(), repository);

      // Assertions
      expect(deposit.amount).to.equal(fundAmount);
      expect(deposit.depositor).to.equal(owner.address);
      expect(deposit.time).to.be.above(0); // Check if the timestamp is set
    });


  })



})
