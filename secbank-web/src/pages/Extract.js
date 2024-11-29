import React, { useState, useEffect } from 'react';
import {
  Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button,
} from '@mui/material';
import { styled } from '@mui/system';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import api from '../api/axiosBase';

const PageContainer = styled(Box)({
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  minHeight: '100vh',
  backgroundColor: '#f5f5f5',
  padding: '20px',
});

const TransactionsPaper = styled(Paper)({
  padding: '20px',
  maxWidth: '800px',
  borderRadius: '12px',
  boxShadow: '0px 10px 30px rgba(0,0,0,0.1)',
  backgroundColor: '#ffffff',
});

const Extract = () => {
  const [transactions, setTransactions] = useState([]);
  const [totalBalance, setTotalBalance] = useState(0);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchExtractData = async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        toast.error('Token não encontrado!');
        return;
      }

      try {
        const userInfoResponse = await api.get('/customer/info');
        if (!userInfoResponse.ok) {
          toast.error('Erro ao obter informações do usuário');
          return;
        }

        const accountId = userInfoResponse.result.data.ID;

        const extractResponse = await api.get(`/balance/extract/${accountId}`);
        if (!extractResponse.ok) {
          toast.error('Erro ao obter extrato');
          return;
        }

        const transactionList = extractResponse.result.data.map((transaction) => ({
          id: transaction.ID,
          type: transaction.OperationName,
          details: transaction.TransferType,
          amount: transaction.Amount,
        }));

        setTransactions(transactionList);

        const balance = transactionList.reduce((acc, transaction) => acc + transaction.amount, 0);
        setTotalBalance(balance);

        toast.success('Dados carregados com sucesso!');
      } catch (error) {
        let errorMessage = 'Erro inesperado';
        try {
          const parsedError = JSON.parse(error.message);
          errorMessage = parsedError.messageError || errorMessage;
        } catch (parseError) {
          // Mantém o erro genérico
        }
        toast.error(errorMessage);
      } finally {
        setLoading(false);
      }
    };

    fetchExtractData();
  }, []);

  if (loading) {
    return <Typography>Carregando...</Typography>;
  }

  return (
    <PageContainer>
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
      <TransactionsPaper elevation={6}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Extrato de Transferências
        </Typography>

        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Nome da Operação</TableCell>
                <TableCell>Detalhes da Operação</TableCell>
                <TableCell align="right">Valor</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {transactions.map((transaction) => (
                <TableRow key={transaction.id}>
                  <TableCell>{transaction.type}</TableCell>
                  <TableCell>{transaction.details}</TableCell>
                  <TableCell
                    align="right"
                    sx={{
                      color: transaction.details.toLowerCase().includes('enviado')
                        ? 'red'
                        : transaction.details.toLowerCase().includes('recebido')
                        ? 'green'
                        : 'inherit',
                    }}
                  >
                    {transaction.amount < 0 ? `- R$ ${Math.abs(transaction.amount).toFixed(2)}` : `R$ ${transaction.amount.toFixed(2)}`}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        <Button variant="contained" color="primary" fullWidth sx={{ mt: 3 }}>
          Voltar ao Dashboard
        </Button>
      </TransactionsPaper>
    </PageContainer>
  );
};

export default Extract;
