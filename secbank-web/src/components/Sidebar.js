import React, { useState } from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Drawer, Typography, Box } from '@mui/material';
import { Home, AccountBalance, Payment, ExitToApp } from '@mui/icons-material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useNavigate } from 'react-router-dom';
import { styled } from '@mui/system';

const DrawerContainer = styled(Drawer)(({ theme }) => ({
  '& .MuiDrawer-paper': {
    width: 280,
    boxSizing: 'border-box',
    backgroundColor: '#f5f5f5', // Fundo claro
    color: '#5c5c5c', // Texto em cinza escuro
    borderRight: '1px solid #e0e0e0',
    paddingTop: theme.spacing(2),
  },
}));

const ListItemStyled = styled(ListItem)(({ theme }) => ({
  marginBottom: theme.spacing(2),
  borderRadius: theme.shape.borderRadius,
  transition: 'background-color 0.3s ease', // Suaviza a mudança no item ativo e no hover
  '&:hover': {
    backgroundColor: '#dcdcdc', // Escurece levemente no hover
  },
}));

const Sidebar = () => {
  const navigate = useNavigate();
  const [selectedItem, setSelectedItem] = useState('Home'); // "Home" selecionado por padrão

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
      <Box sx={{ textAlign: 'center', padding: '20px 0' }}>
        <Typography variant="h6" fontWeight="bold" color="#000">
          Sec Bank
        </Typography>
      </Box>
      <List sx={{ padding: '0 20px' }}>
        {menuItems.map((item) => (
          <ListItemStyled
            button
            key={item.text}
            onClick={() => {
              setSelectedItem(item.text);
              navigate(item.path);
            }}
            sx={{
              color: selectedItem === item.text ? '#000' : '#5c5c5c', // Texto preto no item ativo
            }}
          >
            <ListItemIcon
              sx={{
                color: selectedItem === item.text ? '#000' : '#5c5c5c', // Ícones do item ativo em preto
              }}
            >
              {item.icon}
            </ListItemIcon>
            <ListItemText
              primary={item.text}
              primaryTypographyProps={{
                fontWeight: selectedItem === item.text ? 'bold' : 'normal',
              }}
            />
          </ListItemStyled>
        ))}
      </List>
    </DrawerContainer>
  );
};

export default Sidebar;
