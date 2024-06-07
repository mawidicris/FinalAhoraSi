package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"proyecto-final/controllers"

	"github.com/gorilla/mux"
	//"proyecto-final/models"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la ruta
	vars := mux.Vars(r)
	id := vars["id"]

	// Decodificar el cuerpo de la solicitud en un usuario
	//var updatedUser models.Usuario
	//err := json.NewDecoder(r.Body).Decode(&updatedUser)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	// Convertir la estructura de usuario a []byte
	//userBytes, err := json.Marshal(updatedUser)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Llamar al controlador para actualizar el usuario
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un usuario, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.controller.ActualizarUnUsuario(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al actualizar un usuario, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la ruta
	vars := mux.Vars(r)
	id := vars["id"]

	// Llamar al controlador para eliminar el usuario
	err := h.controller.EliminarUnUsuario(id)
	if err != nil {
		log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
		http.Error(w, fmt.Sprintf("fallo al eliminar un comentario, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) TraerUsuario(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la ruta
	vars := mux.Vars(r)
	id := vars["id"]

	// Llamar al controlador para leer el usuario
	usuario, err := h.controller.LeerUnUsuario(id)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("el cusuario con id %s no se pudo encontrar", id)))
		return
	}

	// Serializar la respuesta en JSON
	//jsonResponse, err := json.Marshal(user)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Establecer el encabezado y escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usuario)
}

func (h *Handler) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Llamar al controlador para listar usuarios
	usuarios, err := h.controller.ListarUsuarios()
	if err != nil {
		log.Printf("fallo al leer los usuarios, con error: %s", err.Error())
		http.Error(w, "fallo al leer los usuarios", http.StatusInternalServerError)
		return
	}

	// Serializar la respuesta en JSON
	//jsonResponse, err := json.Marshal(users)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Establecer el encabezado y escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usuarios)
}

func (h *Handler) CrearUsuario(writer http.ResponseWriter, req *http.Request) {
	// Decodificar el cuerpo de la solicitud en un nuevo usuario
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo usuario", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Llamar al controlador para crear el usuario
	nuevoID, err := h.controller.CrearUsuario(body)
	if err != nil {
		log.Println("fallo al crear un nuevo usuario, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo usuario", http.StatusInternalServerError)
		return
	}

	// Responder con el ID del nuevo usuario creado
	//response := map[string]int64{"id": userID}
	//jsonResponse, err := json.Marshal(response)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nuevo usuario: %d", nuevoID)))
}
