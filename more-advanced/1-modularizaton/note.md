## Comandos úteis

```bash
go work init

go work use ./module-name

go mod init github.com/username/projectname
```

Use `go work init` quando quiser criar um espaço de trabalho Go. Espaços de trabalho são úteis quando você precisa trabalhar com vários módulos no mesmo projeto ou ambiente de desenvolvimento.

O comando cria um arquivo `go.work`.

Use `go work use` para adicionar um módulo ao espaço de trabalho atual.

Por exemplo:

```bash
go work use ./api
go work use ./shared
```

Isso atualiza o arquivo `go.work` e permite que os módulos no espaço de trabalho façam referência uns aos outros localmente, sem precisar publicá-los antes.

Use `go mod init` para inicializar um novo módulo Go. Esse comando cria um arquivo `go.mod` que define o caminho do módulo e gerencia suas dependências.

Recomenda-se usar um caminho de módulo que corresponda ao repositório onde o projeto será hospedado, por exemplo:

```bash
go mod init github.com/username/projectname
```

Isso permite que outros projetos Go importem seu módulo usando o caminho do repositório.
