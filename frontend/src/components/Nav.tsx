"use client"

import Provider from '@/app/(providers)/Wallet';
import { Button } from '@/components/ui/button';
import { DynamicWidget } from '@dynamic-labs/sdk-react-core';

import Image from 'next/image';

function Nav() {
    return (
        <nav className="flex justify-between py-8 px-24 items-center">
            <Image src="/logo.svg" alt="logo" width={112} height={40} />
            <div className="flex space-between items-center">
                <Button variant="link">Git Started</Button>
                <Button variant="link">Git Help</Button>
                <Provider>
                    <DynamicWidget />
                </Provider>
            </div>
        </nav>
    )
}

export default Nav