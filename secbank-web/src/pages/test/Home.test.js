import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import Home from '../Home';
import api from '../../api/axiosBase';

// Mock da API
jest.mock('../../api/axiosBase', () => ({
  get: jest.fn(),
}));

// Função auxiliar para renderizar com roteamento
const renderWithRouter = (ui) => {
  return render(
    <MemoryRouter>
      {ui}
    </MemoryRouter>
  );
};

describe('Home Component', () => {
  let consoleErrorSpy;

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.setItem('token', 'mock-token');

    // Ignorar mensagens de erro no console
    consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
    consoleErrorSpy.mockRestore();
  });

  test('renders loading state initially', async () => {
    renderWithRouter(
      <>
        <Home />
        <ToastContainer />
      </>
    );

    expect(screen.getByRole('progressbar')).toBeInTheDocument();
  });

  test('handles customer info API error', async () => {
    api.get.mockResolvedValueOnce({
      ok: false, // Simula falha no `/customer/info`
    });

    renderWithRouter(
      <>
        <Home />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter informações do cliente/i)).toBeInTheDocument();
    });
  });

  test('handles account info API error', async () => {
    api.get
      .mockResolvedValueOnce({
        ok: true,
        result: { data: { ID: 123 } }, // Simula sucesso no `/customer/info`
      })
      .mockResolvedValueOnce({
        ok: false, // Simula falha no `/account/:id/information`
      });

    renderWithRouter(
      <>
        <Home />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter informações da conta/i)).toBeInTheDocument();
    });
  });

  test('handles balance API error', async () => {
    api.get
      .mockResolvedValueOnce({
        ok: true,
        result: { data: { ID: 123 } }, // Simula sucesso no `/customer/info`
      })
      .mockResolvedValueOnce({
        ok: true,
        result: { data: { IDAccount: 456, AccountNumber: '123456-7' } }, // Simula sucesso no `/account/:id/information`
      })
      .mockResolvedValueOnce({
        ok: false, // Simula falha no `/balance/:id`
      });

    renderWithRouter(
      <>
        <Home />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter saldo/i)).toBeInTheDocument();
    });
  });

  test('handles unexpected error gracefully', async () => {
    api.get.mockRejectedValueOnce(new Error(JSON.stringify({ messageError: 'Erro inesperado' })));

    renderWithRouter(
      <>
        <Home />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro inesperado/i)).toBeInTheDocument();
    });
  });
});
