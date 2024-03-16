"use client"

import { EthereumWalletConnectors } from "@dynamic-labs/ethereum";
import { DynamicContextProvider } from '@dynamic-labs/sdk-react-core';
import { useRouter } from "next/navigation";
import { DynamicWagmiConnector, EthersExtension } from "./Dynamic";

const evmNetworks = [
    {
        blockExplorerUrls: ['https://sepolia.arbiscan.io/'],
        chainId: 421614,
        chainName: 'Arbitrum Sepolia',
        iconUrls: ['https://app.dynamic.xyz/assets/networks/arbitrum.svg'],
        name: 'Arbitrum Sepolia',
        nativeCurrency: {
            decimals: 18,
            name: 'Arbitrum Sepolia Ether',
            symbol: 'SEP',
        },
        networkId: 421614,
        rpcUrls: ['https://arb-sepolia.g.alchemy.com/v2/Z8Y0CZXvhPgiTt8akdr4Z_dS03C2-H0X'],
        vanityName: 'Arbitrum Sepolia',
    },
];

const Provider = ({ children }: any) => {
    const router = useRouter()

    return (
        <DynamicContextProvider
            settings={{
                evmNetworks,
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
                walletConnectorExtensions: [EthersExtension],
                environmentId: '46c4b660-c6c1-462e-817e-1cf4459ac07f',
                walletConnectors: [EthereumWalletConnectors],
            }}>
            <DynamicWagmiConnector>
                {children}
            </DynamicWagmiConnector>
        </DynamicContextProvider>
    )

};

export default Provider;