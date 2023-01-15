parser grammar DMGParser;

options { tokenVocab=DMGLexer; }

select_statement
    : select_clause from_clause
    ;

select_clause
	: SELECT_TOKEN select_expression
	;

select_expression
	: identification_variable
    ;

from_clause
	: FROM_TOKEN identification_variable_declaration
	;

identification_variable_declaration
	: abstract_schema_name AS_TOKEN? identification_variable
	;

abstract_schema_name
	: IDENTIFICATION_VARIABLE_EXPRESSION
	;

identification_variable
	: IDENTIFICATION_VARIABLE_EXPRESSION
	;