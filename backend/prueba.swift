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

print("")
print("=======================================================================")
print("=================================WHILE=================================")
print("=======================================================================")
var index = 0
while index >= 0 {
    if index == 0 {
        index = index + 100
    } else if index > 50 {
        index = index / 2 - 25
    } else {
        index = (index / 2) - 1
    }
    print(index)
}

print("")
print("=======================================================================")
print("==================================FOR===================================")
print("=======================================================================")
for i in 0...9 {
    var output = ""
    for j in 0...(10 - i) {
        output = output + " "
    }
    for k in 0...i {
        output = output + "* "
    }
    print(output)
}

print("")
print("=======================================================================")
print("=================================GUARD=================================")
print("=======================================================================")
var i = 2
while (i <= 10) {
    guard i % 2 == 0 else {
        i = i + 1
        continue
    }
    print(i)
    i = i + 1
}

print("")
print("=======================================================================")
print("================================SWITCH=================================")
print("=======================================================================")
let numero = 2
switch numero {
    case 1:
        print("Uno")
    case 2:
        print("Dos")
    case 3:
        print("Tres")
    default:
        print("Invalid day")
}