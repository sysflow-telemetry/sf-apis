# Generated from sfql.g4 by ANTLR 4.9.2
from antlr4 import *
if __name__ is not None and "." in __name__:
    from .sfqlParser import sfqlParser
else:
    from sfqlParser import sfqlParser

# This class defines a complete listener for a parse tree produced by sfqlParser.
class sfqlListener(ParseTreeListener):

    # Enter a parse tree produced by sfqlParser#definitions.
    def enterDefinitions(self, ctx:sfqlParser.DefinitionsContext):
        pass

    # Exit a parse tree produced by sfqlParser#definitions.
    def exitDefinitions(self, ctx:sfqlParser.DefinitionsContext):
        pass


    # Enter a parse tree produced by sfqlParser#f_query.
    def enterF_query(self, ctx:sfqlParser.F_queryContext):
        pass

    # Exit a parse tree produced by sfqlParser#f_query.
    def exitF_query(self, ctx:sfqlParser.F_queryContext):
        pass


    # Enter a parse tree produced by sfqlParser#f_macro.
    def enterF_macro(self, ctx:sfqlParser.F_macroContext):
        pass

    # Exit a parse tree produced by sfqlParser#f_macro.
    def exitF_macro(self, ctx:sfqlParser.F_macroContext):
        pass


    # Enter a parse tree produced by sfqlParser#f_list.
    def enterF_list(self, ctx:sfqlParser.F_listContext):
        pass

    # Exit a parse tree produced by sfqlParser#f_list.
    def exitF_list(self, ctx:sfqlParser.F_listContext):
        pass


    # Enter a parse tree produced by sfqlParser#expression.
    def enterExpression(self, ctx:sfqlParser.ExpressionContext):
        pass

    # Exit a parse tree produced by sfqlParser#expression.
    def exitExpression(self, ctx:sfqlParser.ExpressionContext):
        pass


    # Enter a parse tree produced by sfqlParser#or_expression.
    def enterOr_expression(self, ctx:sfqlParser.Or_expressionContext):
        pass

    # Exit a parse tree produced by sfqlParser#or_expression.
    def exitOr_expression(self, ctx:sfqlParser.Or_expressionContext):
        pass


    # Enter a parse tree produced by sfqlParser#and_expression.
    def enterAnd_expression(self, ctx:sfqlParser.And_expressionContext):
        pass

    # Exit a parse tree produced by sfqlParser#and_expression.
    def exitAnd_expression(self, ctx:sfqlParser.And_expressionContext):
        pass


    # Enter a parse tree produced by sfqlParser#term.
    def enterTerm(self, ctx:sfqlParser.TermContext):
        pass

    # Exit a parse tree produced by sfqlParser#term.
    def exitTerm(self, ctx:sfqlParser.TermContext):
        pass


    # Enter a parse tree produced by sfqlParser#items.
    def enterItems(self, ctx:sfqlParser.ItemsContext):
        pass

    # Exit a parse tree produced by sfqlParser#items.
    def exitItems(self, ctx:sfqlParser.ItemsContext):
        pass


    # Enter a parse tree produced by sfqlParser#var.
    def enterVar(self, ctx:sfqlParser.VarContext):
        pass

    # Exit a parse tree produced by sfqlParser#var.
    def exitVar(self, ctx:sfqlParser.VarContext):
        pass


    # Enter a parse tree produced by sfqlParser#atom.
    def enterAtom(self, ctx:sfqlParser.AtomContext):
        pass

    # Exit a parse tree produced by sfqlParser#atom.
    def exitAtom(self, ctx:sfqlParser.AtomContext):
        pass


    # Enter a parse tree produced by sfqlParser#binary_operator.
    def enterBinary_operator(self, ctx:sfqlParser.Binary_operatorContext):
        pass

    # Exit a parse tree produced by sfqlParser#binary_operator.
    def exitBinary_operator(self, ctx:sfqlParser.Binary_operatorContext):
        pass


    # Enter a parse tree produced by sfqlParser#unary_operator.
    def enterUnary_operator(self, ctx:sfqlParser.Unary_operatorContext):
        pass

    # Exit a parse tree produced by sfqlParser#unary_operator.
    def exitUnary_operator(self, ctx:sfqlParser.Unary_operatorContext):
        pass



del sfqlParser