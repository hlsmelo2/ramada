# Ramada Atacadista - Full Stack Test

## Ferramentas utilizadas

* Go Lang
* Laravel
* MySQL
* Docker

## Contexto

Back: Criar em Go Lang uma API com um CRUD de produtos contendo filtragem de listas e autenticação
Front: deve ser feito em Laravel

## Execução

Após clonar o repositório deve-se e abrir o **diretório raiz do projeto (default: "ramada")**

Neste serão encontrados outros 2 diretórios: **/back** e **/front**, neste diretórios se encontram todos os arquivos referentes a cada projeto, e estes podem ser explorados da forma forma desejada

___

Estando no **diretório raiz (/)** deve ser executado o comando a seguir:

```sh
docker-compose up -d
```

Este comando instala todas as dependências das camadas de **frontend** e **backend**, e constrói todo o ambiente necessário para a execução de todas camadas e módulos do projeto. Finalizada esta etapa, o projeto já estará sendo executado e pode ser acessado através dos endereços a seguir:

**Frontend** - http://localhost:8000
---
**Backend (API)** - http://localhost:3001
---

Apesar do projeto já poder ser acessado nesta etapa, não é aconselhado fazê-lo, pois neste estágio o projeto ainda não possui nenhum dado armazenado

Porém, se desejar criar tudo do zero, comece criando usuários em **http://localhost:8000/usuarios/create**

Mas se criar os dados manualmente não for o caso, podemos resolver isso com o comando a seguir:

```sh
docker container exec -it ramada_back ./api migrate seed
```
___

Agora sim!
--
Já podemos acessar o projeto com alguns dados inseridos. Utilize o login e senha a seguir para logar na aplicação:

Login: user1@example.com
--
Password: Password@123
--
