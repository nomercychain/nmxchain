import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import AddIcon from '@mui/icons-material/Add';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import LocalFireDepartmentIcon from '@mui/icons-material/LocalFireDepartment';
import VerifiedUserIcon from '@mui/icons-material/VerifiedUser';
import {
  Alert,
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Chip,
  CircularProgress,
  Container,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  Grid,
  InputAdornment,
  LinearProgress,
  Paper,
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
import React, { useContext, useEffect, useState } from 'react';
import { blockchainApi } from '../api/blockchain';
import { fromBaseUnits, msgCreators } from '../api/transactions';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ValidatorCard = styled(Card)(({ theme }) => ({
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

const Staking = () => {
  const { account, balance, client, sendTransaction } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [stakeDialogOpen, setStakeDialogOpen] = useState(false);
  const [selectedValidator, setSelectedValidator] = useState(null);
  const [stakeAmount, setStakeAmount] = useState('');
  const [validators, setValidators] = useState([]);
  const [userStakes, setUserStakes] = useState([]);
  const [rewards, setRewards] = useState(null);
  const [loading, setLoading] = useState(true);
  const [stakingError, setStakingError] = useState(null);
  const [totalStaked, setTotalStaked] = useState('0');
  const [averageApy, setAverageApy] = useState('0');
  const [transactionInProgress, setTransactionInProgress] = useState(false);

  // Fetch validators and delegations
  useEffect(() => {
    const fetchStakingData = async () => {
      setLoading(true);
      try {
        // Fetch validators
        const validatorsResponse = await blockchainApi.getValidators();
        if (validatorsResponse && validatorsResponse.validators) {
          const formattedValidators = validatorsResponse.validators.map(validator => ({
            id: validator.operator_address,
            name: validator.description.moniker,
            address: validator.operator_address,
            votingPower: `${fromBaseUnits(validator.tokens)} NMX`,
            commission: `${(parseFloat(validator.commission.commission_rates.rate) * 100).toFixed(2)}%`,
            uptime: 99.9, // This would need to come from a separate API
            status: validator.status === 3 ? 'active' : 'inactive',
            delegators: parseInt(validator.delegator_shares) || 0,
            apy: 12.0, // This would need to be calculated based on rewards
            identity: validator.description.identity,
            website: validator.description.website,
            details: validator.description.details,
            tokens: validator.tokens,
            jailed: validator.jailed,
          }));
          
          setValidators(formattedValidators);
        }
        
        // Fetch delegations if account is connected
        if (account) {
          const delegationsResponse = await blockchainApi.getDelegations(account);
          const rewardsResponse = await blockchainApi.getDelegationRewards(account);
          
          if (delegationsResponse && delegationsResponse.delegation_responses) {
            // Calculate total staked
            let totalStakedAmount = 0;
            
            const formattedDelegations = delegationsResponse.delegation_responses.map(delegation => {
              const validatorInfo = validatorsResponse.validators.find(
                v => v.operator_address === delegation.delegation.validator_address
              );
              
              const amount = parseFloat(fromBaseUnits(delegation.balance.amount));
              totalStakedAmount += amount;
              
              return {
                id: `${delegation.delegation.validator_address}-${delegation.delegation.delegator_address}`,
                validator: validatorInfo ? validatorInfo.description.moniker : delegation.delegation.validator_address,
                validatorAddress: delegation.delegation.validator_address,
                amount: `${amount.toFixed(2)} NMX`,
                rawAmount: delegation.balance.amount,
                rewards: '0', // Will be updated with rewards data
                apy: '12.0%', // This would need to be calculated
                since: 'N/A', // This information is not directly available
              };
            });
            
            // Update rewards information
            if (rewardsResponse && rewardsResponse.rewards) {
              let totalRewards = 0;
              
              rewardsResponse.rewards.forEach(reward => {
                const delegation = formattedDelegations.find(
                  d => d.validatorAddress === reward.validator_address
                );
                
                if (delegation && reward.reward && reward.reward.length > 0) {
                  const nmxReward = reward.reward.find(r => r.denom === 'unmx');
                  if (nmxReward) {
                    const rewardAmount = parseFloat(fromBaseUnits(nmxReward.amount));
                    totalRewards += rewardAmount;
                    delegation.rewards = `${rewardAmount.toFixed(2)} NMX`;
                  }
                }
              });
              
              setRewards(totalRewards.toFixed(2));
            }
            
            setUserStakes(formattedDelegations);
            setTotalStaked(totalStakedAmount.toFixed(2));
            
            // Calculate average APY (this would be more accurate with real data)
            if (formattedDelegations.length > 0) {
              const avgApy = 12.0; // Placeholder
              setAverageApy(avgApy.toFixed(1));
            }
          }
        }
      } catch (error) {
        console.error('Error fetching staking data:', error);
      } finally {
        setLoading(false);
      }
    };
    
    fetchStakingData();
  }, [account]);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleStakeDialogOpen = (validator) => {
    setSelectedValidator(validator);
    setStakeDialogOpen(true);
    setStakingError(null);
  };

  const handleStakeDialogClose = () => {
    setStakeDialogOpen(false);
    setSelectedValidator(null);
    setStakeAmount('');
    setStakingError(null);
  };

  const handleStake = async () => {
    if (!account || !client || !selectedValidator) {
      setStakingError('Wallet not connected or validator not selected');
      return;
    }
    
    if (!stakeAmount || parseFloat(stakeAmount) <= 0) {
      setStakingError('Please enter a valid amount');
      return;
    }
    
    if (parseFloat(stakeAmount) > parseFloat(balance)) {
      setStakingError('Insufficient balance');
      return;
    }
    
    setTransactionInProgress(true);
    setStakingError(null);
    
    try {
      // Create delegation message
      const delegateMsg = msgCreators.staking.delegate(
        account,
        selectedValidator.address,
        stakeAmount
      );
      
      // Send transaction
      const result = await sendTransaction(delegateMsg);
      
      if (result && result.code === 0) {
        // Transaction successful
        console.log(`Successfully staked ${stakeAmount} NMX to validator ${selectedValidator.name}`);
        
        // Refresh staking data
        const delegationsResponse = await blockchainApi.getDelegations(account);
        if (delegationsResponse && delegationsResponse.delegation_responses) {
          const formattedDelegations = delegationsResponse.delegation_responses.map(delegation => {
            const validatorInfo = validators.find(
              v => v.address === delegation.delegation.validator_address
            );
            
            return {
              id: `${delegation.delegation.validator_address}-${delegation.delegation.delegator_address}`,
              validator: validatorInfo ? validatorInfo.name : delegation.delegation.validator_address,
              validatorAddress: delegation.delegation.validator_address,
              amount: `${fromBaseUnits(delegation.balance.amount)} NMX`,
              rawAmount: delegation.balance.amount,
              rewards: '0', // Will be updated with rewards data
              apy: '12.0%', // This would need to be calculated
              since: new Date().toISOString().split('T')[0], // Today's date
            };
          });
          
          setUserStakes(formattedDelegations);
          
          // Update total staked
          const totalStakedAmount = formattedDelegations.reduce(
            (total, delegation) => total + parseFloat(fromBaseUnits(delegation.rawAmount)),
            0
          );
          setTotalStaked(totalStakedAmount.toFixed(2));
        }
        
        handleStakeDialogClose();
      } else {
        // Transaction failed
        setStakingError(`Transaction failed: ${result.rawLog || 'Unknown error'}`);
      }
    } catch (error) {
      console.error('Error staking tokens:', error);
      setStakingError(`Error: ${error.message}`);
    } finally {
      setTransactionInProgress(false);
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          Staking
        </Typography>
        <Button 
          variant="contained" 
          startIcon={<AddIcon />}
          onClick={() => setStakeDialogOpen(true)}
          disabled={!account}
        >
          Stake Tokens
        </Button>
      </Box>

      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 8 }}>
          <CircularProgress />
        </Box>
      ) : (
        <>
          <Grid container spacing={4}>
            <Grid item xs={12} md={8}>
              <StyledPaper>
                <Typography variant="h6" gutterBottom>
                  Staking Overview
                </Typography>
                <Typography variant="body1" paragraph>
                  Stake your NMX tokens to validators to secure the network and earn rewards. 
                  The current network-wide staking APY is approximately 12.4%.
                </Typography>
                <Grid container spacing={3} sx={{ mt: 2 }}>
                  <Grid item xs={12} sm={4}>
                    <Box sx={{ textAlign: 'center' }}>
                      <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                        <LocalFireDepartmentIcon fontSize="large" />
                      </Avatar>
                      <Typography variant="h5" gutterBottom>
                        {account ? `${totalStaked} NMX` : '0 NMX'}
                      </Typography>
                      <Typography variant="body2" color="text.secondary">
                        Total Staked
                      </Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={12} sm={4}>
                    <Box sx={{ textAlign: 'center' }}>
                      <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                        <AccountBalanceIcon fontSize="large" />
                      </Avatar>
                      <Typography variant="h5" gutterBottom>
                        {account ? `${rewards || '0'} NMX` : '0 NMX'}
                      </Typography>
                      <Typography variant="body2" color="text.secondary">
                        Unclaimed Rewards
                      </Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={12} sm={4}>
                    <Box sx={{ textAlign: 'center' }}>
                      <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                        <VerifiedUserIcon fontSize="large" />
                      </Avatar>
                      <Typography variant="h5" gutterBottom>
                        {account ? `${averageApy}%` : '0%'}
                      </Typography>
                      <Typography variant="body2" color="text.secondary">
                        Average APY
                      </Typography>
                    </Box>
                  </Grid>
                </Grid>
              </StyledPaper>
            </Grid>
            <Grid item xs={12} md={4}>
              <StyledPaper>
                <Typography variant="h6" gutterBottom>
                  Quick Actions
                </Typography>
                <Divider sx={{ my: 2 }} />
                <Button 
                  variant="outlined" 
                  fullWidth 
                  sx={{ mb: 2 }}
                  disabled={!account || !rewards || parseFloat(rewards) <= 0}
                >
                  Claim All Rewards
                </Button>
                <Button 
                  variant="outlined" 
                  fullWidth 
                  sx={{ mb: 2 }}
                  disabled={!account || !rewards || parseFloat(rewards) <= 0}
                >
                  Restake All Rewards
                </Button>
                <Button 
                  variant="outlined" 
                  fullWidth
                  disabled={!account || !totalStaked || parseFloat(totalStaked) <= 0}
                >
                  Unstake Tokens
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
              <StyledTab label="My Stakes" />
              <StyledTab label="Validators" />
              <StyledTab label="Staking Stats" />
            </Tabs>

            {tabValue === 0 && (
              <StyledPaper>
                <Typography variant="h6" gutterBottom>
                  Your Staked Tokens
                </Typography>
                <Divider sx={{ my: 2 }} />
                {account ? (
                  userStakes.length > 0 ? (
                    <TableContainer>
                      <Table>
                        <TableHead>
                          <TableRow>
                            <StyledTableCell>Validator</StyledTableCell>
                            <StyledTableCell align="right">Amount</StyledTableCell>
                            <StyledTableCell align="right">Rewards</StyledTableCell>
                            <StyledTableCell align="right">APY</StyledTableCell>
                            <StyledTableCell align="right">Since</StyledTableCell>
                            <StyledTableCell align="right">Actions</StyledTableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>
                          {userStakes.map((stake) => (
                            <TableRow key={stake.id}>
                              <StyledTableCell>{stake.validator}</StyledTableCell>
                              <StyledTableCell align="right">{stake.amount}</StyledTableCell>
                              <StyledTableCell align="right">{stake.rewards}</StyledTableCell>
                              <StyledTableCell align="right">{stake.apy}</StyledTableCell>
                              <StyledTableCell align="right">{stake.since}</StyledTableCell>
                              <StyledTableCell align="right">
                                <Button size="small" sx={{ mr: 1 }}>Claim</Button>
                                <Button size="small">Unstake</Button>
                              </StyledTableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  ) : (
                    <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                      You don't have any staked tokens yet. Stake some tokens to earn rewards!
                    </Typography>
                  )
                ) : (
                  <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                    Connect your wallet to see your staked tokens
                  </Typography>
                )}
              </StyledPaper>
            )}

            {tabValue === 1 && (
              <Grid container spacing={4}>
                {validators.length > 0 ? (
                  validators.map((validator) => (
                    <Grid item xs={12} sm={6} key={validator.id}>
                      <ValidatorCard>
                        <CardContent>
                          <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                            <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                              {validator.name.charAt(0)}
                            </Avatar>
                            <Box>
                              <Typography variant="h6">
                                {validator.name}
                              </Typography>
                              <Typography variant="caption" color="text.secondary">
                                {validator.address.substring(0, 10)}...{validator.address.substring(validator.address.length - 4)}
                              </Typography>
                            </Box>
                            {validator.status === 'active' && (
                              <Chip 
                                icon={<CheckCircleIcon />} 
                                label="Active" 
                                size="small" 
                                color="success" 
                                sx={{ ml: 'auto' }} 
                              />
                            )}
                          </Box>
                          <Divider sx={{ my: 2 }} />
                          <Grid container spacing={2}>
                            <Grid item xs={6}>
                              <Typography variant="caption" color="text.secondary" display="block">
                                Voting Power
                              </Typography>
                              <Typography variant="body2">
                                {validator.votingPower}
                              </Typography>
                            </Grid>
                            <Grid item xs={6}>
                              <Typography variant="caption" color="text.secondary" display="block">
                                Commission
                              </Typography>
                              <Typography variant="body2">
                                {validator.commission}
                              </Typography>
                            </Grid>
                            <Grid item xs={6}>
                              <Typography variant="caption" color="text.secondary" display="block">
                                Delegators
                              </Typography>
                              <Typography variant="body2">
                                {validator.delegators}
                              </Typography>
                            </Grid>
                            <Grid item xs={6}>
                              <Typography variant="caption" color="text.secondary" display="block">
                                APY
                              </Typography>
                              <Typography variant="body2">
                                {validator.apy}%
                              </Typography>
                            </Grid>
                          </Grid>
                          <Box sx={{ mt: 2 }}>
                            <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                              Uptime: {validator.uptime}%
                            </Typography>
                            <LinearProgress 
                              variant="determinate" 
                              value={validator.uptime} 
                              sx={{ 
                                height: 6, 
                                borderRadius: 3,
                                bgcolor: 'rgba(255,255,255,0.1)',
                                '& .MuiLinearProgress-bar': {
                                  bgcolor: 'success.main',
                                }
                              }} 
                            />
                          </Box>
                        </CardContent>
                        <CardActions sx={{ p: 2 }}>
                          <Button 
                            variant="contained" 
                            fullWidth
                            onClick={() => handleStakeDialogOpen(validator)}
                            disabled={!account}
                          >
                            Stake
                          </Button>
                        </CardActions>
                      </ValidatorCard>
                    </Grid>
                  ))
                ) : (
                  <Grid item xs={12}>
                    <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                      No validators found
                    </Typography>
                  </Grid>
                )}
              </Grid>
            )}

            {tabValue === 2 && (
              <StyledPaper>
                <Typography variant="h6" gutterBottom>
                  Network Staking Statistics
                </Typography>
                <Divider sx={{ my: 2 }} />
                <Grid container spacing={3}>
                  <Grid item xs={12} sm={4}>
                    <Typography variant="body2" color="text.secondary">
                      Total Staked NMX
                    </Typography>
                    <Typography variant="h6">
                      {validators.reduce((total, validator) => {
                        const tokens = parseFloat(fromBaseUnits(validator.tokens || '0'));
                        return total + tokens;
                      }, 0).toLocaleString()} NMX
                    </Typography>
                  </Grid>
                  <Grid item xs={12} sm={4}>
                    <Typography variant="body2" color="text.secondary">
                      Staking Ratio
                    </Typography>
                    <Typography variant="h6">
                      {/* This would need to be calculated based on total supply */}
                      64.5%
                    </Typography>
                  </Grid>
                  <Grid item xs={12} sm={4}>
                    <Typography variant="body2" color="text.secondary">
                      Active Validators
                    </Typography>
                    <Typography variant="h6">
                      {validators.filter(v => v.status === 'active').length} / {validators.length}
                    </Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography variant="body2" color="text.secondary" gutterBottom>
                      Network Security
                    </Typography>
                    <LinearProgress 
                      variant="determinate" 
                      value={64.5} 
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
                      64.5% - Strong Security
                    </Typography>
                  </Grid>
                </Grid>
              </StyledPaper>
            )}
          </Box>
        </>
      )}

      {/* Stake Dialog */}
      <Dialog open={stakeDialogOpen} onClose={handleStakeDialogClose} maxWidth="sm" fullWidth>
        <DialogTitle>
          {selectedValidator ? `Stake to ${selectedValidator.name}` : 'Stake Tokens'}
        </DialogTitle>
        <DialogContent>
          {selectedValidator && (
            <Box sx={{ mb: 3 }}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Validator
              </Typography>
              <Box sx={{ display: 'flex', alignItems: 'center' }}>
                <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                  {selectedValidator.name.charAt(0)}
                </Avatar>
                <Box>
                  <Typography variant="body1">
                    {selectedValidator.name}
                  </Typography>
                  <Typography variant="caption" color="text.secondary">
                    Commission: {selectedValidator.commission} | APY: {selectedValidator.apy}%
                  </Typography>
                </Box>
              </Box>
            </Box>
          )}
          
          {stakingError && (
            <Alert severity="error" sx={{ mb: 2 }}>
              {stakingError}
            </Alert>
          )}
          
          <TextField
            label="Amount to Stake"
            fullWidth
            margin="normal"
            variant="outlined"
            type="number"
            value={stakeAmount}
            onChange={(e) => setStakeAmount(e.target.value)}
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  NMX
                </InputAdornment>
              ),
            }}
            disabled={transactionInProgress}
          />
          <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
            <Typography variant="body2" color="text.secondary">
              Available Balance
            </Typography>
            <Typography variant="body2">
              {balance || 0} NMX
            </Typography>
          </Box>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 1 }}>
            <Typography variant="body2" color="text.secondary">
              Estimated Annual Rewards
            </Typography>
            <Typography variant="body2">
              {stakeAmount && selectedValidator ? (parseFloat(stakeAmount) * selectedValidator.apy / 100).toFixed(2) : '0'} NMX
            </Typography>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleStakeDialogClose} disabled={transactionInProgress}>
            Cancel
          </Button>
          <Button 
            variant="contained" 
            onClick={handleStake}
            disabled={
              transactionInProgress || 
              !stakeAmount || 
              parseFloat(stakeAmount) <= 0 || 
              (balance && parseFloat(stakeAmount) > parseFloat(balance))
            }
          >
            {transactionInProgress ? (
              <>
                <CircularProgress size={24} sx={{ mr: 1 }} />
                Staking...
              </>
            ) : (
              'Stake'
            )}
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default Staking;