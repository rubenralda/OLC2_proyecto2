const lineNumbers = document.querySelector('.line-numbers')
const btnEjecutar = document.getElementById("btn");
const codigo = document.getElementById("codigo");//textarea
const lineNumbersConsole = document.querySelector('.line-numbers-consola')
const consola = document.getElementById("textConsola");//textarea
const btnTabla = document.getElementById("tablaSimbolos");
const btnAbrir = document.getElementById("abriArchivo");
const inputFile = document.createElement('input');
const btnErrores = document.getElementById('reporteErrores');
//const listArchivos = document.getElementById("listDeArchivos");
inputFile.type = 'file';
inputFile.accept = '.swift';
inputFile.style.display = 'none';
document.body.appendChild(inputFile);

document.addEventListener("DOMContentLoaded", (e) => {
    e.preventDefault;
    if (localStorage.getItem("texto") != null) {
        codigo.value = localStorage.getItem("texto")
        let dispararEvento = new Event("keyup")
        codigo.dispatchEvent(dispararEvento)
    }
    e.stopPropagation;
});

codigo.addEventListener("input", (e) =>{
    e.preventDefault;
    localStorage.setItem("texto", codigo.value)
    e.stopPropagation;
});

codigo.addEventListener("keyup", event => {
    const numberOfLines = event.target.value.split("\n").length;
    lineNumbers.innerHTML = Array(numberOfLines)
    .fill('<span></span>')
    .join('')
});

codigo.addEventListener("keydown", function(event) {
    if (event.key === "Tab") {
      event.preventDefault();  // Evita que el foco cambie al siguiente elemento
      const tabCharacter = "    "; // 4 espacios para simular un tabulador
      const start = this.selectionStart;
      const end = this.selectionEnd;
      this.value = this.value.substring(0, start) + tabCharacter + this.value.substring(end);
      this.selectionStart = this.selectionEnd = start + tabCharacter.length;
    }
});

consola.addEventListener("keyup", event => {
    const numberOfLines = event.target.value.split("\n").length;
    lineNumbersConsole.innerHTML = Array(numberOfLines)
    .fill('<span></span>')
    .join('')
});

inputFile.addEventListener('change', () => {
    const file = inputFile.files[0];
    const reader = new FileReader();
    reader.readAsText(file);
    reader.onload = () => {
      const fileContent = reader.result;
      codigo.value = fileContent
      localStorage.setItem("texto", fileContent)
      let dispararEvento = new Event("keyup")
      codigo.dispatchEvent(dispararEvento)
    };
});

btnAbrir.addEventListener("click", (e) =>{
    e.preventDefault();
    inputFile.click();
    e.stopPropagation();
});

btnEjecutar.addEventListener("click", eventbtnEjecutar);
async function eventbtnEjecutar() {
    const ruta = `http://localhost:3000/ejecutar`;
    let bodyJson = {
        Codigo : codigo.value
    }
    const respuesta = await fetch(ruta,{
        method: 'POST',
        body: JSON.stringify(bodyJson), // data can be `string` or {object}!
        headers:{
        'Content-Type': 'application/json'
        }
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    console.log(respuesta)
    if (respuesta.Err) {
        alert("Ocurrio un error")
        //errores = respuesta.Reporte_error
        return
    }
    consola.value = respuesta.Salida
    let dispararEvento = new Event("keyup")
    consola.dispatchEvent(dispararEvento)
}

btnTabla.addEventListener("click", obtenerTabla);
async function obtenerTabla() {
    const ruta = `http://localhost:3000/tablaSimbolos`;
    const respuesta = await fetch(ruta,{
        method: 'GET',
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    localStorage.setItem("reporte", respuesta.Salida)
    window.open(`./html/reporte.html`, "_blank");
    console.log(respuesta)
}

btnErrores.addEventListener("click", async (e) =>{
    const ruta = `http://localhost:3000/errores`;
    const respuesta = await fetch(ruta,{
        method: 'GET',
    })
    .then((res)=> res.json())
    .then((data) => {
        return data
    })
    localStorage.setItem("errores", respuesta.Salida)
    window.open(`./html/errores.html`, "_blank");
    console.log(respuesta)
});