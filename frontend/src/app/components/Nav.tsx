import { Button } from '@/components/ui/button';

import Image from 'next/image';
import Provider from '../(providers)/Wallet';

function Nav() {
    return (
        <nav className="flex justify-between py-8 px-24">
            <Image src="/logo.svg" alt="logo" width={112} height={40} />
            <div className="flex space-between">
                <Button variant="link">Git Started</Button>
                <Button variant="link">Git Help</Button>
                <Provider />
            </div>
        </nav>
    )
}

export default Nav