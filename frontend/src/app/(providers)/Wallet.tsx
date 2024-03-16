"use client"

import { EthereumWalletConnectors } from "@dynamic-labs/ethereum";
import { DynamicContextProvider } from '@dynamic-labs/sdk-react-core';
import { useRouter } from "next/navigation";

const Provider = ({ children }: any) => {
    const router = useRouter()

    return (
        <DynamicContextProvider
            settings={{
                onboardingImageUrl: 'https://i.imgur.com/3g7nmJC.png',
                eventsCallbacks: {
                    onAuthFlowClose: () => {
                        console.log('in onAuthFlowClose');
                    },
                    onAuthFlowOpen: () => {
                        console.log('in onAuthFlowOpen');
                    },
                    onAuthSuccess: (e: any) => {
                        // Set the github username in the url
                        const githubUsername = e.user.verifiedCredentials.find((g: { oauthProvider: string; }) => g.oauthProvider === "github")?.oauthUsername
                        console.log(githubUsername)
                        return router.push(`/repositories?username=${githubUsername}`)
                    },
                    onLogout: () => {
                        console.log('in onLogout');
                    },
                },
                environmentId: '46c4b660-c6c1-462e-817e-1cf4459ac07f',
                walletConnectors: [EthereumWalletConnectors],
            }}>
            {children}
        </DynamicContextProvider>
    )

};

export default Provider;