import React, { useState } from 'react';
import {
  Box, Button, TextField, Typography, Paper, Stepper, Step, StepLabel,
  MenuItem, Select, FormControl, InputLabel
} from '@mui/material';
import { useNavigate } from 'react-router-dom'; // Para redirecionar à página de extrato
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar';

const TransferPage = () => {
  const [activeStep, setActiveStep] = useState(0);
  const [transferData, setTransferData] = useState({
    amount: '',
    fromAccount: '',
    toAccount: '',
    toAccountName: '',
  });

  const navigate = useNavigate();

  // Simular as contas do usuário atual
  const userAccounts = [
    { number: '12345', digit: '1', balance: 1000.0 },
    { number: '67890', digit: '9', balance: 500.0 },
    { number: '54321', digit: '3', balance: 300.0 },
  ];

  const steps = ['Valor e Conta de Origem', 'Conta de Destino', 'Resumo da Transferência'];

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  const handleChange = (e) => {
    setTransferData({ ...transferData, [e.target.name]: e.target.value });
  };

  const handleConfirm = () => {
    alert('Transferência realizada com sucesso!');
    navigate('/extract');
  };

  const renderStepContent = (step) => {
    switch (step) {
      case 0:
        return (
          <>
            <TextField
              label="Valor da Transferência"
              variant="outlined"
              fullWidth
              margin="normal"
              name="amount"
              value={transferData.amount}
              onChange={handleChange}
              type="number"
              required
            />
          </>
        );
      case 1:
        return (
          <>
            <TextField
              label="Número e digito da Conta de Destino"
              variant="outlined"
              fullWidth
              margin="normal"
              name="toAccount"
              value={transferData.toAccount}
              onChange={handleChange}
              required
            />
          </>
        );
      case 2:
        return (
          <Box>
            <Typography variant="h6" gutterBottom>
              Resumo da Transferência
            </Typography>
            <Typography variant="body1">
              <strong>De:</strong> Conta {transferData.fromAccount}
            </Typography>
            <Typography variant="body1">
              <strong>Para:</strong> Conta {transferData.toAccount} ({transferData.toAccountName})
            </Typography>
            <Typography variant="body1">
              <strong>Valor:</strong> R$ {parseFloat(transferData.amount).toFixed(2)}
            </Typography>
          </Box>
        );
      default:
        return 'Etapa desconhecida';
    }
  };

  return (
    <Box
      sx={{
        width: '100%',
        height: '100vh',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#f5f5f5', // Cor de fundo neutra
      }}
    >
      <Navbar />
      <Sidebar />
      <Paper sx={{ p: 4, maxWidth: 600 }}>
        <Stepper activeStep={activeStep} alternativeLabel>
          {steps.map((label, index) => (
            <Step key={index}>
              <StepLabel>{label}</StepLabel>
            </Step>
          ))}
        </Stepper>

        <Box sx={{ mt: 3 }}>
          {renderStepContent(activeStep)}

          <Box sx={{ mt: 3 }}>
            {activeStep > 0 && (
              <Button onClick={handleBack} sx={{ mr: 1 }}>
                Voltar
              </Button>
            )}
            {activeStep < steps.length - 1 ? (
              <Button variant="contained" color="primary" onClick={handleNext}>
                Próximo
              </Button>
            ) : (
              <Button variant="contained" color="primary" onClick={handleConfirm}>
                Confirmar Transferência
              </Button>
            )}
          </Box>
        </Box>
      </Paper>
    </Box>
  );
};

export default TransferPage;
