import AddIcon from '@mui/icons-material/Add';
import PauseIcon from '@mui/icons-material/Pause';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import SettingsIcon from '@mui/icons-material/Settings';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import StorefrontIcon from '@mui/icons-material/Storefront';
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
    LinearProgress,
    MenuItem,
    Paper,
    Select,
    Tab,
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

const AgentCard = styled(Card)(({ theme }) => ({
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

const AgentAvatar = styled(Avatar)(({ theme, agentType }) => ({
  width: 56,
  height: 56,
  margin: '0 auto 16px',
  background: agentType === 'trading' ? 'linear-gradient(135deg, #6200ea 0%, #3700b3 100%)' :
              agentType === 'governance' ? 'linear-gradient(135deg, #00e676 0%, #00c853 100%)' :
              agentType === 'staking' ? 'linear-gradient(135deg, #ff9800 0%, #f57c00 100%)' :
              'linear-gradient(135deg, #f50057 0%, #c51162 100%)',
}));

const StyledTab = styled(Tab)(({ theme }) => ({
  textTransform: 'none',
  fontWeight: 600,
  fontSize: '1rem',
}));

const AIAgents = () => {
  const { account } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [marketplaceDialogOpen, setMarketplaceDialogOpen] = useState(false);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  // Mock AI agents data
  const myAgents = [
    {
      id: 'agent1',
      name: 'Trading Bot Alpha',
      description: 'Automated trading agent for DeFi protocols',
      type: 'trading',
      status: 'active',
      createdAt: '2023-05-10',
      intelligence: 85,
      tasks: 124,
      successRate: 92,
    },
    {
      id: 'agent2',
      name: 'Governance Advisor',
      description: 'Analyzes proposals and suggests voting decisions',
      type: 'governance',
      status: 'inactive',
      createdAt: '2023-04-22',
      intelligence: 78,
      tasks: 45,
      successRate: 89,
    },
    {
      id: 'agent3',
      name: 'Yield Optimizer',
      description: 'Automatically manages staking positions for optimal returns',
      type: 'staking',
      status: 'active',
      createdAt: '2023-06-01',
      intelligence: 92,
      tasks: 67,
      successRate: 95,
    },
  ];

  // Mock marketplace agents
  const marketplaceAgents = [
    {
      id: 'market1',
      name: 'DeFi Master',
      description: 'Advanced trading bot with multi-protocol support',
      type: 'trading',
      creator: '0x1234...5678',
      price: '500 NMX',
      rating: 4.8,
      sales: 156,
    },
    {
      id: 'market2',
      name: 'Proposal Analyzer Pro',
      description: 'Deep analysis of governance proposals with ML-based predictions',
      type: 'governance',
      creator: '0x8765...4321',
      price: '350 NMX',
      rating: 4.6,
      sales: 89,
    },
    {
      id: 'market3',
      name: 'Staking Sentinel',
      description: 'Monitors validator performance and automatically redistributes stakes',
      type: 'staking',
      creator: '0x2468...1357',
      price: '450 NMX',
      rating: 4.9,
      sales: 212,
    },
    {
      id: 'market4',
      name: 'NFT Trader',
      description: 'Analyzes NFT markets and executes trades based on price predictions',
      type: 'trading',
      creator: '0x1357...2468',
      price: '600 NMX',
      rating: 4.7,
      sales: 78,
    },
  ];

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          AI Agents
        </Typography>
        <Box>
          <Button 
            variant="contained" 
            startIcon={<StorefrontIcon />} 
            sx={{ mr: 2 }}
            onClick={() => setMarketplaceDialogOpen(true)}
          >
            Marketplace
          </Button>
          <Button 
            variant="contained" 
            startIcon={<AddIcon />}
            onClick={() => setCreateDialogOpen(true)}
          >
            Create Agent
          </Button>
        </Box>
      </Box>

      <StyledPaper sx={{ mb: 4 }}>
        <Typography variant="h6" gutterBottom>
          What are AI Agents?
        </Typography>
        <Typography variant="body1" paragraph>
          AI Agents are autonomous digital entities that can perform tasks on your behalf on the NoMercyChain blockchain. 
          They can trade assets, participate in governance, manage staking positions, and more.
        </Typography>
        <Typography variant="body1">
          Each agent is powered by advanced AI models and can be trained to suit your specific needs. 
          You can create your own agents or purchase pre-trained ones from the marketplace.
        </Typography>
      </StyledPaper>

      <Box sx={{ mb: 4 }}>
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
          <StyledTab label="My Agents" />
          <StyledTab label="Agent Activity" />
          <StyledTab label="Training Data" />
        </Tabs>

        {tabValue === 0 && (
          <>
            {account ? (
              <Grid container spacing={4}>
                {myAgents.map((agent) => (
                  <Grid item xs={12} sm={6} md={4} key={agent.id}>
                    <AgentCard>
                      <CardContent sx={{ flexGrow: 1 }}>
                        <AgentAvatar agentType={agent.type}>
                          <SmartToyIcon />
                        </AgentAvatar>
                        <Typography variant="h6" align="center" gutterBottom>
                          {agent.name}
                        </Typography>
                        <Chip 
                          label={agent.type} 
                          size="small" 
                          sx={{ 
                            mb: 2,
                            bgcolor: agent.type === 'trading' ? 'primary.main' : 
                                    agent.type === 'governance' ? 'success.main' : 
                                    agent.type === 'staking' ? 'warning.main' : 'error.main',
                            color: 'white'
                          }} 
                        />
                        <Typography variant="body2" color="text.secondary" paragraph>
                          {agent.description}
                        </Typography>
                        <Divider sx={{ my: 2 }} />
                        <Grid container spacing={2}>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Intelligence
                            </Typography>
                            <Typography variant="body2">
                              {agent.intelligence}%
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Tasks
                            </Typography>
                            <Typography variant="body2">
                              {agent.tasks}
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Success Rate
                            </Typography>
                            <Typography variant="body2">
                              {agent.successRate}%
                            </Typography>
                          </Grid>
                        </Grid>
                        <Box sx={{ mt: 2 }}>
                          <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                            Status: {agent.status === 'active' ? 'Active' : 'Inactive'}
                          </Typography>
                          <LinearProgress 
                            variant="determinate" 
                            value={agent.status === 'active' ? 100 : 0} 
                            sx={{ 
                              height: 6, 
                              borderRadius: 3,
                              bgcolor: 'rgba(255,255,255,0.1)',
                              '& .MuiLinearProgress-bar': {
                                bgcolor: agent.status === 'active' ? 'success.main' : 'text.disabled',
                              }
                            }} 
                          />
                        </Box>
                      </CardContent>
                      <CardActions sx={{ justifyContent: 'space-between', p: 2 }}>
                        <Box>
                          <IconButton 
                            size="small" 
                            sx={{ 
                              mr: 1,
                              color: agent.status === 'active' ? 'error.main' : 'success.main'
                            }}
                          >
                            {agent.status === 'active' ? <PauseIcon /> : <PlayArrowIcon />}
                          </IconButton>
                          <IconButton size="small">
                            <SettingsIcon />
                          </IconButton>
                        </Box>
                        <Button size="small" variant="outlined">
                          View Details
                        </Button>
                      </CardActions>
                    </AgentCard>
                  </Grid>
                ))}
              </Grid>
            ) : (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  Connect Your Wallet
                </Typography>
                <Typography variant="body1" color="text.secondary" paragraph>
                  Please connect your wallet to view and manage your AI agents.
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
              Recent Agent Activity
            </Typography>
            <Divider sx={{ my: 2 }} />
            {account ? (
              <Box>
                <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                    <SmartToyIcon />
                  </Avatar>
                  <Box sx={{ flexGrow: 1 }}>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                      <Typography variant="body1">
                        Trading Bot Alpha
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        10 minutes ago
                      </Typography>
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      Executed swap: 50 NMX → 0.025 ETH
                    </Typography>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
                  <Avatar sx={{ bgcolor: 'success.main', mr: 2 }}>
                    <SmartToyIcon />
                  </Avatar>
                  <Box sx={{ flexGrow: 1 }}>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                      <Typography variant="body1">
                        Governance Advisor
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        2 hours ago
                      </Typography>
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      Analyzed proposal #45: Recommended 'Yes' vote
                    </Typography>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                  <Avatar sx={{ bgcolor: 'warning.main', mr: 2 }}>
                    <SmartToyIcon />
                  </Avatar>
                  <Box sx={{ flexGrow: 1 }}>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                      <Typography variant="body1">
                        Yield Optimizer
                      </Typography>
                      <Typography variant="caption" color="text.secondary">
                        5 hours ago
                      </Typography>
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      Restaked 100 NMX to validator 'Cosmos Sentinel'
                    </Typography>
                  </Box>
                </Box>
              </Box>
            ) : (
              <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                Connect your wallet to see your agent activity
              </Typography>
            )}
          </StyledPaper>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Training Data
            </Typography>
            <Divider sx={{ my: 2 }} />
            {account ? (
              <Box>
                <Typography variant="body1" paragraph>
                  Improve your AI agents by providing training data. The more data you provide, the more intelligent your agents become.
                </Typography>
                <Button variant="contained" startIcon={<AddIcon />}>
                  Upload Training Data
                </Button>
              </Box>
            ) : (
              <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                Connect your wallet to manage training data
              </Typography>
            )}
          </StyledPaper>
        )}
      </Box>

      {/* Create Agent Dialog */}
      <Dialog open={createDialogOpen} onClose={() => setCreateDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create New AI Agent</DialogTitle>
        <DialogContent>
          <TextField
            label="Agent Name"
            fullWidth
            margin="normal"
            variant="outlined"
            placeholder="Enter a name for your agent"
          />
          <TextField
            label="Description"
            fullWidth
            margin="normal"
            variant="outlined"
            placeholder="Describe what your agent will do"
            multiline
            rows={3}
          />
          <FormControl fullWidth margin="normal">
            <InputLabel>Agent Type</InputLabel>
            <Select
              label="Agent Type"
              defaultValue=""
            >
              <MenuItem value="trading">Trading</MenuItem>
              <MenuItem value="governance">Governance</MenuItem>
              <MenuItem value="staking">Staking</MenuItem>
              <MenuItem value="gaming">Gaming</MenuItem>
              <MenuItem value="custom">Custom</MenuItem>
            </Select>
          </FormControl>
          <FormControl fullWidth margin="normal">
            <InputLabel>AI Model</InputLabel>
            <Select
              label="AI Model"
              defaultValue=""
            >
              <MenuItem value="basic">Basic (Free)</MenuItem>
              <MenuItem value="advanced">Advanced (50 NMX)</MenuItem>
              <MenuItem value="expert">Expert (150 NMX)</MenuItem>
              <MenuItem value="custom">Custom Upload</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialogOpen(false)}>Cancel</Button>
          <Button variant="contained">Create Agent</Button>
        </DialogActions>
      </Dialog>

      {/* Marketplace Dialog */}
      <Dialog open={marketplaceDialogOpen} onClose={() => setMarketplaceDialogOpen(false)} maxWidth="md" fullWidth>
        <DialogTitle>AI Agent Marketplace</DialogTitle>
        <DialogContent>
          <Grid container spacing={3}>
            {marketplaceAgents.map((agent) => (
              <Grid item xs={12} sm={6} key={agent.id}>
                <Card sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
                  <CardContent>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                      <Avatar 
                        sx={{ 
                          bgcolor: agent.type === 'trading' ? 'primary.main' : 
                                  agent.type === 'governance' ? 'success.main' : 
                                  agent.type === 'staking' ? 'warning.main' : 'error.main',
                          mr: 2
                        }}
                      >
                        <SmartToyIcon />
                      </Avatar>
                      <Box>
                        <Typography variant="h6">{agent.name}</Typography>
                        <Chip 
                          label={agent.type} 
                          size="small" 
                          sx={{ 
                            bgcolor: agent.type === 'trading' ? 'primary.main' : 
                                    agent.type === 'governance' ? 'success.main' : 
                                    agent.type === 'staking' ? 'warning.main' : 'error.main',
                            color: 'white'
                          }} 
                        />
                      </Box>
                    </Box>
                    <Typography variant="body2" color="text.secondary" paragraph>
                      {agent.description}
                    </Typography>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
                      <Typography variant="body2">
                        Creator: {agent.creator}
                      </Typography>
                      <Typography variant="body2">
                        Rating: {agent.rating}/5 ({agent.sales} sales)
                      </Typography>
                    </Box>
                  </CardContent>
                  <Box sx={{ flexGrow: 1 }} />
                  <CardActions sx={{ justifyContent: 'space-between', p: 2 }}>
                    <Typography variant="h6" color="primary">
                      {agent.price}
                    </Typography>
                    <Button variant="contained" size="small">
                      Purchase
                    </Button>
                  </CardActions>
                </Card>
              </Grid>
            ))}
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setMarketplaceDialogOpen(false)}>Close</Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default AIAgents;