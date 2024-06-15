#### Descrição

Esta é uma API de Produtos desenvolvida em Go utilizando o framework Gin. A API permite listar produtos, criar novos produtos e buscar produtos por ID.

#### Estrutura do Projeto

```
.
├── cmd
│   └── main.go
├── controller
│   └── product_controller.go
├── model
│   └── product.go
├── db
│   └── conn.go
├── repository
│   └── product_repository.go
├── usecase
│   └── product_usecase.go
├── .env
├── docker-compose.yml
├── Dockerfile
└── README.md
```

#### Configuração

1. **Clonar o repositório**:

   ```sh
   git clone https://github.com/seu-usuario/go-api-products.git
   cd go-api-products
   ```

2. **Configurar variáveis de ambiente**:

   Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

   ```env
   DB_HOST=go_db
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=1234
   DB_NAME=postgres
   ```

3. **Construir e iniciar os contêineres**:

   Certifique-se de estar no diretório do projeto e execute:

   ```sh
   docker-compose up --build
   ```

   Isso irá construir a imagem do aplicativo e iniciar os contêineres do banco de dados e da aplicação.

#### Endpoints

1. **Listar todos os produtos**

   - **Método:** GET
   - **Rota:** `/products`
   - **Resposta de Sucesso:**
     ```json
     [
       {
         "id": 1,
         "name": "Teclado Gamer",
         "price": 399.99
       },
       {
         "id": 2,
         "name": "Mouse Gamer",
         "price": 199.99
       }
     ]
     ```

2. **Criar um novo produto**

   - **Método:** POST
   - **Rota:** `/products`
   - **Payload:**
     ```json
     {
       "name": "Teclado Gamer",
       "price": 399.99
     }
     ```
   - **Resposta de Sucesso:**
     ```json
     {
       "id": 1,
       "name": "Teclado Gamer",
       "price": 399.99
     }
     ```

3. **Buscar produto por ID**

   - **Método:** GET
   - **Rota:** `/product/:productId`
   - **Parâmetro de Rota:** `productId` (ID do produto a ser buscado)
   - **Resposta de Sucesso:**
     ```json
     {
       "id": 1,
       "name": "Teclado Gamer",
       "price": 399.99
     }
     ```

#### Implementação das Camadas

- **`cmd/main.go`**: Ponto de entrada da aplicação.
  
- **`controllers/`**: Responsável por lidar com as requisições HTTP e chamar os casos de uso correspondentes.

- **`models/`**: Definição dos modelos de dados, como `Product`.

- **`db/`**: Configuração e conexão com o banco de dados PostgreSQL.

- **`repository/`**: Implementação do repositório para acesso aos dados do banco.

- **`usecase/`**: Implementação dos casos de uso que contêm a lógica de negócios da aplicação.