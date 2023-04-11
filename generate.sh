#!/bin/sh
alias antlr4='java -Xmx500M -cp "./antlr4-4.12.1-SNAPSHOT-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o parser -visitor -listener MyRule.g4