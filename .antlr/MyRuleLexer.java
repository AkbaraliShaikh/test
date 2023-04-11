// Generated from /Users/akabaralishaikh/work/github/test/MyRule.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MyRuleLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, EQ=3, LT=4, GT=5, LTEQ=6, GTEQ=7, NEQ=8, IN=9, NIN=10, 
		HALL=11, HNONE=12, HANY=13, OBRACKET=14, CBRACKET=15, OPARAM=16, CPARAM=17, 
		AND=18, OR=19, VNT=20, VNTE=21, VOT=22, VOTE=23, VE=24, WAO=25, INT=26, 
		FLOAT=27, STRING=28, TRUE=29, FALSE=30, SPACE=31, WS=32, ATTRNAME=33;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "EQ", "LT", "GT", "LTEQ", "GTEQ", "NEQ", "IN", "NIN", 
			"HALL", "HNONE", "HANY", "OBRACKET", "CBRACKET", "OPARAM", "CPARAM", 
			"AND", "OR", "VNT", "VNTE", "VOT", "VOTE", "VE", "WAO", "INT", "FLOAT", 
			"STRING", "TRUE", "FALSE", "SPACE", "WS", "ESC", "UNICODE", "HEX", "ATTRNAME", 
			"ATTR_NAME_CHAR", "DIGIT", "ALPHA"
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


	public MyRuleLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "MyRule.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2#\u015e\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3"+
		"\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3"+
		"\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3"+
		"\33\7\33\u0109\n\33\f\33\16\33\u010c\13\33\5\33\u010e\n\33\3\34\5\34\u0111"+
		"\n\34\3\34\6\34\u0114\n\34\r\34\16\34\u0115\3\34\3\34\7\34\u011a\n\34"+
		"\f\34\16\34\u011d\13\34\3\35\3\35\3\35\7\35\u0122\n\35\f\35\16\35\u0125"+
		"\13\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3 \6 \u0135\n \r \16 \u0136\3 \3 \3!\6!\u013c\n!\r!\16!\u013d\3!\3!\3"+
		"\"\3\"\3\"\5\"\u0145\n\"\3#\3#\3#\3#\3#\3#\3$\3$\3%\3%\7%\u0151\n%\f%"+
		"\16%\u0154\13%\3&\3&\3&\5&\u0159\n&\3\'\3\'\3(\3(\2\2)\3\3\5\4\7\5\t\6"+
		"\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24"+
		"\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C\2E\2G\2"+
		"I#K\2M\2O\2\3\2\f\3\2\63;\3\2\62;\4\2--//\4\2$$^^\4\2\13\13\"\"\4\2\13"+
		"\f\17\17\n\2$$\61\61^^ddhhppttvv\5\2\62;CHch\4\2//aa\4\2C\\c|\2\u0164"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2I\3\2\2\2\3Q\3\2\2\2\5_\3\2\2"+
		"\2\7a\3\2\2\2\tc\3\2\2\2\13e\3\2\2\2\rg\3\2\2\2\17j\3\2\2\2\21m\3\2\2"+
		"\2\23r\3\2\2\2\25u\3\2\2\2\27|\3\2\2\2\31\u0084\3\2\2\2\33\u008d\3\2\2"+
		"\2\35\u0095\3\2\2\2\37\u0097\3\2\2\2!\u0099\3\2\2\2#\u009b\3\2\2\2%\u009d"+
		"\3\2\2\2\'\u00a1\3\2\2\2)\u00a4\3\2\2\2+\u00af\3\2\2\2-\u00c6\3\2\2\2"+
		"/\u00d1\3\2\2\2\61\u00e8\3\2\2\2\63\u00f7\3\2\2\2\65\u010d\3\2\2\2\67"+
		"\u0110\3\2\2\29\u011e\3\2\2\2;\u0128\3\2\2\2=\u012d\3\2\2\2?\u0134\3\2"+
		"\2\2A\u013b\3\2\2\2C\u0141\3\2\2\2E\u0146\3\2\2\2G\u014c\3\2\2\2I\u014e"+
		"\3\2\2\2K\u0158\3\2\2\2M\u015a\3\2\2\2O\u015c\3\2\2\2QR\7k\2\2RS\7p\2"+
		"\2ST\7/\2\2TU\7v\2\2UV\7k\2\2VW\7o\2\2WX\7g\2\2XY\7/\2\2YZ\7t\2\2Z[\7"+
		"c\2\2[\\\7p\2\2\\]\7i\2\2]^\7g\2\2^\4\3\2\2\2_`\7.\2\2`\6\3\2\2\2ab\7"+
		"?\2\2b\b\3\2\2\2cd\7>\2\2d\n\3\2\2\2ef\7@\2\2f\f\3\2\2\2gh\7>\2\2hi\7"+
		"?\2\2i\16\3\2\2\2jk\7@\2\2kl\7?\2\2l\20\3\2\2\2mn\7p\2\2no\7q\2\2op\7"+
		"v\2\2pq\7?\2\2q\22\3\2\2\2rs\7k\2\2st\7p\2\2t\24\3\2\2\2uv\7p\2\2vw\7"+
		"q\2\2wx\7v\2\2xy\7/\2\2yz\7k\2\2z{\7p\2\2{\26\3\2\2\2|}\7j\2\2}~\7c\2"+
		"\2~\177\7u\2\2\177\u0080\7/\2\2\u0080\u0081\7c\2\2\u0081\u0082\7n\2\2"+
		"\u0082\u0083\7n\2\2\u0083\30\3\2\2\2\u0084\u0085\7j\2\2\u0085\u0086\7"+
		"c\2\2\u0086\u0087\7u\2\2\u0087\u0088\7/\2\2\u0088\u0089\7p\2\2\u0089\u008a"+
		"\7q\2\2\u008a\u008b\7p\2\2\u008b\u008c\7g\2\2\u008c\32\3\2\2\2\u008d\u008e"+
		"\7j\2\2\u008e\u008f\7c\2\2\u008f\u0090\7u\2\2\u0090\u0091\7/\2\2\u0091"+
		"\u0092\7c\2\2\u0092\u0093\7p\2\2\u0093\u0094\7{\2\2\u0094\34\3\2\2\2\u0095"+
		"\u0096\7]\2\2\u0096\36\3\2\2\2\u0097\u0098\7_\2\2\u0098 \3\2\2\2\u0099"+
		"\u009a\7*\2\2\u009a\"\3\2\2\2\u009b\u009c\7+\2\2\u009c$\3\2\2\2\u009d"+
		"\u009e\7c\2\2\u009e\u009f\7p\2\2\u009f\u00a0\7f\2\2\u00a0&\3\2\2\2\u00a1"+
		"\u00a2\7q\2\2\u00a2\u00a3\7t\2\2\u00a3(\3\2\2\2\u00a4\u00a5\7p\2\2\u00a5"+
		"\u00a6\7g\2\2\u00a6\u00a7\7y\2\2\u00a7\u00a8\7g\2\2\u00a8\u00a9\7t\2\2"+
		"\u00a9\u00aa\7/\2\2\u00aa\u00ab\7v\2\2\u00ab\u00ac\7j\2\2\u00ac\u00ad"+
		"\7c\2\2\u00ad\u00ae\7p\2\2\u00ae*\3\2\2\2\u00af\u00b0\7p\2\2\u00b0\u00b1"+
		"\7g\2\2\u00b1\u00b2\7y\2\2\u00b2\u00b3\7g\2\2\u00b3\u00b4\7t\2\2\u00b4"+
		"\u00b5\7/\2\2\u00b5\u00b6\7v\2\2\u00b6\u00b7\7j\2\2\u00b7\u00b8\7c\2\2"+
		"\u00b8\u00b9\7p\2\2\u00b9\u00ba\7/\2\2\u00ba\u00bb\7q\2\2\u00bb\u00bc"+
		"\7t\2\2\u00bc\u00bd\7/\2\2\u00bd\u00be\7g\2\2\u00be\u00bf\7s\2\2\u00bf"+
		"\u00c0\7w\2\2\u00c0\u00c1\7c\2\2\u00c1\u00c2\7n\2\2\u00c2\u00c3\7/\2\2"+
		"\u00c3\u00c4\7v\2\2\u00c4\u00c5\7q\2\2\u00c5,\3\2\2\2\u00c6\u00c7\7q\2"+
		"\2\u00c7\u00c8\7n\2\2\u00c8\u00c9\7f\2\2\u00c9\u00ca\7g\2\2\u00ca\u00cb"+
		"\7t\2\2\u00cb\u00cc\7/\2\2\u00cc\u00cd\7v\2\2\u00cd\u00ce\7j\2\2\u00ce"+
		"\u00cf\7c\2\2\u00cf\u00d0\7p\2\2\u00d0.\3\2\2\2\u00d1\u00d2\7q\2\2\u00d2"+
		"\u00d3\7n\2\2\u00d3\u00d4\7f\2\2\u00d4\u00d5\7g\2\2\u00d5\u00d6\7t\2\2"+
		"\u00d6\u00d7\7/\2\2\u00d7\u00d8\7v\2\2\u00d8\u00d9\7j\2\2\u00d9\u00da"+
		"\7c\2\2\u00da\u00db\7p\2\2\u00db\u00dc\7/\2\2\u00dc\u00dd\7q\2\2\u00dd"+
		"\u00de\7t\2\2\u00de\u00df\7/\2\2\u00df\u00e0\7g\2\2\u00e0\u00e1\7s\2\2"+
		"\u00e1\u00e2\7w\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7n\2\2\u00e4\u00e5"+
		"\7/\2\2\u00e5\u00e6\7v\2\2\u00e6\u00e7\7q\2\2\u00e7\60\3\2\2\2\u00e8\u00e9"+
		"\7g\2\2\u00e9\u00ea\7s\2\2\u00ea\u00eb\7w\2\2\u00eb\u00ec\7c\2\2\u00ec"+
		"\u00ed\7n\2\2\u00ed\u00ee\7u\2\2\u00ee\u00ef\7/\2\2\u00ef\u00f0\7x\2\2"+
		"\u00f0\u00f1\7g\2\2\u00f1\u00f2\7t\2\2\u00f2\u00f3\7u\2\2\u00f3\u00f4"+
		"\7k\2\2\u00f4\u00f5\7q\2\2\u00f5\u00f6\7p\2\2\u00f6\62\3\2\2\2\u00f7\u00f8"+
		"\7y\2\2\u00f8\u00f9\7k\2\2\u00f9\u00fa\7v\2\2\u00fa\u00fb\7j\2\2\u00fb"+
		"\u00fc\7k\2\2\u00fc\u00fd\7p\2\2\u00fd\u00fe\7/\2\2\u00fe\u00ff\7c\2\2"+
		"\u00ff\u0100\7p\2\2\u0100\u0101\7{\2\2\u0101\u0102\7/\2\2\u0102\u0103"+
		"\7q\2\2\u0103\u0104\7h\2\2\u0104\64\3\2\2\2\u0105\u010e\7\62\2\2\u0106"+
		"\u010a\t\2\2\2\u0107\u0109\t\3\2\2\u0108\u0107\3\2\2\2\u0109\u010c\3\2"+
		"\2\2\u010a\u0108\3\2\2\2\u010a\u010b\3\2\2\2\u010b\u010e\3\2\2\2\u010c"+
		"\u010a\3\2\2\2\u010d\u0105\3\2\2\2\u010d\u0106\3\2\2\2\u010e\66\3\2\2"+
		"\2\u010f\u0111\t\4\2\2\u0110\u010f\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u0113"+
		"\3\2\2\2\u0112\u0114\5M\'\2\u0113\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115"+
		"\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116\u0117\3\2\2\2\u0117\u011b\7\60"+
		"\2\2\u0118\u011a\5M\'\2\u0119\u0118\3\2\2\2\u011a\u011d\3\2\2\2\u011b"+
		"\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c8\3\2\2\2\u011d\u011b\3\2\2\2"+
		"\u011e\u0123\7$\2\2\u011f\u0122\5C\"\2\u0120\u0122\n\5\2\2\u0121\u011f"+
		"\3\2\2\2\u0121\u0120\3\2\2\2\u0122\u0125\3\2\2\2\u0123\u0121\3\2\2\2\u0123"+
		"\u0124\3\2\2\2\u0124\u0126\3\2\2\2\u0125\u0123\3\2\2\2\u0126\u0127\7$"+
		"\2\2\u0127:\3\2\2\2\u0128\u0129\7v\2\2\u0129\u012a\7t\2\2\u012a\u012b"+
		"\7w\2\2\u012b\u012c\7g\2\2\u012c<\3\2\2\2\u012d\u012e\7h\2\2\u012e\u012f"+
		"\7c\2\2\u012f\u0130\7n\2\2\u0130\u0131\7u\2\2\u0131\u0132\7g\2\2\u0132"+
		">\3\2\2\2\u0133\u0135\t\6\2\2\u0134\u0133\3\2\2\2\u0135\u0136\3\2\2\2"+
		"\u0136\u0134\3\2\2\2\u0136\u0137\3\2\2\2\u0137\u0138\3\2\2\2\u0138\u0139"+
		"\b \2\2\u0139@\3\2\2\2\u013a\u013c\t\7\2\2\u013b\u013a\3\2\2\2\u013c\u013d"+
		"\3\2\2\2\u013d\u013b\3\2\2\2\u013d\u013e\3\2\2\2\u013e\u013f\3\2\2\2\u013f"+
		"\u0140\b!\2\2\u0140B\3\2\2\2\u0141\u0144\7^\2\2\u0142\u0145\t\b\2\2\u0143"+
		"\u0145\5E#\2\u0144\u0142\3\2\2\2\u0144\u0143\3\2\2\2\u0145D\3\2\2\2\u0146"+
		"\u0147\7w\2\2\u0147\u0148\5G$\2\u0148\u0149\5G$\2\u0149\u014a\5G$\2\u014a"+
		"\u014b\5G$\2\u014bF\3\2\2\2\u014c\u014d\t\t\2\2\u014dH\3\2\2\2\u014e\u0152"+
		"\5O(\2\u014f\u0151\5K&\2\u0150\u014f\3\2\2\2\u0151\u0154\3\2\2\2\u0152"+
		"\u0150\3\2\2\2\u0152\u0153\3\2\2\2\u0153J\3\2\2\2\u0154\u0152\3\2\2\2"+
		"\u0155\u0159\t\n\2\2\u0156\u0159\5M\'\2\u0157\u0159\5O(\2\u0158\u0155"+
		"\3\2\2\2\u0158\u0156\3\2\2\2\u0158\u0157\3\2\2\2\u0159L\3\2\2\2\u015a"+
		"\u015b\4\62;\2\u015bN\3\2\2\2\u015c\u015d\t\13\2\2\u015dP\3\2\2\2\17\2"+
		"\u010a\u010d\u0110\u0115\u011b\u0121\u0123\u0136\u013d\u0144\u0152\u0158"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}