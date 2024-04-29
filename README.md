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

* na pasta raiz, que mode ter varios modulos, rodar `go work init` Para criar o arquivo go.work

* executar: `go 1/main.go`

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
* Propriedades abertas e fechdas
    A diferença está na letra minuscula(não exporta) e maiuscula(exporta)
