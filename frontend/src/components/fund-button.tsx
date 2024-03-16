
import { Button } from "@/components/ui/button"
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Loader2 } from "lucide-react"
import { useState } from "react"
import { erc20ABI, useAccount, useContractRead, useContractWrite, useNetwork, usePrepareContractWrite, useWaitForTransaction } from 'wagmi'
import { DropdownMenuTokens } from "./dropdown-menu"
import { ToastAction } from "./ui/toast"
import { toast } from "./ui/use-toast"

const MAX_UINT256 = BigInt("115792089237316195423570985008687907853269984665640564039457584007913129639935")

const GIT_PAID_ABI = [{ "inputs": [{ "internalType": "address", "name": "_updaterAddress", "type": "address" }], "stateMutability": "nonpayable", "type": "constructor" }, { "inputs": [{ "internalType": "uint256", "name": "requested", "type": "uint256" }, { "internalType": "uint256", "name": "available", "type": "uint256" }], "name": "InsufficientFunds", "type": "error" }, { "inputs": [{ "internalType": "uint256", "name": "requestedTime", "type": "uint256" }, { "internalType": "uint256", "name": "currentTime", "type": "uint256" }], "name": "LockTimeNotMet", "type": "error" }, { "inputs": [{ "internalType": "address", "name": "owner", "type": "address" }], "name": "OwnableInvalidOwner", "type": "error" }, { "inputs": [{ "internalType": "address", "name": "account", "type": "address" }], "name": "OwnableUnauthorizedAccount", "type": "error" }, { "inputs": [], "name": "TransferFailed", "type": "error" }, { "inputs": [], "name": "Unauthorized", "type": "error" }, { "anonymous": false, "inputs": [{ "indexed": false, "internalType": "address", "name": "token", "type": "address" }, { "indexed": false, "internalType": "string", "name": "repository", "type": "string" }, { "indexed": false, "internalType": "address", "name": "user", "type": "address" }, { "indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "Funded", "type": "event" }, { "anonymous": false, "inputs": [{ "indexed": true, "internalType": "address", "name": "previousOwner", "type": "address" }, { "indexed": true, "internalType": "address", "name": "newOwner", "type": "address" }], "name": "OwnershipTransferred", "type": "event" }, { "anonymous": false, "inputs": [{ "indexed": false, "internalType": "address", "name": "token", "type": "address" }, { "indexed": false, "internalType": "string", "name": "repository", "type": "string" }, { "indexed": false, "internalType": "address", "name": "to", "type": "address" }, { "indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "PaymentMade", "type": "event" }, { "anonymous": false, "inputs": [{ "indexed": false, "internalType": "address", "name": "token", "type": "address" }, { "indexed": false, "internalType": "string", "name": "repository", "type": "string" }, { "indexed": false, "internalType": "address", "name": "user", "type": "address" }, { "indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "Withdrawn", "type": "event" }, { "inputs": [{ "internalType": "address", "name": "tokenAddress", "type": "address" }, { "internalType": "string", "name": "repository", "type": "string" }, { "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "fund", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [], "name": "lockDuration", "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }], "stateMutability": "view", "type": "function" }, { "inputs": [], "name": "owner", "outputs": [{ "internalType": "address", "name": "", "type": "address" }], "stateMutability": "view", "type": "function" }, { "inputs": [{ "internalType": "address", "name": "tokenAddress", "type": "address" }, { "internalType": "string", "name": "repository", "type": "string" }, { "internalType": "address", "name": "to", "type": "address" }, { "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "pay", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [], "name": "renounceOwnership", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [{ "internalType": "address", "name": "", "type": "address" }, { "internalType": "string", "name": "", "type": "string" }], "name": "repositoryMap", "outputs": [{ "internalType": "uint256", "name": "amount", "type": "uint256" }, { "internalType": "uint256", "name": "time", "type": "uint256" }, { "internalType": "address", "name": "depositor", "type": "address" }], "stateMutability": "view", "type": "function" }, { "inputs": [{ "internalType": "address", "name": "newUpdater", "type": "address" }], "name": "setUpdaterAddress", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [], "name": "totalFunded", "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }], "stateMutability": "view", "type": "function" }, { "inputs": [], "name": "totalPaid", "outputs": [{ "internalType": "uint256", "name": "", "type": "uint256" }], "stateMutability": "view", "type": "function" }, { "inputs": [{ "internalType": "address", "name": "newOwner", "type": "address" }], "name": "transferOwnership", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [], "name": "updaterAddress", "outputs": [{ "internalType": "address", "name": "", "type": "address" }], "stateMutability": "view", "type": "function" }, { "inputs": [{ "internalType": "address", "name": "tokenAddress", "type": "address" }, { "internalType": "string", "name": "repository", "type": "string" }, { "internalType": "uint256", "name": "amount", "type": "uint256" }], "name": "withdraw", "outputs": [], "stateMutability": "nonpayable", "type": "function" }]

export function FundButton({ repository, connected }: { repository: string, connected: boolean }) {
    const [amount, setAmount] = useState(1000)
    const [paymentTokenAddress, setPaymentTokenAddress] = useState<undefined | `0x${string}`>("0x1B2F2eed297d6257E9F966E3f375a4e450f4032A")
    const { isConnected, address } = useAccount()
    const { chain } = useNetwork()
    console.log(`Connected: ${isConnected}, chain: ${chain?.name}`)

    const { data: allowance, isSuccess: fetchAllowingSuccess, refetch: refetchAllowance } = useContractRead({
        address: paymentTokenAddress as `0x${string}`,
        abi: erc20ABI,
        functionName: "allowance",
        enabled: Boolean(isConnected && paymentTokenAddress),
        args: [address as `0x${string}`, process.env.NEXT_PUBLIC_CONTRACT_ADDRESS as `0x${string}`],
    });
    console.log(`Allowance: ${allowance}`)
    const amountInWei = BigInt(amount) * BigInt(10 ** 18)

    const { config: approvalConfig } = usePrepareContractWrite({
        address: paymentTokenAddress,
        abi: erc20ABI,
        functionName: "approve",
        enabled: Boolean(paymentTokenAddress && isConnected && allowance === 0n),
        args: [process.env.NEXT_PUBLIC_CONTRACT_ADDRESS as `0x${string}`, MAX_UINT256],
    });

    const {
        data: writeContractResult,
        write: approvalWrite,
    } = useContractWrite(approvalConfig);

    const { isLoading: isApproving, isSuccess: approvalSuccess } = useWaitForTransaction({
        hash: writeContractResult ? writeContractResult.hash : undefined,
        timeout: 15000,
        onSuccess: async () => {
            await refetchAllowance();
            toast({
                title: "Tokens Approved",
                description: "You can now fund your repository",
            })
        },
    });


    const { config, data: isAbleToFund } = usePrepareContractWrite({
        address: process.env.NEXT_PUBLIC_CONTRACT_ADDRESS as `0x${string}`,
        abi: GIT_PAID_ABI,
        functionName: 'fund',
        args: [paymentTokenAddress, repository, amountInWei],
        enabled: Boolean(isConnected && allowance && allowance >= amountInWei && paymentTokenAddress && repository && amountInWei),
    })

    const { write: fundRepository, data: fundRepoData } = useContractWrite(config)
    const { isSuccess: isFundSuccess, isLoading: isFundLoading } = useWaitForTransaction({
        enabled: Boolean(fundRepoData?.hash),
        hash: fundRepoData?.hash,
        onSuccess: async () => {
            await refetchAllowance();
            toast({
                title: "Repository Funded!",
                description: "GLHF",
                action: <ToastAction altText="Play">Repository Funded!</ToastAction>,
            })
        }
    })

    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button disabled={!connected} variant="secondary">Fund</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-md">
                <DialogHeader>
                    <DialogTitle>Fund this repository</DialogTitle>
                    <DialogDescription>
                        You can fund a repository so that the maintainers can allocate funds to issues.
                    </DialogDescription>
                </DialogHeader>
                <div className="flex items-center space-x-2">
                    <div className="grid flex-1 gap-2">
                        <DropdownMenuTokens />
                        <Label htmlFor="link" className="sr-only">
                            Amount
                        </Label>
                        <Input onChange={(e) => setAmount(Number(e.target.value))} type="number" placeholder="Amount" id="link" />
                    </div>

                </div>
                <DialogFooter className="sm:justify-start">
                    <DialogClose asChild>
                        <Button type="button" variant="ghost">
                            Close
                        </Button>
                    </DialogClose>
                    <Button variant={'secondary'} onClick={(e) => {
                        e.preventDefault()
                        if (allowance === 0n) {
                            approvalWrite?.()
                            return
                        }
                        fundRepository?.()
                    }}>

                        {isApproving || isFundLoading ? <Loader2 className="animate-spin w-6 h-6" /> : allowance === 0n ? "Approve" : fundRepository ? "Fund" : "???"}
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}


