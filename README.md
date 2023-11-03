# PROYECTO 2, ORGANIZACION DE LENGUAJES Y COMPILADORES 2
# Compilador para el lenguaje swift

comando para generar el parser de la gramatica:
antlr4 -no-listener -visitor -Dlanguage=Go -o parser ./parser/T_swift.g4