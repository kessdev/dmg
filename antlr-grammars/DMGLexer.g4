lexer grammar DMGLexer;

SELECT_TOKEN
	: 'select'
	;

DOT_TOKEN
	: '.'
	;

FROM_TOKEN
	: 'from'
	;

AS_TOKEN
	: 'as'
	;

IDENTIFICATION_VARIABLE_EXPRESSION
	: ('a' .. 'z' | 'A' .. 'Z' | '_') ('a' .. 'z' | 'A' .. 'Z' | '0' .. '9' | '_')*
	;