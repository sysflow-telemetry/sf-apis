// Generated from /Users/faraujo/workspace/research/sysflow/sf-ql/src/grammar/sfql.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class sfqlParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		QUERY=1, MACRO=2, LIST=3, ITEMS=4, COND=5, AND=6, OR=7, NOT=8, LT=9, LE=10, 
		GT=11, GE=12, EQ=13, NEQ=14, IN=15, CONTAINS=16, ICONTAINS=17, STARTSWITH=18, 
		PMATCH=19, EXISTS=20, LBRACK=21, RBRACK=22, LPAREN=23, RPAREN=24, LISTSEP=25, 
		DECL=26, DEF=27, SEVERITY=28, ID=29, NUMBER=30, PATH=31, STRING=32, WS=33, 
		NL=34, COMMENT=35, ANY=36;
	public static final int
		RULE_definitions = 0, RULE_f_query = 1, RULE_f_macro = 2, RULE_f_list = 3, 
		RULE_expression = 4, RULE_or_expression = 5, RULE_and_expression = 6, 
		RULE_term = 7, RULE_items = 8, RULE_var = 9, RULE_atom = 10, RULE_binary_operator = 11, 
		RULE_unary_operator = 12;
	public static final String[] ruleNames = {
		"definitions", "f_query", "f_macro", "f_list", "expression", "or_expression", 
		"and_expression", "term", "items", "var", "atom", "binary_operator", "unary_operator"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'sfql'", "'macro'", "'list'", "'items'", "'condition'", "'and'", 
		"'or'", "'not'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'in'", 
		"'contains'", "'icontains'", "'startswith'", "'pmatch'", "'exists'", "'['", 
		"']'", "'('", "')'", "','", "'-'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "QUERY", "MACRO", "LIST", "ITEMS", "COND", "AND", "OR", "NOT", "LT", 
		"LE", "GT", "GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH", 
		"PMATCH", "EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "LISTSEP", 
		"DECL", "DEF", "SEVERITY", "ID", "NUMBER", "PATH", "STRING", "WS", "NL", 
		"COMMENT", "ANY"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "sfql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public sfqlParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class DefinitionsContext extends ParserRuleContext {
		public List<F_macroContext> f_macro() {
			return getRuleContexts(F_macroContext.class);
		}
		public F_macroContext f_macro(int i) {
			return getRuleContext(F_macroContext.class,i);
		}
		public List<F_listContext> f_list() {
			return getRuleContexts(F_listContext.class);
		}
		public F_listContext f_list(int i) {
			return getRuleContext(F_listContext.class,i);
		}
		public F_queryContext f_query() {
			return getRuleContext(F_queryContext.class,0);
		}
		public TerminalNode EOF() { return getToken(sfqlParser.EOF, 0); }
		public DefinitionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definitions; }
	}

	public final DefinitionsContext definitions() throws RecognitionException {
		DefinitionsContext _localctx = new DefinitionsContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_definitions);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					setState(28);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
					case 1:
						{
						setState(26);
						f_macro();
						}
						break;
					case 2:
						{
						setState(27);
						f_list();
						}
						break;
					}
					} 
				}
				setState(32);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			setState(34);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DECL) {
				{
				setState(33);
				f_query();
				}
			}

			setState(37);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(36);
				match(EOF);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class F_queryContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(sfqlParser.DECL, 0); }
		public TerminalNode QUERY() { return getToken(sfqlParser.QUERY, 0); }
		public TerminalNode DEF() { return getToken(sfqlParser.DEF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public F_queryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_f_query; }
	}

	public final F_queryContext f_query() throws RecognitionException {
		F_queryContext _localctx = new F_queryContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_f_query);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			match(DECL);
			setState(40);
			match(QUERY);
			setState(41);
			match(DEF);
			setState(42);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class F_macroContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(sfqlParser.DECL, 0); }
		public TerminalNode MACRO() { return getToken(sfqlParser.MACRO, 0); }
		public List<TerminalNode> DEF() { return getTokens(sfqlParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(sfqlParser.DEF, i);
		}
		public TerminalNode ID() { return getToken(sfqlParser.ID, 0); }
		public TerminalNode COND() { return getToken(sfqlParser.COND, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public F_macroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_f_macro; }
	}

	public final F_macroContext f_macro() throws RecognitionException {
		F_macroContext _localctx = new F_macroContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_f_macro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(DECL);
			setState(45);
			match(MACRO);
			setState(46);
			match(DEF);
			setState(47);
			match(ID);
			setState(48);
			match(COND);
			setState(49);
			match(DEF);
			setState(50);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class F_listContext extends ParserRuleContext {
		public TerminalNode DECL() { return getToken(sfqlParser.DECL, 0); }
		public TerminalNode LIST() { return getToken(sfqlParser.LIST, 0); }
		public List<TerminalNode> DEF() { return getTokens(sfqlParser.DEF); }
		public TerminalNode DEF(int i) {
			return getToken(sfqlParser.DEF, i);
		}
		public TerminalNode ID() { return getToken(sfqlParser.ID, 0); }
		public TerminalNode ITEMS() { return getToken(sfqlParser.ITEMS, 0); }
		public ItemsContext items() {
			return getRuleContext(ItemsContext.class,0);
		}
		public F_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_f_list; }
	}

	public final F_listContext f_list() throws RecognitionException {
		F_listContext _localctx = new F_listContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_f_list);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			match(DECL);
			setState(53);
			match(LIST);
			setState(54);
			match(DEF);
			setState(55);
			match(ID);
			setState(56);
			match(ITEMS);
			setState(57);
			match(DEF);
			setState(58);
			items();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public Or_expressionContext or_expression() {
			return getRuleContext(Or_expressionContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			or_expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Or_expressionContext extends ParserRuleContext {
		public List<And_expressionContext> and_expression() {
			return getRuleContexts(And_expressionContext.class);
		}
		public And_expressionContext and_expression(int i) {
			return getRuleContext(And_expressionContext.class,i);
		}
		public List<TerminalNode> OR() { return getTokens(sfqlParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(sfqlParser.OR, i);
		}
		public Or_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_or_expression; }
	}

	public final Or_expressionContext or_expression() throws RecognitionException {
		Or_expressionContext _localctx = new Or_expressionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_or_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			and_expression();
			setState(67);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(63);
				match(OR);
				setState(64);
				and_expression();
				}
				}
				setState(69);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class And_expressionContext extends ParserRuleContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(sfqlParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(sfqlParser.AND, i);
		}
		public And_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_and_expression; }
	}

	public final And_expressionContext and_expression() throws RecognitionException {
		And_expressionContext _localctx = new And_expressionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_and_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			term();
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(71);
				match(AND);
				setState(72);
				term();
				}
				}
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermContext extends ParserRuleContext {
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public TerminalNode NOT() { return getToken(sfqlParser.NOT, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public Unary_operatorContext unary_operator() {
			return getRuleContext(Unary_operatorContext.class,0);
		}
		public Binary_operatorContext binary_operator() {
			return getRuleContext(Binary_operatorContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(sfqlParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(sfqlParser.RPAREN, 0); }
		public TerminalNode IN() { return getToken(sfqlParser.IN, 0); }
		public TerminalNode PMATCH() { return getToken(sfqlParser.PMATCH, 0); }
		public List<ItemsContext> items() {
			return getRuleContexts(ItemsContext.class);
		}
		public ItemsContext items(int i) {
			return getRuleContext(ItemsContext.class,i);
		}
		public List<TerminalNode> LISTSEP() { return getTokens(sfqlParser.LISTSEP); }
		public TerminalNode LISTSEP(int i) {
			return getToken(sfqlParser.LISTSEP, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_term);
		int _la;
		try {
			setState(111);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				var();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				match(NOT);
				setState(80);
				term();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(81);
				atom();
				setState(82);
				unary_operator();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(84);
				atom();
				setState(85);
				binary_operator();
				setState(86);
				atom();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(88);
				atom();
				setState(89);
				_la = _input.LA(1);
				if ( !(_la==IN || _la==PMATCH) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(90);
				match(LPAREN);
				setState(93);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LT:
				case GT:
				case ID:
				case NUMBER:
				case PATH:
				case STRING:
					{
					setState(91);
					atom();
					}
					break;
				case LBRACK:
					{
					setState(92);
					items();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==LISTSEP) {
					{
					{
					setState(95);
					match(LISTSEP);
					setState(98);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case LT:
					case GT:
					case ID:
					case NUMBER:
					case PATH:
					case STRING:
						{
						setState(96);
						atom();
						}
						break;
					case LBRACK:
						{
						setState(97);
						items();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					}
					setState(104);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(105);
				match(RPAREN);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(107);
				match(LPAREN);
				setState(108);
				expression();
				setState(109);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ItemsContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(sfqlParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(sfqlParser.RBRACK, 0); }
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public List<TerminalNode> LISTSEP() { return getTokens(sfqlParser.LISTSEP); }
		public TerminalNode LISTSEP(int i) {
			return getToken(sfqlParser.LISTSEP, i);
		}
		public ItemsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_items; }
	}

	public final ItemsContext items() throws RecognitionException {
		ItemsContext _localctx = new ItemsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_items);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			match(LBRACK);
			setState(122);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << ID) | (1L << NUMBER) | (1L << PATH) | (1L << STRING))) != 0)) {
				{
				setState(114);
				atom();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==LISTSEP) {
					{
					{
					setState(115);
					match(LISTSEP);
					setState(116);
					atom();
					}
					}
					setState(121);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(124);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(sfqlParser.ID, 0); }
		public VarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var; }
	}

	public final VarContext var() throws RecognitionException {
		VarContext _localctx = new VarContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(sfqlParser.ID, 0); }
		public TerminalNode PATH() { return getToken(sfqlParser.PATH, 0); }
		public TerminalNode NUMBER() { return getToken(sfqlParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(sfqlParser.STRING, 0); }
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_atom);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << ID) | (1L << NUMBER) | (1L << PATH) | (1L << STRING))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Binary_operatorContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(sfqlParser.LT, 0); }
		public TerminalNode LE() { return getToken(sfqlParser.LE, 0); }
		public TerminalNode GT() { return getToken(sfqlParser.GT, 0); }
		public TerminalNode GE() { return getToken(sfqlParser.GE, 0); }
		public TerminalNode EQ() { return getToken(sfqlParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(sfqlParser.NEQ, 0); }
		public TerminalNode CONTAINS() { return getToken(sfqlParser.CONTAINS, 0); }
		public TerminalNode ICONTAINS() { return getToken(sfqlParser.ICONTAINS, 0); }
		public TerminalNode STARTSWITH() { return getToken(sfqlParser.STARTSWITH, 0); }
		public Binary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binary_operator; }
	}

	public final Binary_operatorContext binary_operator() throws RecognitionException {
		Binary_operatorContext _localctx = new Binary_operatorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_binary_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << LE) | (1L << GT) | (1L << GE) | (1L << EQ) | (1L << NEQ) | (1L << CONTAINS) | (1L << ICONTAINS) | (1L << STARTSWITH))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Unary_operatorContext extends ParserRuleContext {
		public TerminalNode EXISTS() { return getToken(sfqlParser.EXISTS, 0); }
		public Unary_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unary_operator; }
	}

	public final Unary_operatorContext unary_operator() throws RecognitionException {
		Unary_operatorContext _localctx = new Unary_operatorContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_unary_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(EXISTS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3&\u0089\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\7\2\37\n\2\f\2\16\2\"\13\2\3\2\5"+
		"\2%\n\2\3\2\5\2(\n\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\7\3\7\3\7\7\7D\n\7\f\7\16"+
		"\7G\13\7\3\b\3\b\3\b\7\bL\n\b\f\b\16\bO\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t`\n\t\3\t\3\t\3\t\5\te\n\t\7\tg"+
		"\n\t\f\t\16\tj\13\t\3\t\3\t\3\t\3\t\3\t\3\t\5\tr\n\t\3\n\3\n\3\n\3\n\7"+
		"\nx\n\n\f\n\16\n{\13\n\5\n}\n\n\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16"+
		"\3\16\3\16\2\2\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\5\4\2\21\21\25\25"+
		"\5\2\13\13\r\r\37\"\4\2\13\20\22\24\2\u008b\2 \3\2\2\2\4)\3\2\2\2\6.\3"+
		"\2\2\2\b\66\3\2\2\2\n>\3\2\2\2\f@\3\2\2\2\16H\3\2\2\2\20q\3\2\2\2\22s"+
		"\3\2\2\2\24\u0080\3\2\2\2\26\u0082\3\2\2\2\30\u0084\3\2\2\2\32\u0086\3"+
		"\2\2\2\34\37\5\6\4\2\35\37\5\b\5\2\36\34\3\2\2\2\36\35\3\2\2\2\37\"\3"+
		"\2\2\2 \36\3\2\2\2 !\3\2\2\2!$\3\2\2\2\" \3\2\2\2#%\5\4\3\2$#\3\2\2\2"+
		"$%\3\2\2\2%\'\3\2\2\2&(\7\2\2\3\'&\3\2\2\2\'(\3\2\2\2(\3\3\2\2\2)*\7\34"+
		"\2\2*+\7\3\2\2+,\7\35\2\2,-\5\n\6\2-\5\3\2\2\2./\7\34\2\2/\60\7\4\2\2"+
		"\60\61\7\35\2\2\61\62\7\37\2\2\62\63\7\7\2\2\63\64\7\35\2\2\64\65\5\n"+
		"\6\2\65\7\3\2\2\2\66\67\7\34\2\2\678\7\5\2\289\7\35\2\29:\7\37\2\2:;\7"+
		"\6\2\2;<\7\35\2\2<=\5\22\n\2=\t\3\2\2\2>?\5\f\7\2?\13\3\2\2\2@E\5\16\b"+
		"\2AB\7\t\2\2BD\5\16\b\2CA\3\2\2\2DG\3\2\2\2EC\3\2\2\2EF\3\2\2\2F\r\3\2"+
		"\2\2GE\3\2\2\2HM\5\20\t\2IJ\7\b\2\2JL\5\20\t\2KI\3\2\2\2LO\3\2\2\2MK\3"+
		"\2\2\2MN\3\2\2\2N\17\3\2\2\2OM\3\2\2\2Pr\5\24\13\2QR\7\n\2\2Rr\5\20\t"+
		"\2ST\5\26\f\2TU\5\32\16\2Ur\3\2\2\2VW\5\26\f\2WX\5\30\r\2XY\5\26\f\2Y"+
		"r\3\2\2\2Z[\5\26\f\2[\\\t\2\2\2\\_\7\31\2\2]`\5\26\f\2^`\5\22\n\2_]\3"+
		"\2\2\2_^\3\2\2\2`h\3\2\2\2ad\7\33\2\2be\5\26\f\2ce\5\22\n\2db\3\2\2\2"+
		"dc\3\2\2\2eg\3\2\2\2fa\3\2\2\2gj\3\2\2\2hf\3\2\2\2hi\3\2\2\2ik\3\2\2\2"+
		"jh\3\2\2\2kl\7\32\2\2lr\3\2\2\2mn\7\31\2\2no\5\n\6\2op\7\32\2\2pr\3\2"+
		"\2\2qP\3\2\2\2qQ\3\2\2\2qS\3\2\2\2qV\3\2\2\2qZ\3\2\2\2qm\3\2\2\2r\21\3"+
		"\2\2\2s|\7\27\2\2ty\5\26\f\2uv\7\33\2\2vx\5\26\f\2wu\3\2\2\2x{\3\2\2\2"+
		"yw\3\2\2\2yz\3\2\2\2z}\3\2\2\2{y\3\2\2\2|t\3\2\2\2|}\3\2\2\2}~\3\2\2\2"+
		"~\177\7\30\2\2\177\23\3\2\2\2\u0080\u0081\7\37\2\2\u0081\25\3\2\2\2\u0082"+
		"\u0083\t\3\2\2\u0083\27\3\2\2\2\u0084\u0085\t\4\2\2\u0085\31\3\2\2\2\u0086"+
		"\u0087\7\26\2\2\u0087\33\3\2\2\2\16\36 $\'EM_dhqy|";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}