import React, { useState } from 'react';
import {
  Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TableFooter, Button
} from '@mui/material';
import { styled } from '@mui/system';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';

// Estilo para o container da página
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
  // Simulação de transferências realizadas
  const [transactions, setTransactions] = useState([
    { id: 1, type: 'Envio de Transferência', details: 'Enviado para João da Silva - Conta 54321', amount: -150 },
    { id: 2, type: 'Recebimento de Transferência', details: 'Recebido de Maria Oliveira - Conta 12345', amount: 200 },
    { id: 3, type: 'Envio de Transferência', details: 'Enviado para Lucas Lima - Conta 98765', amount: -50 },
    { id: 4, type: 'Recebimento de Transferência', details: 'Recebido de Ana Paula - Conta 87654', amount: 300 },
  ]);

  // Cálculo do valor total da conta
  const totalBalance = transactions.reduce((acc, transaction) => acc + transaction.amount, 0);

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
