import React, { useState, useEffect } from 'react';
import { Box, Typography, Button, Paper, Grid, IconButton } from '@mui/material';
import { styled } from '@mui/system';
import { Line } from 'react-chartjs-2';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import SendIcon from '@mui/icons-material/Send';
import EditIcon from '@mui/icons-material/Edit';
import { useNavigate } from 'react-router-dom';
import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale } from 'chart.js';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';

// Registrar os componentes do Chart.js
Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale);

const MainContent = styled(Box)({
  background: '#f5f5f5',
  minHeight: '100vh',
  padding: '40px',
  marginLeft: '240px', // Considera o espaço da sidebar
  marginTop: '64px', // Considera o espaço da navbar
});

const BalancePaper = styled(Paper)({
  padding: '30px',
  background: '#ffffff',
  textAlign: 'center',
  borderRadius: '12px',
  boxShadow: '0px 4px 10px rgba(0,0,0,0.1)',
  marginBottom: '30px',
});

const StyledButton = styled(Button)({
  marginTop: '20px',
  padding: '10px',
  borderRadius: '10px',
  color: '#fff',
  textTransform: 'none',
  transition: '0.3s',
  '&:hover': {
    background: '#303f9f',
  },
});

const CardPaper = styled(Paper)({
  padding: '20px',
  textAlign: 'center',
  borderRadius: '12px',
  boxShadow: '0px 4px 10px rgba(0,0,0,0.1)',
  height: '200px',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'center',
});

const Home = () => {
  const [balance, setBalance] = useState(0.0); // Saldo inicial
  const [balanceData, setBalanceData] = useState({
    labels: [],
    datasets: [],
  });
  const navigate = useNavigate();

  useEffect(() => {
    // Simular uma chamada para buscar o saldo do banco
    setTimeout(() => {
      setBalance(1200.50); // Valor simulado
    }, 1000);

    // Simular a busca do histórico de saldo
    const fetchData = async () => {
      const data = await getBalanceHistoryData();
      setBalanceData({
        labels: data.map((item) => item.date),
        datasets: [
          {
            label: 'Saldo ao longo do tempo',
            data: data.map((item) => item.balance),
            borderColor: 'rgba(75, 192, 192, 1)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            fill: true,
            tension: 0.4,
          },
        ],
      });
    };

    fetchData();
  }, []);

  const getBalanceHistoryData = () => {
    // Simular dados de histórico de saldo
    return [
      { date: '01/10', balance: 1000 },
      { date: '05/10', balance: 1200 },
      { date: '10/10', balance: 800 },
      { date: '15/10', balance: 1500 },
      { date: '20/10', balance: 1300 },
    ];
  };

  const handleNavigate = (path) => {
    navigate(path);
  };

  return (
    <MainContent>
            
      <Navbar />
      <Sidebar />
      {/* Mostrando o saldo disponível */}
      <BalancePaper elevation={3}>
        <Typography variant="h6" color="textSecondary" sx={{ mb: 2 }}>
          Saldo Disponível
        </Typography>
        <Typography variant="h3" color="primary" fontWeight="bold">
          R$ {balance.toFixed(2)}
        </Typography>
        {/* Gráfico de Histórico de Saldo */}
        <Box sx={{ mt: 4 }}>
          <Typography variant="h6" color="textSecondary" sx={{ mb: 2 }}>
            Histórico de Saldo
          </Typography>
          <Paper elevation={3} sx={{ padding: '20px', maxWidth: '600px', margin: '0 auto' }}>
            <Line data={balanceData} width={400} height={300} />
          </Paper>
        </Box>
      </BalancePaper>

      {/* Funções principais */}
      <Grid container spacing={3}>
        {/* Botão Extrato */}
        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton color="primary" onClick={() => handleNavigate('/extract')}>
              <AccountBalanceIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Extrato</Typography>
            <StyledButton
              fullWidth
              variant="contained"
              color="primary"
              onClick={() => handleNavigate('/extract')}
            >
              Acessar Extrato
            </StyledButton>
          </CardPaper>
        </Grid>

        {/* Botão Transferência */}
        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton color="primary" onClick={() => handleNavigate('/transfer')}>
              <SendIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Transferir</Typography>
            <StyledButton
              fullWidth
              variant="contained"
              color="primary"
              onClick={() => handleNavigate('/transfer')}
            >
              Fazer Transferência
            </StyledButton>
          </CardPaper>
        </Grid>

        {/* Botão Editar Perfil */}
        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton color="primary" onClick={() => handleNavigate('/user/edit')}>
              <EditIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Editar Perfil</Typography>
            <StyledButton
              fullWidth
              variant="contained"
              color="primary"
              onClick={() => handleNavigate('/user/edit')}
            >
              Editar Perfil
            </StyledButton>
          </CardPaper>
        </Grid>
      </Grid>
    </MainContent>
  );
};

export default Home;
