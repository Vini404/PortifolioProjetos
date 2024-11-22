import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Paper,
  Grid,
  IconButton,
  CircularProgress,
  Tooltip,
} from '@mui/material';
import { styled } from '@mui/system';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import SendIcon from '@mui/icons-material/Send';
import EditIcon from '@mui/icons-material/Edit';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import { useNavigate } from 'react-router-dom';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import api from '../api/axiosBase';

const MainContent = styled(Box)({
  background: '#f5f5f5',
  minHeight: '100vh',
  padding: '40px',
  marginLeft: '240px',
  marginTop: '64px',
});

const BalancePaper = styled(Paper)({
  padding: '30px',
  background: '#ffffff',
  textAlign: 'center',
  borderRadius: '12px',
  boxShadow: '0px 4px 10px rgba(0,0,0,0.1)',
  marginBottom: '30px',
});

const AccountNumberContainer = styled(Box)({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  marginTop: '20px',
});

const StyledButton = styled(Button)({
  marginTop: '20px',
  padding: '10px',
  borderRadius: '10px',
  color: '#ffffff',
  backgroundColor: '#303f9f', // Azul padrão do botão
  textTransform: 'none',
  transition: '0.3s',
  '&:hover': {
    backgroundColor: '#283593', // Azul mais escuro no hover
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
  const [balance, setBalance] = useState(0.0);
  const [accountNumber, setAccountNumber] = useState('');
  const [loading, setLoading] = useState(true);
  const [copySuccess, setCopySuccess] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAccountInfo = async () => {
      setLoading(true);
      try {
        const customerResponse = await api.get('/customer/info');
        if (!customerResponse.ok) {
          toast.error('Erro ao obter informações do cliente');
          return;
        }
        const customerID = customerResponse.result.data.ID;

        const accountResponse = await api.get(`/account/${customerID}/information`);
        if (!accountResponse.ok) {
          toast.error('Erro ao obter informações da conta');
          return;
        }
        const accountID = accountResponse.result.data.IDAccount;
        const accountNumber = accountResponse.result.data.AccountNumber;
        setAccountNumber(accountNumber);

        const balanceResponse = await api.get(`/balance/${accountID}`);
        if (!balanceResponse.ok) {
          toast.error('Erro ao obter saldo');
          return;
        }
        setBalance(balanceResponse.result.data.Amount);

        toast.success('Dados carregados com sucesso!');
      } catch (error) {
        const errorMessage = JSON.parse(error.message).messageError || 'Erro ao buscar o saldo';
        toast.error(errorMessage);
      } finally {
        setLoading(false);
      }
    };

    fetchAccountInfo();
  }, []);

  const handleCopy = () => {
    navigator.clipboard.writeText(accountNumber).then(() => {
      setCopySuccess(true);
      toast.success('Número da conta copiado!');
      setTimeout(() => setCopySuccess(false), 2000);
    });
  };

  const handleNavigate = (path) => {
    navigate(path);
  };

  return (
    <MainContent>
      <ToastContainer
        position="top-right"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      <Navbar />
      <Sidebar />
      <BalancePaper elevation={3}>
        <Typography variant="h6" color="textSecondary" sx={{ mb: 2 }}>
          Saldo Disponível
        </Typography>
        {loading ? (
          <CircularProgress color="primary" />
        ) : (
          <>
            <Typography variant="h3" sx={{ color: '#303f9f', fontWeight: 'bold' }}>
              R$ {balance.toFixed(2)}
            </Typography>
            <AccountNumberContainer>
              <Typography variant="h4" sx={{ color: '#303f9f', fontWeight: 'bold', mr: 2 }}>
                Conta: {accountNumber}
              </Typography>
              <Tooltip title={copySuccess ? 'Copiado!' : 'Copiar'}>
                <IconButton onClick={handleCopy} sx={{ color: '#303f9f' }}>
                  <ContentCopyIcon />
                </IconButton>
              </Tooltip>
            </AccountNumberContainer>
          </>
        )}
      </BalancePaper>

      <Grid container spacing={3}>
        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton sx={{ color: '#303f9f' }} onClick={() => handleNavigate('/extract')}>
              <AccountBalanceIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Extrato</Typography>
            <StyledButton
              fullWidth
              variant="contained"
              onClick={() => handleNavigate('/extract')}
            >
              Acessar Extrato
            </StyledButton>
          </CardPaper>
        </Grid>

        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton sx={{ color: '#303f9f' }} onClick={() => handleNavigate('/transfer')}>
              <SendIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Transferir</Typography>
            <StyledButton
              fullWidth
              variant="contained"
              onClick={() => handleNavigate('/transfer')}
            >
              Fazer Transferência
            </StyledButton>
          </CardPaper>
        </Grid>

        <Grid item xs={12} sm={4}>
          <CardPaper elevation={3}>
            <IconButton sx={{ color: '#303f9f' }} onClick={() => handleNavigate('/user/edit')}>
              <EditIcon sx={{ fontSize: 50 }} />
            </IconButton>
            <Typography variant="h6" sx={{ mt: 2 }}>Editar Perfil</Typography>
            <StyledButton
              fullWidth
              variant="contained"
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
