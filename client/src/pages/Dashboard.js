import React, { useContext } from 'react';
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
  LinearProgress,
  Avatar
} from '@mui/material';
import { styled } from '@mui/material/styles';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import PeopleIcon from '@mui/icons-material/People';
import LayersIcon from '@mui/icons-material/Layers';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import { WalletContext } from '../context/WalletContext';
import { Link } from 'react-router-dom';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  height: '100%',
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const StatCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(2),
  display: 'flex',
  alignItems: 'center',
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const IconWrapper = styled(Box)(({ theme, color }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  width: 56,
  height: 56,
  borderRadius: 12,
  marginRight: theme.spacing(2),
  backgroundColor: color,
}));

const FeatureCard = styled(Card)(({ theme }) => ({
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const Dashboard = () => {
  const { account, connectWallet } = useContext(WalletContext);

  // Mock data
  const stats = [
    { 
      title: 'NMX Price', 
      value: '$2.45', 
      change: '+5.2%', 
      icon: <TrendingUpIcon />, 
      color: 'rgba(98, 0, 234, 0.2)' 
    },
    { 
      title: 'Active Users', 
      value: '12,456', 
      change: '+12.3%', 
      icon: <PeopleIcon />, 
      color: 'rgba(0, 230, 118, 0.2)' 
    },
    { 
      title: 'HyperChains', 
      value: '28', 
      change: '+3', 
      icon: <LayersIcon />, 
      color: 'rgba(255, 87, 34, 0.2)' 
    },
    { 
      title: 'AI Agents', 
      value: '5,234', 
      change: '+432', 
      icon: <SmartToyIcon />, 
      color: 'rgba(3, 169, 244, 0.2)' 
    },
  ];

  const features = [
    {
      title: 'DynaContracts',
      description: 'AI-powered smart contracts that adapt based on external data and AI models.',
      link: '/smart-contracts'
    },
    {
      title: 'DeAI Agents',
      description: 'Create and manage your own AI agents that can perform actions on the blockchain.',
      link: '/ai-agents'
    },
    {
      title: 'TruthGPT Oracle',
      description: 'Decentralized oracle network powered by AI for verifying information and detecting misinformation.',
      link: '/oracle'
    },
    {
      title: 'HyperChains',
      description: 'Create your own Layer 3 chains using natural language prompts and AI.',
      link: '/hyperchains'
    },
  ];

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      {!account && (
        <StyledPaper sx={{ mb: 4, p: 4, textAlign: 'center' }}>
          <Typography variant="h4" gutterBottom>
            Welcome to NoMercyChain
          </Typography>
          <Typography variant="body1" color="text.secondary" paragraph>
            The next-generation blockchain platform with AI-powered smart contracts, decentralized AI agents, and Layer 3 solutions.
          </Typography>
          <Button 
            variant="contained" 
            size="large" 
            onClick={connectWallet}
            sx={{ 
              mt: 2, 
              background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
              '&:hover': {
                background: 'linear-gradient(45deg, #5000d6 30%, #00c060 90%)',
              },
            }}
          >
            Connect Wallet to Get Started
          </Button>
        </StyledPaper>
      )}

      <Grid container spacing={4}>
        {stats.map((stat, index) => (
          <Grid item xs={12} sm={6} md={3} key={index}>
            <StatCard>
              <IconWrapper color={stat.color}>
                {stat.icon}
              </IconWrapper>
              <Box>
                <Typography variant="body2" color="text.secondary">
                  {stat.title}
                </Typography>
                <Typography variant="h5" sx={{ fontWeight: 600 }}>
                  {stat.value}
                </Typography>
                <Typography variant="body2" color="success.main">
                  {stat.change}
                </Typography>
              </Box>
            </StatCard>
          </Grid>
        ))}
      </Grid>

      <Grid container spacing={4} sx={{ mt: 2 }}>
        <Grid item xs={12} md={8}>
          <StyledPaper>
            <Typography variant="h5" gutterBottom>
              Network Activity
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Box sx={{ height: 300, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
              <Typography variant="body1" color="text.secondary">
                Chart will be displayed here
              </Typography>
            </Box>
          </StyledPaper>
        </Grid>
        <Grid item xs={12} md={4}>
          <StyledPaper>
            <Typography variant="h5" gutterBottom>
              Recent Transactions
            </Typography>
            <Divider sx={{ my: 2 }} />
            {account ? (
              <Box>
                <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mr: 2, width: 32, height: 32 }}>Tx</Avatar>
                  <Box>
                    <Typography variant="body2">
                      Staked 100 NMX
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      2 minutes ago
                    </Typography>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mr: 2, width: 32, height: 32 }}>Tx</Avatar>
                  <Box>
                    <Typography variant="body2">
                      Created AI Agent
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      1 hour ago
                    </Typography>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                  <Avatar sx={{ bgcolor: 'error.main', mr: 2, width: 32, height: 32 }}>Tx</Avatar>
                  <Box>
                    <Typography variant="body2">
                      Deployed Smart Contract
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      3 hours ago
                    </Typography>
                  </Box>
                </Box>
              </Box>
            ) : (
              <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                Connect your wallet to see your transactions
              </Typography>
            )}
          </StyledPaper>
        </Grid>
      </Grid>

      <Typography variant="h5" sx={{ mt: 4, mb: 3 }}>
        Key Features
      </Typography>
      <Grid container spacing={4}>
        {features.map((feature, index) => (
          <Grid item xs={12} sm={6} md={3} key={index}>
            <FeatureCard>
              <CardContent sx={{ flexGrow: 1 }}>
                <Typography variant="h6" gutterBottom>
                  {feature.title}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {feature.description}
                </Typography>
              </CardContent>
              <CardActions>
                <Button 
                  size="small" 
                  color="primary" 
                  component={Link} 
                  to={feature.link}
                >
                  Explore
                </Button>
              </CardActions>
            </FeatureCard>
          </Grid>
        ))}
      </Grid>

      <StyledPaper sx={{ mt: 4, p: 3 }}>
        <Typography variant="h5" gutterBottom>
          Network Status
        </Typography>
        <Divider sx={{ my: 2 }} />
        <Grid container spacing={3}>
          <Grid item xs={12} sm={4}>
            <Typography variant="body2" color="text.secondary">
              Block Height
            </Typography>
            <Typography variant="h6">
              1,234,567
            </Typography>
          </Grid>
          <Grid item xs={12} sm={4}>
            <Typography variant="body2" color="text.secondary">
              Validators
            </Typography>
            <Typography variant="h6">
              100 / 100
            </Typography>
          </Grid>
          <Grid item xs={12} sm={4}>
            <Typography variant="body2" color="text.secondary">
              Transactions Per Second
            </Typography>
            <Typography variant="h6">
              1,245
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Typography variant="body2" color="text.secondary" gutterBottom>
              Network Load
            </Typography>
            <LinearProgress 
              variant="determinate" 
              value={65} 
              sx={{ 
                height: 10, 
                borderRadius: 5,
                bgcolor: 'rgba(255,255,255,0.1)',
                '& .MuiLinearProgress-bar': {
                  background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
                }
              }} 
            />
            <Typography variant="caption" color="text.secondary" sx={{ mt: 1, display: 'block' }}>
              65% - Normal Operation
            </Typography>
          </Grid>
        </Grid>
      </StyledPaper>
    </Container>
  );
};

export default Dashboard;