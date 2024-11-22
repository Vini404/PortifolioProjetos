import React from 'react';
import { AppBar, Toolbar, Box } from '@mui/material';

const Navbar = () => {
  return (
    <AppBar
      position="fixed"
      sx={{
        zIndex: (theme) => theme.zIndex.drawer + 1,
        backgroundColor: '#303f9f',
        boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.2)',
        height: 80, // Aumenta a altura do Navbar para acomodar a logo maior
        display: 'flex',
        justifyContent: 'center',
      }}
    >
      <Toolbar
        sx={{
          display: 'flex',
          justifyContent: 'center', // Centraliza o conteúdo horizontalmente
          alignItems: 'center', // Centraliza o conteúdo verticalmente
        }}
      >
        <Box
          component="img"
          src="/images/logo/logo.png" // Caminho relativo à pasta public
          alt="Sec Bank Logo"
          sx={{
            width: 200, // Aumenta a largura da logo
            height: 60, // Aumenta a altura da logo
            objectFit: 'contain',
          }}
        />
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
