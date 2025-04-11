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
  Stepper,
  Step,
  StepLabel
} from '@mui/material';
import { styled } from '@mui/material/styles';
import LayersIcon from '@mui/icons-material/Layers';
import AddIcon from '@mui/icons-material/Add';
import SettingsIcon from '@mui/icons-material/Settings';
import StorageIcon from '@mui/icons-material/Storage';
import AutoFixHighIcon from '@mui/icons-material/AutoFixHigh';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import BarChartIcon from '@mui/icons-material/BarChart';
import RocketLaunchIcon from '@mui/icons-material/RocketLaunch';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ChainCard = styled(Card)(({ theme }) => ({
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

const HyperChains = () => {
  const { account, balance } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [activeStep, setActiveStep] = useState(0);
  const [chainName, setChainName] = useState('');
  const [chainDescription, setChainDescription] = useState('');
  const [chainType, setChainType] = useState('');
  const [chainModules, setChainModules] = useState([]);
  const [chainPrompt, setChainPrompt] = useState('');

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleCreateDialogOpen = () => {
    setCreateDialogOpen(true);
  };

  const handleCreateDialogClose = () => {
    setCreateDialogOpen(false);
    setActiveStep(0);
    setChainName('');
    setChainDescription('');
    setChainType('');
    setChainModules([]);
    setChainPrompt('');
  };

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  const handleCreateChain = () => {
    // Mock chain creation
    console.log(`Created HyperChain: ${chainName}`);
    handleCreateDialogClose();
  };

  const handleModuleChange = (event) => {
    const {
      target: { value },
    } = event;
    setChainModules(
      // On autofill we get a stringified value.
      typeof value === 'string' ? value.split(',') : value,
    );
  };

  // Mock hyperchains data
  const hyperchains = [
    {
      id: 'chain1',
      name: 'DefiChain',
      description: 'Specialized chain for DeFi applications with optimized transaction processing',
      type: 'application',
      status: 'active',
      createdAt: '2023-06-10',
      validators: 15,
      transactions: 1245,
      tps: 500,
      modules: ['defi', 'token', 'oracle'],
    },
    {
      id: 'chain2',
      name: 'GameVerse',
      description: 'Gaming-focused chain with NFT support and low latency',
      type: 'application',
      status: 'active',
      createdAt: '2023-05-22',
      validators: 10,
      transactions: 876,
      tps: 450,
      modules: ['nft', 'token', 'marketplace'],
    },
    {
      id: 'chain3',
      name: 'DataChain',
      description: 'Enterprise data chain with privacy features and high throughput',
      type: 'enterprise',
      status: 'active',
      createdAt: '2023-06-01',
      validators: 5,
      transactions: 567,
      tps: 600,
      modules: ['data', 'privacy', 'identity'],
    },
  ];

  // Available modules
  const availableModules = [
    'token', 'defi', 'nft', 'marketplace', 'oracle', 'governance', 'staking', 
    'identity', 'privacy', 'data', 'ai', 'gaming'
  ];

  // Chain types
  const chainTypes = [
    { value: 'application', label: 'Application Chain' },
    { value: 'enterprise', label: 'Enterprise Chain' },
    { value: 'gaming', label: 'Gaming Chain' },
    { value: 'defi', label: 'DeFi Chain' },
    { value: 'data', label: 'Data Chain' },
  ];

  // Steps for creating a HyperChain
  const steps = ['Basic Information', 'Chain Configuration', 'AI Prompt', 'Review'];

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          HyperChains
        </Typography>
        <Button 
          variant="contained" 
          startIcon={<AddIcon />}
          onClick={handleCreateDialogOpen}
          disabled={!account}
        >
          Create HyperChain
        </Button>
      </Box>

      <Grid container spacing={4}>
        <Grid item xs={12} md={8}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              HyperChains: Layer 3 Solutions
            </Typography>
            <Typography variant="body1" paragraph>
              HyperChains are specialized Layer 3 blockchains that can be created on top of NoMercyChain.
              They provide customized environments for specific applications, with their own validators, consensus mechanisms, and governance.
            </Typography>
            <Grid container spacing={3} sx={{ mt: 2 }}>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <RocketLaunchIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Scalable
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Up to 10,000 TPS per chain
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <AutoFixHighIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    Customizable
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Tailored for specific use cases
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <SmartToyIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    AI-Generated
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Created using natural language
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </StyledPaper>
        </Grid>
        <Grid item xs={12} md={4}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Network Stats
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Total HyperChains
              </Typography>
              <Typography variant="h5">
                28
              </Typography>
            </Box>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Active Validators
              </Typography>
              <Typography variant="h5">
                156
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Total TPS
              </Typography>
              <Typography variant="h5">
                12,450
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
          <StyledTab label="Your HyperChains" />
          <StyledTab label="Explorer" />
          <StyledTab label="Analytics" />
        </Tabs>

        {tabValue === 0 && (
          <>
            {account ? (
              <Grid container spacing={4}>
                {hyperchains.map((chain) => (
                  <Grid item xs={12} sm={6} key={chain.id}>
                    <ChainCard>
                      <CardContent sx={{ flexGrow: 1 }}>
                        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                          <Avatar sx={{ 
                            bgcolor: chain.type === 'application' ? 'primary.main' : 'secondary.main',
                            mr: 2
                          }}>
                            <LayersIcon />
                          </Avatar>
                          <Typography variant="h6">
                            {chain.name}
                          </Typography>
                        </Box>
                        <Chip 
                          label={chain.type.charAt(0).toUpperCase() + chain.type.slice(1)} 
                          size="small" 
                          sx={{ 
                            mb: 2,
                            bgcolor: chain.type === 'application' ? 'primary.main' : 'secondary.main',
                            color: 'white'
                          }} 
                        />
                        <Typography variant="body2" color="text.secondary" paragraph>
                          {chain.description}
                        </Typography>
                        <Box sx={{ mb: 2 }}>
                          <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                            Modules
                          </Typography>
                          <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                            {chain.modules.map((module) => (
                              <Chip 
                                key={module} 
                                label={module} 
                                size="small" 
                                sx={{ mr: 0.5, mb: 0.5 }} 
                              />
                            ))}
                          </Box>
                        </Box>
                        <Divider sx={{ my: 2 }} />
                        <Grid container spacing={2}>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Validators
                            </Typography>
                            <Typography variant="body2">
                              {chain.validators}
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Transactions
                            </Typography>
                            <Typography variant="body2">
                              {chain.transactions}
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              TPS
                            </Typography>
                            <Typography variant="body2">
                              {chain.tps}
                            </Typography>
                          </Grid>
                        </Grid>
                        <Box sx={{ mt: 2 }}>
                          <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                            Status: {chain.status === 'active' ? 'Active' : 'Inactive'}
                          </Typography>
                          <LinearProgress 
                            variant="determinate" 
                            value={chain.status === 'active' ? 100 : 0} 
                            sx={{ 
                              height: 6, 
                              borderRadius: 3,
                              bgcolor: 'rgba(255,255,255,0.1)',
                              '& .MuiLinearProgress-bar': {
                                bgcolor: chain.status === 'active' ? 'success.main' : 'text.disabled',
                              }
                            }} 
                          />
                        </Box>
                      </CardContent>
                      <CardActions sx={{ justifyContent: 'space-between', p: 2 }}>
                        <IconButton size="small">
                          <SettingsIcon />
                        </IconButton>
                        <Button size="small" variant="outlined">
                          Manage
                        </Button>
                      </CardActions>
                    </ChainCard>
                  </Grid>
                ))}
              </Grid>
            ) : (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  Connect Your Wallet
                </Typography>
                <Typography variant="body1" color="text.secondary" paragraph>
                  Please connect your wallet to view and manage your HyperChains.
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
              HyperChain Explorer
            </Typography>
            <Divider sx={{ my: 2 }} />
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <StyledTableCell>Name</StyledTableCell>
                    <StyledTableCell>Type</StyledTableCell>
                    <StyledTableCell>Validators</StyledTableCell>
                    <StyledTableCell>TPS</StyledTableCell>
                    <StyledTableCell>Status</StyledTableCell>
                    <StyledTableCell align="right">Actions</StyledTableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {hyperchains.map((chain) => (
                    <TableRow key={chain.id}>
                      <StyledTableCell>
                        <Box sx={{ display: 'flex', alignItems: 'center' }}>
                          <Avatar sx={{ 
                            bgcolor: chain.type === 'application' ? 'primary.main' : 'secondary.main',
                            mr: 2,
                            width: 32,
                            height: 32
                          }}>
                            <LayersIcon fontSize="small" />
                          </Avatar>
                          <Box>
                            <Typography variant="body2">{chain.name}</Typography>
                            <Typography variant="caption" color="text.secondary">{chain.createdAt}</Typography>
                          </Box>
                        </Box>
                      </StyledTableCell>
                      <StyledTableCell>
                        <Chip 
                          label={chain.type.charAt(0).toUpperCase() + chain.type.slice(1)} 
                          size="small" 
                        />
                      </StyledTableCell>
                      <StyledTableCell>{chain.validators}</StyledTableCell>
                      <StyledTableCell>{chain.tps}</StyledTableCell>
                      <StyledTableCell>
                        <Chip 
                          label={chain.status.charAt(0).toUpperCase() + chain.status.slice(1)} 
                          size="small"
                          color={chain.status === 'active' ? 'success' : 'default'}
                        />
                      </StyledTableCell>
                      <StyledTableCell align="right">
                        <Button size="small">Explore</Button>
                      </StyledTableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </StyledPaper>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              HyperChain Analytics
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Grid container spacing={3}>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1 }}>
                    <BarChartIcon />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    28
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Total HyperChains
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1 }}>
                    <StorageIcon />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    12,450
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Transactions Per Second
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1 }}>
                    <LayersIcon />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    156
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Active Validators
                  </Typography>
                </Box>
              </Grid>
            </Grid>
            <Box sx={{ mt: 4, height: 300, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
              <Typography variant="body1" color="text.secondary">
                Analytics charts will be displayed here
              </Typography>
            </Box>
          </StyledPaper>
        )}
      </Box>

      {/* Create HyperChain Dialog */}
      <Dialog open={createDialogOpen} onClose={handleCreateDialogClose} maxWidth="md" fullWidth>
        <DialogTitle>
          Create HyperChain
        </DialogTitle>
        <DialogContent>
          <Stepper activeStep={activeStep} sx={{ mt: 2, mb: 4 }}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          
          {activeStep === 0 && (
            <Box>
              <Typography variant="body2" color="text.secondary" paragraph>
                Enter the basic information for your HyperChain.
              </Typography>
              <TextField
                label="Chain Name"
                fullWidth
                margin="normal"
                variant="outlined"
                value={chainName}
                onChange={(e) => setChainName(e.target.value)}
              />
              <TextField
                label="Description"
                fullWidth
                margin="normal"
                variant="outlined"
                multiline
                rows={3}
                value={chainDescription}
                onChange={(e) => setChainDescription(e.target.value)}
              />
              <FormControl fullWidth margin="normal">
                <InputLabel>Chain Type</InputLabel>
                <Select
                  value={chainType}
                  label="Chain Type"
                  onChange={(e) => setChainType(e.target.value)}
                >
                  {chainTypes.map((type) => (
                    <MenuItem key={type.value} value={type.value}>{type.label}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Box>
          )}

          {activeStep === 1 && (
            <Box>
              <Typography variant="body2" color="text.secondary" paragraph>
                Configure the modules and parameters for your HyperChain.
              </Typography>
              <FormControl fullWidth margin="normal">
                <InputLabel>Modules</InputLabel>
                <Select
                  multiple
                  value={chainModules}
                  onChange={handleModuleChange}
                  label="Modules"
                  renderValue={(selected) => (
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                      {selected.map((value) => (
                        <Chip key={value} label={value} />
                      ))}
                    </Box>
                  )}
                >
                  {availableModules.map((module) => (
                    <MenuItem key={module} value={module}>
                      {module}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
              <Grid container spacing={2} sx={{ mt: 1 }}>
                <Grid item xs={12} sm={6}>
                  <FormControl fullWidth margin="normal">
                    <InputLabel>Validators</InputLabel>
                    <Select
                      value="10"
                      label="Validators"
                    >
                      <MenuItem value="5">5</MenuItem>
                      <MenuItem value="10">10</MenuItem>
                      <MenuItem value="15">15</MenuItem>
                      <MenuItem value="20">20</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} sm={6}>
                  <FormControl fullWidth margin="normal">
                    <InputLabel>Consensus</InputLabel>
                    <Select
                      value="pos"
                      label="Consensus"
                    >
                      <MenuItem value="pos">Proof of Stake</MenuItem>
                      <MenuItem value="poa">Proof of Authority</MenuItem>
                      <MenuItem value="dpos">Delegated Proof of Stake</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
              </Grid>
            </Box>
          )}

          {activeStep === 2 && (
            <Box>
              <Typography variant="body2" color="text.secondary" paragraph>
                Describe your HyperChain in natural language. Our AI will generate the optimal configuration based on your description.
              </Typography>
              <TextField
                label="AI Prompt"
                fullWidth
                margin="normal"
                variant="outlined"
                multiline
                rows={6}
                placeholder="Example: Create a high-performance chain for DeFi applications with support for automated market makers, lending protocols, and cross-chain bridges. It should prioritize security and transaction throughput."
                value={chainPrompt}
                onChange={(e) => setChainPrompt(e.target.value)}
              />
              <Typography variant="caption" color="text.secondary">
                The more detailed your description, the better the AI can optimize your HyperChain.
              </Typography>
            </Box>
          )}

          {activeStep === 3 && (
            <Box>
              <Typography variant="body2" color="text.secondary" paragraph>
                Review your HyperChain configuration before creation.
              </Typography>
              <StyledPaper sx={{ p: 2, mb: 2 }}>
                <Grid container spacing={2}>
                  <Grid item xs={12} sm={6}>
                    <Typography variant="caption" color="text.secondary" display="block">
                      Chain Name
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                      {chainName}
                    </Typography>
                  </Grid>
                  <Grid item xs={12} sm={6}>
                    <Typography variant="caption" color="text.secondary" display="block">
                      Chain Type
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                      {chainType && chainType.charAt(0).toUpperCase() + chainType.slice(1)}
                    </Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography variant="caption" color="text.secondary" display="block">
                      Description
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                      {chainDescription}
                    </Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography variant="caption" color="text.secondary" display="block">
                      Modules
                    </Typography>
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5, mt: 1 }}>
                      {chainModules.map((module) => (
                        <Chip key={module} label={module} size="small" />
                      ))}
                    </Box>
                  </Grid>
                </Grid>
              </StyledPaper>
              <Typography variant="body2" color="text.secondary" paragraph>
                Creating a HyperChain requires a deposit of 10,000 NMX tokens. This deposit is used to secure the chain and can be reclaimed if you decide to shut down the chain.
              </Typography>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                <Typography variant="body2" color="text.secondary">
                  Required Deposit
                </Typography>
                <Typography variant="body2">
                  10,000 NMX
                </Typography>
              </Box>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
                <Typography variant="body2" color="text.secondary">
                  Your Balance
                </Typography>
                <Typography variant="body2">
                  {balance || 0} NMX
                </Typography>
              </Box>
            </Box>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCreateDialogClose}>Cancel</Button>
          {activeStep > 0 && (
            <Button onClick={handleBack}>
              Back
            </Button>
          )}
          {activeStep < steps.length - 1 ? (
            <Button 
              variant="contained" 
              onClick={handleNext}
              disabled={
                (activeStep === 0 && (!chainName || !chainDescription || !chainType)) ||
                (activeStep === 1 && chainModules.length === 0) ||
                (activeStep === 2 && !chainPrompt)
              }
            >
              Next
            </Button>
          ) : (
            <Button 
              variant="contained" 
              onClick={handleCreateChain}
              disabled={!account || (balance && parseFloat(balance) < 10000)}
            >
              Create HyperChain
            </Button>
          )}
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default HyperChains;