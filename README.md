# Gemini API Base Module (Go)

Este repositório fornece uma **implementação base dos serviços da API do Google Gemini**, desenvolvida em **Go (Golang)**, com o objetivo de servir como **módulo reutilizável** para qualquer aplicação que precise integrar recursos de IA generativa.

---

## 📌 Funcionalidades

* Conexão simplificada com a API do Gemini.
* Suporte a envio de **texto, imagens e arquivos**.
* Estrutura de **módulo independente**, fácil de integrar em outros projetos Go.
* Configuração via variáveis de ambiente.
* Pronto para expansão e customização conforme as necessidades da aplicação.

---

## 📂 Estrutura do Projeto

```
.
├── backend/                 # Implementação do backend em Go
│   ├── config/               # Configuração e carregamento de variáveis de ambiente
│   ├── controllers/          # Controladores que recebem requisições e chamam os serviços
│   ├── docs/                  # Documentação e especificações da API (ex: Swagger)
│   ├── models/                # Estruturas de dados e entidades
│   ├── routes/                # Definição das rotas da API
│   ├── services/              # Implementações de chamadas à API do Gemini
│   ├── .env.example           # Exemplo de configuração de variáveis
│   ├── main.go                # Ponto de entrada do backend
│   ├── go.mod / go.sum        # Dependências do Go
│   └── README.md
│
└── frontend/                  # Interface opcional (pode ser adaptada ou removida)
    └── README.md
```

---

## 🚀 Como usar

### 1️⃣ Clonar o repositório

```bash
git clone https://github.com/seu-usuario/dev_gemini_api_modules.git
cd dev_gemini_api_modules
```

### 2️⃣ Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
ENV=

GEMINI_API_KEY=

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

PORT=
```

### 3️⃣ Instalar dependências

```bash
go mod tidy
```
---

## 📖 Exemplos

No diretório [`util_examples`](./backend/services/util_examples/) há códigos prontos para:

* Enviar prompts de texto.
* Processar imagens.
* Trabalhar com múltiplos arquivos.
* Uso de **streaming** de resposta.

---

## 🛠 Tecnologias utilizadas

* **Go** (>= 1.22)
* **Google Gemini API**
* **godotenv** (para carregar variáveis de ambiente)