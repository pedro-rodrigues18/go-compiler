// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseGrammarListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseGrammarListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterListaFuncoes is called when production listaFuncoes is entered.
func (s *BaseGrammarListener) EnterListaFuncoes(ctx *ListaFuncoesContext) {}

// ExitListaFuncoes is called when production listaFuncoes is exited.
func (s *BaseGrammarListener) ExitListaFuncoes(ctx *ListaFuncoesContext) {}

// EnterDecFuncao is called when production decFuncao is entered.
func (s *BaseGrammarListener) EnterDecFuncao(ctx *DecFuncaoContext) {}

// ExitDecFuncao is called when production decFuncao is exited.
func (s *BaseGrammarListener) ExitDecFuncao(ctx *DecFuncaoContext) {}

// EnterTipoRetorno is called when production tipoRetorno is entered.
func (s *BaseGrammarListener) EnterTipoRetorno(ctx *TipoRetornoContext) {}

// ExitTipoRetorno is called when production tipoRetorno is exited.
func (s *BaseGrammarListener) ExitTipoRetorno(ctx *TipoRetornoContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterTipoBase is called when production tipoBase is entered.
func (s *BaseGrammarListener) EnterTipoBase(ctx *TipoBaseContext) {}

// ExitTipoBase is called when production tipoBase is exited.
func (s *BaseGrammarListener) ExitTipoBase(ctx *TipoBaseContext) {}

// EnterDimensao is called when production dimensao is entered.
func (s *BaseGrammarListener) EnterDimensao(ctx *DimensaoContext) {}

// ExitDimensao is called when production dimensao is exited.
func (s *BaseGrammarListener) ExitDimensao(ctx *DimensaoContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseGrammarListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseGrammarListener) ExitParametros(ctx *ParametrosContext) {}

// EnterListaParametros is called when production listaParametros is entered.
func (s *BaseGrammarListener) EnterListaParametros(ctx *ListaParametrosContext) {}

// ExitListaParametros is called when production listaParametros is exited.
func (s *BaseGrammarListener) ExitListaParametros(ctx *ListaParametrosContext) {}

// EnterPrincipal is called when production principal is entered.
func (s *BaseGrammarListener) EnterPrincipal(ctx *PrincipalContext) {}

// ExitPrincipal is called when production principal is exited.
func (s *BaseGrammarListener) ExitPrincipal(ctx *PrincipalContext) {}

// EnterBloco is called when production bloco is entered.
func (s *BaseGrammarListener) EnterBloco(ctx *BlocoContext) {}

// ExitBloco is called when production bloco is exited.
func (s *BaseGrammarListener) ExitBloco(ctx *BlocoContext) {}

// EnterListaVariaveis is called when production listaVariaveis is entered.
func (s *BaseGrammarListener) EnterListaVariaveis(ctx *ListaVariaveisContext) {}

// ExitListaVariaveis is called when production listaVariaveis is exited.
func (s *BaseGrammarListener) ExitListaVariaveis(ctx *ListaVariaveisContext) {}

// EnterListaId is called when production listaId is entered.
func (s *BaseGrammarListener) EnterListaId(ctx *ListaIdContext) {}

// ExitListaId is called when production listaId is exited.
func (s *BaseGrammarListener) ExitListaId(ctx *ListaIdContext) {}

// EnterComandos is called when production comandos is entered.
func (s *BaseGrammarListener) EnterComandos(ctx *ComandosContext) {}

// ExitComandos is called when production comandos is exited.
func (s *BaseGrammarListener) ExitComandos(ctx *ComandosContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseGrammarListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseGrammarListener) ExitComando(ctx *ComandoContext) {}

// EnterLeitura is called when production leitura is entered.
func (s *BaseGrammarListener) EnterLeitura(ctx *LeituraContext) {}

// ExitLeitura is called when production leitura is exited.
func (s *BaseGrammarListener) ExitLeitura(ctx *LeituraContext) {}

// EnterTermoLeitura is called when production termoLeitura is entered.
func (s *BaseGrammarListener) EnterTermoLeitura(ctx *TermoLeituraContext) {}

// ExitTermoLeitura is called when production termoLeitura is exited.
func (s *BaseGrammarListener) ExitTermoLeitura(ctx *TermoLeituraContext) {}

// EnterNovoTermoLeitura is called when production novoTermoLeitura is entered.
func (s *BaseGrammarListener) EnterNovoTermoLeitura(ctx *NovoTermoLeituraContext) {}

// ExitNovoTermoLeitura is called when production novoTermoLeitura is exited.
func (s *BaseGrammarListener) ExitNovoTermoLeitura(ctx *NovoTermoLeituraContext) {}

// EnterDimensao2 is called when production dimensao2 is entered.
func (s *BaseGrammarListener) EnterDimensao2(ctx *Dimensao2Context) {}

// ExitDimensao2 is called when production dimensao2 is exited.
func (s *BaseGrammarListener) ExitDimensao2(ctx *Dimensao2Context) {}

// EnterEscrita is called when production escrita is entered.
func (s *BaseGrammarListener) EnterEscrita(ctx *EscritaContext) {}

// ExitEscrita is called when production escrita is exited.
func (s *BaseGrammarListener) ExitEscrita(ctx *EscritaContext) {}

// EnterTermoEscrita is called when production termoEscrita is entered.
func (s *BaseGrammarListener) EnterTermoEscrita(ctx *TermoEscritaContext) {}

// ExitTermoEscrita is called when production termoEscrita is exited.
func (s *BaseGrammarListener) ExitTermoEscrita(ctx *TermoEscritaContext) {}

// EnterNovoTermoEscrita is called when production novoTermoEscrita is entered.
func (s *BaseGrammarListener) EnterNovoTermoEscrita(ctx *NovoTermoEscritaContext) {}

// ExitNovoTermoEscrita is called when production novoTermoEscrita is exited.
func (s *BaseGrammarListener) ExitNovoTermoEscrita(ctx *NovoTermoEscritaContext) {}

// EnterSelecao is called when production selecao is entered.
func (s *BaseGrammarListener) EnterSelecao(ctx *SelecaoContext) {}

// ExitSelecao is called when production selecao is exited.
func (s *BaseGrammarListener) ExitSelecao(ctx *SelecaoContext) {}

// EnterSenao is called when production senao is entered.
func (s *BaseGrammarListener) EnterSenao(ctx *SenaoContext) {}

// ExitSenao is called when production senao is exited.
func (s *BaseGrammarListener) ExitSenao(ctx *SenaoContext) {}

// EnterEnquanto is called when production enquanto is entered.
func (s *BaseGrammarListener) EnterEnquanto(ctx *EnquantoContext) {}

// ExitEnquanto is called when production enquanto is exited.
func (s *BaseGrammarListener) ExitEnquanto(ctx *EnquantoContext) {}

// EnterAtribuicao is called when production atribuicao is entered.
func (s *BaseGrammarListener) EnterAtribuicao(ctx *AtribuicaoContext) {}

// ExitAtribuicao is called when production atribuicao is exited.
func (s *BaseGrammarListener) ExitAtribuicao(ctx *AtribuicaoContext) {}

// EnterComplemento is called when production complemento is entered.
func (s *BaseGrammarListener) EnterComplemento(ctx *ComplementoContext) {}

// ExitComplemento is called when production complemento is exited.
func (s *BaseGrammarListener) ExitComplemento(ctx *ComplementoContext) {}

// EnterFuncao is called when production funcao is entered.
func (s *BaseGrammarListener) EnterFuncao(ctx *FuncaoContext) {}

// ExitFuncao is called when production funcao is exited.
func (s *BaseGrammarListener) ExitFuncao(ctx *FuncaoContext) {}

// EnterArgumentos is called when production argumentos is entered.
func (s *BaseGrammarListener) EnterArgumentos(ctx *ArgumentosContext) {}

// ExitArgumentos is called when production argumentos is exited.
func (s *BaseGrammarListener) ExitArgumentos(ctx *ArgumentosContext) {}

// EnterNovoArgumento is called when production novoArgumento is entered.
func (s *BaseGrammarListener) EnterNovoArgumento(ctx *NovoArgumentoContext) {}

// ExitNovoArgumento is called when production novoArgumento is exited.
func (s *BaseGrammarListener) ExitNovoArgumento(ctx *NovoArgumentoContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BaseGrammarListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BaseGrammarListener) ExitRetorno(ctx *RetornoContext) {}

// EnterExpressao is called when production expressao is entered.
func (s *BaseGrammarListener) EnterExpressao(ctx *ExpressaoContext) {}

// ExitExpressao is called when production expressao is exited.
func (s *BaseGrammarListener) ExitExpressao(ctx *ExpressaoContext) {}

// EnterExprOu is called when production exprOu is entered.
func (s *BaseGrammarListener) EnterExprOu(ctx *ExprOuContext) {}

// ExitExprOu is called when production exprOu is exited.
func (s *BaseGrammarListener) ExitExprOu(ctx *ExprOuContext) {}

// EnterExprOu2 is called when production exprOu2 is entered.
func (s *BaseGrammarListener) EnterExprOu2(ctx *ExprOu2Context) {}

// ExitExprOu2 is called when production exprOu2 is exited.
func (s *BaseGrammarListener) ExitExprOu2(ctx *ExprOu2Context) {}

// EnterExprE is called when production exprE is entered.
func (s *BaseGrammarListener) EnterExprE(ctx *ExprEContext) {}

// ExitExprE is called when production exprE is exited.
func (s *BaseGrammarListener) ExitExprE(ctx *ExprEContext) {}

// EnterExprE2 is called when production exprE2 is entered.
func (s *BaseGrammarListener) EnterExprE2(ctx *ExprE2Context) {}

// ExitExprE2 is called when production exprE2 is exited.
func (s *BaseGrammarListener) ExitExprE2(ctx *ExprE2Context) {}

// EnterExprRelacional is called when production exprRelacional is entered.
func (s *BaseGrammarListener) EnterExprRelacional(ctx *ExprRelacionalContext) {}

// ExitExprRelacional is called when production exprRelacional is exited.
func (s *BaseGrammarListener) ExitExprRelacional(ctx *ExprRelacionalContext) {}

// EnterExprRelacional2 is called when production exprRelacional2 is entered.
func (s *BaseGrammarListener) EnterExprRelacional2(ctx *ExprRelacional2Context) {}

// ExitExprRelacional2 is called when production exprRelacional2 is exited.
func (s *BaseGrammarListener) ExitExprRelacional2(ctx *ExprRelacional2Context) {}

// EnterExprAditiva is called when production exprAditiva is entered.
func (s *BaseGrammarListener) EnterExprAditiva(ctx *ExprAditivaContext) {}

// ExitExprAditiva is called when production exprAditiva is exited.
func (s *BaseGrammarListener) ExitExprAditiva(ctx *ExprAditivaContext) {}

// EnterExprAditiva2 is called when production exprAditiva2 is entered.
func (s *BaseGrammarListener) EnterExprAditiva2(ctx *ExprAditiva2Context) {}

// ExitExprAditiva2 is called when production exprAditiva2 is exited.
func (s *BaseGrammarListener) ExitExprAditiva2(ctx *ExprAditiva2Context) {}

// EnterOpAditivo is called when production opAditivo is entered.
func (s *BaseGrammarListener) EnterOpAditivo(ctx *OpAditivoContext) {}

// ExitOpAditivo is called when production opAditivo is exited.
func (s *BaseGrammarListener) ExitOpAditivo(ctx *OpAditivoContext) {}

// EnterExprMultiplicativa is called when production exprMultiplicativa is entered.
func (s *BaseGrammarListener) EnterExprMultiplicativa(ctx *ExprMultiplicativaContext) {}

// ExitExprMultiplicativa is called when production exprMultiplicativa is exited.
func (s *BaseGrammarListener) ExitExprMultiplicativa(ctx *ExprMultiplicativaContext) {}

// EnterExprMultiplicativa2 is called when production exprMultiplicativa2 is entered.
func (s *BaseGrammarListener) EnterExprMultiplicativa2(ctx *ExprMultiplicativa2Context) {}

// ExitExprMultiplicativa2 is called when production exprMultiplicativa2 is exited.
func (s *BaseGrammarListener) ExitExprMultiplicativa2(ctx *ExprMultiplicativa2Context) {}

// EnterOpMultiplicativo is called when production opMultiplicativo is entered.
func (s *BaseGrammarListener) EnterOpMultiplicativo(ctx *OpMultiplicativoContext) {}

// ExitOpMultiplicativo is called when production opMultiplicativo is exited.
func (s *BaseGrammarListener) ExitOpMultiplicativo(ctx *OpMultiplicativoContext) {}

// EnterFator is called when production fator is entered.
func (s *BaseGrammarListener) EnterFator(ctx *FatorContext) {}

// ExitFator is called when production fator is exited.
func (s *BaseGrammarListener) ExitFator(ctx *FatorContext) {}

// EnterTermo is called when production termo is entered.
func (s *BaseGrammarListener) EnterTermo(ctx *TermoContext) {}

// ExitTermo is called when production termo is exited.
func (s *BaseGrammarListener) ExitTermo(ctx *TermoContext) {}

// EnterSinal is called when production sinal is entered.
func (s *BaseGrammarListener) EnterSinal(ctx *SinalContext) {}

// ExitSinal is called when production sinal is exited.
func (s *BaseGrammarListener) ExitSinal(ctx *SinalContext) {}
