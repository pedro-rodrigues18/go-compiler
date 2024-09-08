// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // MyGrammar

import "github.com/antlr4-go/antlr/v4"

// BaseMyGrammarListener is a complete listener for a parse tree produced by MyGrammarParser.
type BaseMyGrammarListener struct{}

var _ MyGrammarListener = &BaseMyGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseMyGrammarListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseMyGrammarListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterListaFuncoes is called when production listaFuncoes is entered.
func (s *BaseMyGrammarListener) EnterListaFuncoes(ctx *ListaFuncoesContext) {}

// ExitListaFuncoes is called when production listaFuncoes is exited.
func (s *BaseMyGrammarListener) ExitListaFuncoes(ctx *ListaFuncoesContext) {}

// EnterDecFuncao is called when production decFuncao is entered.
func (s *BaseMyGrammarListener) EnterDecFuncao(ctx *DecFuncaoContext) {}

// ExitDecFuncao is called when production decFuncao is exited.
func (s *BaseMyGrammarListener) ExitDecFuncao(ctx *DecFuncaoContext) {}

// EnterTipoRetorno is called when production tipoRetorno is entered.
func (s *BaseMyGrammarListener) EnterTipoRetorno(ctx *TipoRetornoContext) {}

// ExitTipoRetorno is called when production tipoRetorno is exited.
func (s *BaseMyGrammarListener) ExitTipoRetorno(ctx *TipoRetornoContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseMyGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseMyGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterTipoBase is called when production tipoBase is entered.
func (s *BaseMyGrammarListener) EnterTipoBase(ctx *TipoBaseContext) {}

// ExitTipoBase is called when production tipoBase is exited.
func (s *BaseMyGrammarListener) ExitTipoBase(ctx *TipoBaseContext) {}

// EnterDimensao is called when production dimensao is entered.
func (s *BaseMyGrammarListener) EnterDimensao(ctx *DimensaoContext) {}

// ExitDimensao is called when production dimensao is exited.
func (s *BaseMyGrammarListener) ExitDimensao(ctx *DimensaoContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseMyGrammarListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseMyGrammarListener) ExitParametros(ctx *ParametrosContext) {}

// EnterListaParametros is called when production listaParametros is entered.
func (s *BaseMyGrammarListener) EnterListaParametros(ctx *ListaParametrosContext) {}

// ExitListaParametros is called when production listaParametros is exited.
func (s *BaseMyGrammarListener) ExitListaParametros(ctx *ListaParametrosContext) {}

// EnterPrincipal is called when production principal is entered.
func (s *BaseMyGrammarListener) EnterPrincipal(ctx *PrincipalContext) {}

// ExitPrincipal is called when production principal is exited.
func (s *BaseMyGrammarListener) ExitPrincipal(ctx *PrincipalContext) {}

// EnterBloco is called when production bloco is entered.
func (s *BaseMyGrammarListener) EnterBloco(ctx *BlocoContext) {}

// ExitBloco is called when production bloco is exited.
func (s *BaseMyGrammarListener) ExitBloco(ctx *BlocoContext) {}

// EnterListaVariaveis is called when production listaVariaveis is entered.
func (s *BaseMyGrammarListener) EnterListaVariaveis(ctx *ListaVariaveisContext) {}

// ExitListaVariaveis is called when production listaVariaveis is exited.
func (s *BaseMyGrammarListener) ExitListaVariaveis(ctx *ListaVariaveisContext) {}

// EnterListaId is called when production listaId is entered.
func (s *BaseMyGrammarListener) EnterListaId(ctx *ListaIdContext) {}

// ExitListaId is called when production listaId is exited.
func (s *BaseMyGrammarListener) ExitListaId(ctx *ListaIdContext) {}

// EnterComandos is called when production comandos is entered.
func (s *BaseMyGrammarListener) EnterComandos(ctx *ComandosContext) {}

// ExitComandos is called when production comandos is exited.
func (s *BaseMyGrammarListener) ExitComandos(ctx *ComandosContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseMyGrammarListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseMyGrammarListener) ExitComando(ctx *ComandoContext) {}

// EnterLeitura is called when production leitura is entered.
func (s *BaseMyGrammarListener) EnterLeitura(ctx *LeituraContext) {}

// ExitLeitura is called when production leitura is exited.
func (s *BaseMyGrammarListener) ExitLeitura(ctx *LeituraContext) {}

// EnterTermoLeitura is called when production termoLeitura is entered.
func (s *BaseMyGrammarListener) EnterTermoLeitura(ctx *TermoLeituraContext) {}

// ExitTermoLeitura is called when production termoLeitura is exited.
func (s *BaseMyGrammarListener) ExitTermoLeitura(ctx *TermoLeituraContext) {}

// EnterNovoTermoLeitura is called when production novoTermoLeitura is entered.
func (s *BaseMyGrammarListener) EnterNovoTermoLeitura(ctx *NovoTermoLeituraContext) {}

// ExitNovoTermoLeitura is called when production novoTermoLeitura is exited.
func (s *BaseMyGrammarListener) ExitNovoTermoLeitura(ctx *NovoTermoLeituraContext) {}

// EnterDimensao2 is called when production dimensao2 is entered.
func (s *BaseMyGrammarListener) EnterDimensao2(ctx *Dimensao2Context) {}

// ExitDimensao2 is called when production dimensao2 is exited.
func (s *BaseMyGrammarListener) ExitDimensao2(ctx *Dimensao2Context) {}

// EnterEscrita is called when production escrita is entered.
func (s *BaseMyGrammarListener) EnterEscrita(ctx *EscritaContext) {}

// ExitEscrita is called when production escrita is exited.
func (s *BaseMyGrammarListener) ExitEscrita(ctx *EscritaContext) {}

// EnterTermoEscrita is called when production termoEscrita is entered.
func (s *BaseMyGrammarListener) EnterTermoEscrita(ctx *TermoEscritaContext) {}

// ExitTermoEscrita is called when production termoEscrita is exited.
func (s *BaseMyGrammarListener) ExitTermoEscrita(ctx *TermoEscritaContext) {}

// EnterNovoTermoEscrita is called when production novoTermoEscrita is entered.
func (s *BaseMyGrammarListener) EnterNovoTermoEscrita(ctx *NovoTermoEscritaContext) {}

// ExitNovoTermoEscrita is called when production novoTermoEscrita is exited.
func (s *BaseMyGrammarListener) ExitNovoTermoEscrita(ctx *NovoTermoEscritaContext) {}

// EnterSelecao is called when production selecao is entered.
func (s *BaseMyGrammarListener) EnterSelecao(ctx *SelecaoContext) {}

// ExitSelecao is called when production selecao is exited.
func (s *BaseMyGrammarListener) ExitSelecao(ctx *SelecaoContext) {}

// EnterSenao is called when production senao is entered.
func (s *BaseMyGrammarListener) EnterSenao(ctx *SenaoContext) {}

// ExitSenao is called when production senao is exited.
func (s *BaseMyGrammarListener) ExitSenao(ctx *SenaoContext) {}

// EnterEnquanto is called when production enquanto is entered.
func (s *BaseMyGrammarListener) EnterEnquanto(ctx *EnquantoContext) {}

// ExitEnquanto is called when production enquanto is exited.
func (s *BaseMyGrammarListener) ExitEnquanto(ctx *EnquantoContext) {}

// EnterAtribuicao is called when production atribuicao is entered.
func (s *BaseMyGrammarListener) EnterAtribuicao(ctx *AtribuicaoContext) {}

// ExitAtribuicao is called when production atribuicao is exited.
func (s *BaseMyGrammarListener) ExitAtribuicao(ctx *AtribuicaoContext) {}

// EnterComplemento is called when production complemento is entered.
func (s *BaseMyGrammarListener) EnterComplemento(ctx *ComplementoContext) {}

// ExitComplemento is called when production complemento is exited.
func (s *BaseMyGrammarListener) ExitComplemento(ctx *ComplementoContext) {}

// EnterFuncao is called when production funcao is entered.
func (s *BaseMyGrammarListener) EnterFuncao(ctx *FuncaoContext) {}

// ExitFuncao is called when production funcao is exited.
func (s *BaseMyGrammarListener) ExitFuncao(ctx *FuncaoContext) {}

// EnterArgumentos is called when production argumentos is entered.
func (s *BaseMyGrammarListener) EnterArgumentos(ctx *ArgumentosContext) {}

// ExitArgumentos is called when production argumentos is exited.
func (s *BaseMyGrammarListener) ExitArgumentos(ctx *ArgumentosContext) {}

// EnterNovoArgumento is called when production novoArgumento is entered.
func (s *BaseMyGrammarListener) EnterNovoArgumento(ctx *NovoArgumentoContext) {}

// ExitNovoArgumento is called when production novoArgumento is exited.
func (s *BaseMyGrammarListener) ExitNovoArgumento(ctx *NovoArgumentoContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BaseMyGrammarListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BaseMyGrammarListener) ExitRetorno(ctx *RetornoContext) {}

// EnterExpressao is called when production expressao is entered.
func (s *BaseMyGrammarListener) EnterExpressao(ctx *ExpressaoContext) {}

// ExitExpressao is called when production expressao is exited.
func (s *BaseMyGrammarListener) ExitExpressao(ctx *ExpressaoContext) {}

// EnterExprOu is called when production exprOu is entered.
func (s *BaseMyGrammarListener) EnterExprOu(ctx *ExprOuContext) {}

// ExitExprOu is called when production exprOu is exited.
func (s *BaseMyGrammarListener) ExitExprOu(ctx *ExprOuContext) {}

// EnterExprOu2 is called when production exprOu2 is entered.
func (s *BaseMyGrammarListener) EnterExprOu2(ctx *ExprOu2Context) {}

// ExitExprOu2 is called when production exprOu2 is exited.
func (s *BaseMyGrammarListener) ExitExprOu2(ctx *ExprOu2Context) {}

// EnterExprE is called when production exprE is entered.
func (s *BaseMyGrammarListener) EnterExprE(ctx *ExprEContext) {}

// ExitExprE is called when production exprE is exited.
func (s *BaseMyGrammarListener) ExitExprE(ctx *ExprEContext) {}

// EnterExprE2 is called when production exprE2 is entered.
func (s *BaseMyGrammarListener) EnterExprE2(ctx *ExprE2Context) {}

// ExitExprE2 is called when production exprE2 is exited.
func (s *BaseMyGrammarListener) ExitExprE2(ctx *ExprE2Context) {}

// EnterExprRelacional is called when production exprRelacional is entered.
func (s *BaseMyGrammarListener) EnterExprRelacional(ctx *ExprRelacionalContext) {}

// ExitExprRelacional is called when production exprRelacional is exited.
func (s *BaseMyGrammarListener) ExitExprRelacional(ctx *ExprRelacionalContext) {}

// EnterExprRelacional2 is called when production exprRelacional2 is entered.
func (s *BaseMyGrammarListener) EnterExprRelacional2(ctx *ExprRelacional2Context) {}

// ExitExprRelacional2 is called when production exprRelacional2 is exited.
func (s *BaseMyGrammarListener) ExitExprRelacional2(ctx *ExprRelacional2Context) {}

// EnterExprAditiva is called when production exprAditiva is entered.
func (s *BaseMyGrammarListener) EnterExprAditiva(ctx *ExprAditivaContext) {}

// ExitExprAditiva is called when production exprAditiva is exited.
func (s *BaseMyGrammarListener) ExitExprAditiva(ctx *ExprAditivaContext) {}

// EnterExprAditiva2 is called when production exprAditiva2 is entered.
func (s *BaseMyGrammarListener) EnterExprAditiva2(ctx *ExprAditiva2Context) {}

// ExitExprAditiva2 is called when production exprAditiva2 is exited.
func (s *BaseMyGrammarListener) ExitExprAditiva2(ctx *ExprAditiva2Context) {}

// EnterOpAditivo is called when production opAditivo is entered.
func (s *BaseMyGrammarListener) EnterOpAditivo(ctx *OpAditivoContext) {}

// ExitOpAditivo is called when production opAditivo is exited.
func (s *BaseMyGrammarListener) ExitOpAditivo(ctx *OpAditivoContext) {}

// EnterExprMultiplicativa is called when production exprMultiplicativa is entered.
func (s *BaseMyGrammarListener) EnterExprMultiplicativa(ctx *ExprMultiplicativaContext) {}

// ExitExprMultiplicativa is called when production exprMultiplicativa is exited.
func (s *BaseMyGrammarListener) ExitExprMultiplicativa(ctx *ExprMultiplicativaContext) {}

// EnterExprMultiplicativa2 is called when production exprMultiplicativa2 is entered.
func (s *BaseMyGrammarListener) EnterExprMultiplicativa2(ctx *ExprMultiplicativa2Context) {}

// ExitExprMultiplicativa2 is called when production exprMultiplicativa2 is exited.
func (s *BaseMyGrammarListener) ExitExprMultiplicativa2(ctx *ExprMultiplicativa2Context) {}

// EnterOpMultiplicativo is called when production opMultiplicativo is entered.
func (s *BaseMyGrammarListener) EnterOpMultiplicativo(ctx *OpMultiplicativoContext) {}

// ExitOpMultiplicativo is called when production opMultiplicativo is exited.
func (s *BaseMyGrammarListener) ExitOpMultiplicativo(ctx *OpMultiplicativoContext) {}

// EnterFator is called when production fator is entered.
func (s *BaseMyGrammarListener) EnterFator(ctx *FatorContext) {}

// ExitFator is called when production fator is exited.
func (s *BaseMyGrammarListener) ExitFator(ctx *FatorContext) {}

// EnterTermo is called when production termo is entered.
func (s *BaseMyGrammarListener) EnterTermo(ctx *TermoContext) {}

// ExitTermo is called when production termo is exited.
func (s *BaseMyGrammarListener) ExitTermo(ctx *TermoContext) {}

// EnterSinal is called when production sinal is entered.
func (s *BaseMyGrammarListener) EnterSinal(ctx *SinalContext) {}

// ExitSinal is called when production sinal is exited.
func (s *BaseMyGrammarListener) ExitSinal(ctx *SinalContext) {}
