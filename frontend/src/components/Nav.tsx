"use client"

import Provider from '@/app/(providers)/Wallet';
import { DynamicWidget } from '@dynamic-labs/sdk-react-core';
import Image from 'next/image';
import Link from 'next/link';
import NavLinks from './NavLinks';

function Nav() {
    return (
        <nav className="flex justify-between py-8 px-24 items-center">
            <Link legacyBehavior={true} href="/">
                <Image src="/logo.svg" alt="logo" width={112} height={40} />
            </Link>
            <div className="flex space-between items-center space-x-4">
                <Provider>
                    <NavLinks />
                    <DynamicWidget />
                </Provider>
            </div>
        </nav>
    )
}

export default Nav