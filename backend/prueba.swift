print("--------------------------")
print("----ARCHIVO INTERMEDIO----")
print("----------10 pts----------")
print("--------------------------")

print("")
print("=======================================================================")
print("=============================IFs ANIDADOS==============================")
print("=======================================================================")
var a = 909
var aux = 10
if aux > 0 {
	print("PRIMER IF CORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
} else if aux <= 3 {
    print("PRIMER IF INCORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
} else if aux == a {
    print("PRIMER IF INCORRECTO")
    if true && (aux == 1) {
        print("SEGUNDO IF INCORRECTO")
    } else if aux > 10 {
        print("SEGUNDO IF INCORRECTO")
    } else {
        print("SEGUNDO IF CORRECTO")
    }
}