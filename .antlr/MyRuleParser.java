// Generated from /Users/akabaralishaikh/work/github/test/MyRule.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MyRuleParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, EQ=3, LT=4, GT=5, LTEQ=6, GTEQ=7, NEQ=8, IN=9, NIN=10, 
		HALL=11, HNONE=12, HANY=13, OBRACKET=14, CBRACKET=15, OPARAM=16, CPARAM=17, 
		AND=18, OR=19, VNT=20, VNTE=21, VOT=22, VOTE=23, VE=24, WAO=25, INT=26, 
		FLOAT=27, STRING=28, TRUE=29, FALSE=30, SPACE=31, WS=32, ATTRNAME=33;
	public static final int
		RULE_query = 0, RULE_leftexpr = 1, RULE_listStrings = 2, RULE_stringElements = 3, 
		RULE_stringElement = 4, RULE_listInts = 5, RULE_intElements = 6, RULE_intElement = 7, 
		RULE_listOfLists = 8, RULE_list = 9, RULE_floatElement = 10, RULE_listOfTimeLists = 11, 
		RULE_timeList = 12, RULE_timeStringElement = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"query", "leftexpr", "listStrings", "stringElements", "stringElement", 
			"listInts", "intElements", "intElement", "listOfLists", "list", "floatElement", 
			"listOfTimeLists", "timeList", "timeStringElement"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'in-time-range'", "','", "'='", "'<'", "'>'", "'<='", "'>='", 
			"'not='", "'in'", "'not-in'", "'has-all'", "'has-none'", "'has-any'", 
			"'['", "']'", "'('", "')'", "'and'", "'or'", "'newer-than'", "'newer-than-or-equal-to'", 
			"'older-than'", "'older-than-or-equal-to'", "'equals-version'", "'within-any-of'", 
			null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "EQ", "LT", "GT", "LTEQ", "GTEQ", "NEQ", "IN", "NIN", 
			"HALL", "HNONE", "HANY", "OBRACKET", "CBRACKET", "OPARAM", "CPARAM", 
			"AND", "OR", "VNT", "VNTE", "VOT", "VOTE", "VE", "WAO", "INT", "FLOAT", 
			"STRING", "TRUE", "FALSE", "SPACE", "WS", "ATTRNAME"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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
	public String getGrammarFileName() { return "MyRule.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public MyRuleParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class QueryContext extends ParserRuleContext {
		public QueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query; }
	 
		public QueryContext() { }
		public void copyFrom(QueryContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class HasIntsExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListIntsContext listInts() {
			return getRuleContext(ListIntsContext.class,0);
		}
		public TerminalNode HALL() { return getToken(MyRuleParser.HALL, 0); }
		public TerminalNode HANY() { return getToken(MyRuleParser.HANY, 0); }
		public TerminalNode HNONE() { return getToken(MyRuleParser.HNONE, 0); }
		public HasIntsExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class CompareIntExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public TerminalNode INT() { return getToken(MyRuleParser.INT, 0); }
		public TerminalNode GT() { return getToken(MyRuleParser.GT, 0); }
		public TerminalNode LT() { return getToken(MyRuleParser.LT, 0); }
		public TerminalNode GTEQ() { return getToken(MyRuleParser.GTEQ, 0); }
		public TerminalNode LTEQ() { return getToken(MyRuleParser.LTEQ, 0); }
		public CompareIntExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class WithinExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListOfListsContext listOfLists() {
			return getRuleContext(ListOfListsContext.class,0);
		}
		public TerminalNode WAO() { return getToken(MyRuleParser.WAO, 0); }
		public WithinExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class OrExprContext extends QueryContext {
		public List<QueryContext> query() {
			return getRuleContexts(QueryContext.class);
		}
		public QueryContext query(int i) {
			return getRuleContext(QueryContext.class,i);
		}
		public TerminalNode OR() { return getToken(MyRuleParser.OR, 0); }
		public OrExprContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class EqualStringExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public TerminalNode STRING() { return getToken(MyRuleParser.STRING, 0); }
		public TerminalNode EQ() { return getToken(MyRuleParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(MyRuleParser.NEQ, 0); }
		public EqualStringExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class VersionExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public TerminalNode STRING() { return getToken(MyRuleParser.STRING, 0); }
		public TerminalNode VNT() { return getToken(MyRuleParser.VNT, 0); }
		public TerminalNode VNTE() { return getToken(MyRuleParser.VNTE, 0); }
		public TerminalNode VOT() { return getToken(MyRuleParser.VOT, 0); }
		public TerminalNode VOTE() { return getToken(MyRuleParser.VOTE, 0); }
		public TerminalNode VE() { return getToken(MyRuleParser.VE, 0); }
		public VersionExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class TimeRangeExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListOfTimeListsContext listOfTimeLists() {
			return getRuleContext(ListOfTimeListsContext.class,0);
		}
		public TimeRangeExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class EmbbedExprContext extends QueryContext {
		public TerminalNode OPARAM() { return getToken(MyRuleParser.OPARAM, 0); }
		public QueryContext query() {
			return getRuleContext(QueryContext.class,0);
		}
		public TerminalNode CPARAM() { return getToken(MyRuleParser.CPARAM, 0); }
		public EmbbedExprContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class InIntsExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListIntsContext listInts() {
			return getRuleContext(ListIntsContext.class,0);
		}
		public TerminalNode IN() { return getToken(MyRuleParser.IN, 0); }
		public TerminalNode NIN() { return getToken(MyRuleParser.NIN, 0); }
		public InIntsExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class EqualIntExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public TerminalNode INT() { return getToken(MyRuleParser.INT, 0); }
		public TerminalNode EQ() { return getToken(MyRuleParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(MyRuleParser.NEQ, 0); }
		public EqualIntExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class HasStringsExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListStringsContext listStrings() {
			return getRuleContext(ListStringsContext.class,0);
		}
		public TerminalNode HALL() { return getToken(MyRuleParser.HALL, 0); }
		public TerminalNode HANY() { return getToken(MyRuleParser.HANY, 0); }
		public TerminalNode HNONE() { return getToken(MyRuleParser.HNONE, 0); }
		public HasStringsExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class InStringsExpContext extends QueryContext {
		public Token op;
		public LeftexprContext leftexpr() {
			return getRuleContext(LeftexprContext.class,0);
		}
		public ListStringsContext listStrings() {
			return getRuleContext(ListStringsContext.class,0);
		}
		public TerminalNode IN() { return getToken(MyRuleParser.IN, 0); }
		public TerminalNode NIN() { return getToken(MyRuleParser.NIN, 0); }
		public InStringsExpContext(QueryContext ctx) { copyFrom(ctx); }
	}
	public static class AndExprContext extends QueryContext {
		public List<QueryContext> query() {
			return getRuleContexts(QueryContext.class);
		}
		public QueryContext query(int i) {
			return getRuleContext(QueryContext.class,i);
		}
		public TerminalNode AND() { return getToken(MyRuleParser.AND, 0); }
		public AndExprContext(QueryContext ctx) { copyFrom(ctx); }
	}

	public final QueryContext query() throws RecognitionException {
		return query(0);
	}

	private QueryContext query(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		QueryContext _localctx = new QueryContext(_ctx, _parentState);
		QueryContext _prevctx = _localctx;
		int _startState = 0;
		enterRecursionRule(_localctx, 0, RULE_query, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				_localctx = new EmbbedExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(29);
				match(OPARAM);
				setState(30);
				query(0);
				setState(31);
				match(CPARAM);
				}
				break;
			case 2:
				{
				_localctx = new EqualIntExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(33);
				leftexpr();
				setState(34);
				((EqualIntExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==EQ || _la==NEQ) ) {
					((EqualIntExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(35);
				match(INT);
				}
				break;
			case 3:
				{
				_localctx = new EqualStringExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(37);
				leftexpr();
				setState(38);
				((EqualStringExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==EQ || _la==NEQ) ) {
					((EqualStringExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(39);
				match(STRING);
				}
				break;
			case 4:
				{
				_localctx = new CompareIntExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(41);
				leftexpr();
				setState(42);
				((CompareIntExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LTEQ) | (1L << GTEQ))) != 0)) ) {
					((CompareIntExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(43);
				match(INT);
				}
				break;
			case 5:
				{
				_localctx = new InIntsExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(45);
				leftexpr();
				setState(46);
				((InIntsExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==IN || _la==NIN) ) {
					((InIntsExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(47);
				listInts();
				}
				break;
			case 6:
				{
				_localctx = new InStringsExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(49);
				leftexpr();
				setState(50);
				((InStringsExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==IN || _la==NIN) ) {
					((InStringsExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(51);
				listStrings();
				}
				break;
			case 7:
				{
				_localctx = new HasIntsExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(53);
				leftexpr();
				setState(54);
				((HasIntsExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << HALL) | (1L << HNONE) | (1L << HANY))) != 0)) ) {
					((HasIntsExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(55);
				listInts();
				}
				break;
			case 8:
				{
				_localctx = new HasStringsExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(57);
				leftexpr();
				setState(58);
				((HasStringsExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << HALL) | (1L << HNONE) | (1L << HANY))) != 0)) ) {
					((HasStringsExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(59);
				listStrings();
				}
				break;
			case 9:
				{
				_localctx = new WithinExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(61);
				leftexpr();
				setState(62);
				((WithinExpContext)_localctx).op = match(WAO);
				setState(63);
				listOfLists();
				}
				break;
			case 10:
				{
				_localctx = new TimeRangeExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(65);
				leftexpr();
				setState(66);
				((TimeRangeExpContext)_localctx).op = match(T__0);
				setState(67);
				listOfTimeLists();
				}
				break;
			case 11:
				{
				_localctx = new VersionExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(69);
				leftexpr();
				setState(70);
				((VersionExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VNT) | (1L << VNTE) | (1L << VOT) | (1L << VOTE) | (1L << VE))) != 0)) ) {
					((VersionExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(71);
				match(STRING);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(83);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(81);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new AndExprContext(new QueryContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_query);
						setState(75);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(76);
						match(AND);
						setState(77);
						query(14);
						}
						break;
					case 2:
						{
						_localctx = new OrExprContext(new QueryContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_query);
						setState(78);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(79);
						match(OR);
						setState(80);
						query(13);
						}
						break;
					}
					} 
				}
				setState(85);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class LeftexprContext extends ParserRuleContext {
		public TerminalNode ATTRNAME() { return getToken(MyRuleParser.ATTRNAME, 0); }
		public LeftexprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_leftexpr; }
	}

	public final LeftexprContext leftexpr() throws RecognitionException {
		LeftexprContext _localctx = new LeftexprContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_leftexpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			match(ATTRNAME);
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

	public static class ListStringsContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public StringElementsContext stringElements() {
			return getRuleContext(StringElementsContext.class,0);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public ListStringsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStrings; }
	}

	public final ListStringsContext listStrings() throws RecognitionException {
		ListStringsContext _localctx = new ListStringsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_listStrings);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(OBRACKET);
			setState(89);
			stringElements();
			setState(90);
			match(CBRACKET);
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

	public static class StringElementsContext extends ParserRuleContext {
		public List<StringElementContext> stringElement() {
			return getRuleContexts(StringElementContext.class);
		}
		public StringElementContext stringElement(int i) {
			return getRuleContext(StringElementContext.class,i);
		}
		public List<TerminalNode> SPACE() { return getTokens(MyRuleParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(MyRuleParser.SPACE, i);
		}
		public StringElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringElements; }
	}

	public final StringElementsContext stringElements() throws RecognitionException {
		StringElementsContext _localctx = new StringElementsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_stringElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			stringElement();
			setState(103);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(93);
				match(T__1);
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SPACE) {
					{
					{
					setState(94);
					match(SPACE);
					}
					}
					setState(99);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(100);
				stringElement();
				}
				}
				setState(105);
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

	public static class StringElementContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(MyRuleParser.STRING, 0); }
		public StringElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringElement; }
	}

	public final StringElementContext stringElement() throws RecognitionException {
		StringElementContext _localctx = new StringElementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_stringElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			match(STRING);
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

	public static class ListIntsContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public IntElementsContext intElements() {
			return getRuleContext(IntElementsContext.class,0);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public ListIntsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listInts; }
	}

	public final ListIntsContext listInts() throws RecognitionException {
		ListIntsContext _localctx = new ListIntsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_listInts);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(OBRACKET);
			setState(109);
			intElements();
			setState(110);
			match(CBRACKET);
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

	public static class IntElementsContext extends ParserRuleContext {
		public List<IntElementContext> intElement() {
			return getRuleContexts(IntElementContext.class);
		}
		public IntElementContext intElement(int i) {
			return getRuleContext(IntElementContext.class,i);
		}
		public List<TerminalNode> SPACE() { return getTokens(MyRuleParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(MyRuleParser.SPACE, i);
		}
		public IntElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intElements; }
	}

	public final IntElementsContext intElements() throws RecognitionException {
		IntElementsContext _localctx = new IntElementsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_intElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			intElement();
			setState(123);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(113);
				match(T__1);
				setState(117);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SPACE) {
					{
					{
					setState(114);
					match(SPACE);
					}
					}
					setState(119);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(120);
				intElement();
				}
				}
				setState(125);
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

	public static class IntElementContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(MyRuleParser.INT, 0); }
		public IntElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intElement; }
	}

	public final IntElementContext intElement() throws RecognitionException {
		IntElementContext _localctx = new IntElementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_intElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(INT);
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

	public static class ListOfListsContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public List<ListContext> list() {
			return getRuleContexts(ListContext.class);
		}
		public ListContext list(int i) {
			return getRuleContext(ListContext.class,i);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public List<TerminalNode> SPACE() { return getTokens(MyRuleParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(MyRuleParser.SPACE, i);
		}
		public ListOfListsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listOfLists; }
	}

	public final ListOfListsContext listOfLists() throws RecognitionException {
		ListOfListsContext _localctx = new ListOfListsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_listOfLists);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(OBRACKET);
			setState(129);
			list();
			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(130);
				match(T__1);
				setState(134);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SPACE) {
					{
					{
					setState(131);
					match(SPACE);
					}
					}
					setState(136);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(137);
				list();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(143);
			match(CBRACKET);
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

	public static class ListContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public List<FloatElementContext> floatElement() {
			return getRuleContexts(FloatElementContext.class);
		}
		public FloatElementContext floatElement(int i) {
			return getRuleContext(FloatElementContext.class,i);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public List<TerminalNode> SPACE() { return getTokens(MyRuleParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(MyRuleParser.SPACE, i);
		}
		public ListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list; }
	}

	public final ListContext list() throws RecognitionException {
		ListContext _localctx = new ListContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			match(OBRACKET);
			setState(146);
			floatElement();
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(147);
				match(T__1);
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SPACE) {
					{
					{
					setState(148);
					match(SPACE);
					}
					}
					setState(153);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(154);
				floatElement();
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(160);
			match(CBRACKET);
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

	public static class FloatElementContext extends ParserRuleContext {
		public TerminalNode FLOAT() { return getToken(MyRuleParser.FLOAT, 0); }
		public TerminalNode INT() { return getToken(MyRuleParser.INT, 0); }
		public FloatElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatElement; }
	}

	public final FloatElementContext floatElement() throws RecognitionException {
		FloatElementContext _localctx = new FloatElementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_floatElement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_la = _input.LA(1);
			if ( !(_la==INT || _la==FLOAT) ) {
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

	public static class ListOfTimeListsContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public List<TimeListContext> timeList() {
			return getRuleContexts(TimeListContext.class);
		}
		public TimeListContext timeList(int i) {
			return getRuleContext(TimeListContext.class,i);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public ListOfTimeListsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listOfTimeLists; }
	}

	public final ListOfTimeListsContext listOfTimeLists() throws RecognitionException {
		ListOfTimeListsContext _localctx = new ListOfTimeListsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_listOfTimeLists);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(OBRACKET);
			setState(165);
			timeList();
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(166);
				match(T__1);
				setState(167);
				timeList();
				}
				}
				setState(172);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(173);
			match(CBRACKET);
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

	public static class TimeListContext extends ParserRuleContext {
		public TerminalNode OBRACKET() { return getToken(MyRuleParser.OBRACKET, 0); }
		public List<TimeStringElementContext> timeStringElement() {
			return getRuleContexts(TimeStringElementContext.class);
		}
		public TimeStringElementContext timeStringElement(int i) {
			return getRuleContext(TimeStringElementContext.class,i);
		}
		public TerminalNode CBRACKET() { return getToken(MyRuleParser.CBRACKET, 0); }
		public List<TerminalNode> SPACE() { return getTokens(MyRuleParser.SPACE); }
		public TerminalNode SPACE(int i) {
			return getToken(MyRuleParser.SPACE, i);
		}
		public TimeListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_timeList; }
	}

	public final TimeListContext timeList() throws RecognitionException {
		TimeListContext _localctx = new TimeListContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_timeList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(OBRACKET);
			setState(176);
			timeStringElement();
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==STRING || _la==SPACE) {
				{
				{
				setState(180);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SPACE) {
					{
					{
					setState(177);
					match(SPACE);
					}
					}
					setState(182);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(183);
				timeStringElement();
				}
				}
				setState(188);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(189);
			match(CBRACKET);
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

	public static class TimeStringElementContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(MyRuleParser.STRING, 0); }
		public TimeStringElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_timeStringElement; }
	}

	public final TimeStringElementContext timeStringElement() throws RecognitionException {
		TimeStringElementContext _localctx = new TimeStringElementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_timeStringElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
			match(STRING);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 0:
			return query_sempred((QueryContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean query_sempred(QueryContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 13);
		case 1:
			return precpred(_ctx, 12);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#\u00c4\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\5\2L\n\2\3\2\3\2\3\2\3\2\3\2\3\2\7\2T\n\2\f\2\16\2W\13\2\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\7\5b\n\5\f\5\16\5e\13\5\3\5\7\5h\n"+
		"\5\f\5\16\5k\13\5\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\7\bv\n\b\f\b\16"+
		"\by\13\b\3\b\7\b|\n\b\f\b\16\b\177\13\b\3\t\3\t\3\n\3\n\3\n\3\n\7\n\u0087"+
		"\n\n\f\n\16\n\u008a\13\n\3\n\7\n\u008d\n\n\f\n\16\n\u0090\13\n\3\n\3\n"+
		"\3\13\3\13\3\13\3\13\7\13\u0098\n\13\f\13\16\13\u009b\13\13\3\13\7\13"+
		"\u009e\n\13\f\13\16\13\u00a1\13\13\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\r\7"+
		"\r\u00ab\n\r\f\r\16\r\u00ae\13\r\3\r\3\r\3\16\3\16\3\16\7\16\u00b5\n\16"+
		"\f\16\16\16\u00b8\13\16\3\16\7\16\u00bb\n\16\f\16\16\16\u00be\13\16\3"+
		"\16\3\16\3\17\3\17\3\17\2\3\2\20\2\4\6\b\n\f\16\20\22\24\26\30\32\34\2"+
		"\b\4\2\5\5\n\n\3\2\6\t\3\2\13\f\3\2\r\17\3\2\26\32\3\2\34\35\2\u00cc\2"+
		"K\3\2\2\2\4X\3\2\2\2\6Z\3\2\2\2\b^\3\2\2\2\nl\3\2\2\2\fn\3\2\2\2\16r\3"+
		"\2\2\2\20\u0080\3\2\2\2\22\u0082\3\2\2\2\24\u0093\3\2\2\2\26\u00a4\3\2"+
		"\2\2\30\u00a6\3\2\2\2\32\u00b1\3\2\2\2\34\u00c1\3\2\2\2\36\37\b\2\1\2"+
		"\37 \7\22\2\2 !\5\2\2\2!\"\7\23\2\2\"L\3\2\2\2#$\5\4\3\2$%\t\2\2\2%&\7"+
		"\34\2\2&L\3\2\2\2\'(\5\4\3\2()\t\2\2\2)*\7\36\2\2*L\3\2\2\2+,\5\4\3\2"+
		",-\t\3\2\2-.\7\34\2\2.L\3\2\2\2/\60\5\4\3\2\60\61\t\4\2\2\61\62\5\f\7"+
		"\2\62L\3\2\2\2\63\64\5\4\3\2\64\65\t\4\2\2\65\66\5\6\4\2\66L\3\2\2\2\67"+
		"8\5\4\3\289\t\5\2\29:\5\f\7\2:L\3\2\2\2;<\5\4\3\2<=\t\5\2\2=>\5\6\4\2"+
		">L\3\2\2\2?@\5\4\3\2@A\7\33\2\2AB\5\22\n\2BL\3\2\2\2CD\5\4\3\2DE\7\3\2"+
		"\2EF\5\30\r\2FL\3\2\2\2GH\5\4\3\2HI\t\6\2\2IJ\7\36\2\2JL\3\2\2\2K\36\3"+
		"\2\2\2K#\3\2\2\2K\'\3\2\2\2K+\3\2\2\2K/\3\2\2\2K\63\3\2\2\2K\67\3\2\2"+
		"\2K;\3\2\2\2K?\3\2\2\2KC\3\2\2\2KG\3\2\2\2LU\3\2\2\2MN\f\17\2\2NO\7\24"+
		"\2\2OT\5\2\2\20PQ\f\16\2\2QR\7\25\2\2RT\5\2\2\17SM\3\2\2\2SP\3\2\2\2T"+
		"W\3\2\2\2US\3\2\2\2UV\3\2\2\2V\3\3\2\2\2WU\3\2\2\2XY\7#\2\2Y\5\3\2\2\2"+
		"Z[\7\20\2\2[\\\5\b\5\2\\]\7\21\2\2]\7\3\2\2\2^i\5\n\6\2_c\7\4\2\2`b\7"+
		"!\2\2a`\3\2\2\2be\3\2\2\2ca\3\2\2\2cd\3\2\2\2df\3\2\2\2ec\3\2\2\2fh\5"+
		"\n\6\2g_\3\2\2\2hk\3\2\2\2ig\3\2\2\2ij\3\2\2\2j\t\3\2\2\2ki\3\2\2\2lm"+
		"\7\36\2\2m\13\3\2\2\2no\7\20\2\2op\5\16\b\2pq\7\21\2\2q\r\3\2\2\2r}\5"+
		"\20\t\2sw\7\4\2\2tv\7!\2\2ut\3\2\2\2vy\3\2\2\2wu\3\2\2\2wx\3\2\2\2xz\3"+
		"\2\2\2yw\3\2\2\2z|\5\20\t\2{s\3\2\2\2|\177\3\2\2\2}{\3\2\2\2}~\3\2\2\2"+
		"~\17\3\2\2\2\177}\3\2\2\2\u0080\u0081\7\34\2\2\u0081\21\3\2\2\2\u0082"+
		"\u0083\7\20\2\2\u0083\u008e\5\24\13\2\u0084\u0088\7\4\2\2\u0085\u0087"+
		"\7!\2\2\u0086\u0085\3\2\2\2\u0087\u008a\3\2\2\2\u0088\u0086\3\2\2\2\u0088"+
		"\u0089\3\2\2\2\u0089\u008b\3\2\2\2\u008a\u0088\3\2\2\2\u008b\u008d\5\24"+
		"\13\2\u008c\u0084\3\2\2\2\u008d\u0090\3\2\2\2\u008e\u008c\3\2\2\2\u008e"+
		"\u008f\3\2\2\2\u008f\u0091\3\2\2\2\u0090\u008e\3\2\2\2\u0091\u0092\7\21"+
		"\2\2\u0092\23\3\2\2\2\u0093\u0094\7\20\2\2\u0094\u009f\5\26\f\2\u0095"+
		"\u0099\7\4\2\2\u0096\u0098\7!\2\2\u0097\u0096\3\2\2\2\u0098\u009b\3\2"+
		"\2\2\u0099\u0097\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u009c\3\2\2\2\u009b"+
		"\u0099\3\2\2\2\u009c\u009e\5\26\f\2\u009d\u0095\3\2\2\2\u009e\u00a1\3"+
		"\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a2\3\2\2\2\u00a1"+
		"\u009f\3\2\2\2\u00a2\u00a3\7\21\2\2\u00a3\25\3\2\2\2\u00a4\u00a5\t\7\2"+
		"\2\u00a5\27\3\2\2\2\u00a6\u00a7\7\20\2\2\u00a7\u00ac\5\32\16\2\u00a8\u00a9"+
		"\7\4\2\2\u00a9\u00ab\5\32\16\2\u00aa\u00a8\3\2\2\2\u00ab\u00ae\3\2\2\2"+
		"\u00ac\u00aa\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00af\3\2\2\2\u00ae\u00ac"+
		"\3\2\2\2\u00af\u00b0\7\21\2\2\u00b0\31\3\2\2\2\u00b1\u00b2\7\20\2\2\u00b2"+
		"\u00bc\5\34\17\2\u00b3\u00b5\7!\2\2\u00b4\u00b3\3\2\2\2\u00b5\u00b8\3"+
		"\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8"+
		"\u00b6\3\2\2\2\u00b9\u00bb\5\34\17\2\u00ba\u00b6\3\2\2\2\u00bb\u00be\3"+
		"\2\2\2\u00bc\u00ba\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bf\3\2\2\2\u00be"+
		"\u00bc\3\2\2\2\u00bf\u00c0\7\21\2\2\u00c0\33\3\2\2\2\u00c1\u00c2\7\36"+
		"\2\2\u00c2\35\3\2\2\2\20KSUciw}\u0088\u008e\u0099\u009f\u00ac\u00b6\u00bc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}