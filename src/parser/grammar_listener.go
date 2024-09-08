// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterListaFuncoes is called when entering the listaFuncoes production.
	EnterListaFuncoes(c *ListaFuncoesContext)

	// EnterDecFuncao is called when entering the decFuncao production.
	EnterDecFuncao(c *DecFuncaoContext)

	// EnterTipoRetorno is called when entering the tipoRetorno production.
	EnterTipoRetorno(c *TipoRetornoContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterTipoBase is called when entering the tipoBase production.
	EnterTipoBase(c *TipoBaseContext)

	// EnterDimensao is called when entering the dimensao production.
	EnterDimensao(c *DimensaoContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterListaParametros is called when entering the listaParametros production.
	EnterListaParametros(c *ListaParametrosContext)

	// EnterPrincipal is called when entering the principal production.
	EnterPrincipal(c *PrincipalContext)

	// EnterBloco is called when entering the bloco production.
	EnterBloco(c *BlocoContext)

	// EnterListaVariaveis is called when entering the listaVariaveis production.
	EnterListaVariaveis(c *ListaVariaveisContext)

	// EnterListaId is called when entering the listaId production.
	EnterListaId(c *ListaIdContext)

	// EnterComandos is called when entering the comandos production.
	EnterComandos(c *ComandosContext)

	// EnterComando is called when entering the comando production.
	EnterComando(c *ComandoContext)

	// EnterLeitura is called when entering the leitura production.
	EnterLeitura(c *LeituraContext)

	// EnterTermoLeitura is called when entering the termoLeitura production.
	EnterTermoLeitura(c *TermoLeituraContext)

	// EnterNovoTermoLeitura is called when entering the novoTermoLeitura production.
	EnterNovoTermoLeitura(c *NovoTermoLeituraContext)

	// EnterDimensao2 is called when entering the dimensao2 production.
	EnterDimensao2(c *Dimensao2Context)

	// EnterEscrita is called when entering the escrita production.
	EnterEscrita(c *EscritaContext)

	// EnterTermoEscrita is called when entering the termoEscrita production.
	EnterTermoEscrita(c *TermoEscritaContext)

	// EnterNovoTermoEscrita is called when entering the novoTermoEscrita production.
	EnterNovoTermoEscrita(c *NovoTermoEscritaContext)

	// EnterSelecao is called when entering the selecao production.
	EnterSelecao(c *SelecaoContext)

	// EnterSenao is called when entering the senao production.
	EnterSenao(c *SenaoContext)

	// EnterEnquanto is called when entering the enquanto production.
	EnterEnquanto(c *EnquantoContext)

	// EnterAtribuicao is called when entering the atribuicao production.
	EnterAtribuicao(c *AtribuicaoContext)

	// EnterComplemento is called when entering the complemento production.
	EnterComplemento(c *ComplementoContext)

	// EnterFuncao is called when entering the funcao production.
	EnterFuncao(c *FuncaoContext)

	// EnterArgumentos is called when entering the argumentos production.
	EnterArgumentos(c *ArgumentosContext)

	// EnterNovoArgumento is called when entering the novoArgumento production.
	EnterNovoArgumento(c *NovoArgumentoContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// EnterExpressao is called when entering the expressao production.
	EnterExpressao(c *ExpressaoContext)

	// EnterExprOu is called when entering the exprOu production.
	EnterExprOu(c *ExprOuContext)

	// EnterExprOu2 is called when entering the exprOu2 production.
	EnterExprOu2(c *ExprOu2Context)

	// EnterExprE is called when entering the exprE production.
	EnterExprE(c *ExprEContext)

	// EnterExprE2 is called when entering the exprE2 production.
	EnterExprE2(c *ExprE2Context)

	// EnterExprRelacional is called when entering the exprRelacional production.
	EnterExprRelacional(c *ExprRelacionalContext)

	// EnterExprRelacional2 is called when entering the exprRelacional2 production.
	EnterExprRelacional2(c *ExprRelacional2Context)

	// EnterExprAditiva is called when entering the exprAditiva production.
	EnterExprAditiva(c *ExprAditivaContext)

	// EnterExprAditiva2 is called when entering the exprAditiva2 production.
	EnterExprAditiva2(c *ExprAditiva2Context)

	// EnterOpAditivo is called when entering the opAditivo production.
	EnterOpAditivo(c *OpAditivoContext)

	// EnterExprMultiplicativa is called when entering the exprMultiplicativa production.
	EnterExprMultiplicativa(c *ExprMultiplicativaContext)

	// EnterExprMultiplicativa2 is called when entering the exprMultiplicativa2 production.
	EnterExprMultiplicativa2(c *ExprMultiplicativa2Context)

	// EnterOpMultiplicativo is called when entering the opMultiplicativo production.
	EnterOpMultiplicativo(c *OpMultiplicativoContext)

	// EnterFator is called when entering the fator production.
	EnterFator(c *FatorContext)

	// EnterTermo is called when entering the termo production.
	EnterTermo(c *TermoContext)

	// EnterSinal is called when entering the sinal production.
	EnterSinal(c *SinalContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitListaFuncoes is called when exiting the listaFuncoes production.
	ExitListaFuncoes(c *ListaFuncoesContext)

	// ExitDecFuncao is called when exiting the decFuncao production.
	ExitDecFuncao(c *DecFuncaoContext)

	// ExitTipoRetorno is called when exiting the tipoRetorno production.
	ExitTipoRetorno(c *TipoRetornoContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitTipoBase is called when exiting the tipoBase production.
	ExitTipoBase(c *TipoBaseContext)

	// ExitDimensao is called when exiting the dimensao production.
	ExitDimensao(c *DimensaoContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitListaParametros is called when exiting the listaParametros production.
	ExitListaParametros(c *ListaParametrosContext)

	// ExitPrincipal is called when exiting the principal production.
	ExitPrincipal(c *PrincipalContext)

	// ExitBloco is called when exiting the bloco production.
	ExitBloco(c *BlocoContext)

	// ExitListaVariaveis is called when exiting the listaVariaveis production.
	ExitListaVariaveis(c *ListaVariaveisContext)

	// ExitListaId is called when exiting the listaId production.
	ExitListaId(c *ListaIdContext)

	// ExitComandos is called when exiting the comandos production.
	ExitComandos(c *ComandosContext)

	// ExitComando is called when exiting the comando production.
	ExitComando(c *ComandoContext)

	// ExitLeitura is called when exiting the leitura production.
	ExitLeitura(c *LeituraContext)

	// ExitTermoLeitura is called when exiting the termoLeitura production.
	ExitTermoLeitura(c *TermoLeituraContext)

	// ExitNovoTermoLeitura is called when exiting the novoTermoLeitura production.
	ExitNovoTermoLeitura(c *NovoTermoLeituraContext)

	// ExitDimensao2 is called when exiting the dimensao2 production.
	ExitDimensao2(c *Dimensao2Context)

	// ExitEscrita is called when exiting the escrita production.
	ExitEscrita(c *EscritaContext)

	// ExitTermoEscrita is called when exiting the termoEscrita production.
	ExitTermoEscrita(c *TermoEscritaContext)

	// ExitNovoTermoEscrita is called when exiting the novoTermoEscrita production.
	ExitNovoTermoEscrita(c *NovoTermoEscritaContext)

	// ExitSelecao is called when exiting the selecao production.
	ExitSelecao(c *SelecaoContext)

	// ExitSenao is called when exiting the senao production.
	ExitSenao(c *SenaoContext)

	// ExitEnquanto is called when exiting the enquanto production.
	ExitEnquanto(c *EnquantoContext)

	// ExitAtribuicao is called when exiting the atribuicao production.
	ExitAtribuicao(c *AtribuicaoContext)

	// ExitComplemento is called when exiting the complemento production.
	ExitComplemento(c *ComplementoContext)

	// ExitFuncao is called when exiting the funcao production.
	ExitFuncao(c *FuncaoContext)

	// ExitArgumentos is called when exiting the argumentos production.
	ExitArgumentos(c *ArgumentosContext)

	// ExitNovoArgumento is called when exiting the novoArgumento production.
	ExitNovoArgumento(c *NovoArgumentoContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)

	// ExitExpressao is called when exiting the expressao production.
	ExitExpressao(c *ExpressaoContext)

	// ExitExprOu is called when exiting the exprOu production.
	ExitExprOu(c *ExprOuContext)

	// ExitExprOu2 is called when exiting the exprOu2 production.
	ExitExprOu2(c *ExprOu2Context)

	// ExitExprE is called when exiting the exprE production.
	ExitExprE(c *ExprEContext)

	// ExitExprE2 is called when exiting the exprE2 production.
	ExitExprE2(c *ExprE2Context)

	// ExitExprRelacional is called when exiting the exprRelacional production.
	ExitExprRelacional(c *ExprRelacionalContext)

	// ExitExprRelacional2 is called when exiting the exprRelacional2 production.
	ExitExprRelacional2(c *ExprRelacional2Context)

	// ExitExprAditiva is called when exiting the exprAditiva production.
	ExitExprAditiva(c *ExprAditivaContext)

	// ExitExprAditiva2 is called when exiting the exprAditiva2 production.
	ExitExprAditiva2(c *ExprAditiva2Context)

	// ExitOpAditivo is called when exiting the opAditivo production.
	ExitOpAditivo(c *OpAditivoContext)

	// ExitExprMultiplicativa is called when exiting the exprMultiplicativa production.
	ExitExprMultiplicativa(c *ExprMultiplicativaContext)

	// ExitExprMultiplicativa2 is called when exiting the exprMultiplicativa2 production.
	ExitExprMultiplicativa2(c *ExprMultiplicativa2Context)

	// ExitOpMultiplicativo is called when exiting the opMultiplicativo production.
	ExitOpMultiplicativo(c *OpMultiplicativoContext)

	// ExitFator is called when exiting the fator production.
	ExitFator(c *FatorContext)

	// ExitTermo is called when exiting the termo production.
	ExitTermo(c *TermoContext)

	// ExitSinal is called when exiting the sinal production.
	ExitSinal(c *SinalContext)
}
