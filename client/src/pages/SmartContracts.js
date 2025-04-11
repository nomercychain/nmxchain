import AddIcon from '@mui/icons-material/Add';
import AutoFixHighIcon from '@mui/icons-material/AutoFixHigh';
import CodeIcon from '@mui/icons-material/Code';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import SettingsIcon from '@mui/icons-material/Settings';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import StorageIcon from '@mui/icons-material/Storage';
import {
    Avatar,
    Box,
    Button,
    Card,
    CardActions,
    CardContent,
    Chip,
    Container,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Divider,
    FormControl,
    Grid,
    IconButton,
    InputLabel,
    MenuItem,
    Paper,
    Select,
    Tab,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Tabs,
    TextField,
    Typography
} from '@mui/material';
import { styled } from '@mui/material/styles';
import React, { useContext, useState } from 'react';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ContractCard = styled(Card)(({ theme }) => ({
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
  transition: 'transform 0.2s',
  '&:hover': {
    transform: 'translateY(-4px)',
  },
}));

const StyledTab = styled(Tab)(({ theme }) => ({
  textTransform: 'none',
  fontWeight: 600,
  fontSize: '1rem',
}));

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  borderBottom: `1px solid ${theme.palette.divider}`,
}));

const CodeBlock = styled(Box)(({ theme }) => ({
  backgroundColor: 'rgba(0, 0, 0, 0.2)',
  borderRadius: 8,
  padding: theme.spacing(2),
  fontFamily: 'monospace',
  fontSize: '0.875rem',
  overflowX: 'auto',
  position: 'relative',
}));

const SmartContracts = () => {
  const { account } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [contractType, setContractType] = useState('');
  const [contractName, setContractName] = useState('');
  const [contractDescription, setContractDescription] = useState('');

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleCreateDialogOpen = () => {
    setCreateDialogOpen(true);
  };

  const handleCreateDialogClose = () => {
    setCreateDialogOpen(false);
    setContractType('');
    setContractName('');
    setContractDescription('');
  };

  const handleCreateContract = () => {
    // Mock contract creation
    console.log(`Created ${contractType} contract: ${contractName}`);
    handleCreateDialogClose();
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    // You could add a toast notification here
  };

  // Mock contracts data
  const contracts = [
    {
      id: 'contract1',
      name: 'TokenSwap DynaContract',
      description: 'AI-powered token swap contract that adapts to market conditions',
      type: 'dynacontract',
      address: '0x1234...5678',
      createdAt: '2023-06-10',
      status: 'active',
      transactions: 124,
      aiModel: 'GPT-4',
    },
    {
      id: 'contract2',
      name: 'Governance Voting',
      description: 'Standard voting contract for governance proposals',
      type: 'standard',
      address: '0x8765...4321',
      createdAt: '2023-05-22',
      status: 'active',
      transactions: 45,
    },
    {
      id: 'contract3',
      name: 'NFT Marketplace',
      description: 'Contract for buying and selling NFTs with dynamic pricing',
      type: 'dynacontract',
      address: '0x2468...1357',
      createdAt: '2023-06-01',
      status: 'active',
      transactions: 67,
      aiModel: 'GPT-4',
    },
  ];

  // Mock contract interactions
  const contractInteractions = [
    {
      id: 'tx1',
      contract: 'TokenSwap DynaContract',
      function: 'swap',
      params: 'tokenA: NMX, tokenB: ETH, amount: 50',
      result: 'Success',
      time: '2023-06-15 14:30',
    },
    {
      id: 'tx2',
      contract: 'Governance Voting',
      function: 'castVote',
      params: 'proposalId: 5, vote: yes',
      result: 'Success',
      time: '2023-06-14 10:15',
    },
    {
      id: 'tx3',
      contract: 'NFT Marketplace',
      function: 'buyNFT',
      params: 'tokenId: 123, price: 50 NMX',
      result: 'Success',
      time: '2023-06-13 09:45',
    },
  ];

  // Sample DynaContract code
  const sampleDynaContractCode = `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@nomercychain/dynacontracts/DynaContract.sol";

contract TokenSwapDynaContract is DynaContract {
    // State variables
    address public owner;
    mapping(address => mapping(address => uint256)) public liquidity;
    
    // AI model configuration
    AIModel public priceModel;
    
    constructor() {
        owner = msg.sender;
        // Initialize AI model for price prediction
        priceModel = registerAIModel("price-prediction", "GPT-4");
    }
    
    // Dynamic swap function that adapts based on AI predictions
    function swap(address tokenA, address tokenB, uint256 amount) 
        external 
        returns (uint256) {
        
        // Get price prediction from AI model
        bytes memory input = abi.encode(tokenA, tokenB, amount);
        bytes memory output = priceModel.predict(input);
        uint256 expectedReturn = abi.decode(output, (uint256));
        
        // Execute swap with dynamic pricing
        // ... swap logic here ...
        
        return expectedReturn;
    }
}`;

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          Smart Contracts
        </Typography>
        <Button 
          variant="contained" 
          startIcon={<AddIcon />}
          onClick={handleCreateDialogOpen}
          disabled={!account}
        >
          Create Contract
        </Button>
      </Box>

      <Grid container spacing={4}>
        <Grid item xs={12} md={8}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              DynaContracts: AI-Powered Smart Contracts
            </Typography>
            <Typography variant="body1" paragraph>
              DynaContracts are a new type of smart contract that can adapt their behavior based on external data and AI models.
              They can make decisions, optimize parameters, and evolve over time without requiring manual updates.
            </Typography>
            <Grid container spacing={3} sx={{ mt: 2 }}>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <AutoFixHighIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Self-Adapting
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Contracts that evolve based on conditions
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <SmartToyIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    AI-Powered
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Leverages advanced AI models
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <StorageIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Secure
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Built on proven blockchain security
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </StyledPaper>
        </Grid>
        <Grid item xs={12} md={4}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Contract Templates
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Button 
              variant="outlined" 
              fullWidth 
              sx={{ mb: 2 }}
              startIcon={<AutoFixHighIcon />}
              onClick={() => {
                setContractType('dynacontract');
                setCreateDialogOpen(true);
              }}
              disabled={!account}
            >
              DynaContract
            </Button>
            <Button 
              variant="outlined" 
              fullWidth 
              sx={{ mb: 2 }}
              startIcon={<CodeIcon />}
              onClick={() => {
                setContractType('standard');
                setCreateDialogOpen(true);
              }}
              disabled={!account}
            >
              Standard Contract
            </Button>
            <Button 
              variant="outlined" 
              fullWidth
              startIcon={<StorageIcon />}
              onClick={() => {
                setContractType('upgradeable');
                setCreateDialogOpen(true);
              }}
              disabled={!account}
            >
              Upgradeable Contract
            </Button>
          </StyledPaper>
        </Grid>
      </Grid>

      <Box sx={{ mt: 4 }}>
        <Tabs 
          value={tabValue} 
          onChange={handleTabChange} 
          sx={{ 
            mb: 3,
            '& .MuiTabs-indicator': {
              backgroundColor: 'primary.main',
            }
          }}
        >
          <StyledTab label="Your Contracts" />
          <StyledTab label="Contract Interactions" />
          <StyledTab label="Code Examples" />
        </Tabs>

        {tabValue === 0 && (
          <>
            {account ? (
              <Grid container spacing={4}>
                {contracts.map((contract) => (
                  <Grid item xs={12} sm={6} md={4} key={contract.id}>
                    <ContractCard>
                      <CardContent sx={{ flexGrow: 1 }}>
                        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                          <Avatar sx={{ 
                            bgcolor: contract.type === 'dynacontract' ? 'primary.main' : 'secondary.main',
                            mr: 2
                          }}>
                            {contract.type === 'dynacontract' ? <AutoFixHighIcon /> : <CodeIcon />}
                          </Avatar>
                          <Typography variant="h6">
                            {contract.name}
                          </Typography>
                        </Box>
                        <Chip 
                          label={contract.type === 'dynacontract' ? 'DynaContract' : 'Standard Contract'} 
                          size="small" 
                          sx={{ 
                            mb: 2,
                            bgcolor: contract.type === 'dynacontract' ? 'primary.main' : 'secondary.main',
                            color: 'white'
                          }} 
                        />
                        <Typography variant="body2" color="text.secondary" paragraph>
                          {contract.description}
                        </Typography>
                        <Divider sx={{ my: 2 }} />
                        <Grid container spacing={2}>
                          <Grid item xs={6}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Address
                            </Typography>
                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                              <Typography variant="body2" sx={{ mr: 1 }}>
                                {contract.address}
                              </Typography>
                              <IconButton 
                                size="small" 
                                onClick={() => copyToClipboard(contract.address)}
                                sx={{ p: 0 }}
                              >
                                <ContentCopyIcon fontSize="small" />
                              </IconButton>
                            </Box>
                          </Grid>
                          <Grid item xs={6}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Transactions
                            </Typography>
                            <Typography variant="body2">
                              {contract.transactions}
                            </Typography>
                          </Grid>
                          {contract.type === 'dynacontract' && (
                            <Grid item xs={6}>
                              <Typography variant="caption" color="text.secondary" display="block">
                                AI Model
                              </Typography>
                              <Typography variant="body2">
                                {contract.aiModel}
                              </Typography>
                            </Grid>
                          )}
                          <Grid item xs={6}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Created
                            </Typography>
                            <Typography variant="body2">
                              {contract.createdAt}
                            </Typography>
                          </Grid>
                        </Grid>
                      </CardContent>
                      <CardActions sx={{ justifyContent: 'space-between', p: 2 }}>
                        <IconButton size="small">
                          <SettingsIcon />
                        </IconButton>
                        <Button size="small" variant="outlined" startIcon={<PlayArrowIcon />}>
                          Interact
                        </Button>
                      </CardActions>
                    </ContractCard>
                  </Grid>
                ))}
              </Grid>
            ) : (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  Connect Your Wallet
                </Typography>
                <Typography variant="body1" color="text.secondary" paragraph>
                  Please connect your wallet to view and manage your smart contracts.
                </Typography>
                <Button 
                  variant="contained" 
                  size="large" 
                  onClick={() => {}}
                  sx={{ 
                    mt: 2, 
                    background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
                    '&:hover': {
                      background: 'linear-gradient(45deg, #5000d6 30%, #00c060 90%)',
                    },
                  }}
                >
                  Connect Wallet
                </Button>
              </StyledPaper>
            )}
          </>
        )}

        {tabValue === 1 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Recent Contract Interactions
            </Typography>
            <Divider sx={{ my: 2 }} />
            {account ? (
              <TableContainer>
                <Table>
                  <TableHead>
                    <TableRow>
                      <StyledTableCell>Contract</StyledTableCell>
                      <StyledTableCell>Function</StyledTableCell>
                      <StyledTableCell>Parameters</StyledTableCell>
                      <StyledTableCell>Result</StyledTableCell>
                      <StyledTableCell>Time</StyledTableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {contractInteractions.map((interaction) => (
                      <TableRow key={interaction.id}>
                        <StyledTableCell>{interaction.contract}</StyledTableCell>
                        <StyledTableCell>{interaction.function}</StyledTableCell>
                        <StyledTableCell>{interaction.params}</StyledTableCell>
                        <StyledTableCell>
                          <Chip 
                            label={interaction.result} 
                            size="small"
                            color={interaction.result === 'Success' ? 'success' : 'error'}
                          />
                        </StyledTableCell>
                        <StyledTableCell>{interaction.time}</StyledTableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            ) : (
              <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                Connect your wallet to see your contract interactions
              </Typography>
            )}
          </StyledPaper>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              DynaContract Example
            </Typography>
            <Typography variant="body2" color="text.secondary" paragraph>
              Below is an example of a DynaContract that uses AI to optimize token swaps:
            </Typography>
            <CodeBlock>
              <Box sx={{ position: 'absolute', top: 8, right: 8 }}>
                <IconButton 
                  size="small" 
                  onClick={() => copyToClipboard(sampleDynaContractCode)}
                  sx={{ color: 'text.secondary' }}
                >
                  <ContentCopyIcon fontSize="small" />
                </IconButton>
              </Box>
              <pre style={{ margin: 0 }}>
                {sampleDynaContractCode}
              </pre>
            </CodeBlock>
            <Box sx={{ mt: 3 }}>
              <Typography variant="body2" paragraph>
                Key features of this DynaContract:
              </Typography>
              <ul>
                <li>
                  <Typography variant="body2" paragraph>
                    Imports the DynaContract base contract which provides AI integration capabilities
                  </Typography>
                </li>
                <li>
                  <Typography variant="body2" paragraph>
                    Registers an AI model for price prediction using the GPT-4 model
                  </Typography>
                </li>
                <li>
                  <Typography variant="body2" paragraph>
                    The swap function uses the AI model to predict optimal pricing for token swaps
                  </Typography>
                </li>
              </ul>
            </Box>
          </StyledPaper>
        )}
      </Box>

      {/* Create Contract Dialog */}
      <Dialog open={createDialogOpen} onClose={handleCreateDialogClose} maxWidth="md" fullWidth>
        <DialogTitle>
          {contractType === 'dynacontract' ? 'Create DynaContract' : 
           contractType === 'standard' ? 'Create Standard Contract' : 
           contractType === 'upgradeable' ? 'Create Upgradeable Contract' : 
           'Create Smart Contract'}
        </DialogTitle>
        <DialogContent>
          {!contractType && (
            <FormControl fullWidth margin="normal">
              <InputLabel>Contract Type</InputLabel>
              <Select
                value={contractType}
                label="Contract Type"
                onChange={(e) => setContractType(e.target.value)}
              >
                <MenuItem value="dynacontract">DynaContract</MenuItem>
                <MenuItem value="standard">Standard Contract</MenuItem>
                <MenuItem value="upgradeable">Upgradeable Contract</MenuItem>
              </Select>
            </FormControl>
          )}
          <TextField
            label="Contract Name"
            fullWidth
            margin="normal"
            variant="outlined"
            value={contractName}
            onChange={(e) => setContractName(e.target.value)}
          />
          <TextField
            label="Contract Description"
            fullWidth
            margin="normal"
            variant="outlined"
            multiline
            rows={3}
            value={contractDescription}
            onChange={(e) => setContractDescription(e.target.value)}
          />
          {contractType === 'dynacontract' && (
            <FormControl fullWidth margin="normal">
              <InputLabel>AI Model</InputLabel>
              <Select
                value="gpt4"
                label="AI Model"
              >
                <MenuItem value="gpt4">GPT-4</MenuItem>
                <MenuItem value="gpt3">GPT-3.5</MenuItem>
                <MenuItem value="custom">Custom Model</MenuItem>
              </Select>
            </FormControl>
          )}
          {contractType && (
            <Box sx={{ mt: 3 }}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Contract Template
              </Typography>
              <CodeBlock sx={{ maxHeight: 200, overflow: 'auto' }}>
                <pre style={{ margin: 0 }}>
                  {contractType === 'dynacontract' ? 
                    `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@nomercychain/dynacontracts/DynaContract.sol";

contract ${contractName || 'MyDynaContract'} is DynaContract {
    // State variables
    address public owner;
    
    // AI model configuration
    AIModel public aiModel;
    
    constructor() {
        owner = msg.sender;
        // Initialize AI model
        aiModel = registerAIModel("my-model", "GPT-4");
    }
    
    // Your contract logic here
}` : 
                    `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ${contractName || 'MyContract'} {
    // State variables
    address public owner;
    
    constructor() {
        owner = msg.sender;
    }
    
    // Your contract logic here
}`}
                </pre>
              </CodeBlock>
            </Box>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCreateDialogClose}>Cancel</Button>
          <Button 
            variant="contained" 
            onClick={handleCreateContract}
            disabled={!contractType || !contractName}
          >
            Create
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default SmartContracts;