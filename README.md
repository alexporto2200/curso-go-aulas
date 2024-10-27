# curso-go-aulas


## Aula 0

* Instalar o go lang no linux pelo site
https://go.dev/dl/

* extrair a pasta go para /usr/local

* adicionar o path no .bashrc o `PATH=$PATH:/usr/local/go/bin`

* testar o go com o comando `go version`

* criar a pasta go no home do usuario `mkdir ~/go; mkdir ~/go/src; mkdir ~/go/bin; mkdir ~/go/pkg`

* rodar `go env`, que são variaveis de ambiente do go

## Configurar vscode
* instalar a extensão do go (go team at google)

* Control+shift+p -> go: install/update tools -> selecionar todos e instalar


* executar: `go run main.go`

## Aula 1

go run 1/*

## Aula 2

* variaveis, constantes, tipos de dados
* inferencia de tipos, inferencia de dados
* escopo de variaveis
* short declaration

## Aula 3
* declaracão de tipos personalizados

## Aula 4
* pacote fmt
* imprimir tipos das variaveis

## Aula 5
* array
* tipo de dados array
* for array

## Aula 6
* slice
* declaracão de slice
* append, len, cap
* corte de slice, em outros slice, utilizando o :

## Aula 7
* hashtable ou map, key-value
* declaracão de map
* range map
* delete map

## Aula 8
* funcoes
* funções com multiplos retornos

## Aula 9
* funções variadicas 
* aceita mais de um parametro

## Aula 10
* funções anonimas
* funções anonimas auto-executaveis

## Aula 11
* structs
* acessando tipos de dados de structs

## Aula 11-1
* composição de structs
* structs com atributos para outras structs

## Aula 11-2
* metodos de structs
* funções que pertencem a structs

## Aula 11-3
* interfaces no go

## Aula 12
* ponteiros
* ponteiros, valores, variaveis

## Aula 13
* copia de variaveis

## Aula 14
* ponteiros e structs

## Aula 15
* interfaces vazias
* codigos legados utilizam
* evite usar

## Aula 16
* conversão de tipos - type assertion

## Aula 17
* generics

## Aula 18
* Mostrar uma falha de importação de modulo

## Aula 18-1
* Importar modulo e usar o go.mod
    (go mod init modulo)
    (go mod tidy)

## Aula 18-2
* Propriedades abertas e fechadas
    A diferença está na letra minuscula(não exporta) e maiuscula(exporta)

## Aula 18-3
* Instalar pacotes com o go
     go get github.com/google/uuid
     go mod tidy (para atualizar o go.mod removendo pacotes não utilizados)
     arquivo: go.mod fica responsavel por gerenciar os pacotes
     arquivo: go.sum garante a integridade dos pacotes

## Aula 19
* For (unico loop em go)

## Aula 20
* Condicionais
    if, else, else if
    switch
    switch com variaveis

## Aula 21
* Compilador do go
    runtime + seu código fonte = unico binário
    comando para compilar:  go build main.go
    comando para executar: ./main

    comando para compilar para windows: GOOS=windows go build main.go
    comando para compilar para linux: GOOS=linux go build main.go
    comando para outras arquiteturas: GOOS=linux GOARCH=x86 go build main.go

    Mostrar todas as arquiteturas: go tool dist list
    Ver o que esta ativo:   go env GOOS GOARCH

    Para compilar o módul0
    precisa ser uma pasta com o nome do módulo e usar o comando `go mod init modulo`
    da de mudar o nome do modulo no build com `go build -o nome_do_modulo`


## Aula 22
 * Lidar com arquivos

## Aula 23 
    * Chamadas http

## Aula 24
    * deffer   
    O deffer faz atrasar o comando. Ele é executado no final da função.
    O deffer sempre vai ser chamado por ultimo, independente de onde ele foi chamado.

## Aula 25
    * Json

## Aula 26
    * criar um sistema de busca de cep

## Aula 27 
    * Criar um servidor http

## Aula 28
    * Http headers

## Aula 29
    *  Http buscacep

## Aula 30
    * Http multplexador

## Aula 31
    * Http Fileserver

## Aula 32
    * Templates

## Aula 33
    * Template Must

## Aula 34
    * HTML template
    https://pkg.go.dev/html/template

## Aula 35
    * HTML webserver

## Aula 36
    * HTML webserver, mutiplos templates

## Aula 37
    * http Client, get e post, CopyBuffer

## Aula 38
    * http Client and Request


## Aula 39
    * http Client Context Control