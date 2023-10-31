print("--------------------------");
print("---FUNCIONES EMBEBIDAS----");
print("----------14 pts----------");
print("--------------------------");

print("");
func suma(_ numero1: Int, _ numero2: Int) -> Int {
    let resultado = numero1 + numero2
    return resultado
}

let resultado = suma(5, 3)
print("La suma es: ", resultado)

print("");
func saludo3() {
    print("saludos!")
}

func saludo2() {
    print("mundo")
    saludo3()
}

func saludo1() {
    print("hola")
    saludo2()
}

saludo1()

print("");
func ejemplo2(verdura v: Int, v verdura: Int ) {
    print(v)
    print(verdura)
}

let precio1: Int = 251
let precio2: Int = 85
ejemplo2(verdura: precio1, v: precio2)

print("");
func duplicar(_ x: inout Int){
    x += x
}

var numero1 = 1
duplicar(&numero1)
print("numero2:", numero1)

print("");