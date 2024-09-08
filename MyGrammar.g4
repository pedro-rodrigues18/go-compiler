grammar MyGrammar;

// Ponto de entrada principal
programa : listaFuncoes principal EOF ;

listaFuncoes : decFuncao listaFuncoes | ;
decFuncao : tipoRetorno ID '(' parametros ')' bloco ;

tipoRetorno : tipo | VOID ;

tipo : tipoBase dimensao ;
tipoBase : INT | FLOAT | CHAR | BOOLEAN ;
dimensao : '[' NUM_INT ']' dimensao | ; // Permitir ausência de dimensão

parametros : tipo ID listaParametros | ; // Permitir ausência de parâmetros
listaParametros : ',' tipo ID listaParametros | ; 

principal : MAIN '(' ')' bloco ;

bloco : LBRACE listaVariaveis comandos RBRACE ;
listaVariaveis : tipo ID listaId SEMI listaVariaveis | ; // Permitir ausência de variáveis
listaId : ',' ID listaId | ;

comandos : comando comandos | ; // Permitir ausência de comandos
comando : leitura | escrita | atribuicao | funcao | selecao | enquanto | retorno ;

leitura : SCANF '(' termoLeitura novoTermoLeitura ')' SEMI ;
termoLeitura : ID dimensao2 ;
novoTermoLeitura : ',' termoLeitura novoTermoLeitura | ;

dimensao2 : '[' exprAditiva ']' dimensao2 | ; // Permitir ausência de dimensão

escrita : PRINTLN '(' termoEscrita novoTermoEscrita ')' SEMI ;
termoEscrita : ID dimensao2 | CONSTANTE | TEXTO ;
novoTermoEscrita : ',' termoEscrita novoTermoEscrita | ;

selecao : IF '(' expressao ')' bloco senao ;
senao : ELSE bloco | ;

enquanto : WHILE '(' expressao ')' bloco ;

atribuicao : ID ASSIGN complemento SEMI ;
complemento : expressao | funcao ;

funcao : 'func' ID '(' argumentos ')' ;
argumentos : expressao novoArgumento | ;
novoArgumento : ',' expressao novoArgumento | ;

retorno : RETURN expressao SEMI ;

expressao : exprOu ;
exprOu : exprE exprOu2 ;
exprOu2 : '||' exprE exprOu2 | ;
exprE : exprRelacional exprE2 ;
exprE2 : '&&' exprRelacional exprE2 | ;
exprRelacional : exprAditiva exprRelacional2 ;
exprRelacional2 : COMP exprAditiva | ;

exprAditiva : exprMultiplicativa exprAditiva2 ;
exprAditiva2 : opAditivo exprMultiplicativa exprAditiva2 | ;
opAditivo : PLUS | MINUS ;
exprMultiplicativa : fator exprMultiplicativa2 ;
exprMultiplicativa2 : opMultiplicativo fator exprMultiplicativa2 | ;
opMultiplicativo : MUL | DIV | MOD ;

fator : sinal termo | TEXTO | '!' fator | LPAREN expressao RPAREN ;
termo : ID dimensao2 | CONSTANTE ;
sinal : PLUS | MINUS | ;

// Regras para tokens primitivos
CONSTANTE : NUM_INT | NUM_DEC ;
NUM_INT : [0-9]+ ;
NUM_DEC : [0-9]+'.'[0-9]+ ;
TEXTO : '"' .*? '"' ;
COMP : '<' | '>' | '==' | '<=' | '>=' | '!=' ;

// Palavras chaves
MAIN : 'main' ;
INT : 'int' ;
FLOAT : 'float' ;
CHAR : 'char' ;
BOOLEAN : 'boolean' ;
VOID : 'void' ;
RETURN : 'return' ;
IF : 'if' ;
ELSE : 'else' ;
WHILE : 'while' ;
SCANF : 'scanf' ;
PRINTLN : 'println' ;

// Definir operadores e símbolos explícitos
LPAREN : '(' ;
RPAREN : ')' ;
LBRACE : '{' ;
RBRACE : '}' ;
SEMI : ';' ;
ASSIGN : '=' ;
PLUS : '+' ;
MINUS : '-' ;
MUL : '*' ;
DIV : '/' ;
MOD : '%' ;

// Regras de identificadores
ID : [a-zA-Z_][a-zA-Z_0-9]* ;

// Ignorar espaços em branco
WS : [ \t\r\n]+ -> skip ;
