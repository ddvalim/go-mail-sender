# go-mail-sender

Go Mail Sender é um projeto em Go que permite enviar e-mails utilizando a API do Gmail. Este projeto utiliza OAuth 2.0 para autenticação e autorização de acesso à conta do Gmail do usuário

# Pré-requisitos

- Go
- Git
- Uma conta no Google Cloud Platform

# Configurando a conta no Google Cloud Platform

- Acesse o Google Cloud Console. 
- Crie um novo projeto ou selecione um projeto existente. 
- Navegue até APIs & Services > Credentials. 
- Clique em Create Credentials e selecione OAuth 2.0 Client IDs. 
- Configure a tela de consentimento OAuth, se ainda não o fez. 
- Em Application type, selecione Web application. 
- Clique em Create e faça o download do arquivo credentials.json. Salve-o no diretório /go-mail-sender/config.

# Clonando o Repositório

```shell
git clone https://github.com/seu-usuario/go-mail-sender.git
cd go-mail-sender
```

# Instalando as Dependências

```shell
go get -u golang.org/x/oauth2
go get -u google.golang.org/api/gmail/v1
```

# Executando o Projeto

```shell
go run main.go
```

# Rotas

- http://localhost:5885/auth
- http://localhost:5885/send

A rota /auth levará ao sistema de autenticação da Google com um callback para a rota /callback, que escreve as credenciais
do usuário no arquivo config/token.json.

Com este arquivo criado, é possível acessar a rota /send.