import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import Extract from '../Extract';
import api from '../../api/axiosBase';

jest.mock('../../api/axiosBase', () => ({
  get: jest.fn(),
}));

const renderWithRouter = (ui) => {
  return render(
    <MemoryRouter>
      {ui}
    </MemoryRouter>
  );
};

describe('Extract Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.setItem('token', 'mock-token');
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test('renders loading state initially', () => {
    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    expect(screen.getByText(/Carregando.../i)).toBeInTheDocument();
  });

  test('handles token not found error', async () => {
    localStorage.removeItem('token'); // Simulate missing token

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Token não encontrado!/i)).toBeInTheDocument();
    });
  });

  test('handles user info API error', async () => {
    api.get.mockResolvedValueOnce({
      ok: false, // `/customer/info` failure
    });

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter informações do usuário/i)).toBeInTheDocument();
    });
  });

  test('handles extract API error', async () => {
    api.get
      .mockResolvedValueOnce({
        ok: true,
        result: { data: { ID: 123 } }, // `/customer/info` success
      })
      .mockResolvedValueOnce({
        ok: false, // `/balance/extract/:id` failure
      });

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter extrato/i)).toBeInTheDocument();
    });
  });

  test('handles unexpected error gracefully', async () => {
    api.get.mockRejectedValueOnce(new Error(JSON.stringify({ messageError: 'Erro inesperado' })));

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro inesperado/i)).toBeInTheDocument();
    });
  });
});
