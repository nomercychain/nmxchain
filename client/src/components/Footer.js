import React from 'react';
import { Box, Typography, Link, Container, Grid, IconButton } from '@mui/material';
import { styled } from '@mui/material/styles';
import TwitterIcon from '@mui/icons-material/Twitter';
import GitHubIcon from '@mui/icons-material/GitHub';
import TelegramIcon from '@mui/icons-material/Telegram';
import RedditIcon from '@mui/icons-material/Reddit';

const StyledFooter = styled(Box)(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  padding: theme.spacing(6, 0),
  marginTop: 'auto',
  borderTop: `1px solid ${theme.palette.divider}`,
}));

const SocialIcon = styled(IconButton)(({ theme }) => ({
  color: theme.palette.text.secondary,
  '&:hover': {
    color: theme.palette.primary.main,
  },
}));

const Footer = () => {
  return (
    <StyledFooter>
      <Container maxWidth="lg">
        <Grid container spacing={4} justifyContent="space-between">
          <Grid item xs={12} sm={4}>
            <Typography variant="h6" color="text.primary" gutterBottom>
              NoMercyChain
            </Typography>
            <Typography variant="body2" color="text.secondary">
              The next-generation blockchain platform with AI-powered smart contracts, 
              decentralized AI agents, and Layer 3 solutions.
            </Typography>
            <Box sx={{ mt: 2 }}>
              <SocialIcon aria-label="twitter">
                <TwitterIcon />
              </SocialIcon>
              <SocialIcon aria-label="github">
                <GitHubIcon />
              </SocialIcon>
              <SocialIcon aria-label="telegram">
                <TelegramIcon />
              </SocialIcon>
              <SocialIcon aria-label="reddit">
                <RedditIcon />
              </SocialIcon>
            </Box>
          </Grid>
          <Grid item xs={6} sm={2}>
            <Typography variant="h6" color="text.primary" gutterBottom>
              Resources
            </Typography>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Documentation
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Whitepaper
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              API Reference
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              GitHub
            </Link>
          </Grid>
          <Grid item xs={6} sm={2}>
            <Typography variant="h6" color="text.primary" gutterBottom>
              Community
            </Typography>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Discord
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Forum
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Events
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Blog
            </Link>
          </Grid>
          <Grid item xs={6} sm={2}>
            <Typography variant="h6" color="text.primary" gutterBottom>
              Company
            </Typography>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              About
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Careers
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Contact
            </Link>
            <Link href="#" color="text.secondary" display="block" sx={{ mb: 1 }}>
              Privacy Policy
            </Link>
          </Grid>
        </Grid>
        <Box mt={5}>
          <Typography variant="body2" color="text.secondary" align="center">
            {'© '}
            {new Date().getFullYear()}
            {' NoMercyChain. All rights reserved.'}
          </Typography>
        </Box>
      </Container>
    </StyledFooter>
  );
};

export default Footer;