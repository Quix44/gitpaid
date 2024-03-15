import { ethers } from "hardhat";

const fs = require("fs");

async function main() {
  const [owner, addr1, addr2] = await ethers.getSigners();

  const GitPaid = await ethers.getContractFactory("GitPaid");
  const ApeCoin = await ethers.getContractFactory("MockERC20");
  const ArbToken = await ethers.getContractFactory("MockERC20");
  const USDC = await ethers.getContractFactory("MockERC20");

  const updaterAddress = "0xb6D9f614907368499bAF7b288b54B839fC891660"

  const gitpaid = await GitPaid.deploy(updaterAddress);
  const apeCoin = await ApeCoin.deploy("ApeCoin", "APE", ethers.parseEther("1000000000"))
  const arbToken = await ArbToken.deploy("Arbitrum Token", "ARB", ethers.parseEther("1000000000"))
  const usdc = await USDC.deploy("USDC", "USDC", ethers.parseEther("1000000000000"))

  fs.writeFileSync("gitpaid.txt", await gitpaid.getAddress());
  fs.writeFileSync("apecoin.txt", await apeCoin.getAddress());
  fs.writeFileSync("arbtoken.txt", await arbToken.getAddress());
  fs.writeFileSync("usdc.txt", await usdc.getAddress());

  console.log("GitPaid deployed to:", await gitpaid.getAddress());
  console.log("ApeCoin deployed to:", await apeCoin.getAddress());
  console.log("ArbToken deployed to:", await arbToken.getAddress());
  console.log("USDC deployed to:", await usdc.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
