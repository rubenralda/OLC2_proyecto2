struct StructArr {
    var datos: Float
}

var arr1: [StructArr] = [StructArr(datos: Float("23.43")), StructArr(datos: Float("96.3"))]

print(arr1[1].datos)