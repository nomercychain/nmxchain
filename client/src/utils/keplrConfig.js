/**
 * Keplr wallet configuration utilities
 */

// Chain information for NoMercyChain
const getChainInfo = () => {
  // Try to parse the chain info from environment variable
  try {
    const chainInfoStr = process.env.REACT_APP_KEPLR_CHAIN_INFO;
    if (chainInfoStr) {
      return JSON.parse(chainInfoStr);
    }
  } catch (error) {
    console.error("Error parsing KEPLR_CHAIN_INFO:", error);
  }

  // Fallback to default configuration
  return {
    chainId: process.env.REACT_APP_CHAIN_ID || "nomercychain-testnet-1",
    chainName: process.env.REACT_APP_CHAIN_NAME || "NoMercyChain Testnet",
    rpc: process.env.REACT_APP_RPC_ENDPOINT || "http://localhost:26657",
    rest: process.env.REACT_APP_API_URL || "http://localhost:1317",
    stakeCurrency: {
      coinDenom: process.env.REACT_APP_DENOM_NAME || "NMX",
      coinMinimalDenom: process.env.REACT_APP_DENOM || "unmx",
      coinDecimals: parseInt(process.env.REACT_APP_DECIMAL_PLACES || "6"),
    },
    bip44: {
      coinType: 118,
    },
    bech32Config: {
      bech32PrefixAccAddr: "nmx",
      bech32PrefixAccPub: "nmxpub",
      bech32PrefixValAddr: "nmxvaloper",
      bech32PrefixValPub: "nmxvaloperpub",
      bech32PrefixConsAddr: "nmxvalcons",
      bech32PrefixConsPub: "nmxvalconspub",
    },
    currencies: [
      {
        coinDenom: process.env.REACT_APP_DENOM_NAME || "NMX",
        coinMinimalDenom: process.env.REACT_APP_DENOM || "unmx",
        coinDecimals: parseInt(process.env.REACT_APP_DECIMAL_PLACES || "6"),
      },
    ],
    feeCurrencies: [
      {
        coinDenom: process.env.REACT_APP_DENOM_NAME || "NMX",
        coinMinimalDenom: process.env.REACT_APP_DENOM || "unmx",
        coinDecimals: parseInt(process.env.REACT_APP_DECIMAL_PLACES || "6"),
      },
    ],
    gasPriceStep: {
      low: 0.01,
      average: 0.025,
      high: 0.04,
    },
  };
};

/**
 * Register NoMercyChain with Keplr wallet
 * @returns {Promise<boolean>} True if registration was successful
 */
export const registerKeplr = async () => {
  if (!window.keplr) {
    alert("Please install Keplr extension");
    return false;
  }

  try {
    // Try to enable the chain
    await window.keplr.enable(getChainInfo().chainId);
    return true;
  } catch (error) {
    console.log("Chain not registered in Keplr, attempting to suggest chain...");
    
    try {
      // Suggest chain to Keplr
      await window.keplr.experimentalSuggestChain(getChainInfo());
      
      // Enable the chain
      await window.keplr.enable(getChainInfo().chainId);
      console.log("Chain successfully registered with Keplr");
      return true;
    } catch (error) {
      console.error("Failed to suggest chain to Keplr", error);
      alert("Failed to register chain with Keplr: " + error.message);
      return false;
    }
  }
};

/**
 * Get Keplr offline signer for NoMercyChain
 * @returns {Promise<Object>} Offline signer object
 */
export const getOfflineSigner = async () => {
  if (!window.keplr) {
    throw new Error("Keplr extension not installed");
  }

  const chainId = getChainInfo().chainId;
  await window.keplr.enable(chainId);
  
  return window.keplr.getOfflineSigner(chainId);
};

export default {
  getChainInfo,
  registerKeplr,
  getOfflineSigner,
};