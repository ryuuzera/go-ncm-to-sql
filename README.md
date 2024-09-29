## NCM to SQL

Este script em Go faz o download da tabela atualizada de NCM (Nomenclatura Comum do Mercosul) diretamente do portal da SEFAZ e cria um banco de dados SQLite. O banco gerado pode ser utilizado por aplicações desenvolvidas em qualquer linguagem de programação.

![NCM Table Example](https://github.com/user-attachments/assets/28d331e9-6ef3-4b7e-afb0-13c66a9a16a4)

## Funcionalidades
- **Download automatizado**: Baixa a tabela de NCM diretamente do portal da SEFAZ.
- **Banco de dados SQLite**: Gera um banco de dados SQLite, facilitando a integração com qualquer linguagem ou aplicação.
- **Portabilidade**: Permite o uso do banco de dados SQLite gerado em projetos Go ou em qualquer outra linguagem compatível com SQLite.

## Como usar

### 1. Clone o repositório:

```
git clone https://github.com/ryuuzera/go-ncm-to-sql.git
cd go-ncm-to-sql
```

### 2. Execute o script Go:
Para executar o script diretamente, utilize o seguinte comando:

```
go run main.go
```

### 3. Ou baixe o executável:
Se preferir, você pode baixar o executável pré-compilado disponível nas [releases](https://github.com/ryuuzera/go-ncm-to-sql/releases/) e executá-lo diretamente no seu sistema.

### 4. Banco de Dados Gerado:
Após a execução, um arquivo `ncm.db` será criado no diretório de trabalho, contendo a tabela de NCM pronta para ser utilizada.
