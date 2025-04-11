import React, { useState, useContext } from 'react';
import { 
  Box, 
  Container, 
  Grid, 
  Paper, 
  Typography, 
  Button, 
  Divider,
  TextField,
  Switch,
  FormControlLabel,
  FormGroup,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  ListItemSecondaryAction,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Alert,
  Tabs,
  Tab
} from '@mui/material';
import { styled } from '@mui/material/styles';
import SettingsIcon from '@mui/icons-material/Settings';
import SecurityIcon from '@mui/icons-material/Security';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import NotificationsIcon from '@mui/icons-material/Notifications';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import LanguageIcon from '@mui/icons-material/Language';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';
import VpnKeyIcon from '@mui/icons-material/VpnKey';
import DeleteIcon from '@mui/icons-material/Delete';
import AddIcon from '@mui/icons-material/Add';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const StyledTab = styled(Tab)(({ theme }) => ({
  textTransform: 'none',
  fontWeight: 600,
  fontSize: '1rem',
}));

const Settings = () => {
  const { account, disconnectWallet } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [darkMode, setDarkMode] = useState(true);
  const [notifications, setNotifications] = useState(true);
  const [language, setLanguage] = useState('en');
  const [gasPreference, setGasPreference] = useState('standard');
  const [showPrivateKey, setShowPrivateKey] = useState(false);
  const [exportDialogOpen, setExportDialogOpen] = useState(false);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [password, setPassword] = useState('');
  const [exportData, setExportData] = useState('');
  const [showExportData, setShowExportData] = useState(false);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  const handleExportWallet = () => {
    // Mock export functionality
    setExportData('{"version":"1.0","address":"' + account + '","privateKey":"xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334"}');
    setShowExportData(true);
  };

  const handleDeleteAccount = () => {
    // Mock delete functionality
    disconnectWallet();
    setDeleteDialogOpen(false);
  };

  // Mock connected apps
  const connectedApps = [
    {
      id: 'app1',
      name: 'DeFi Exchange',
      permissions: ['View address', 'Sign transactions'],
      lastUsed: '2023-06-15',
    },
    {
      id: 'app2',
      name: 'NFT Marketplace',
      permissions: ['View address'],
      lastUsed: '2023-06-10',
    },
    {
      id: 'app3',
      name: 'Governance Portal',
      permissions: ['View address', 'Sign transactions', 'View balance'],
      lastUsed: '2023-06-05',
    },
  ];

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 4 }}>
        <SettingsIcon sx={{ mr: 2, fontSize: 32 }} />
        <Typography variant="h4">
          Settings
        </Typography>
      </Box>

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
        <StyledTab label="General" />
        <StyledTab label="Security" />
        <StyledTab label="Connected Apps" />
        <StyledTab label="Advanced" />
      </Tabs>

      {tabValue === 0 && (
        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Appearance
              </Typography>
              <Divider sx={{ my: 2 }} />
              <FormGroup>
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={darkMode} 
                      onChange={(e) => setDarkMode(e.target.checked)} 
                    />
                  } 
                  label="Dark Mode" 
                />
              </FormGroup>
              <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                Toggle between dark and light theme
              </Typography>
            </StyledPaper>
          </Grid>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Notifications
              </Typography>
              <Divider sx={{ my: 2 }} />
              <FormGroup>
                <FormControlLabel 
                  control={
                    <Switch 
                      checked={notifications} 
                      onChange={(e) => setNotifications(e.target.checked)} 
                    />
                  } 
                  label="Enable Notifications" 
                />
              </FormGroup>
              <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                Receive notifications for transactions, governance proposals, and other important events
              </Typography>
            </StyledPaper>
          </Grid>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Language
              </Typography>
              <Divider sx={{ my: 2 }} />
              <FormControl fullWidth>
                <InputLabel>Language</InputLabel>
                <Select
                  value={language}
                  label="Language"
                  onChange={(e) => setLanguage(e.target.value)}
                >
                  <MenuItem value="en">English</MenuItem>
                  <MenuItem value="es">Español</MenuItem>
                  <MenuItem value="fr">Français</MenuItem>
                  <MenuItem value="de">Deutsch</MenuItem>
                  <MenuItem value="zh">中文</MenuItem>
                  <MenuItem value="ja">日本語</MenuItem>
                  <MenuItem value="ko">한국어</MenuItem>
                </Select>
              </FormControl>
            </StyledPaper>
          </Grid>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Transaction Settings
              </Typography>
              <Divider sx={{ my: 2 }} />
              <FormControl fullWidth sx={{ mb: 2 }}>
                <InputLabel>Gas Preference</InputLabel>
                <Select
                  value={gasPreference}
                  label="Gas Preference"
                  onChange={(e) => setGasPreference(e.target.value)}
                >
                  <MenuItem value="slow">Slow (Cheaper)</MenuItem>
                  <MenuItem value="standard">Standard</MenuItem>
                  <MenuItem value="fast">Fast (More Expensive)</MenuItem>
                </Select>
              </FormControl>
              <Typography variant="body2" color="text.secondary">
                Set your preferred transaction speed and cost
              </Typography>
            </StyledPaper>
          </Grid>
        </Grid>
      )}

      {tabValue === 1 && (
        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Wallet Security
              </Typography>
              <Divider sx={{ my: 2 }} />
              {account ? (
                <>
                  <Box sx={{ mb: 3 }}>
                    <Typography variant="body2" color="text.secondary" gutterBottom>
                      Wallet Address
                    </Typography>
                    <Typography variant="body1">
                      {account}
                    </Typography>
                  </Box>
                  <Box sx={{ mb: 3 }}>
                    <Typography variant="body2" color="text.secondary" gutterBottom>
                      Private Key
                    </Typography>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <TextField
                        type={showPrivateKey ? 'text' : 'password'}
                        value="xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334"
                        fullWidth
                        InputProps={{
                          readOnly: true,
                          endAdornment: (
                            <IconButton onClick={() => setShowPrivateKey(!showPrivateKey)}>
                              {showPrivateKey ? <VisibilityOffIcon /> : <VisibilityIcon />}
                            </IconButton>
                          ),
                        }}
                      />
                    </Box>
                    <Typography variant="caption" color="error" sx={{ mt: 1, display: 'block' }}>
                      Never share your private key with anyone!
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                    <Button 
                      variant="outlined" 
                      startIcon={<VpnKeyIcon />}
                      onClick={() => setExportDialogOpen(true)}
                    >
                      Export Wallet
                    </Button>
                    <Button 
                      variant="outlined" 
                      color="error" 
                      startIcon={<DeleteIcon />}
                      onClick={() => setDeleteDialogOpen(true)}
                    >
                      Delete Wallet
                    </Button>
                  </Box>
                </>
              ) : (
                <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
                  Connect your wallet to manage security settings
                </Typography>
              )}
            </StyledPaper>
          </Grid>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Security Settings
              </Typography>
              <Divider sx={{ my: 2 }} />
              <List>
                <ListItem>
                  <ListItemIcon>
                    <SecurityIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="Auto-Lock Timeout" 
                    secondary="Automatically lock your wallet after a period of inactivity" 
                  />
                  <ListItemSecondaryAction>
                    <FormControl sx={{ minWidth: 120 }}>
                      <Select
                        value="5"
                        size="small"
                      >
                        <MenuItem value="1">1 minute</MenuItem>
                        <MenuItem value="5">5 minutes</MenuItem>
                        <MenuItem value="15">15 minutes</MenuItem>
                        <MenuItem value="30">30 minutes</MenuItem>
                        <MenuItem value="60">1 hour</MenuItem>
                        <MenuItem value="never">Never</MenuItem>
                      </Select>
                    </FormControl>
                  </ListItemSecondaryAction>
                </ListItem>
                <ListItem>
                  <ListItemIcon>
                    <NotificationsIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="Transaction Alerts" 
                    secondary="Get notified for all transactions" 
                  />
                  <ListItemSecondaryAction>
                    <Switch defaultChecked />
                  </ListItemSecondaryAction>
                </ListItem>
                <ListItem>
                  <ListItemIcon>
                    <AccountBalanceWalletIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="Spending Limits" 
                    secondary="Set daily transaction limits" 
                  />
                  <ListItemSecondaryAction>
                    <Switch />
                  </ListItemSecondaryAction>
                </ListItem>
              </List>
            </StyledPaper>
          </Grid>
        </Grid>
      )}

      {tabValue === 2 && (
        <StyledPaper>
          <Typography variant="h6" gutterBottom>
            Connected Applications
          </Typography>
          <Divider sx={{ my: 2 }} />
          {account ? (
            <>
              <List>
                {connectedApps.map((app) => (
                  <ListItem key={app.id}>
                    <ListItemIcon>
                      <AccountBalanceWalletIcon />
                    </ListItemIcon>
                    <ListItemText 
                      primary={app.name} 
                      secondary={
                        <>
                          <Typography variant="body2" component="span">
                            Permissions: {app.permissions.join(', ')}
                          </Typography>
                          <br />
                          <Typography variant="caption" component="span" color="text.secondary">
                            Last used: {app.lastUsed}
                          </Typography>
                        </>
                      } 
                    />
                    <ListItemSecondaryAction>
                      <Button 
                        variant="outlined" 
                        color="error" 
                        size="small"
                      >
                        Disconnect
                      </Button>
                    </ListItemSecondaryAction>
                  </ListItem>
                ))}
              </List>
              <Box sx={{ mt: 2 }}>
                <Alert severity="info">
                  Connected applications can access your wallet address and request transaction signatures. 
                  Review and disconnect any applications you no longer use.
                </Alert>
              </Box>
            </>
          ) : (
            <Typography variant="body1" color="text.secondary" align="center" sx={{ py: 4 }}>
              Connect your wallet to view connected applications
            </Typography>
          )}
        </StyledPaper>
      )}

      {tabValue === 3 && (
        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Network Settings
              </Typography>
              <Divider sx={{ my: 2 }} />
              <FormControl fullWidth sx={{ mb: 3 }}>
                <InputLabel>Network</InputLabel>
                <Select
                  value="mainnet"
                  label="Network"
                >
                  <MenuItem value="mainnet">NoMercyChain Mainnet</MenuItem>
                  <MenuItem value="testnet">NoMercyChain Testnet</MenuItem>
                  <MenuItem value="devnet">Development Network</MenuItem>
                </Select>
              </FormControl>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Custom RPC URL
              </Typography>
              <TextField
                fullWidth
                placeholder="https://rpc.nomercychain.io"
                variant="outlined"
                size="small"
              />
            </StyledPaper>
          </Grid>
          <Grid item xs={12} md={6}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Advanced Settings
              </Typography>
              <Divider sx={{ my: 2 }} />
              <List>
                <ListItem>
                  <ListItemIcon>
                    <LanguageIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="IPFS Gateway" 
                    secondary="Custom IPFS gateway URL" 
                  />
                  <ListItemSecondaryAction>
                    <IconButton size="small">
                      <SettingsIcon fontSize="small" />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                <ListItem>
                  <ListItemIcon>
                    <DarkModeIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="Developer Mode" 
                    secondary="Enable advanced features for developers" 
                  />
                  <ListItemSecondaryAction>
                    <Switch />
                  </ListItemSecondaryAction>
                </ListItem>
              </List>
              <Box sx={{ mt: 3 }}>
                <Button 
                  variant="outlined" 
                  color="error" 
                  fullWidth
                >
                  Reset All Settings
                </Button>
              </Box>
            </StyledPaper>
          </Grid>
          <Grid item xs={12}>
            <StyledPaper>
              <Typography variant="h6" gutterBottom>
                Custom Networks
              </Typography>
              <Divider sx={{ my: 2 }} />
              <List>
                <ListItem>
                  <ListItemIcon>
                    <AccountBalanceWalletIcon />
                  </ListItemIcon>
                  <ListItemText 
                    primary="NoMercyChain Testnet" 
                    secondary="https://testnet-rpc.nomercychain.io" 
                  />
                  <ListItemSecondaryAction>
                    <IconButton>
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
              </List>
              <Box sx={{ mt: 2, display: 'flex', justifyContent: 'center' }}>
                <Button 
                  variant="outlined" 
                  startIcon={<AddIcon />}
                >
                  Add Network
                </Button>
              </Box>
            </StyledPaper>
          </Grid>
        </Grid>
      )}

      {/* Export Wallet Dialog */}
      <Dialog open={exportDialogOpen} onClose={() => setExportDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>
          Export Wallet
        </DialogTitle>
        <DialogContent>
          <Alert severity="warning" sx={{ mb: 3 }}>
            Warning: Exporting your wallet will reveal your private key. Make sure you are in a secure environment and no one is watching your screen.
          </Alert>
          {!showExportData ? (
            <>
              <Typography variant="body2" paragraph>
                Enter your password to export your wallet:
              </Typography>
              <TextField
                label="Password"
                type="password"
                fullWidth
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </>
          ) : (
            <>
              <Typography variant="body2" paragraph>
                Your wallet data:
              </Typography>
              <TextField
                multiline
                rows={4}
                fullWidth
                value={exportData}
                InputProps={{
                  readOnly: true,
                }}
              />
              <Typography variant="caption" color="error" sx={{ mt: 1, display: 'block' }}>
                Store this information securely. Anyone with access to this data can control your wallet.
              </Typography>
            </>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => {
            setExportDialogOpen(false);
            setShowExportData(false);
            setPassword('');
            setExportData('');
          }}>
            Close
          </Button>
          {!showExportData && (
            <Button 
              variant="contained" 
              onClick={handleExportWallet}
              disabled={!password}
            >
              Export
            </Button>
          )}
        </DialogActions>
      </Dialog>

      {/* Delete Wallet Dialog */}
      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>
          Delete Wallet
        </DialogTitle>
        <DialogContent>
          <Alert severity="error" sx={{ mb: 3 }}>
            Warning: This action cannot be undone. Make sure you have backed up your private key before proceeding.
          </Alert>
          <Typography variant="body2" paragraph>
            Enter "DELETE" to confirm:
          </Typography>
          <TextField
            fullWidth
            placeholder="DELETE"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteDialogOpen(false)}>
            Cancel
          </Button>
          <Button 
            variant="contained" 
            color="error"
            onClick={handleDeleteAccount}
          >
            Delete Wallet
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default Settings;