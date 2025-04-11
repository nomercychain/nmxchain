import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import React, { useEffect, useState } from 'react';
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import './App.css';

// Components
import Footer from './components/Footer';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';

// Pages
import AIAgents from './pages/AIAgents';
import Dashboard from './pages/Dashboard';
import Wallet from './pages/Wallet';
import Staking from './pages/Staking';
import Governance from './pages/Governance';
import SmartContracts from './pages/SmartContracts';
import HyperChains from './pages/HyperChains';
import Oracle from './pages/Oracle';
import Settings from './pages/Settings';

// Context
import { WalletContext } from './context/WalletContext';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#6200ea',
    },
    secondary: {
      main: '#00e676',
    },
    background: {
      default: '#121212',
      paper: '#1e1e1e',
    },
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontSize: '2.5rem',
      fontWeight: 700,
    },
    h2: {
      fontSize: '2rem',
      fontWeight: 600,
    },
    h3: {
      fontSize: '1.75rem',
      fontWeight: 600,
    },
    h4: {
      fontSize: '1.5rem',
      fontWeight: 600,
    },
    h5: {
      fontSize: '1.25rem',
      fontWeight: 600,
    },
    h6: {
      fontSize: '1rem',
      fontWeight: 600,
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: 8,
          textTransform: 'none',
          fontWeight: 600,
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: 12,
          boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
        },
      },
    },
  },
});

function App() {
  const [account, setAccount] = useState(null);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [balance, setBalance] = useState(null);
  const [chainId, setChainId] = useState(null);
  const [isConnecting, setIsConnecting] = useState(false);

  const connectWallet = async () => {
    try {
      setIsConnecting(true);
      
      // Simulate connection delay
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Mock wallet data
      const mockAccount = '0x' + Math.random().toString(16).substring(2, 42);
      const mockBalance = (Math.random() * 1000).toFixed(2);
      
      setAccount(mockAccount);
      setBalance(mockBalance);
      setChainId('0x1'); // Ethereum Mainnet
      
      // Save to localStorage for persistence
      localStorage.setItem('walletAccount', mockAccount);
      
      setIsConnecting(false);
    } catch (error) {
      console.error('Error connecting wallet:', error);
      setIsConnecting(false);
    }
  };

  const disconnectWallet = () => {
    setAccount(null);
    setProvider(null);
    setSigner(null);
    setChainId(null);
    setBalance(null);
    
    // Remove from localStorage
    localStorage.removeItem('walletAccount');
  };

  useEffect(() => {
    // Check if wallet was previously connected (for demo purposes)
    const checkPreviousConnection = async () => {
      const savedAccount = localStorage.getItem('walletAccount');
      if (savedAccount) {
        setAccount(savedAccount);
        setBalance((Math.random() * 1000).toFixed(2));
        setChainId('0x1');
      }
    };
    
    checkPreviousConnection();
  }, []);

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <WalletContext.Provider value={{ account, provider, signer, balance, chainId, connectWallet, disconnectWallet, isConnecting }}>
        <Router>
          <div className="app">
            <Navbar />
            <div className="content-container">
              <Sidebar />
              <main className="main-content">
                <Routes>
                  <Route path="/" element={<Dashboard />} />
                  <Route path="/wallet" element={<Wallet />} />
                  <Route path="/staking" element={<Staking />} />
                  <Route path="/governance" element={<Governance />} />
                  <Route path="/ai-agents" element={<AIAgents />} />
                  <Route path="/smart-contracts" element={<SmartContracts />} />
                  <Route path="/hyperchains" element={<HyperChains />} />
                  <Route path="/oracle" element={<Oracle />} />
                  <Route path="/settings" element={<Settings />} />
                </Routes>
              </main>
            </div>
            <Footer />
          </div>
        </Router>
      </WalletContext.Provider>
    </ThemeProvider>
  );
}

export default App;