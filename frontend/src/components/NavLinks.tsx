"use client"

import { useDynamicContext } from '@dynamic-labs/sdk-react-core';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import { Button } from './ui/button';

function NavLinks() {
    const [isLoaded, setIsLoaded] = useState(false)
    const { user, isAuthenticated } = useDynamicContext()
    const githubUsername = user?.verifiedCredentials.find(g => g.oauthProvider === "github")?.oauthUsername

    useEffect(() => {
        setIsLoaded(true)
    }, [isAuthenticated])

    return (
        <>
            {isLoaded && user ? (
                <>
                    {isAuthenticated ? <Button variant="link" asChild>
                        <Link href={`/repositories/?username=${githubUsername}`}>
                            My Repositories
                        </Link>
                    </Button> : null}
                    <Button variant="link" asChild>
                        <Link href="/issues">
                            Open Issues
                        </Link>
                    </Button>
                </>
            ) : null}
        </>
    )
}

export default NavLinks