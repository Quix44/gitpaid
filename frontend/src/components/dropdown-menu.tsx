"use client"

import * as React from "react";

import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";

export function DropdownMenuTokens() {
    const [selectedToken, setSelectedToken] = React.useState<string | undefined>(undefined);

    const handleSelectChange = (value: string) => {
        setSelectedToken(value);
        // Implement any additional logic you need when a new token is selected.
        // For example, update other components or states based on the selected token.
    };

    return (
        <Select onValueChange={handleSelectChange}>
            <SelectTrigger className="w-[280px]">
                <SelectValue placeholder="Select ERC20" />
            </SelectTrigger>
            <SelectContent>
                <SelectGroup>
                    <SelectLabel>ERC20 Token</SelectLabel>
                    <SelectItem value="usdc">USDC</SelectItem>

                    <SelectItem value="ape" >Ape Coin</SelectItem>
                    <SelectItem value="arb">Arb Token</SelectItem>
                </SelectGroup>
            </SelectContent>
        </Select>
    );
}
