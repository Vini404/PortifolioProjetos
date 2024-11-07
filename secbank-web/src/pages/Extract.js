import React, { useState, useEffect } from 'react';
import {
  Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TableFooter, Button
} from '@mui/material';
import { styled } from '@mui/system';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import api from '../api/axiosBase'

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
        alert('Token não encontrado!');
        return;
      }

      try {
        // Primeiro, obtenha as informações do usuário para encontrar o ID da conta
        const userInfoResponse = await api.get('/customer/info');
        if (!userInfoResponse.ok) {
          alert( 'Erro ao obter informações do usuário');
          return;
        }

        const accountId = userInfoResponse.result.data.ID; // Pegue o ID da conta do usuário

        // Agora, obtenha o extrato da conta usando o ID
        const extractResponse = await api.get(`/balance/extract/${accountId}`);

        if (!extractResponse.ok) {
          alert('Erro ao obter extrato');
          return;
        }

        // Formate os dados recebidos para exibir na tabela
        const transactionList = extractResponse.result.data.map((transaction) => ({
          id: transaction.ID,
          type: transaction.OperationName,
          details: transaction.TransferType,
          amount: transaction.Amount,
        }));

        setTransactions(transactionList);

        // Calcule o saldo total
        const balance = transactionList.reduce((acc, transaction) => acc + transaction.amount, 0);
        setTotalBalance(balance);
      } catch (error) {
        const errorMessage = JSON.parse(error.message).messageError
        alert(errorMessage);
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
                  <TableCell align="right" sx={{ color: transaction.amount < 0 ? 'red' : 'green' }}>
                    {transaction.amount < 0 ? `- R$ ${Math.abs(transaction.amount).toFixed(2)}` : `R$ ${transaction.amount.toFixed(2)}`}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TableCell colSpan={2}>
                  <Typography variant="h6">Saldo Total da Conta</Typography>
                </TableCell>
                <TableCell align="right" sx={{ color: totalBalance < 0 ? 'red' : 'green' }}>
                  {totalBalance < 0 ? `- R$ ${Math.abs(totalBalance).toFixed(2)}` : `R$ ${totalBalance.toFixed(2)}`}
                </TableCell>
              </TableRow>
            </TableFooter>
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
