%{
package vdbutil

func setResult(l yyLexer, v Statement) {
  l.(*Tokenizer).Statement = v
}

type KeyValue struct {
  key string
  val string
}

%}

%union{
  statement   Statement
  bytes       []byte
  mapstruct   map[string]string
  str         string
  kv          KeyValue          
}

%token LEX_ERROR

// aliases for what is %union and what production tags
%type <statement> command command_list
%type <statement> create_database_stmt create_server_stmt create_data_wrapper
%type <bytes> value_tag comment_list ident version_tag type_tag
%type <mapstruct> options_tag options_list 
%type <kv> single_option

// tokens are parse file based constants or variables
// %token <str> STRING
%token <bytes> NUMBER LITERAL STRING
%token <bytes> CREATE DATABASE VERSION OPTIONS COMMENT SERVER FOREIGN DATA WRAPPER TRANSLATOR TYPE


%start main

%%

main: command_list

command_list:  command semicolon_opt
  {
    setResult(yylex, $1)
  }
  | 
  command_list semicolon_opt command
  {
    setResult(yylex, $3)
  }
  

semicolon_opt: /*empty*/ {}
  | ';' {}

command: 
  create_database_stmt
  |
  comment_list
    {
      addComment(string($1))
    }
  |
  create_server_stmt
  |
  create_data_wrapper
  
create_database_stmt: CREATE DATABASE ident version_tag options_tag
  {
    setDatabase(string($3), string($4), $5)
  }

version_tag: /* empty */ 
    {
      $$ = []byte("1")
    }
  | VERSION ident
  {
    $$ = $2
  }

create_data_wrapper: CREATE FOREIGN DATA WRAPPER ident type_tag options_tag
  {
    addDataWrapper(string($5), string($6), $7)
  }
  |
  CREATE FOREIGN TRANSLATOR ident type_tag options_tag
  {
    addDataWrapper(string($4), string($5), $6)
  }

create_server_stmt: CREATE SERVER ident type_tag FOREIGN DATA WRAPPER ident options_tag
  {
    addServer(string($3), string($8), $9)
  }
  |
  CREATE SERVER ident type_tag FOREIGN TRANSLATOR ident options_tag
  {
    addServer(string($3), string($7), $8)
  }

type_tag: /* empty */ 
    {
      $$ = []byte("")
    }
  | TYPE ident
  {
    $$ = $2
  }

options_tag: /* empty */ 
  {
    $$ = map[string]string {}
  }
  | OPTIONS '(' options_list ')'
  {
    $$ = $3
  }

options_list:
  {
    $$ = map[string]string {}
  }
  | single_option
  {
    $$ = map[string]string {
      $1.key: $1.val,
    }
  }
  | options_list ',' single_option
  {
    $1[$3.key] = $3.val
    $$ = $1
  }

single_option: STRING value_tag
  {
    $$ = KeyValue{key: string($1), val: string($2)}
  }

ident: STRING 
  {
    $$ = $1
  }
  | '"' STRING '"'
  {
    $$ = $2
  }
  | '\'' STRING '\''
  {
    $$ = $2
  }

value_tag: STRING
  {
    $$ = $1
  }
  | NUMBER
  {
    $$ = $1
  }
  | LITERAL
  {
    $$ = $1
  }

  comment_list:
  {
    $$ = nil
  }
  | comment_list COMMENT
  {
    $$ = []byte(string($1) + string($2))
  }