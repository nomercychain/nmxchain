import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { 
  Box, 
  Drawer, 
  List, 
  ListItem, 
  ListItemIcon, 
  ListItemText, 
  Divider, 
  Typography,
  useMediaQuery
} from '@mui/material';
import { styled, useTheme } from '@mui/material/styles';
import DashboardIcon from '@mui/icons-material/Dashboard';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';
import HowToVoteIcon from '@mui/icons-material/HowToVote';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import CodeIcon from '@mui/icons-material/Code';
import LayersIcon from '@mui/icons-material/Layers';
import VerifiedUserIcon from '@mui/icons-material/VerifiedUser';
import SettingsIcon from '@mui/icons-material/Settings';
import LocalFireDepartmentIcon from '@mui/icons-material/LocalFireDepartment';

const drawerWidth = 240;

const StyledDrawer = styled(Drawer)(({ theme }) => ({
  width: drawerWidth,
  flexShrink: 0,
  '& .MuiDrawer-paper': {
    width: drawerWidth,
    boxSizing: 'border-box',
    background: theme.palette.background.paper,
    borderRight: `1px solid ${theme.palette.divider}`,
  },
}));

const StyledListItem = styled(ListItem)(({ theme, active }) => ({
  borderRadius: theme.shape.borderRadius,
  margin: '4px 8px',
  color: active ? theme.palette.primary.main : theme.palette.text.primary,
  '&:hover': {
    backgroundColor: 'rgba(98, 0, 234, 0.08)',
  },
  ...(active && {
    backgroundColor: 'rgba(98, 0, 234, 0.08)',
    '&::before': {
      content: '""',
      position: 'absolute',
      left: 0,
      top: '25%',
      height: '50%',
      width: 4,
      backgroundColor: theme.palette.primary.main,
      borderRadius: '0 4px 4px 0',
    },
  }),
}));

const menuItems = [
  { text: 'Dashboard', icon: <DashboardIcon />, path: '/' },
  { text: 'Wallet', icon: <AccountBalanceWalletIcon />, path: '/wallet' },
  { text: 'Staking', icon: <LocalFireDepartmentIcon />, path: '/staking' },
  { text: 'Governance', icon: <HowToVoteIcon />, path: '/governance' },
  { text: 'AI Agents', icon: <SmartToyIcon />, path: '/ai-agents' },
  { text: 'Smart Contracts', icon: <CodeIcon />, path: '/smart-contracts' },
  { text: 'HyperChains', icon: <LayersIcon />, path: '/hyperchains' },
  { text: 'TruthGPT Oracle', icon: <VerifiedUserIcon />, path: '/oracle' },
  { text: 'Settings', icon: <SettingsIcon />, path: '/settings' },
];

const Sidebar = ({ mobileOpen, handleDrawerToggle }) => {
  const location = useLocation();
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  const drawer = (
    <>
      <Box sx={{ p: 2, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <Typography variant="h6" sx={{ fontWeight: 700 }}>
          NoMercyChain
        </Typography>
      </Box>
      <Divider />
      <List>
        {menuItems.map((item) => (
          <StyledListItem
            button
            key={item.text}
            component={Link}
            to={item.path}
            active={location.pathname === item.path ? 1 : 0}
            onClick={isMobile ? handleDrawerToggle : undefined}
          >
            <ListItemIcon sx={{ color: 'inherit', minWidth: 40 }}>
              {item.icon}
            </ListItemIcon>
            <ListItemText primary={item.text} />
          </StyledListItem>
        ))}
      </List>
      <Box sx={{ flexGrow: 1 }} />
      <Box sx={{ p: 2 }}>
        <Typography variant="caption" color="text.secondary" align="center" display="block">
          NoMercyChain v0.1.0
        </Typography>
      </Box>
    </>
  );

  return (
    <Box
      component="nav"
      sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
    >
      {/* Mobile drawer */}
      <Drawer
        variant="temporary"
        open={mobileOpen}
        onClose={handleDrawerToggle}
        ModalProps={{
          keepMounted: true, // Better open performance on mobile
        }}
        sx={{
          display: { xs: 'block', sm: 'none' },
          '& .MuiDrawer-paper': { 
            width: drawerWidth,
            boxSizing: 'border-box',
            background: theme.palette.background.paper,
          },
        }}
      >
        {drawer}
      </Drawer>
      
      {/* Desktop drawer */}
      <StyledDrawer
        variant="permanent"
        sx={{
          display: { xs: 'none', sm: 'block' },
        }}
        open
      >
        {drawer}
      </StyledDrawer>
    </Box>
  );
};

export default Sidebar;