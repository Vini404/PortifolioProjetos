import React from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Drawer } from '@mui/material';
import { Home, AccountBalance, Payment, ExitToApp } from '@mui/icons-material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useNavigate } from 'react-router-dom';
import { styled } from '@mui/system';
import UserProfileEdit from '../pages/UserProfileEdit';

const DrawerContainer = styled(Drawer)(({ theme }) => ({
  '& .MuiDrawer-paper': {
    width: 240,
    boxSizing: 'border-box',
    backgroundColor: theme.palette.background.default,
    borderRight: '1px solid #e0e0e0',
    paddingTop: '10px', // Adiciona padding-top de 10 pixels
    marginTop: '50px', // Adiciona margin-top de 50 pixels
  },
}));

const Sidebar = () => {
  const navigate = useNavigate();

  const menuItems = [
    { text: 'Home', icon: <Home />, path: '/home' },
    { text: 'Extrato', icon: <AccountBalance />, path: '/extract' },
    { text: 'Transferir dinheiro', icon: <Payment />, path: '/transfer' },
    { text: 'Perfil', icon: <AccountCircleIcon />, path: '/user/edit' },
    { text: 'Sair', icon: <ExitToApp />, path: '/' },
  ];

  return (
    <DrawerContainer
      variant="permanent"
      sx={{
        flexShrink: 0,
      }}
    >
      <List>
        {menuItems.map((item) => (
          <ListItem
            button
            key={item.text}
            onClick={() => navigate(item.path)}
            sx={{
              '&:hover': {
                backgroundColor: '#f5f5f5', // Cor de fundo ao passar o mouse
              },
            }}
          >
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItem>
        ))}
      </List>
    </DrawerContainer>
  );
};

export default Sidebar;
