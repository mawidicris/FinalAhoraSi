const container = document.querySelector(".container");
const btnSignIn = document.getElementById("btn-sign-in");
const btnSignUp = document.getElementById("btn-sign-up");


btnSignIn.addEventListener("click",()=>{
container.classList.remove("toggle");
});
btnSignUp.addEventListener("click",()=>{
container.classList.add("toggle");
});

// Función para el registro de usuarios
const registerUser = async (data) => {
    try {
        const response = await fetch('http://localhost:8080/usuarios', {
            method: 'POST',mode:"no-cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        console.log("Respuesta del servidor:", response); // Verificar respuesta del servidor

        if (response.status != 0) {
            throw new Error('Error al registrar usuario');
        }
        alert('Usuario registrado correctamente');
    } catch (error) {
        console.error('Error:', error);
        alert('Ocurrió un error al registrar usuario');
    }
};

// Event listener para el botón de registro
document.getElementById("btn-registro").addEventListener("click", function(event) {
    event.preventDefault(); // Evitar envío del formulario

    // Obtener los datos de registro ingresados por el usuario
    const peticion = {
        ID: 0,
        Nombre: document.getElementById("user_name")?.value,
        Correo: document.getElementById("user_email")?.value,
        Password: document.getElementById("user_password")?.value

    } 

    console.log(peticion)
    // Llamar a la función para registrar al usuario
    registerUser(peticion);
});