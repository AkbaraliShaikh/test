grammar MyRule;

// Tokens
EQ: '=';
LT: '<';
GT: '>';
LTEQ: '<=';
GTEQ: '>=';
NEQ: 'not=';

IN: 'in';
NIN: 'not-in';

HALL: 'has-all';
HNONE: 'has-none';
HANY: 'has-any';

OBRACKET: '[';
CBRACKET: ']';
OPARAM: '(';
CPARAM: ')';

AND: 'and';
OR: 'or';

VNT: 'newer-than';
VNTE: 'newer-than-or-equal-to';
VOT: 'older-than';
VOTE: 'older-than-or-equal-to';
VE: 'equals-version';

WAO: 'within-any-of';

// Rules
query:
	query 'and' query                                                 # andExpr
	| query 'or' query                                                # orExpr
	| '(' query ')'                                                   # embbedExpr
	| leftexpr op = ('=' | 'not=') INT                                # equalIntExp
	| leftexpr op = ('=' | 'not=') STRING                             # equalStringExp
	| leftexpr op = ('>' | '<' | '>=' | '<=') INT					  # compareIntExp
	| leftexpr op = ('in' | 'not-in') listInts						  # inIntsExp
	| leftexpr op = ('in' | 'not-in') listStrings					  # inStringsExp
	| leftexpr op = ('has-all' | 'has-any' | 'has-none') listInts     # hasIntsExp
	| leftexpr op = ('has-all' | 'has-any' | 'has-none') listStrings  # hasStringsExp
	| leftexpr op = 'within-any-of' listOfLists                       # withinExp
	| leftexpr op = 'in-time-range' listOfTimeLists                   # timeRangeExp
	| leftexpr op = (
		'newer-than'
		| 'newer-than-or-equal-to'
		| 'older-than'
		| 'older-than-or-equal-to'
		| 'equals-version'
	) STRING                                                          # versionExp;

leftexpr: ATTRNAME;

listStrings: '[' stringElements ']';
stringElements: stringElement (',' SPACE* stringElement)*;
stringElement: STRING;

listInts: '[' intElements ']';
intElements: intElement (',' SPACE* intElement)*;
intElement: INT;

listOfLists: '[' list (',' SPACE* list)* ']';
list: '[' floatElement (',' SPACE* floatElement)* ']';
floatElement: FLOAT | INT;

listOfTimeLists: '[' timeList (',' timeList)* ']';
timeList: '[' timeStringElement (SPACE* timeStringElement)* ']';
timeStringElement: STRING;

// INT no leading zeros.
INT: '0' | [1-9] [0-9]*;

FLOAT: ('+' | '-')? DIGIT+ '.' DIGIT*;

STRING: '"' (ESC | ~ ["\\])* '"';

TRUE: 'true';

FALSE: 'false';

SPACE: [ \t]+ -> skip;

WS: [\t\r\n]+ -> skip;

fragment ESC: '\\' (["\\/bfnrt] | UNICODE);

fragment UNICODE: 'u' HEX HEX HEX HEX;

fragment HEX: [0-9a-fA-F];

ATTRNAME: ALPHA ATTR_NAME_CHAR*;

fragment ATTR_NAME_CHAR: '-' | '_' | DIGIT | ALPHA;

fragment DIGIT: ('0' ..'9');

fragment ALPHA: ( 'A' ..'Z' | 'a' ..'z');
