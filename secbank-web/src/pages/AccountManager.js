import React, { useState } from 'react';
import {
  Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow,
  Button, Checkbox, Modal, TextField, TableFooter, TablePagination
} from '@mui/material';
import { styled } from '@mui/system';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';

// Estilos para o container principal e o modal
const PageContainer = styled(Box)({
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  minHeight: '100vh',
  backgroundColor: '#f5f5f5',
  padding: '20px',
});

const AccountsPaper = styled(Paper)({
  padding: '20px',
  maxWidth: '800px',
  borderRadius: '12px',
  boxShadow: '0px 10px 30px rgba(0,0,0,0.1)',
  backgroundColor: '#ffffff',
});

const ModalContainer = styled(Box)({
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  backgroundColor: 'white',
  padding: '20px',
  boxShadow: '0px 4px 10px rgba(0,0,0,0.2)',
  borderRadius: '8px',
});

const AccountMannager = () => {
  const [accounts, setAccounts] = useState([
    { number: '12345', digit: '1', balance: 1000.0, status: true, description: 'Poupança' },
    { number: '67890', digit: '9', balance: 500.0, status: false, description: 'Corrente' },
  ]);

  const [open, setOpen] = useState(false);
  const [newAccountDescription, setNewAccountDescription] = useState('');
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(5);

  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const handleAddAccount = () => {
    // Lógica para adicionar uma nova conta
    const newAccount = {
      number: Math.floor(10000 + Math.random() * 90000).toString(), // Gerar um número aleatório de conta
      digit: Math.floor(Math.random() * 9).toString(),
      balance: 0,
      status: true,
      description: newAccountDescription,
    };
    setAccounts([...accounts, newAccount]);
    setNewAccountDescription('');
    handleClose(); // Fecha o modal após adicionar
  };

  const handleStatusChange = (index) => {
    const updatedAccounts = [...accounts];
    updatedAccounts[index].status = !updatedAccounts[index].status;
    setAccounts(updatedAccounts);
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  return (
    <PageContainer>
      <Navbar />
      <Sidebar />
      <AccountsPaper elevation={6}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Gerenciamento de Contas
        </Typography>

        <Button variant="contained" color="primary" onClick={handleOpen} sx={{ mb: 2 }}>
          Adicionar Nova Conta
        </Button>

        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Conta</TableCell>
                <TableCell>Valor Guardado</TableCell>
                <TableCell>Situação da Conta</TableCell>
                <TableCell>Finalidade da Conta</TableCell>
                <TableCell>Ativar/Desativar</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {accounts.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((account, index) => (
                <TableRow key={index}>
                  <TableCell>
                    {account.number} - {account.digit}
                  </TableCell>
                  <TableCell>R$ {account.balance.toFixed(2)}</TableCell>
                  <TableCell>{account.status ? 'Ativada' : 'Desativada'}</TableCell>
                  <TableCell>{account.description}</TableCell>
                  <TableCell>
                    <Checkbox
                      checked={account.status}
                      onChange={() => handleStatusChange(index)}
                      color="primary"
                    />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TablePagination
                  rowsPerPageOptions={[5, 10, 25]}
                  count={accounts.length}
                  rowsPerPage={rowsPerPage}
                  page={page}
                  onPageChange={handleChangePage}
                  onRowsPerPageChange={handleChangeRowsPerPage}
                />
              </TableRow>
            </TableFooter>
          </Table>
        </TableContainer>
      </AccountsPaper>

      {/* Modal para adicionar nova conta */}
      <Modal open={open} onClose={handleClose}>
        <ModalContainer>
          <Typography variant="h6" gutterBottom>
            Adicionar Nova Conta
          </Typography>
          <TextField
            label="Finalidade da Conta"
            variant="outlined"
            fullWidth
            margin="normal"
            value={newAccountDescription}
            onChange={(e) => setNewAccountDescription(e.target.value)}
          />
          <Button variant="contained" color="primary" onClick={handleAddAccount} fullWidth>
            OK
          </Button>
        </ModalContainer>
      </Modal>
    </PageContainer>
  );
};

export default AccountMannager;
