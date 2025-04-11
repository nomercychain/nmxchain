import React, { useState, useContext } from 'react';
import { 
  Box, 
  Container, 
  Grid, 
  Paper, 
  Typography, 
  Button, 
  Card, 
  CardContent, 
  CardActions,
  Divider,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Chip,
  Avatar,
  IconButton,
  Tabs,
  Tab,
  LinearProgress,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Rating
} from '@mui/material';
import { styled } from '@mui/material/styles';
import VerifiedUserIcon from '@mui/icons-material/VerifiedUser';
import AddIcon from '@mui/icons-material/Add';
import SearchIcon from '@mui/icons-material/Search';
import SettingsIcon from '@mui/icons-material/Settings';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import FactCheckIcon from '@mui/icons-material/FactCheck';
import DataObjectIcon from '@mui/icons-material/DataObject';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const QueryCard = styled(Card)(({ theme }) => ({
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

const Oracle = () => {
  const { account } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [queryDialogOpen, setQueryDialogOpen] = useState(false);
  const [queryType, setQueryType] = useState('');
  const [queryPrompt, setQueryPrompt] = useState('');
  const [queryResult, setQueryResult] = useState('');
  const [showResult, setShowResult] = useState(false);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleQueryDialogOpen = () => {
    setQueryDialogOpen(true);
    setShowResult(false);
  };

  const handleQueryDialogClose = () => {
    setQueryDialogOpen(false);
    setQueryType('');
    setQueryPrompt('');
    setQueryResult('');
    setShowResult(false);
  };

  const handleSubmitQuery = () => {
    // Mock query submission
    console.log(`Submitted ${queryType} query: ${queryPrompt}`);
    
    // Mock results based on query type
    let result = '';
    if (queryType === 'factCheck') {
      result = `{
  "query": "${queryPrompt}",
  "result": {
    "factual": true,
    "confidence": 0.92,
    "sources": [
      "https://example.com/reliable-source-1",
      "https://example.com/reliable-source-2"
    ],
    "explanation": "The statement is verified by multiple reliable sources with consistent information."
  }
}`;
    } else if (queryType === 'dataFeed') {
      result = `{
  "query": "${queryPrompt}",
  "result": {
    "value": 1842.56,
    "timestamp": "${new Date().toISOString()}",
    "sources": [
      "https://api.example.com/prices",
      "https://data.example.com/crypto"
    ],
    "aggregation": "median"
  }
}`;
    } else if (queryType === 'prediction') {
      result = `{
  "query": "${queryPrompt}",
  "result": {
    "prediction": 0.78,
    "confidence": 0.85,
    "explanation": "Based on historical patterns and current market conditions, there is a 78% probability of this outcome.",
    "methodology": "Ensemble of transformer models with time-series analysis"
  }
}`;
    }
    
    setQueryResult(result);
    setShowResult(true);
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    // You could add a toast notification here
  };

  // Mock oracle queries data
  const queries = [
    {
      id: 'query1',
      type: 'factCheck',
      prompt: 'Is the statement "NoMercyChain can process 10,000 transactions per second" accurate?',
      result: {
        factual: true,
        confidence: 0.92,
        sources: 2,
      },
      status: 'completed',
      timestamp: '2023-06-15 14:30',
      cost: '5 NMX',
    },
    {
      id: 'query2',
      type: 'dataFeed',
      prompt: 'Current ETH/USD price',
      result: {
        value: '1842.56',
        timestamp: '2023-06-15 14:25',
        sources: 5,
      },
      status: 'completed',
      timestamp: '2023-06-15 14:25',
      cost: '2 NMX',
    },
    {
      id: 'query3',
      type: 'prediction',
      prompt: 'What is the probability that ETH will exceed $2000 within the next 7 days?',
      result: {
        prediction: '78%',
        confidence: 0.85,
      },
      status: 'completed',
      timestamp: '2023-06-14 10:15',
      cost: '10 NMX',
    },
  ];

  // Mock oracle providers
  const providers = [
    {
      id: 'provider1',
      name: 'TruthGPT Core',
      description: 'Official NoMercyChain oracle service with high accuracy and reliability',
      accuracy: 0.98,
      queries: 12456,
      types: ['factCheck', 'dataFeed', 'prediction'],
      cost: 'Variable',
    },
    {
      id: 'provider2',
      name: 'CryptoData Oracle',
      description: 'Specialized in cryptocurrency and financial data feeds',
      accuracy: 0.95,
      queries: 8765,
      types: ['dataFeed'],
      cost: '1-5 NMX',
    },
    {
      id: 'provider3',
      name: 'PredictAI',
      description: 'Advanced AI models for predictive analytics and forecasting',
      accuracy: 0.92,
      queries: 6543,
      types: ['prediction'],
      cost: '5-15 NMX',
    },
  ];

  // Sample oracle integration code
  const sampleOracleCode = `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@nomercychain/oracle/TruthGPTConsumer.sol";

contract MyDApp is TruthGPTConsumer {
    // State variables
    mapping(bytes32 => bool) public factChecks;
    
    constructor(address oracleAddress) TruthGPTConsumer(oracleAddress) {
        // Initialize contract
    }
    
    // Request fact checking from the oracle
    function checkFact(string memory statement) external returns (bytes32) {
        bytes32 requestId = requestFactCheck(statement);
        return requestId;
    }
    
    // Oracle callback function
    function fulfillFactCheck(
        bytes32 requestId,
        bool isFactual,
        uint256 confidence
    ) internal override {
        factChecks[requestId] = isFactual;
        // Additional logic based on the result
    }
}`;

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          TruthGPT Oracle
        </Typography>
        <Button 
          variant="contained" 
          startIcon={<SearchIcon />}
          onClick={handleQueryDialogOpen}
          disabled={!account}
        >
          New Query
        </Button>
      </Box>

      <Grid container spacing={4}>
        <Grid item xs={12} md={8}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Decentralized AI Oracle Network
            </Typography>
            <Typography variant="body1" paragraph>
              TruthGPT is a decentralized oracle network powered by AI for verifying information, providing data feeds, and making predictions.
              It enables smart contracts to access external data and AI capabilities with high reliability and accuracy.
            </Typography>
            <Grid container spacing={3} sx={{ mt: 2 }}>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <FactCheckIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Fact Checking
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Verify information accuracy
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <DataObjectIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Data Feeds
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Real-time data from multiple sources
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <SmartToyIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Predictions
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    AI-powered forecasting
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </StyledPaper>
        </Grid>
        <Grid item xs={12} md={4}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Oracle Stats
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Total Queries
              </Typography>
              <Typography variant="h5">
                27,764
              </Typography>
            </Box>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Average Accuracy
              </Typography>
              <Typography variant="h5">
                96.8%
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Active Providers
              </Typography>
              <Typography variant="h5">
                12
              </Typography>
            </Box>
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
          <StyledTab label="Your Queries" />
          <StyledTab label="Oracle Providers" />
          <StyledTab label="Developer Guide" />
        </Tabs>

        {tabValue === 0 && (
          <>
            {account ? (
              <Grid container spacing={4}>
                {queries.map((query) => (
                  <Grid item xs={12} sm={6} md={4} key={query.id}>
                    <QueryCard>
                      <CardContent sx={{ flexGrow: 1 }}>
                        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                          <Avatar sx={{ 
                            bgcolor: query.type === 'factCheck' ? 'primary.main' : 
                                     query.type === 'dataFeed' ? 'secondary.main' : 'success.main',
                            mr: 2
                          }}>
                            {query.type === 'factCheck' ? <FactCheckIcon /> : 
                             query.type === 'dataFeed' ? <DataObjectIcon /> : <SmartToyIcon />}
                          </Avatar>
                          <Typography variant="h6">
                            {query.type === 'factCheck' ? 'Fact Check' : 
                             query.type === 'dataFeed' ? 'Data Feed' : 'Prediction'}
                          </Typography>
                        </Box>
                        <Chip 
                          label={query.status.charAt(0).toUpperCase() + query.status.slice(1)} 
                          size="small" 
                          color={query.status === 'completed' ? 'success' : 'warning'}
                          sx={{ mb: 2 }}
                        />
                        <Typography variant="body2" paragraph>
                          {query.prompt}
                        </Typography>
                        <Divider sx={{ my: 2 }} />
                        <Typography variant="body2" gutterBottom>
                          Result
                        </Typography>
                        {query.type === 'factCheck' && (
                          <Box sx={{ mb: 2 }}>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                              <Typography variant="caption" color="text.secondary">
                                Factual
                              </Typography>
                              <Chip 
                                label={query.result.factual ? 'True' : 'False'} 
                                size="small"
                                color={query.result.factual ? 'success' : 'error'}
                              />
                            </Box>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                              <Typography variant="caption" color="text.secondary">
                                Confidence
                              </Typography>
                              <Typography variant="body2">
                                {(query.result.confidence * 100).toFixed(1)}%
                              </Typography>
                            </Box>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                              <Typography variant="caption" color="text.secondary">
                                Sources
                              </Typography>
                              <Typography variant="body2">
                                {query.result.sources}
                              </Typography>
                            </Box>
                          </Box>
                        )}
                        {query.type === 'dataFeed' && (
                          <Box sx={{ mb: 2 }}>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                              <Typography variant="caption" color="text.secondary">
                                Value
                              </Typography>
                              <Typography variant="body2">
                                {query.result.value}
                              </Typography>
                            </Box>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                              <Typography variant="caption" color="text.secondary">
                                Timestamp
                              </Typography>
                              <Typography variant="body2">
                                {query.result.timestamp}
                              </Typography>
                            </Box>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                              <Typography variant="caption" color="text.secondary">
                                Sources
                              </Typography>
                              <Typography variant="body2">
                                {query.result.sources}
                              </Typography>
                            </Box>
                          </Box>
                        )}
                        {query.type === 'prediction' && (
                          <Box sx={{ mb: 2 }}>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                              <Typography variant="caption" color="text.secondary">
                                Prediction
                              </Typography>
                              <Typography variant="body2">
                                {query.result.prediction}
                              </Typography>
                            </Box>
                            <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                              <Typography variant="caption" color="text.secondary">
                                Confidence
                              </Typography>
                              <Typography variant="body2">
                                {(query.result.confidence * 100).toFixed(1)}%
                              </Typography>
                            </Box>
                          </Box>
                        )}
                        <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                          <Typography variant="caption" color="text.secondary">
                            Cost
                          </Typography>
                          <Typography variant="body2">
                            {query.cost}
                          </Typography>
                        </Box>
                        <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                          <Typography variant="caption" color="text.secondary">
                            Timestamp
                          </Typography>
                          <Typography variant="body2">
                            {query.timestamp}
                          </Typography>
                        </Box>
                      </CardContent>
                      <CardActions sx={{ p: 2 }}>
                        <Button size="small" variant="outlined" fullWidth>
                          View Details
                        </Button>
                      </CardActions>
                    </QueryCard>
                  </Grid>
                ))}
              </Grid>
            ) : (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  Connect Your Wallet
                </Typography>
                <Typography variant="body1" color="text.secondary" paragraph>
                  Please connect your wallet to view your oracle queries.
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
          <Grid container spacing={4}>
            {providers.map((provider) => (
              <Grid item xs={12} sm={6} key={provider.id}>
                <StyledPaper>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                      <VerifiedUserIcon />
                    </Avatar>
                    <Typography variant="h6">
                      {provider.name}
                    </Typography>
                  </Box>
                  <Typography variant="body2" color="text.secondary" paragraph>
                    {provider.description}
                  </Typography>
                  <Divider sx={{ my: 2 }} />
                  <Grid container spacing={2}>
                    <Grid item xs={6}>
                      <Typography variant="caption" color="text.secondary" display="block">
                        Accuracy
                      </Typography>
                      <Box sx={{ display: 'flex', alignItems: 'center' }}>
                        <Typography variant="body2" sx={{ mr: 1 }}>
                          {(provider.accuracy * 100).toFixed(1)}%
                        </Typography>
                        <Rating 
                          value={provider.accuracy * 5 / 100} 
                          precision={0.5} 
                          readOnly 
                          size="small" 
                        />
                      </Box>
                    </Grid>
                    <Grid item xs={6}>
                      <Typography variant="caption" color="text.secondary" display="block">
                        Total Queries
                      </Typography>
                      <Typography variant="body2">
                        {provider.queries.toLocaleString()}
                      </Typography>
                    </Grid>
                    <Grid item xs={6}>
                      <Typography variant="caption" color="text.secondary" display="block">
                        Query Types
                      </Typography>
                      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5, mt: 0.5 }}>
                        {provider.types.map((type) => (
                          <Chip 
                            key={type} 
                            label={type} 
                            size="small" 
                            sx={{ mr: 0.5, mb: 0.5 }} 
                          />
                        ))}
                      </Box>
                    </Grid>
                    <Grid item xs={6}>
                      <Typography variant="caption" color="text.secondary" display="block">
                        Cost
                      </Typography>
                      <Typography variant="body2">
                        {provider.cost}
                      </Typography>
                    </Grid>
                  </Grid>
                  <Box sx={{ mt: 2, display: 'flex', justifyContent: 'flex-end' }}>
                    <Button 
                      variant="outlined" 
                      size="small"
                      onClick={() => {
                        setQueryDialogOpen(true);
                      }}
                      disabled={!account}
                    >
                      Use Provider
                    </Button>
                  </Box>
                </StyledPaper>
              </Grid>
            ))}
          </Grid>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Oracle Integration Guide
            </Typography>
            <Typography variant="body2" color="text.secondary" paragraph>
              Integrate TruthGPT Oracle into your smart contracts to access AI capabilities and external data.
            </Typography>
            <Typography variant="subtitle1" gutterBottom sx={{ mt: 3 }}>
              Sample Contract
            </Typography>
            <CodeBlock>
              <Box sx={{ position: 'absolute', top: 8, right: 8 }}>
                <IconButton 
                  size="small" 
                  onClick={() => copyToClipboard(sampleOracleCode)}
                  sx={{ color: 'text.secondary' }}
                >
                  <ContentCopyIcon fontSize="small" />
                </IconButton>
              </Box>
              <pre style={{ margin: 0 }}>
                {sampleOracleCode}
              </pre>
            </CodeBlock>
            <Typography variant="subtitle1" gutterBottom sx={{ mt: 4 }}>
              Integration Steps
            </Typography>
            <ol>
              <li>
                <Typography variant="body2" paragraph>
                  Import the TruthGPTConsumer contract from the NoMercyChain oracle package
                </Typography>
              </li>
              <li>
                <Typography variant="body2" paragraph>
                  Inherit from TruthGPTConsumer and pass the oracle address to the constructor
                </Typography>
              </li>
              <li>
                <Typography variant="body2" paragraph>
                  Use the provided methods to request data: requestFactCheck(), requestDataFeed(), or requestPrediction()
                </Typography>
              </li>
              <li>
                <Typography variant="body2" paragraph>
                  Override the callback functions to handle the oracle responses
                </Typography>
              </li>
            </ol>
            <Typography variant="subtitle1" gutterBottom sx={{ mt: 4 }}>
              Available Query Types
            </Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <StyledTableCell>Query Type</StyledTableCell>
                    <StyledTableCell>Description</StyledTableCell>
                    <StyledTableCell>Method</StyledTableCell>
                    <StyledTableCell>Callback</StyledTableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  <TableRow>
                    <StyledTableCell>Fact Check</StyledTableCell>
                    <StyledTableCell>Verify the factual accuracy of a statement</StyledTableCell>
                    <StyledTableCell>requestFactCheck(string)</StyledTableCell>
                    <StyledTableCell>fulfillFactCheck(bytes32, bool, uint256)</StyledTableCell>
                  </TableRow>
                  <TableRow>
                    <StyledTableCell>Data Feed</StyledTableCell>
                    <StyledTableCell>Get real-time data from external sources</StyledTableCell>
                    <StyledTableCell>requestDataFeed(string)</StyledTableCell>
                    <StyledTableCell>fulfillDataFeed(bytes32, uint256, uint256)</StyledTableCell>
                  </TableRow>
                  <TableRow>
                    <StyledTableCell>Prediction</StyledTableCell>
                    <StyledTableCell>Get AI-powered predictions for future events</StyledTableCell>
                    <StyledTableCell>requestPrediction(string)</StyledTableCell>
                    <StyledTableCell>fulfillPrediction(bytes32, uint256, uint256)</StyledTableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
          </StyledPaper>
        )}
      </Box>

      {/* Query Dialog */}
      <Dialog open={queryDialogOpen} onClose={handleQueryDialogClose} maxWidth="md" fullWidth>
        <DialogTitle>
          New Oracle Query
        </DialogTitle>
        <DialogContent>
          {!showResult ? (
            <>
              <FormControl fullWidth margin="normal">
                <InputLabel>Query Type</InputLabel>
                <Select
                  value={queryType}
                  label="Query Type"
                  onChange={(e) => setQueryType(e.target.value)}
                >
                  <MenuItem value="factCheck">Fact Check</MenuItem>
                  <MenuItem value="dataFeed">Data Feed</MenuItem>
                  <MenuItem value="prediction">Prediction</MenuItem>
                </Select>
              </FormControl>
              {queryType === 'factCheck' && (
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                  Submit a statement to verify its factual accuracy. The oracle will check multiple sources and return whether the statement is true or false.
                </Typography>
              )}
              {queryType === 'dataFeed' && (
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                  Request real-time data from external sources. The oracle will aggregate data from multiple providers and return the result.
                </Typography>
              )}
              {queryType === 'prediction' && (
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                  Get AI-powered predictions for future events or outcomes. The oracle will analyze historical data and current conditions to make a forecast.
                </Typography>
              )}
              <TextField
                label="Query Prompt"
                fullWidth
                margin="normal"
                variant="outlined"
                multiline
                rows={4}
                value={queryPrompt}
                onChange={(e) => setQueryPrompt(e.target.value)}
                placeholder={
                  queryType === 'factCheck' ? 'Enter a statement to verify (e.g., "NoMercyChain can process 10,000 transactions per second")' :
                  queryType === 'dataFeed' ? 'Specify the data you need (e.g., "Current ETH/USD price")' :
                  queryType === 'prediction' ? 'Describe the prediction you need (e.g., "What is the probability that ETH will exceed $2000 within the next 7 days?")' :
                  'Enter your query prompt'
                }
              />
              <FormControl fullWidth margin="normal">
                <InputLabel>Oracle Provider</InputLabel>
                <Select
                  value="provider1"
                  label="Oracle Provider"
                >
                  <MenuItem value="provider1">TruthGPT Core</MenuItem>
                  <MenuItem value="provider2">CryptoData Oracle</MenuItem>
                  <MenuItem value="provider3">PredictAI</MenuItem>
                </Select>
              </FormControl>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                <Typography variant="body2" color="text.secondary">
                  Estimated Cost
                </Typography>
                <Typography variant="body2">
                  {queryType === 'factCheck' ? '5 NMX' : 
                   queryType === 'dataFeed' ? '2 NMX' : 
                   queryType === 'prediction' ? '10 NMX' : 
                   '0 NMX'}
                </Typography>
              </Box>
            </>
          ) : (
            <>
              <Typography variant="body2" color="text.secondary" paragraph>
                Query Result
              </Typography>
              <CodeBlock>
                <Box sx={{ position: 'absolute', top: 8, right: 8 }}>
                  <IconButton 
                    size="small" 
                    onClick={() => copyToClipboard(queryResult)}
                    sx={{ color: 'text.secondary' }}
                  >
                    <ContentCopyIcon fontSize="small" />
                  </IconButton>
                </Box>
                <pre style={{ margin: 0 }}>
                  {queryResult}
                </pre>
              </CodeBlock>
            </>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleQueryDialogClose}>
            {showResult ? 'Close' : 'Cancel'}
          </Button>
          {!showResult && (
            <Button 
              variant="contained" 
              onClick={handleSubmitQuery}
              disabled={!queryType || !queryPrompt}
            >
              Submit Query
            </Button>
          )}
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default Oracle;