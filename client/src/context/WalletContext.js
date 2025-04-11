import React, { createContext, useEffect, useState } from 'react';
import { connectWallet } from '../api/transactions';
import { blockchainApi } from '../api/blockchain';
import { fromBaseUnits } from '../api/transactions';
import { registerKeplr } from '../utils/keplrConfig';

// Chain configuration
const CHAIN_ID = process.env.REACT_APP_CHAIN_ID || "nomercychain-testnet-1";

export const WalletContext = createContext({
  account: null,
  client: null,
  balance: null,
  chainId: null,
  connectWallet: () => {},
  disconnectWallet: () => {},
  isConnecting: false,
  error: null,
  sendTransaction: () => {},
});

export const WalletProvider = ({ children }) => {
  const [account, setAccount] = useState(null);
  const [client, setClient] = useState(null);
  const [balance, setBalance] = useState(null);
  const [chainId, setChainId] = useState(CHAIN_ID);
  const [isConnecting, setIsConnecting] = useState(false);
  const [error, setError] = useState(null);

  // Connect to Keplr wallet
  const connectKeplrWallet = async () => {
    try {
      setIsConnecting(true);
      setError(null);
      
      // Check if Keplr is installed
      if (!window.keplr) {
        throw new Error("Keplr wallet not found. Please install Keplr extension.");
      }
      
      // Register NoMercyChain with Keplr
      const registered = await registerKeplr();
      if (!registered) {
        throw new Error("Failed to register NoMercyChain with Keplr.");
      }
      
      // Connect to Keplr
      const { client: signingClient, accounts } = await connectWallet(window.keplr);
      const account = accounts[0];
      
      // Set account and client
      setAccount(account.address);
      setClient(signingClient);
      
      // Get account balance
      await updateBalance(account.address);
      
      // Save connection state
      localStorage.setItem('walletConnected', 'true');
      localStorage.setItem('walletAddress', account.address);
      
      setIsConnecting(false);
    } catch (error) {
      console.error('Error connecting wallet:', error);
      setError(error.message);
      setIsConnecting(false);
    }
  };

  // Disconnect wallet
  const disconnectWallet = () => {
    setAccount(null);
    setClient(null);
    setBalance(null);
    
    localStorage.removeItem('walletConnected');
    localStorage.removeItem('walletAddress');
  };

  // Update account balance
  const updateBalance = async (address) => {
    try {
      const balanceResponse = await blockchainApi.getBalance(address);
      const nmxBalance = balanceResponse.balances.find(b => b.denom === "unmx");
      setBalance(nmxBalance ? fromBaseUnits(nmxBalance.amount) : "0");
    } catch (error) {
      console.error('Error fetching balance:', error);
      setBalance("0");
    }
  };

  // Send a transaction
  const sendTransaction = async (msg, memo = "") => {
    if (!client || !account) {
      throw new Error("Wallet not connected");
    }
    
    try {
      // Default fee
      const fee = {
        amount: [{ denom: "unmx", amount: "5000" }],
        gas: "200000",
      };
      
      // Sign and broadcast the transaction
      const result = await client.signAndBroadcast(
        account,
        [msg],
        fee,
        memo
      );
      
      // Update balance after transaction
      await updateBalance(account);
      
      return result;
    } catch (error) {
      console.error("Error sending transaction:", error);
      throw error;
    }
  };

  // Check if wallet was previously connected
  useEffect(() => {
    const checkConnection = async () => {
      const isConnected = localStorage.getItem('walletConnected') === 'true';
      const savedAddress = localStorage.getItem('walletAddress');
      
      if (isConnected && window.keplr && savedAddress) {
        try {
          // Try to reconnect
          await connectKeplrWallet();
        } catch (error) {
          console.error('Error reconnecting wallet:', error);
          // Clear saved connection if reconnection fails
          localStorage.removeItem('walletConnected');
          localStorage.removeItem('walletAddress');
        }
      }
    };
    
    checkConnection();
  }, []);

  // Set up Keplr change listener
  useEffect(() => {
    if (window.keplr) {
      // Listen for account changes
      const handleAccountsChanged = () => {
        // Reconnect with new account
        if (account) {
          connectKeplrWallet();
        }
      };
      
      window.addEventListener("keplr_keystorechange", handleAccountsChanged);
      
      return () => {
        window.removeEventListener("keplr_keystorechange", handleAccountsChanged);
      };
    }
  }, [account]);

  // Periodically update balance
  useEffect(() => {
    if (account) {
      const intervalId = setInterval(() => {
        updateBalance(account);
      }, 30000); // Update every 30 seconds
      
      return () => clearInterval(intervalId);
    }
  }, [account]);

  return (
    <WalletContext.Provider
      value={{
        account,
        client,
        balance,
        chainId,
        connectWallet: connectKeplrWallet,
        disconnectWallet,
        isConnecting,
        error,
        sendTransaction,
      }}
    >
      {children}
    </WalletContext.Provider>
  );
};