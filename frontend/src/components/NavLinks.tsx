import { useDynamicContext } from '@dynamic-labs/sdk-react-core';
import Link from 'next/link';
import { Button } from './ui/button';

function NavLinks() {
    const { user, isAuthenticated } = useDynamicContext()

    // Find the credential with an oauthUsername
    const oauthCredential = user?.verifiedCredentials.find(credential => credential.oauthUsername);

    // Extract the oauthUsername, if available
    const oauthUsername = oauthCredential ? oauthCredential.oauthUsername : null;


    return (
        <>
            {user ? (
                <>
                    <Button variant="link" asChild>
                        <Link href={`/repository/${oauthUsername}`}>
                            My Repositories
                        </Link>

                    </Button>
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