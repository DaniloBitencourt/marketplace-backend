# Sistema de Venda de Ingressos

Este projeto implementa uma API backend de um sistema de venda de ingressos para eventos, utilizando o Go com o ORM GORM para manipulação do banco de dados. Ele oferece funcionalidades de CRUD para gerenciar usuários, eventos, ingressos, compras e avaliações, com persistência em banco de dados SQLite.

## Descrição

O sistema foi projetado para possibilitar a gestão de eventos e ingressos de forma simples e eficiente. Ele permite que os usuários se registrem, comprem ingressos para eventos, e posteriormente, façam avaliações sobre os eventos que participaram. O sistema também inclui a possibilidade de criar e gerenciar eventos, além de controlar os ingressos disponíveis e vendidos.

### Funcionalidades

- **Usuários**:
  - Cadastro, login e gestão de perfis de usuário.
  
- **Eventos**:
  - Criação e gestão de eventos, com informações sobre data, local e disponibilidade de ingressos.
  
- **Ingressos**:
  - Visualização, compra e controle de ingressos disponíveis para eventos.

- **Compras**:
  - Registro de compras de ingressos realizadas pelos usuários.
  
- **Avaliações**:
  - Permite que os usuários avaliem os eventos que participaram, com notas e comentários.

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Banco de Dados**: SQLite
- **ORM**: GORM
- **Outras bibliotecas**:
  - GORM para manipulação do banco de dados.
  - Gin (caso precise para API REST, se você implementar no futuro)

## Como Rodar o Projeto

### Pré-requisitos

- Go (versão 1.18 ou superior)
- SQLite (para persistência no banco de dados)

