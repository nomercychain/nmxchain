import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { 
  AppBar, 
  Toolbar, 
  Typography, 
  Button, 
  Box, 
  IconButton, 
  Avatar, 
  Menu, 
  MenuItem, 
  Tooltip,
  CircularProgress
} from '@mui/material';
import { styled } from '@mui/material/styles';
import MenuIcon from '@mui/icons-material/Menu';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { WalletContext } from '../context/WalletContext';

const StyledAppBar = styled(AppBar)(({ theme }) => ({
  background: 'linear-gradient(90deg, #121212 0%, #1e1e1e 100%)',
  boxShadow: '0 4px 20px rgba(0, 0, 0, 0.15)',
}));

const LogoText = styled(Typography)(({ theme }) => ({
  fontWeight: 700,
  background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
  WebkitBackgroundClip: 'text',
  WebkitTextFillColor: 'transparent',
  marginRight: theme.spacing(2),
}));

const WalletButton = styled(Button)(({ theme }) => ({
  background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
  color: theme.palette.common.white,
  padding: '8px 16px',
  '&:hover': {
    background: 'linear-gradient(45deg, #5000d6 30%, #00c060 90%)',
  },
}));

const WalletInfo = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  background: 'rgba(255, 255, 255, 0.05)',
  borderRadius: theme.shape.borderRadius,
  padding: '4px 12px',
  marginRight: theme.spacing(2),
}));

const Navbar = () => {
  const { account, balance, connectWallet, disconnectWallet, isConnecting } = useContext(WalletContext);
  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const formatAddress = (address) => {
    if (!address) return '';
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
  };

  return (
    <StyledAppBar position="static">
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2, display: { xs: 'block', md: 'none' } }}
        >
          <MenuIcon />
        </IconButton>
        
        <LogoText variant="h5" component={Link} to="/" sx={{ textDecoration: 'none' }}>
          NoMercyChain
        </LogoText>
        
        <Box sx={{ flexGrow: 1 }} />
        
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <Tooltip title="Notifications">
            <IconButton color="inherit" sx={{ mr: 2 }}>
              <NotificationsIcon />
            </IconButton>
          </Tooltip>
          
          {account ? (
            <>
              <WalletInfo>
                <Typography variant="body2" sx={{ mr: 1 }}>
                  {formatAddress(account)}
                </Typography>
                <Typography variant="body2" color="secondary">
                  {parseFloat(balance).toFixed(4)} NMX
                </Typography>
              </WalletInfo>
              
              <Tooltip title="Account settings">
                <IconButton onClick={handleMenu} size="small">
                  <Avatar sx={{ width: 32, height: 32, background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)' }}>
                    {account ? account.substring(2, 4).toUpperCase() : ''}
                  </Avatar>
                </IconButton>
              </Tooltip>
              
              <Menu
                anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={handleClose}
                transformOrigin={{ horizontal: 'right', vertical: 'top' }}
                anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
              >
                <MenuItem component={Link} to="/wallet" onClick={handleClose}>My Wallet</MenuItem>
                <MenuItem component={Link} to="/settings" onClick={handleClose}>Settings</MenuItem>
                <MenuItem onClick={() => { disconnectWallet(); handleClose(); }}>Disconnect</MenuItem>
              </Menu>
            </>
          ) : (
            <WalletButton 
              variant="contained" 
              onClick={connectWallet}
              disabled={isConnecting}
              startIcon={isConnecting ? <CircularProgress size={20} color="inherit" /> : null}
            >
              {isConnecting ? 'Connecting...' : 'Connect Wallet'}
            </WalletButton>
          )}
        </Box>
      </Toolbar>
    </StyledAppBar>
  );
};

export default Navbar;