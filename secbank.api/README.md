## Requisitos

- **Go** 1.20+
- Banco de dados **PostgreSQL**

---

## Como Rodar o Projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/Vini404/PortifolioProjetos.git
   cd PortifolioProjetos
   ```

2. Configure as variáveis de ambiente no arquivo `.env`:
   ```bash
    PORT=8080
    DB_HOST="HOST URL"
    DB_PORT=5432
    DB_DATABASE=postgres
    DB_USERNAME=postgres
    DB_PASSWORD=YOUR_PASSWORD
    DB_SCHEMA=public
    JWT_SECRET=<YOUR SECRET JWT>
   ```

3. Execute o projeto localmente:
   ```bash
   go run main.go
   ```
---

## Contribuição

1. Faça um **fork** do repositório.
2. Crie um branch para sua feature/bugfix:
   ```bash
   git checkout -b minha-feature
   ```
3. Faça o commit de suas alterações:
   ```bash
   git commit -m "Descrição da alteração"
   ```
4. Envie as alterações para o repositório remoto:
   ```bash
   git push origin minha-feature
   ```
5. Crie um **pull request** explicando suas alterações.

---

## Contato

Para dúvidas ou sugestões, entre em contato pelo e-mail: **vinicius404contato@gmail.com**.
