function antlr41 { java -jar .\antlr-4.13.0-complete.jar $args }
antlr41 -o TestRig T_swift.g4
javac -cp ".\antlr-4.13.0-complete.jar" -g -Xlint TestRig/T_swift*.java
java -cp ".\antlr-4.13.0-complete.jar;./TestRig" org.antlr.v4.gui.TestRig T_swift inicio -gui .\entrada.swift


## para mostrar tokens -tokens
## java -cp "C:\Users\Damihan\Downloads\antlr-4.13.0-complete.jar;./TestRig" org.antlr.v4.runtime.misc.TestRig Lab0_Parser s -tokens TestRig/entrada.txt
## java -cp "C:\Users\Damihan\Downloads\antlr-4.13.0-complete.jar;./TestRig" org.antlr.v4.gui.TestRig Lab0_Parser s -tree TestRig/entrada.txt
## -tokens 
## -tokens -trace
## -tree 
## -tree -trace
## https://github.com/antlr/antlr4-tools
## https://tomassetti.me/antlr-mega-tutorial/
## https://github.com/antlr/grammars-v4
## https://github.com/antlr/antlr4/blob/master/doc/getting-started.md



## ejecutar en powershell con permisos de administrador
## Set-ExecutionPolicy RemoteSigned -Scope CurrentUser  
## aceptamos