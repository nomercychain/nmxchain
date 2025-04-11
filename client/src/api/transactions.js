import { SigningStargateClient, GasPrice, coins } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";

// Chain configuration
const CHAIN_ID = process.env.REACT_APP_CHAIN_ID || "nomercychain-testnet-1";
const RPC_ENDPOINT = process.env.REACT_APP_RPC_ENDPOINT || "http://localhost:26657";
const GAS_PRICE = process.env.REACT_APP_GAS_PRICE || "0.025unmx";
const DENOM = process.env.REACT_APP_DENOM || "unmx";

// Create registry with custom message types
const registry = new Registry();

// Register custom message types for NoMercyChain modules
// These will need to be updated with the actual protobuf types once implemented
registry.register("/nomercychain.nmxchain.dynacontract.MsgCreateDynaContract", {});
registry.register("/nomercychain.nmxchain.dynacontract.MsgUpdateDynaContract", {});
registry.register("/nomercychain.nmxchain.dynacontract.MsgExecuteDynaContract", {});
registry.register("/nomercychain.nmxchain.dynacontract.MsgAddLearningData", {});

registry.register("/nomercychain.nmxchain.hyperchain.MsgCreateChain", {});
registry.register("/nomercychain.nmxchain.hyperchain.MsgUpdateChain", {});
registry.register("/nomercychain.nmxchain.hyperchain.MsgJoinChain", {});

registry.register("/nomercychain.nmxchain.truthgpt.MsgSubmitOracleQuery", {});
registry.register("/nomercychain.nmxchain.truthgpt.MsgRegisterOracleProvider", {});
registry.register("/nomercychain.nmxchain.truthgpt.MsgVerifyOracleResponse", {});

registry.register("/nomercychain.nmxchain.deai.MsgCreateAIAgent", {});
registry.register("/nomercychain.nmxchain.deai.MsgUpdateAIAgent", {});
registry.register("/nomercychain.nmxchain.deai.MsgExecuteAIAgent", {});

/**
 * Connect to the blockchain using Keplr wallet
 * @param {Object} keplr - Keplr wallet instance
 * @returns {Promise<Object>} Client and accounts
 */
export const connectWallet = async (keplr) => {
  if (!keplr) {
    throw new Error("Keplr wallet not found");
  }
  
  // Enable the chain in Keplr
  await keplr.enable(CHAIN_ID);
  
  // Get the offline signer
  const offlineSigner = keplr.getOfflineSigner(CHAIN_ID);
  const accounts = await offlineSigner.getAccounts();
  
  // Create the signing client
  const client = await SigningStargateClient.connectWithSigner(
    RPC_ENDPOINT,
    offlineSigner,
    { 
      registry, 
      gasPrice: GasPrice.fromString(GAS_PRICE) 
    }
  );
  
  return { client, accounts };
};

/**
 * Send a transaction to the blockchain
 * @param {Object} client - SigningStargateClient instance
 * @param {Object} sender - Sender account
 * @param {Object} msg - Transaction message
 * @param {string} memo - Transaction memo
 * @returns {Promise<Object>} Transaction result
 */
export const sendTransaction = async (client, sender, msg, memo = "") => {
  // Default fee
  const fee = {
    amount: coins(5000, DENOM),
    gas: "200000",
  };
  
  try {
    // Sign and broadcast the transaction
    const result = await client.signAndBroadcast(
      sender.address,
      [msg],
      fee,
      memo
    );
    
    return result;
  } catch (error) {
    console.error("Error sending transaction:", error);
    throw error;
  }
};

/**
 * Convert token amount to base units (e.g., NMX to unmx)
 * @param {string|number} amount - Token amount
 * @returns {string} Amount in base units
 */
export const toBaseUnits = (amount) => {
  return (parseFloat(amount) * 1000000).toString();
};

/**
 * Convert base units to token amount (e.g., unmx to NMX)
 * @param {string|number} amount - Amount in base units
 * @returns {string} Token amount
 */
export const fromBaseUnits = (amount) => {
  return (parseInt(amount) / 1000000).toString();
};

// ==================== Transaction Message Creators ====================

/**
 * Create messages for common blockchain operations
 */
export const msgCreators = {
  // Bank module messages
  bank: {
    /**
     * Create a send message
     * @param {string} fromAddress - Sender address
     * @param {string} toAddress - Recipient address
     * @param {string|number} amount - Amount to send
     * @returns {Object} MsgSend message
     */
    send: (fromAddress, toAddress, amount) => ({
      typeUrl: "/cosmos.bank.v1beta1.MsgSend",
      value: {
        fromAddress,
        toAddress,
        amount: [
          {
            denom: DENOM,
            amount: toBaseUnits(amount),
          },
        ],
      },
    }),
  },
  
  // Staking module messages
  staking: {
    /**
     * Create a delegate message
     * @param {string} delegatorAddress - Delegator address
     * @param {string} validatorAddress - Validator address
     * @param {string|number} amount - Amount to delegate
     * @returns {Object} MsgDelegate message
     */
    delegate: (delegatorAddress, validatorAddress, amount) => ({
      typeUrl: "/cosmos.staking.v1beta1.MsgDelegate",
      value: {
        delegatorAddress,
        validatorAddress,
        amount: {
          denom: DENOM,
          amount: toBaseUnits(amount),
        },
      },
    }),
    
    /**
     * Create an undelegate message
     * @param {string} delegatorAddress - Delegator address
     * @param {string} validatorAddress - Validator address
     * @param {string|number} amount - Amount to undelegate
     * @returns {Object} MsgUndelegate message
     */
    undelegate: (delegatorAddress, validatorAddress, amount) => ({
      typeUrl: "/cosmos.staking.v1beta1.MsgUndelegate",
      value: {
        delegatorAddress,
        validatorAddress,
        amount: {
          denom: DENOM,
          amount: toBaseUnits(amount),
        },
      },
    }),
    
    /**
     * Create a redelegate message
     * @param {string} delegatorAddress - Delegator address
     * @param {string} validatorSrcAddress - Source validator address
     * @param {string} validatorDstAddress - Destination validator address
     * @param {string|number} amount - Amount to redelegate
     * @returns {Object} MsgBeginRedelegate message
     */
    redelegate: (delegatorAddress, validatorSrcAddress, validatorDstAddress, amount) => ({
      typeUrl: "/cosmos.staking.v1beta1.MsgBeginRedelegate",
      value: {
        delegatorAddress,
        validatorSrcAddress,
        validatorDstAddress,
        amount: {
          denom: DENOM,
          amount: toBaseUnits(amount),
        },
      },
    }),
  },
  
  // Distribution module messages
  distribution: {
    /**
     * Create a withdraw rewards message
     * @param {string} delegatorAddress - Delegator address
     * @param {string} validatorAddress - Validator address
     * @returns {Object} MsgWithdrawDelegatorReward message
     */
    withdrawRewards: (delegatorAddress, validatorAddress) => ({
      typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
      value: {
        delegatorAddress,
        validatorAddress,
      },
    }),
    
    /**
     * Create a withdraw all rewards message
     * @param {string} delegatorAddress - Delegator address
     * @param {Array<string>} validatorAddresses - List of validator addresses
     * @returns {Array<Object>} Array of MsgWithdrawDelegatorReward messages
     */
    withdrawAllRewards: (delegatorAddress, validatorAddresses) => {
      return validatorAddresses.map(validatorAddress => ({
        typeUrl: "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
        value: {
          delegatorAddress,
          validatorAddress,
        },
      }));
    },
  },
  
  // Governance module messages
  governance: {
    /**
     * Create a submit proposal message
     * @param {string} proposer - Proposer address
     * @param {string} title - Proposal title
     * @param {string} description - Proposal description
     * @param {string|number} deposit - Initial deposit amount
     * @returns {Object} MsgSubmitProposal message
     */
    submitProposal: (proposer, title, description, deposit) => ({
      typeUrl: "/cosmos.gov.v1beta1.MsgSubmitProposal",
      value: {
        content: {
          typeUrl: "/cosmos.gov.v1beta1.TextProposal",
          value: {
            title,
            description,
          },
        },
        proposer,
        initialDeposit: [
          {
            denom: DENOM,
            amount: toBaseUnits(deposit),
          },
        ],
      },
    }),
    
    /**
     * Create a deposit message
     * @param {string} depositor - Depositor address
     * @param {string|number} proposalId - Proposal ID
     * @param {string|number} amount - Deposit amount
     * @returns {Object} MsgDeposit message
     */
    deposit: (depositor, proposalId, amount) => ({
      typeUrl: "/cosmos.gov.v1beta1.MsgDeposit",
      value: {
        proposalId,
        depositor,
        amount: [
          {
            denom: DENOM,
            amount: toBaseUnits(amount),
          },
        ],
      },
    }),
    
    /**
     * Create a vote message
     * @param {string} voter - Voter address
     * @param {string|number} proposalId - Proposal ID
     * @param {number} option - Vote option (1=Yes, 2=Abstain, 3=No, 4=NoWithVeto)
     * @returns {Object} MsgVote message
     */
    vote: (voter, proposalId, option) => ({
      typeUrl: "/cosmos.gov.v1beta1.MsgVote",
      value: {
        proposalId,
        voter,
        option,
      },
    }),
  },
  
  // DynaContract module messages
  dynacontract: {
    /**
     * Create a create contract message
     * @param {string} creator - Creator address
     * @param {string} name - Contract name
     * @param {string} code - Contract code
     * @param {string} description - Contract description
     * @param {string} aiModel - AI model to use
     * @returns {Object} MsgCreateDynaContract message
     */
    createContract: (creator, name, code, description, aiModel) => ({
      typeUrl: "/nomercychain.nmxchain.dynacontract.MsgCreateDynaContract",
      value: {
        creator,
        name,
        code,
        description,
        aiModel,
        contractType: "1", // Standard contract
        deposit: {
          denom: DENOM,
          amount: toBaseUnits(100), // Fixed deposit of 100 NMX
        },
      },
    }),
    
    /**
     * Create an execute contract message
     * @param {string} sender - Sender address
     * @param {string} contractId - Contract ID
     * @param {string} functionName - Function to execute
     * @param {Object} params - Function parameters
     * @param {string|number} amount - Amount to send with execution
     * @returns {Object} MsgExecuteDynaContract message
     */
    executeContract: (sender, contractId, functionName, params, amount = 0) => ({
      typeUrl: "/nomercychain.nmxchain.dynacontract.MsgExecuteDynaContract",
      value: {
        sender,
        contractId,
        functionName,
        params: JSON.stringify(params),
        amount: {
          denom: DENOM,
          amount: toBaseUnits(amount),
        },
      },
    }),
    
    /**
     * Create an add learning data message
     * @param {string} sender - Sender address
     * @param {string} contractId - Contract ID
     * @param {string} dataType - Data type (TRAINING, FEEDBACK, etc.)
     * @param {Object} data - Learning data
     * @returns {Object} MsgAddLearningData message
     */
    addLearningData: (sender, contractId, dataType, data) => ({
      typeUrl: "/nomercychain.nmxchain.dynacontract.MsgAddLearningData",
      value: {
        sender,
        contractId,
        dataType,
        data: JSON.stringify(data),
        source: "client",
      },
    }),
  },
  
  // HyperChain module messages
  hyperchain: {
    /**
     * Create a create hyperchain message
     * @param {string} creator - Creator address
     * @param {string} name - Chain name
     * @param {string} description - Chain description
     * @param {string} chainType - Chain type
     * @param {Array<string>} modules - Chain modules
     * @param {string} aiPrompt - AI prompt for chain configuration
     * @returns {Object} MsgCreateChain message
     */
    createHyperChain: (creator, name, description, chainType, modules, aiPrompt) => ({
      typeUrl: "/nomercychain.nmxchain.hyperchain.MsgCreateChain",
      value: {
        creator,
        name,
        description,
        chainType,
        modules,
        aiPrompt,
        deposit: {
          denom: DENOM,
          amount: toBaseUnits(10000), // Fixed deposit of 10,000 NMX
        },
      },
    }),
    
    /**
     * Create a join hyperchain message
     * @param {string} validator - Validator address
     * @param {string} chainId - HyperChain ID
     * @param {string|number} stake - Stake amount
     * @returns {Object} MsgJoinChain message
     */
    joinHyperChain: (validator, chainId, stake) => ({
      typeUrl: "/nomercychain.nmxchain.hyperchain.MsgJoinChain",
      value: {
        validator,
        chainId,
        stake: {
          denom: DENOM,
          amount: toBaseUnits(stake),
        },
      },
    }),
  },
  
  // TruthGPT Oracle module messages
  truthgpt: {
    /**
     * Create a submit query message
     * @param {string} sender - Sender address
     * @param {string} queryType - Query type (factCheck, dataFeed, prediction)
     * @param {string} queryPrompt - Query prompt
     * @param {string} provider - Oracle provider address
     * @returns {Object} MsgSubmitOracleQuery message
     */
    submitQuery: (sender, queryType, queryPrompt, provider) => ({
      typeUrl: "/nomercychain.nmxchain.truthgpt.MsgSubmitOracleQuery",
      value: {
        sender,
        queryType,
        queryPrompt,
        provider,
        fee: {
          denom: DENOM,
          amount: toBaseUnits(queryType === 'factCheck' ? 5 : queryType === 'dataFeed' ? 2 : 10),
        },
      },
    }),
    
    /**
     * Create a verify oracle response message
     * @param {string} verifier - Verifier address
     * @param {string} queryId - Query ID
     * @param {boolean} isValid - Whether the response is valid
     * @param {string} feedback - Verification feedback
     * @returns {Object} MsgVerifyOracleResponse message
     */
    verifyResponse: (verifier, queryId, isValid, feedback) => ({
      typeUrl: "/nomercychain.nmxchain.truthgpt.MsgVerifyOracleResponse",
      value: {
        verifier,
        queryId,
        isValid,
        feedback,
      },
    }),
  },
  
  // DeAI module messages
  deai: {
    /**
     * Create a create AI agent message
     * @param {string} creator - Creator address
     * @param {string} name - Agent name
     * @param {string} description - Agent description
     * @param {string} aiModel - AI model to use
     * @param {string} prompt - Agent prompt/instructions
     * @returns {Object} MsgCreateAIAgent message
     */
    createAIAgent: (creator, name, description, aiModel, prompt) => ({
      typeUrl: "/nomercychain.nmxchain.deai.MsgCreateAIAgent",
      value: {
        creator,
        name,
        description,
        aiModel,
        prompt,
        deposit: {
          denom: DENOM,
          amount: toBaseUnits(100), // Fixed deposit of 100 NMX
        },
      },
    }),
    
    /**
     * Create an execute AI agent message
     * @param {string} sender - Sender address
     * @param {string} agentId - Agent ID
     * @param {string} input - Input for the agent
     * @returns {Object} MsgExecuteAIAgent message
     */
    executeAIAgent: (sender, agentId, input) => ({
      typeUrl: "/nomercychain.nmxchain.deai.MsgExecuteAIAgent",
      value: {
        sender,
        agentId,
        input,
        fee: {
          denom: DENOM,
          amount: toBaseUnits(1), // Fixed fee of 1 NMX
        },
      },
    }),
  },
};

export default {
  connectWallet,
  sendTransaction,
  toBaseUnits,
  fromBaseUnits,
  msgCreators,
};