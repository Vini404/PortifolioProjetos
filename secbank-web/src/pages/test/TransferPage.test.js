import React from 'react';
import { render, screen, fireEvent, waitFor, within } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import TransferPage from '../TransferPage';
import api from '../../api/axiosBase';

// Mock da API e Webcam
jest.mock('../../api/axiosBase', () => ({
  post: jest.fn(),
}));
jest.mock('react-webcam', () => jest.fn(() => <div data-testid="webcam">Webcam</div>));

// Função auxiliar para renderizar com roteamento
const renderWithRouter = (ui) => {
  return render(
    <MemoryRouter>
      {ui}
    </MemoryRouter>
  );
};

describe('TransferPage Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.setItem('token', 'mock-token');
  });

  test('renders the transfer page with initial state', () => {
    renderWithRouter(
      <>
        <TransferPage />
        <ToastContainer />
      </>
    );

    expect(screen.getByText(/Valor e Conta de Origem/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /Próximo/i })).toBeInTheDocument();
  });

  test('progresses through steps and shows Confirmar Transferência button in final step', async () => {
    renderWithRouter(
      <>
        <TransferPage />
        <ToastContainer />
      </>
    );

    // Passo 1
    fireEvent.change(screen.getByLabelText(/Valor da Transferência/i), { target: { value: '1000' } });
    fireEvent.click(screen.getByRole('button', { name: /Próximo/i }));

    // Passo 2
    fireEvent.change(screen.getByLabelText(/Número da Conta de Destino/i), { target: { value: '1234567' } });
    fireEvent.change(screen.getByLabelText(/Dígito/i), { target: { value: '1' } });
    fireEvent.click(screen.getByRole('button', { name: /Próximo/i }));

    // Passo 3
    const step3 = screen.getByTestId('summary-step');
    expect(within(step3).getByText(/Resumo da Transferência/i)).toBeInTheDocument();
    fireEvent.click(screen.getByRole('button', { name: /Próximo/i }));

    // Passo 4
    const step4 = screen.getByTestId('identity-step');
    expect(within(step4).getByText(/Verificação de Identidade/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /Confirmar Transferência/i })).toBeInTheDocument();
  });
});
