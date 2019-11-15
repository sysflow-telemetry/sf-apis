grammar sfql;

QUERY: 'sfql';
MACRO: 'macro';
LIST: 'list';
ITEMS: 'items';
COND: 'condition';

definitions
	: (f_macro | f_list)* (f_query)? (EOF)?
	;

f_query
	: DECL QUERY DEF expression
	;

f_macro
	: DECL MACRO DEF ID COND DEF expression
	;

f_list
	: DECL LIST DEF ID ITEMS DEF items 
	;
	
expression 
	: or_expression 
	;

or_expression 
	: and_expression (OR and_expression)*
	;

and_expression 
	: term (AND term)*
	;

term 
	: var
	| NOT term
	| atom unary_operator 
	| atom binary_operator atom 
	| atom (IN|PMATCH) LPAREN (atom|items) (LISTSEP (atom|items))* RPAREN 
	| LPAREN expression RPAREN
	;

items 
	: LBRACK (atom (LISTSEP atom)*)? RBRACK
	;
	
var
	: ID
	;		

atom 
	: ID
	| PATH
	| NUMBER
	| STRING
	| '<' /* event direction */
	| '>' /* event direction */
	;

binary_operator 
	: LT 
	| LE 
	| GT 
	| GE 
	| EQ 
	| NEQ 
	| CONTAINS 
	| ICONTAINS
	| STARTSWITH
	;

unary_operator 
	: EXISTS
	;

AND 
	: 'and'
	;

OR 
	: 'or'
	;

NOT 
	: 'not'
	;

LT 
	: '<'
	;

LE 
	: '<='
	;

GT 
	: '>'
	;

GE 
	: '>='
	;

EQ 
	: '='
	;

NEQ 
	: '!='
	;

IN 
	: 'in'
	;

CONTAINS 
	: 'contains'
	;

ICONTAINS 
	: 'icontains'
	;
	
STARTSWITH 
	: 'startswith'
	;
	
PMATCH
	: 'pmatch'
	;

EXISTS 
	: 'exists'
	;

LBRACK 
	: '['
	;

RBRACK 
	: ']'
	;

LPAREN 
	: '('
	;

RPAREN 
	: ')'
	;

LISTSEP 
	: ','
	;

DECL 
	: '-'
	;
	
DEF
	: ':' ((' ')* '>')? 
	;

SEVERITY
	: 'DEBUG'
	| 'INFO'
	| 'NOTICE'
	| 'WARNING'
	| 'ERROR'
	;

ID
	:  ('a'..'z' | 'A'..'Z' | '0'..'9' | '_') ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | ':'? '[' (NUMBER|PATH) (':' PATH)* ']' | '*' )*	
	;
	
NUMBER 
	: ('0'..'9')+ ('.' ('0'..'9')+)?
	;
	
PATH
	:  ('a'..'z' | 'A'..'Z' | '/' ) ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-' | '.' | '/' | '*' )*	
	;

STRING 
    : '"' (STRING|STRLIT) '"' 
    | '\'' (STRING|STRLIT) '\''
    | '\\"' (STRING|STRLIT) '\\"'
    | '\'\'' (STRING|STRLIT) '\'\''
    ;

fragment STRLIT 
    //: .*? 
    : ~[\r\n]*?
	;
	
fragment ESC : '\\"' | '\'\'' ;
		
WS
	: [ \t\r\n\u000C]+ -> channel(HIDDEN)
	;
	
NL
	: '\r'? '\n' -> channel(HIDDEN)
	;
	
COMMENT 
	: '#' ~[\r\n]* -> channel(HIDDEN)
	;
	
ANY : . ;
