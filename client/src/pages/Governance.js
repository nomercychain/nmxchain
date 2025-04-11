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
  LinearProgress,
  FormControl,
  FormControlLabel,
  RadioGroup,
  Radio,
  Tabs,
  Tab
} from '@mui/material';
import { styled } from '@mui/material/styles';
import HowToVoteIcon from '@mui/icons-material/HowToVote';
import AddIcon from '@mui/icons-material/Add';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import CancelIcon from '@mui/icons-material/Cancel';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import ThumbDownIcon from '@mui/icons-material/ThumbDown';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ProposalCard = styled(Card)(({ theme }) => ({
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

const Governance = () => {
  const { account, balance } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [voteDialogOpen, setVoteDialogOpen] = useState(false);
  const [createProposalDialogOpen, setCreateProposalDialogOpen] = useState(false);
  const [selectedProposal, setSelectedProposal] = useState(null);
  const [voteOption, setVoteOption] = useState('yes');
  const [proposalTitle, setProposalTitle] = useState('');
  const [proposalDescription, setProposalDescription] = useState('');

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleVoteDialogOpen = (proposal) => {
    setSelectedProposal(proposal);
    setVoteDialogOpen(true);
  };

  const handleVoteDialogClose = () => {
    setVoteDialogOpen(false);
    setSelectedProposal(null);
    setVoteOption('yes');
  };

  const handleCreateProposalDialogOpen = () => {
    setCreateProposalDialogOpen(true);
  };

  const handleCreateProposalDialogClose = () => {
    setCreateProposalDialogOpen(false);
    setProposalTitle('');
    setProposalDescription('');
  };

  const handleVote = () => {
    // Mock voting functionality
    console.log(`Voted ${voteOption} on proposal ${selectedProposal.id}`);
    handleVoteDialogClose();
  };

  const handleCreateProposal = () => {
    // Mock proposal creation
    console.log(`Created proposal: ${proposalTitle}`);
    handleCreateProposalDialogClose();
  };

  // Mock proposals data
  const proposals = [
    {
      id: 'prop1',
      title: 'Increase Validator Set to 150',
      description: 'Proposal to increase the maximum number of validators from 100 to 150 to improve decentralization.',
      proposer: '0x1234...5678',
      status: 'active',
      votingEndTime: '2023-07-15',
      votingPower: {
        yes: 65,
        no: 10,
        abstain: 5,
        noWithVeto: 2,
      },
      quorum: 40,
      threshold: 50,
    },
    {
      id: 'prop2',
      title: 'Reduce Transaction Fees by 20%',
      description: 'Proposal to reduce the base transaction fee by 20% to encourage more network usage.',
      proposer: '0x8765...4321',
      status: 'active',
      votingEndTime: '2023-07-10',
      votingPower: {
        yes: 45,
        no: 35,
        abstain: 8,
        noWithVeto: 5,
      },
      quorum: 40,
      threshold: 50,
    },
    {
      id: 'prop3',
      title: 'Add New AI Oracle Module',
      description: 'Proposal to add a new AI-powered oracle module for providing off-chain data to smart contracts.',
      proposer: '0x2468...1357',
      status: 'passed',
      votingEndTime: '2023-06-30',
      votingPower: {
        yes: 75,
        no: 15,
        abstain: 5,
        noWithVeto: 1,
      },
      quorum: 40,
      threshold: 50,
    },
    {
      id: 'prop4',
      title: 'Increase Community Pool Allocation',
      description: 'Proposal to increase the community pool allocation from 2% to 5% of block rewards.',
      proposer: '0x1357...2468',
      status: 'rejected',
      votingEndTime: '2023-06-25',
      votingPower: {
        yes: 30,
        no: 60,
        abstain: 5,
        noWithVeto: 3,
      },
      quorum: 40,
      threshold: 50,
    },
  ];

  // Mock user votes
  const userVotes = [
    {
      proposalId: 'prop3',
      proposalTitle: 'Add New AI Oracle Module',
      vote: 'yes',
      votingPower: '500 NMX',
      time: '2023-06-28',
    },
    {
      proposalId: 'prop4',
      proposalTitle: 'Increase Community Pool Allocation',
      vote: 'no',
      votingPower: '500 NMX',
      time: '2023-06-23',
    },
  ];

  const getStatusColor = (status) => {
    switch (status) {
      case 'active':
        return 'primary';
      case 'passed':
        return 'success';
      case 'rejected':
        return 'error';
      default:
        return 'default';
    }
  };

  const getStatusIcon = (status) => {
    switch (status) {
      case 'active':
        return <HowToVoteIcon />;
      case 'passed':
        return <CheckCircleIcon />;
      case 'rejected':
        return <CancelIcon />;
      default:
        return null;
    }
  };

  const getVoteIcon = (vote) => {
    switch (vote) {
      case 'yes':
        return <ThumbUpIcon fontSize="small" />;
      case 'no':
        return <ThumbDownIcon fontSize="small" />;
      default:
        return null;
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          Governance
        </Typography>
        <Button 
          variant="contained" 
          startIcon={<AddIcon />}
          onClick={handleCreateProposalDialogOpen}
          disabled={!account}
        >
          Create Proposal
        </Button>
      </Box>

      <Grid container spacing={4}>
        <Grid item xs={12} md={8}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Governance Overview
            </Typography>
            <Typography variant="body1" paragraph>
              Participate in the decentralized governance of NoMercyChain by voting on proposals. 
              Your voting power is proportional to your staked NMX tokens.
            </Typography>
            <Grid container spacing={3} sx={{ mt: 2 }}>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'primary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <HowToVoteIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    {account ? '800 NMX' : '0 NMX'}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Your Voting Power
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'secondary.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <CheckCircleIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    2
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Active Proposals
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} sm={4}>
                <Box sx={{ textAlign: 'center' }}>
                  <Avatar sx={{ bgcolor: 'success.main', mx: 'auto', mb: 1, width: 56, height: 56 }}>
                    <HowToVoteIcon fontSize="large" />
                  </Avatar>
                  <Typography variant="h5" gutterBottom>
                    {account ? '2' : '0'}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Your Votes
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </StyledPaper>
        </Grid>
        <Grid item xs={12} md={4}>
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Governance Parameters
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Voting Period
              </Typography>
              <Typography variant="body1">
                14 days
              </Typography>
            </Box>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Quorum
              </Typography>
              <Typography variant="body1">
                40%
              </Typography>
            </Box>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body2" color="text.secondary">
                Threshold
              </Typography>
              <Typography variant="body1">
                50%
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Veto Threshold
              </Typography>
              <Typography variant="body1">
                33.4%
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
          <StyledTab label="Active Proposals" />
          <StyledTab label="Proposal History" />
          <StyledTab label="Your Votes" />
        </Tabs>

        {tabValue === 0 && (
          <Grid container spacing={4}>
            {proposals.filter(p => p.status === 'active').map((proposal) => (
              <Grid item xs={12} key={proposal.id}>
                <ProposalCard>
                  <CardContent>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                      <Avatar sx={{ bgcolor: 'primary.main', mr: 2 }}>
                        <HowToVoteIcon />
                      </Avatar>
                      <Box sx={{ flexGrow: 1 }}>
                        <Typography variant="h6">
                          {proposal.title}
                        </Typography>
                        <Typography variant="caption" color="text.secondary">
                          Proposed by {proposal.proposer} • Ends on {proposal.votingEndTime}
                        </Typography>
                      </Box>
                      <Chip 
                        icon={getStatusIcon(proposal.status)} 
                        label={proposal.status.charAt(0).toUpperCase() + proposal.status.slice(1)} 
                        color={getStatusColor(proposal.status)} 
                        size="small" 
                      />
                    </Box>
                    <Typography variant="body2" paragraph>
                      {proposal.description}
                    </Typography>
                    <Divider sx={{ my: 2 }} />
                    <Typography variant="body2" gutterBottom>
                      Voting Results
                    </Typography>
                    <Grid container spacing={2} sx={{ mb: 2 }}>
                      <Grid item xs={3}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          Yes
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.yes}%
                        </Typography>
                      </Grid>
                      <Grid item xs={3}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          No
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.no}%
                        </Typography>
                      </Grid>
                      <Grid item xs={3}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          Abstain
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.abstain}%
                        </Typography>
                      </Grid>
                      <Grid item xs={3}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          No with Veto
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.noWithVeto}%
                        </Typography>
                      </Grid>
                    </Grid>
                    <Box sx={{ mb: 1 }}>
                      <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                        <Typography variant="caption" color="text.secondary">
                          Turnout
                        </Typography>
                        <Typography variant="caption" color="text.secondary">
                          {proposal.votingPower.yes + proposal.votingPower.no + proposal.votingPower.abstain + proposal.votingPower.noWithVeto}%
                        </Typography>
                      </Box>
                      <LinearProgress 
                        variant="determinate" 
                        value={proposal.votingPower.yes + proposal.votingPower.no + proposal.votingPower.abstain + proposal.votingPower.noWithVeto} 
                        sx={{ 
                          height: 6, 
                          borderRadius: 3,
                          bgcolor: 'rgba(255,255,255,0.1)',
                          '& .MuiLinearProgress-bar': {
                            bgcolor: 'info.main',
                          }
                        }} 
                      />
                    </Box>
                    <Box>
                      <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                        <Typography variant="caption" color="text.secondary">
                          Yes Votes
                        </Typography>
                        <Typography variant="caption" color="text.secondary">
                          {proposal.votingPower.yes}%
                        </Typography>
                      </Box>
                      <LinearProgress 
                        variant="determinate" 
                        value={proposal.votingPower.yes} 
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
                      onClick={() => handleVoteDialogOpen(proposal)}
                      disabled={!account}
                    >
                      Vote
                    </Button>
                  </CardActions>
                </ProposalCard>
              </Grid>
            ))}
          </Grid>
        )}

        {tabValue === 1 && (
          <Grid container spacing={4}>
            {proposals.filter(p => p.status !== 'active').map((proposal) => (
              <Grid item xs={12} sm={6} key={proposal.id}>
                <ProposalCard>
                  <CardContent>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                      <Avatar sx={{ 
                        bgcolor: proposal.status === 'passed' ? 'success.main' : 'error.main', 
                        mr: 2 
                      }}>
                        {getStatusIcon(proposal.status)}
                      </Avatar>
                      <Box sx={{ flexGrow: 1 }}>
                        <Typography variant="h6">
                          {proposal.title}
                        </Typography>
                        <Typography variant="caption" color="text.secondary">
                          Ended on {proposal.votingEndTime}
                        </Typography>
                      </Box>
                      <Chip 
                        icon={getStatusIcon(proposal.status)} 
                        label={proposal.status.charAt(0).toUpperCase() + proposal.status.slice(1)} 
                        color={getStatusColor(proposal.status)} 
                        size="small" 
                      />
                    </Box>
                    <Typography variant="body2" paragraph>
                      {proposal.description}
                    </Typography>
                    <Divider sx={{ my: 2 }} />
                    <Typography variant="body2" gutterBottom>
                      Final Results
                    </Typography>
                    <Grid container spacing={2}>
                      <Grid item xs={6}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          Yes
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.yes}%
                        </Typography>
                      </Grid>
                      <Grid item xs={6}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          No
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.no}%
                        </Typography>
                      </Grid>
                      <Grid item xs={6}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          Abstain
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.abstain}%
                        </Typography>
                      </Grid>
                      <Grid item xs={6}>
                        <Typography variant="caption" color="text.secondary" display="block">
                          No with Veto
                        </Typography>
                        <Typography variant="body2">
                          {proposal.votingPower.noWithVeto}%
                        </Typography>
                      </Grid>
                    </Grid>
                  </CardContent>
                  <CardActions sx={{ p: 2 }}>
                    <Button 
                      variant="outlined" 
                      fullWidth
                    >
                      View Details
                    </Button>
                  </CardActions>
                </ProposalCard>
              </Grid>
            ))}
          </Grid>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              Your Voting History
            </Typography>
            <Divider sx={{ my: 2 }} />
            {account ? (
              <TableContainer>
                <Table>
                  <TableHead>
                    <TableRow>
                      <StyledTableCell>Proposal</StyledTableCell>
                      <StyledTableCell>Your Vote</StyledTableCell>
                      <StyledTableCell align="right">Voting Power</StyledTableCell>
                      <StyledTableCell align="right">Time</StyledTableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {userVotes.map((vote) => (
                      <TableRow key={vote.proposalId}>
                        <StyledTableCell>{vote.proposalTitle}</StyledTableCell>
                        <StyledTableCell>
                          <Chip 
                            icon={getVoteIcon(vote.vote)} 
                            label={vote.vote.charAt(0).toUpperCase() + vote.vote.slice(1)} 
                            size="small"
                            color={vote.vote === 'yes' ? 'success' : vote.vote === 'no' ? 'error' : 'default'}
                          />
                        </StyledTableCell>
                        <StyledTableCell align="right">{vote.votingPower}</StyledTableCell>
                        <StyledTableCell align="right">{vote.time}</StyledTableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            ) : (
              <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                Connect your wallet to see your voting history
              </Typography>
            )}
          </StyledPaper>
        )}
      </Box>

      {/* Vote Dialog */}
      <Dialog open={voteDialogOpen} onClose={handleVoteDialogClose} maxWidth="sm" fullWidth>
        <DialogTitle>
          {selectedProposal ? `Vote on Proposal: ${selectedProposal.title}` : 'Vote on Proposal'}
        </DialogTitle>
        <DialogContent>
          {selectedProposal && (
            <Box sx={{ mb: 3 }}>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Proposal Description
              </Typography>
              <Typography variant="body1" paragraph>
                {selectedProposal.description}
              </Typography>
              <Divider sx={{ my: 2 }} />
            </Box>
          )}
          <FormControl component="fieldset">
            <Typography variant="body2" color="text.secondary" gutterBottom>
              Your Vote
            </Typography>
            <RadioGroup
              aria-label="vote"
              name="vote"
              value={voteOption}
              onChange={(e) => setVoteOption(e.target.value)}
            >
              <FormControlLabel value="yes" control={<Radio />} label="Yes" />
              <FormControlLabel value="no" control={<Radio />} label="No" />
              <FormControlLabel value="abstain" control={<Radio />} label="Abstain" />
              <FormControlLabel value="noWithVeto" control={<Radio />} label="No with Veto" />
            </RadioGroup>
          </FormControl>
          <Box sx={{ mt: 3 }}>
            <Typography variant="body2" color="text.secondary" gutterBottom>
              Your Voting Power
            </Typography>
            <Typography variant="body1">
              {account ? '800 NMX' : '0 NMX'}
            </Typography>
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleVoteDialogClose}>Cancel</Button>
          <Button 
            variant="contained" 
            onClick={handleVote}
            disabled={!account}
          >
            Submit Vote
          </Button>
        </DialogActions>
      </Dialog>

      {/* Create Proposal Dialog */}
      <Dialog open={createProposalDialogOpen} onClose={handleCreateProposalDialogClose} maxWidth="md" fullWidth>
        <DialogTitle>
          Create Governance Proposal
        </DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="text.secondary" paragraph sx={{ mt: 1 }}>
            Creating a proposal requires a deposit of 1,000 NMX tokens. This deposit will be returned if the proposal reaches quorum, regardless of the outcome.
          </Typography>
          <TextField
            label="Proposal Title"
            fullWidth
            margin="normal"
            variant="outlined"
            value={proposalTitle}
            onChange={(e) => setProposalTitle(e.target.value)}
          />
          <TextField
            label="Proposal Description"
            fullWidth
            margin="normal"
            variant="outlined"
            multiline
            rows={6}
            value={proposalDescription}
            onChange={(e) => setProposalDescription(e.target.value)}
          />
          <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 2 }}>
            <Typography variant="body2" color="text.secondary">
              Required Deposit
            </Typography>
            <Typography variant="body2">
              1,000 NMX
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
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCreateProposalDialogClose}>Cancel</Button>
          <Button 
            variant="contained" 
            onClick={handleCreateProposal}
            disabled={!proposalTitle || !proposalDescription || (balance && parseFloat(balance) < 1000)}
          >
            Create Proposal
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default Governance;