import { ethers } from "hardhat";
import { GitPaid, MockERC20 } from "../typechain-types";

const fs = require("fs");

async function main() {
  const [owner, addr1, addr2] = await ethers.getSigners();

  const GitPaid = await ethers.getContractFactory("GitPaid");
  const USDC = await ethers.getContractFactory("MockERC20")

  const usdcContractAddress = fs.readFileSync("usdc.txt", "utf8");
  const mockUSDC = USDC.attach(usdcContractAddress) as MockERC20
  const gitpaidContract = GitPaid.attach(usdcContractAddress) as GitPaid

  // Fund with 1000000000 USDC
  const approveTxn = await mockUSDC.approve(await gitpaidContract.getAddress(), "10000000000000000000000");
  await approveTxn.wait(1)
  const txn = await gitpaidContract.fund(usdcContractAddress, "", "1000000000000000");
  await txn.wait(1)
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
