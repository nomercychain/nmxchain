import axios from 'axios';

// Get API URL from environment variable or use default
const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:1317';

/**
 * Blockchain API client for NoMercyChain
 * Provides methods to interact with the blockchain's REST API
 */
export const blockchainApi = {
  // ==================== Account Endpoints ====================
  
  /**
   * Get account information
   * @param {string} address - Account address
   * @returns {Promise<Object>} Account information
   */
  getAccount: async (address) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/auth/v1beta1/accounts/${address}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching account:', error);
      throw error;
    }
  },
  
  /**
   * Get account balances
   * @param {string} address - Account address
   * @returns {Promise<Object>} Account balances
   */
  getBalance: async (address) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/bank/v1beta1/balances/${address}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching balance:', error);
      throw error;
    }
  },
  
  // ==================== Staking Endpoints ====================
  
  /**
   * Get all validators
   * @param {string} status - Optional validator status filter (bonded, unbonded, unbonding)
   * @returns {Promise<Object>} List of validators
   */
  getValidators: async (status = '') => {
    try {
      const url = status 
        ? `${API_BASE_URL}/cosmos/staking/v1beta1/validators?status=${status}`
        : `${API_BASE_URL}/cosmos/staking/v1beta1/validators`;
      const response = await axios.get(url);
      return response.data;
    } catch (error) {
      console.error('Error fetching validators:', error);
      throw error;
    }
  },
  
  /**
   * Get validator details
   * @param {string} validatorAddr - Validator address
   * @returns {Promise<Object>} Validator details
   */
  getValidator: async (validatorAddr) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/staking/v1beta1/validators/${validatorAddr}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching validator:', error);
      throw error;
    }
  },
  
  /**
   * Get delegations for an address
   * @param {string} address - Delegator address
   * @returns {Promise<Object>} List of delegations
   */
  getDelegations: async (address) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/staking/v1beta1/delegations/${address}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching delegations:', error);
      throw error;
    }
  },
  
  /**
   * Get delegation rewards
   * @param {string} address - Delegator address
   * @returns {Promise<Object>} Delegation rewards
   */
  getDelegationRewards: async (address) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/distribution/v1beta1/delegators/${address}/rewards`);
      return response.data;
    } catch (error) {
      console.error('Error fetching delegation rewards:', error);
      throw error;
    }
  },
  
  // ==================== Governance Endpoints ====================
  
  /**
   * Get governance proposals
   * @param {string} status - Optional proposal status filter
   * @returns {Promise<Object>} List of proposals
   */
  getProposals: async (status = '') => {
    try {
      const url = status 
        ? `${API_BASE_URL}/cosmos/gov/v1beta1/proposals?proposal_status=${status}`
        : `${API_BASE_URL}/cosmos/gov/v1beta1/proposals`;
      const response = await axios.get(url);
      return response.data;
    } catch (error) {
      console.error('Error fetching proposals:', error);
      throw error;
    }
  },
  
  /**
   * Get proposal details
   * @param {string} proposalId - Proposal ID
   * @returns {Promise<Object>} Proposal details
   */
  getProposal: async (proposalId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/gov/v1beta1/proposals/${proposalId}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching proposal:', error);
      throw error;
    }
  },
  
  /**
   * Get votes for a proposal
   * @param {string} proposalId - Proposal ID
   * @returns {Promise<Object>} List of votes
   */
  getProposalVotes: async (proposalId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/gov/v1beta1/proposals/${proposalId}/votes`);
      return response.data;
    } catch (error) {
      console.error('Error fetching proposal votes:', error);
      throw error;
    }
  },
  
  /**
   * Get vote by voter
   * @param {string} proposalId - Proposal ID
   * @param {string} voter - Voter address
   * @returns {Promise<Object>} Vote details
   */
  getVote: async (proposalId, voter) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/gov/v1beta1/proposals/${proposalId}/votes/${voter}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching vote:', error);
      throw error;
    }
  },
  
  // ==================== DynaContract Endpoints ====================
  
  /**
   * Get all DynaContracts
   * @returns {Promise<Object>} List of DynaContracts
   */
  getDynaContracts: async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/dynacontract/dyna-contracts`);
      return response.data;
    } catch (error) {
      console.error('Error fetching DynaContracts:', error);
      throw error;
    }
  },
  
  /**
   * Get DynaContract details
   * @param {string} contractId - Contract ID
   * @returns {Promise<Object>} Contract details
   */
  getDynaContract: async (contractId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/dynacontract/dyna-contract/${contractId}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching DynaContract:', error);
      throw error;
    }
  },
  
  /**
   * Get DynaContracts owned by an address
   * @param {string} owner - Owner address
   * @returns {Promise<Object>} List of contracts
   */
  getDynaContractsByOwner: async (owner) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/dynacontract/dyna-contracts-by-owner/${owner}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching DynaContracts by owner:', error);
      throw error;
    }
  },
  
  // ==================== HyperChain Endpoints ====================
  
  /**
   * Get all HyperChains
   * @returns {Promise<Object>} List of HyperChains
   */
  getHyperChains: async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/hyperchain/chains`);
      return response.data;
    } catch (error) {
      console.error('Error fetching HyperChains:', error);
      throw error;
    }
  },
  
  /**
   * Get HyperChain details
   * @param {string} chainId - HyperChain ID
   * @returns {Promise<Object>} HyperChain details
   */
  getHyperChain: async (chainId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/hyperchain/chain/${chainId}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching HyperChain:', error);
      throw error;
    }
  },
  
  /**
   * Get HyperChains owned by an address
   * @param {string} owner - Owner address
   * @returns {Promise<Object>} List of HyperChains
   */
  getHyperChainsByOwner: async (owner) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/hyperchain/chains-by-owner/${owner}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching HyperChains by owner:', error);
      throw error;
    }
  },
  
  // ==================== TruthGPT Oracle Endpoints ====================
  
  /**
   * Get all oracle queries
   * @returns {Promise<Object>} List of oracle queries
   */
  getOracleQueries: async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/truthgpt/oracle-queries`);
      return response.data;
    } catch (error) {
      console.error('Error fetching oracle queries:', error);
      throw error;
    }
  },
  
  /**
   * Get oracle queries by address
   * @param {string} address - User address
   * @returns {Promise<Object>} List of oracle queries
   */
  getOracleQueriesByAddress: async (address) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/truthgpt/oracle-queries-by-address/${address}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching oracle queries by address:', error);
      throw error;
    }
  },
  
  /**
   * Get oracle query details
   * @param {string} queryId - Query ID
   * @returns {Promise<Object>} Query details
   */
  getOracleQuery: async (queryId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/truthgpt/oracle-query/${queryId}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching oracle query:', error);
      throw error;
    }
  },
  
  // ==================== DeAI Agents Endpoints ====================
  
  /**
   * Get all AI agents
   * @returns {Promise<Object>} List of AI agents
   */
  getAIAgents: async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/deai/ai-agents`);
      return response.data;
    } catch (error) {
      console.error('Error fetching AI agents:', error);
      throw error;
    }
  },
  
  /**
   * Get AI agents owned by an address
   * @param {string} owner - Owner address
   * @returns {Promise<Object>} List of AI agents
   */
  getAIAgentsByOwner: async (owner) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/deai/ai-agents-by-owner/${owner}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching AI agents by owner:', error);
      throw error;
    }
  },
  
  /**
   * Get AI agent details
   * @param {string} agentId - Agent ID
   * @returns {Promise<Object>} Agent details
   */
  getAIAgent: async (agentId) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/nomercychain/nmxchain/deai/ai-agent/${agentId}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching AI agent:', error);
      throw error;
    }
  },
  
  // ==================== Transaction Endpoints ====================
  
  /**
   * Get transaction details
   * @param {string} hash - Transaction hash
   * @returns {Promise<Object>} Transaction details
   */
  getTransaction: async (hash) => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cosmos/tx/v1beta1/txs/${hash}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching transaction:', error);
      throw error;
    }
  },
  
  /**
   * Get transactions for an address
   * @param {string} address - Account address
   * @param {string} type - Transaction type (send, receive, etc.)
   * @returns {Promise<Object>} List of transactions
   */
  getTransactionsByAddress: async (address, type = '') => {
    try {
      let events = [`message.sender='${address}'`];
      if (type === 'receive') {
        events = [`transfer.recipient='${address}'`];
      }
      
      const eventsParam = events.map(e => `events=${encodeURIComponent(e)}`).join('&');
      const response = await axios.get(`${API_BASE_URL}/cosmos/tx/v1beta1/txs?${eventsParam}&pagination.limit=100`);
      return response.data;
    } catch (error) {
      console.error('Error fetching transactions by address:', error);
      throw error;
    }
  }
};

export default blockchainApi;