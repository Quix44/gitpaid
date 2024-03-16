"use client"

import { DropdownMenuCheckboxItemProps } from "@radix-ui/react-dropdown-menu"
import * as React from "react"

import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuCheckboxItem,
    DropdownMenuContent,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

type Checked = DropdownMenuCheckboxItemProps["checked"]

export function DropdownMenuTokens({ setter }: any) {
    const [showStatusBar, setShowStatusBar] = React.useState<Checked>(true)
    const [showActivityBar, setShowActivityBar] = React.useState<Checked>(false)
    const [showPanel, setShowPanel] = React.useState<Checked>(false)

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="outline">Open</Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56" >
                <DropdownMenuLabel>Select Token</DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuCheckboxItem
                    checked={showStatusBar}
                    onCheckedChange={setShowStatusBar}
                >
                    USDC
                </DropdownMenuCheckboxItem>
                <DropdownMenuCheckboxItem
                    checked={showActivityBar}
                    onCheckedChange={setShowActivityBar}
                    disabled
                >
                    Ape Coin
                </DropdownMenuCheckboxItem>
                <DropdownMenuCheckboxItem
                    checked={showPanel}
                    onCheckedChange={setShowPanel}
                >
                    Arb Token
                </DropdownMenuCheckboxItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}