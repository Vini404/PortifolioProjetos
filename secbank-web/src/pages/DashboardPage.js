import React, { useEffect, useState } from 'react';
import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Box,
  Drawer,
  List,
  ListItem,
  ListItemText,
  CssBaseline,
  CircularProgress,
} from '@mui/material';
import { styled } from '@mui/system';
import MenuIcon from '@mui/icons-material/Menu';
import TransferPage from './TransferPage';
import api from '../api/axiosBase';

const drawerWidth = 240;

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: `-${drawerWidth}px`,
    ...(open && {
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
      marginLeft: 0,
    }),
  })
);

const DashboardPage = () => {
  const [open, setOpen] = useState(false);
  const [currentPage, setCurrentPage] = useState('Transfer');
  const [balance, setBalance] = useState(null);
  const [loading, setLoading] = useState(true);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const handleNavigation = (page) => {
    setCurrentPage(page);
    setOpen(false);
  };

  const renderPage = () => {
    switch (currentPage) {
      case 'Transfer':
        return <TransferPage />;
      default:
        return <TransferPage />;
    }
  };

  useEffect(() => {
    const fetchBalance = async () => {
      setLoading(true);
      try {
        // Obter o CustomerID
        const customerResponse = await api.get('/customer/info');
        if (!customerResponse.data.success) {
          alert('Erro ao obter informações do cliente');
          return;
        }

        const customerID = customerResponse.data.data.ID;

        // Obter o AccountID
        const accountResponse = await api.get(`/account/${customerID}/information`);
        if (!accountResponse.data.success) {
          alert('Erro ao obter informações da conta');
          return;
        }

        const accountID = accountResponse.data.data.IDAccount;

        // Obter o saldo
        const balanceResponse = await api.get(`/balance/${accountID}`);
        if (!balanceResponse.data.success) {
          alert('Erro ao obter saldo');
          return;
        }

        setBalance(balanceResponse.data.data.Amount);
      } catch (error) {
        alert(error.response?.data?.messageError || 'Erro ao buscar o saldo');
      } finally {
        setLoading(false);
      }
    };

    fetchBalance();
  }, []);

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar position="fixed">
        <Toolbar>
          <IconButton edge="start" color="inherit" aria-label="menu" onClick={handleDrawerOpen}>
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div">
            Dashboard
          </Typography>
          <Box sx={{ marginLeft: 'auto' }}>
            {loading ? (
              <CircularProgress color="inherit" size={24} />
            ) : (
              <Typography variant="h6">Saldo disponível: R$ {balance}</Typography>
            )}
          </Box>
        </Toolbar>
      </AppBar>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            backgroundColor: '#ffffff',
            borderRight: '1px solid #e0e0e0',
          },
        }}
        variant="persistent"
        anchor="left"
        open={open}
      >
        <Toolbar />
        <List>
          <ListItem button onClick={() => handleNavigation('Transfer')}>
            <ListItemText primary="Transferir" />
          </ListItem>
        </List>
      </Drawer>
      <Main open={open}>
        <Toolbar />
        {renderPage()}
      </Main>
    </Box>
  );
};

export default DashboardPage;
