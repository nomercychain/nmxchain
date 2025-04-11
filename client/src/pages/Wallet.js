import React, { useContext, useState } from 'react';
import { 
  Box, 
  Container, 
  Grid, 
  Paper, 
  Typography, 
  Button, 
  Tabs, 
  Tab, 
  TextField, 
  Divider,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  IconButton,
  Chip,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  InputAdornment
} from '@mui/material';
import { styled } from '@mui/material/styles';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import SendIcon from '@mui/icons-material/Send';
import ReceiveIcon from '@mui/icons-material/CallReceived';
import SwapHorizIcon from '@mui/icons-material/SwapHoriz';
import QrCodeIcon from '@mui/icons-material/QrCode';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const WalletCard = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, #6200ea 0%, #3700b3 100%)',
  color: theme.palette.common.white,
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ActionButton = styled(Button)(({ theme }) => ({
  borderRadius: 8,
  padding: theme.spacing(1, 2),
  backgroundColor: 'rgba(255, 255, 255, 0.1)',
  color: theme.palette.common.white,
  '&:hover': {
    backgroundColor: 'rgba(255, 255, 255, 0.2)',
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

const Wallet = () => {
  const { account, balance } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [sendDialogOpen, setSendDialogOpen] = useState(false);
  const [receiveDialogOpen, setReceiveDialogOpen] = useState(false);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const formatAddress = (address) => {
    if (!address) return '';
    return `${address.substring(0, 10)}...${address.substring(address.length - 8)}`;
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    // You could add a toast notification here
  };

  // Mock transaction data
  const transactions = [
    {
      id: 'tx1',
      type: 'send',
      amount: '50 NMX',
      to: '0x1234...5678',
      from: account,
      date: '2023-06-15 14:30',
      status: 'completed'
    },
    {
      id: 'tx2',
      type: 'receive',
      amount: '100 NMX',
      to: account,
      from: '0x8765...4321',
      date: '2023-06-14 10:15',
      status: 'completed'
    },
    {
      id: 'tx3',
      type: 'swap',
      amount: '25 NMX',
      to: 'ETH',
      from: 'NMX',
      date: '2023-06-13 09:45',
      status: 'completed'
    },
    {
      id: 'tx4',
      type: 'stake',
      amount: '200 NMX',
      to: 'Validator',
      from: account,
      date: '2023-06-12 16:20',
      status: 'completed'
    },
  ];

  // Mock token data
  const tokens = [
    { symbol: 'NMX', name: 'NoMercyChain', balance: parseFloat(balance || 0), value: parseFloat(balance || 0) * 2.45 },
    { symbol: 'ETH', name: 'Ethereum', balance: 0.5, value: 0.5 * 1800 },
    { symbol: 'USDT', name: 'Tether', balance: 500, value: 500 },
    { symbol: 'USDC', name: 'USD Coin', balance: 750, value: 750 },
  ];

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      {account ? (
        <>
          <Grid container spacing={4}>
            <Grid item xs={12} md={8}>
              <WalletCard>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
                  <Box>
                    <Typography variant="h5" gutterBottom>
                      My Wallet
                    </Typography>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                      <Typography variant="body2" sx={{ mr: 1 }}>
                        {formatAddress(account)}
                      </Typography>
                      <IconButton 
                        size="small" 
                        onClick={() => copyToClipboard(account)}
                        sx={{ color: 'rgba(255, 255, 255, 0.7)' }}
                      >
                        <ContentCopyIcon fontSize="small" />
                      </IconButton>
                    </Box>
                  </Box>
                  <Box>
                    <Typography variant="h4" sx={{ fontWeight: 700 }}>
                      ${(parseFloat(balance || 0) * 2.45).toFixed(2)}
                    </Typography>
                    <Typography variant="body2" align="right">
                      {parseFloat(balance || 0).toFixed(4)} NMX
                    </Typography>
                  </Box>
                </Box>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', mt: 4 }}>
                  <ActionButton startIcon={<SendIcon />} onClick={() => setSendDialogOpen(true)}>
                    Send
                  </ActionButton>
                  <ActionButton startIcon={<ReceiveIcon />} onClick={() => setReceiveDialogOpen(true)}>
                    Receive
                  </ActionButton>
                  <ActionButton startIcon={<SwapHorizIcon />}>
                    Swap
                  </ActionButton>
                  <ActionButton startIcon={<QrCodeIcon />}>
                    Scan
                  </ActionButton>
                </Box>
              </WalletCard>
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
                >
                  Buy NMX
                </Button>
                <Button 
                  variant="outlined" 
                  fullWidth 
                  sx={{ mb: 2 }}
                >
                  Bridge Assets
                </Button>
                <Button 
                  variant="outlined" 
                  fullWidth
                >
                  Connect to DApp
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
              <StyledTab label="Assets" />
              <StyledTab label="Transactions" />
              <StyledTab label="NFTs" />
            </Tabs>

            {tabValue === 0 && (
              <StyledPaper>
                <TableContainer>
                  <Table>
                    <TableHead>
                      <TableRow>
                        <StyledTableCell>Asset</StyledTableCell>
                        <StyledTableCell align="right">Balance</StyledTableCell>
                        <StyledTableCell align="right">Value (USD)</StyledTableCell>
                        <StyledTableCell align="right">Actions</StyledTableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {tokens.map((token) => (
                        <TableRow key={token.symbol}>
                          <StyledTableCell>
                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                              <Box 
                                sx={{ 
                                  width: 32, 
                                  height: 32, 
                                  borderRadius: '50%', 
                                  bgcolor: 'primary.main', 
                                  display: 'flex', 
                                  alignItems: 'center', 
                                  justifyContent: 'center',
                                  mr: 2
                                }}
                              >
                                {token.symbol.substring(0, 1)}
                              </Box>
                              <Box>
                                <Typography variant="body2">{token.name}</Typography>
                                <Typography variant="caption" color="text.secondary">{token.symbol}</Typography>
                              </Box>
                            </Box>
                          </StyledTableCell>
                          <StyledTableCell align="right">{token.balance.toFixed(4)}</StyledTableCell>
                          <StyledTableCell align="right">${token.value.toFixed(2)}</StyledTableCell>
                          <StyledTableCell align="right">
                            <Button size="small" sx={{ mr: 1 }}>Send</Button>
                            <Button size="small">Swap</Button>
                          </StyledTableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </StyledPaper>
            )}

            {tabValue === 1 && (
              <StyledPaper>
                <TableContainer>
                  <Table>
                    <TableHead>
                      <TableRow>
                        <StyledTableCell>Type</StyledTableCell>
                        <StyledTableCell>Amount</StyledTableCell>
                        <StyledTableCell>From/To</StyledTableCell>
                        <StyledTableCell>Date</StyledTableCell>
                        <StyledTableCell>Status</StyledTableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {transactions.map((tx) => (
                        <TableRow key={tx.id}>
                          <StyledTableCell>
                            <Chip 
                              label={tx.type} 
                              size="small"
                              sx={{ 
                                bgcolor: tx.type === 'receive' ? 'success.main' : 
                                         tx.type === 'send' ? 'error.main' : 
                                         tx.type === 'swap' ? 'info.main' : 'warning.main',
                                color: 'white'
                              }}
                            />
                          </StyledTableCell>
                          <StyledTableCell>{tx.amount}</StyledTableCell>
                          <StyledTableCell>
                            {tx.type === 'swap' ? `${tx.from} → ${tx.to}` : 
                             tx.type === 'receive' ? `From: ${tx.from}` : 
                             `To: ${tx.to}`}
                          </StyledTableCell>
                          <StyledTableCell>{tx.date}</StyledTableCell>
                          <StyledTableCell>
                            <Chip 
                              label={tx.status} 
                              size="small"
                              sx={{ 
                                bgcolor: tx.status === 'completed' ? 'success.main' : 
                                         tx.status === 'pending' ? 'warning.main' : 'error.main',
                                color: 'white'
                              }}
                            />
                          </StyledTableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </StyledPaper>
            )}

            {tabValue === 2 && (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  No NFTs Found
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  You don't have any NFTs in your wallet yet.
                </Typography>
                <Button 
                  variant="contained" 
                  sx={{ mt: 2 }}
                >
                  Browse NFT Marketplace
                </Button>
              </StyledPaper>
            )}
          </Box>

          {/* Send Dialog */}
          <Dialog open={sendDialogOpen} onClose={() => setSendDialogOpen(false)} maxWidth="sm" fullWidth>
            <DialogTitle>Send Tokens</DialogTitle>
            <DialogContent>
              <TextField
                label="Recipient Address"
                fullWidth
                margin="normal"
                variant="outlined"
                placeholder="Enter recipient address"
              />
              <TextField
                label="Amount"
                fullWidth
                margin="normal"
                variant="outlined"
                placeholder="0.00"
                InputProps={{
                  endAdornment: (
                    <InputAdornment position="end">
                      NMX
                    </InputAdornment>
                  ),
                }}
              />
              <TextField
                label="Note (Optional)"
                fullWidth
                margin="normal"
                variant="outlined"
                placeholder="Add a note to this transaction"
              />
            </DialogContent>
            <DialogActions>
              <Button onClick={() => setSendDialogOpen(false)}>Cancel</Button>
              <Button variant="contained">Send</Button>
            </DialogActions>
          </Dialog>

          {/* Receive Dialog */}
          <Dialog open={receiveDialogOpen} onClose={() => setReceiveDialogOpen(false)} maxWidth="sm" fullWidth>
            <DialogTitle>Receive Tokens</DialogTitle>
            <DialogContent sx={{ textAlign: 'center' }}>
              <Box 
                sx={{ 
                  width: 200, 
                  height: 200, 
                  bgcolor: 'background.paper', 
                  mx: 'auto', 
                  my: 2, 
                  display: 'flex', 
                  alignItems: 'center', 
                  justifyContent: 'center' 
                }}
              >
                <Typography variant="body2" color="text.secondary">
                  QR Code will be displayed here
                </Typography>
              </Box>
              <TextField
                label="Your Wallet Address"
                fullWidth
                margin="normal"
                variant="outlined"
                value={account || ''}
                InputProps={{
                  readOnly: true,
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton onClick={() => copyToClipboard(account)}>
                        <ContentCopyIcon />
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
              />
              <Typography variant="body2" color="text.secondary" sx={{ mt: 2 }}>
                Send only NMX and NMX-based tokens to this address. Sending other tokens may result in permanent loss.
              </Typography>
            </DialogContent>
            <DialogActions>
              <Button onClick={() => setReceiveDialogOpen(false)}>Close</Button>
            </DialogActions>
          </Dialog>
        </>
      ) : (
        <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
          <Typography variant="h5" gutterBottom>
            Connect Your Wallet
          </Typography>
          <Typography variant="body1" color="text.secondary" paragraph>
            Please connect your wallet to view your assets and transactions.
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
    </Container>
  );
};

export default Wallet;