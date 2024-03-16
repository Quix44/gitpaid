import { ethers } from "hardhat";
import { GitPaid, MockERC20 } from "../typechain-types";

const fs = require("fs");

async function main() {
  const [owner, addr1, addr2] = await ethers.getSigners();

  const GitPaid = await ethers.getContractFactory("GitPaid");
  const USDC = await ethers.getContractFactory("MockERC20")

  const usdcContractAddress = fs.readFileSync("usdc.txt", "utf8");
  const gitpaidAddress = fs.readFileSync("gitpaid.txt", "utf8");

  const mockUSDC = USDC.attach(usdcContractAddress) as MockERC20
  const gitpaidContract = GitPaid.attach(gitpaidAddress) as GitPaid

  // Fund with 1000000000 USDC
  const fundAmount = ethers.parseEther("100000");

  console.log(`Approving ${gitpaidAddress} to spend 1000000000 USDC`)
  const approveTxn = await mockUSDC.connect(owner).approve(gitpaidAddress, fundAmount);
  await approveTxn.wait(1)
  console.log('Approved: ', owner)

  // Get the balance of the owner
  const balance = await mockUSDC.balanceOf(owner.address);
  console.log('Balance: ', balance.toString())

  // get the allowance of the owner
  const allowance = await mockUSDC.allowance(owner.address, gitpaidAddress);
  console.log('Allowance: ', allowance.toString())

  const txn = await gitpaidContract.connect(owner).fund(usdcContractAddress, "772761887", fundAmount);
  console.log("Funded with 1000000000 USDC");
  await txn.wait(1)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
