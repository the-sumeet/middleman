grammar Expression;

// Parser rules
expr:   expr AND expr                # AndExpr
    |   expr OR expr                 # OrExpr
    |   NOT expr                     # NotExpr
    |   comparison                   # ComparisonExpr
    |   '(' expr ')'                 # ParenExpr
    ;

comparison: IDENTIFIER op=EQUALS operand          # EqualsExpr
          | IDENTIFIER op=NOT_EQUALS operand      # NotEqualsExpr
          | IDENTIFIER op=CONTAINS operand        # ContainsExpr
          ;


operand : BOOLEAN           # BooleanOprand
        | STRING            # StringOprand
        ;

// Lexer rules
AND         : 'and';
OR          : 'or';
NOT         : 'not';
EQUALS      : '=';
NOT_EQUALS  : '!=';
CONTAINS    : 'contains';
BOOLEAN     : 'true' | 'false';
IDENTIFIER  : [bB][oO][dD][yY] | [uU][rR][lL] | [hH][eE][aA][dD][eE][rR];   // Matches variable names
STRING      : '"' (~["\r\n])* '"';      // Matches string literals
LPAREN      : '(';
RPAREN      : ')';

// Skip whitespace
WS          : [ \t\r\n]+ -> skip;
