import { useDynamicContext } from '@dynamic-labs/sdk-react-core'
import Link from 'next/link'
import { Button } from './ui/button'

function NavLinks() {
    const { user, isAuthenticated } = useDynamicContext()
    return (
        <>
            {user ? (
                <>
                    <Button variant="link" asChild>
                        <Link href="/repositories">
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