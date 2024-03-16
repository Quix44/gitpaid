"use client"
import { EthereumWalletConnectors } from "@dynamic-labs/ethereum";
import { DynamicContextProvider, DynamicWidget } from '@dynamic-labs/sdk-react-core';

const Provider = () => (
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
                onAuthSuccess: () => {
                    console.log('in onAuthSuccess');
                    window.location.href = "/dashboard"
                },
                onLogout: () => {
                    console.log('in onLogout');
                },
            },
            environmentId: '46c4b660-c6c1-462e-817e-1cf4459ac07f',
            walletConnectors: [EthereumWalletConnectors],
        }}>
        <DynamicWidget />
    </DynamicContextProvider>
);

export default Provider;