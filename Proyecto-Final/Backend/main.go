package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"proyecto-final/controllers"
	"proyecto-final/handlers"
	"proyecto-final/models"
	repositorio "proyecto-final/repository"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Car struct {
	Nombre       string `json:"nombre"`
	Marca        string `json:"marca"`
	Combustible  string `json:"combustible"`
	Transmision  string `json:"transmision"`
	Carroceria   string `json:"carroceria"`
	Modelo       int    `json:"modelo"`
	PrecioPorDia int    `json:"precio_por_dia"`
	Imagen       string `json:"imagen"`
}

var cars = []Car{
	{"Clase A 250e Progressive Line Advanced 8G-DCT", "Mercedes-Benz", "Híbrido", "Automática", "Sedán", 2020, 450000, "/Backend/Autos/Clase A 250e Progressive Line Advanced 8G-DCT 4.png"},
	{"Clase GLA 45 S AMG 4Matic 8G-DCT", "Mercedes-Benz", "Gasolina", "Automática", "SUV", 2016, 600000, "/Backend/Autos/Clase GLA 45 S AMG 4Matic 8G-DCT 5.png"},
	{"Clase GLB 200 7G-DCT", "Mercedes-Benz", "Gasolina", "Automática", "SUV", 2021, 500000, "/Backend/Autos/Clase GLB 200 7G-DCT 5.png"},
	{"MERCEDES-BENZ EQB 350 4Matic", "Mercedes-Benz", "Eléctrico", "Automática", "SUV", 2022, 700000, "/Backend/Autos/MERCEDES-BENZ EQB 350 4Matic 5.png"},
	{"EQE AMG 53 4Matic Edition", "Mercedes-Benz", "Eléctrico", "Automática", "Sedán", 2019, 800000, "/Backend/Autos/EQE AMG 53 4Matic Edition 4 Puertas.png"},
	{"BMW X7 2024", "BMW", "Híbrido", "Automática", "SUV", 2024, 900000, "/Backend/Autos/BMW X7 2024.png"},
	{"BMW X6 2024", "BMW", "Híbrido", "Automática", "SUV", 2024, 850000, "/Backend/Autos/BMW X6 2024.png"},
	{"BMW X3 2024", "BMW", "Eléctrico", "Automática", "SUV", 2024, 750000, "/Backend/Autos/BMW X3 2024.png"},
	{"BMW i5 2024", "BMW", "Eléctrico", "Automática", "Sedán", 2024, 800000, "/Backend/Autos/BMW i5 2024.png"},
	{"BMW M440i Convertible", "BMW", "Gasolina", "Automática", "Sedán", 2022, 700000, "/Backend/Autos/BMW M440i Convertible.png"},
	{"Mazda Sedan 2", "Mazda", "Gasolina", "Manual", "Sedán", 2012, 150000, "/Backend/Autos/Mazda Sedan 2.png"},
	{"Mazda CX-90 Híbrida", "Mazda", "Híbrido", "Automática", "SUV", 2024, 700000, "/Backend/Autos/Mazda CX-90 híbrida.png"},
	{"Mazda MX-30 Eléctrico", "Mazda", "Eléctrico", "Automática", "SUV", 2024, 650000, "/Backend/Autos/Mazda MX-30 eléctrico.png"},
	{"Mazda 3 2.0 Touring", "Mazda", "Gasolina", "Manual", "Sedán", 2017, 200000, "/Backend/Autos/Mazda Mazda 3 2.0 Touring.png"},
	{"Mazda CX-30 Touring Mecánica", "Mazda", "Gasolina", "Manual", "SUV", 2019, 250000, "/Backend/Autos/Mazda CX-30 Touring Mecánica.png"},
	{"Nissan Qashqai 2.0 Sense 140 hp", "Nissan", "Gasolina", "Manual", "SUV", 2024, 350000, "/Backend/Autos/Nissan Qashqai 2.0 Sense 140 hp.png"},
	{"Nissan Versa 1.6 Exclusive", "Nissan", "Gasolina", "Automática", "Sedán", 2022, 250000, "/Backend/Autos/Nissan Versa 1.6 Exclusive.png"},
	{"Nissan Patrol", "Nissan", "Gasolina", "Automática", "SUV", 2019, 400000, "/Backend/Autos/NISSAN PATROL.png"},
	{"Nissan Tiida 1.8 Miio", "Nissan", "Gasolina", "Manual", "Sedán", 2013, 150000, "/Backend/Autos/nissan-tiida-2008-732x488.png"},
	{"Nissan Murano Exclusive", "Nissan", "Gasolina", "Automática", "SUV", 2018, 350000, "/Backend/Autos/NISSAN MURANO EXCLUSIVE.png"},
	{"Audi Q5 45 TFSI MHEV quattro S line", "Audi", "Híbrido", "Automática", "SUV", 2024, 700000, "/Backend/Autos/Audi Q5 45 TFSI MHEV quattro S line.png"},
	{"Audi A4 40 TFSI Prestige Black", "Audi", "Gasolina", "Automática", "Sedán", 2014, 250000, "/Backend/Autos/Audi A4 40 TFSI Prestige Black.png"},
	{"Audi Q5 TFSI quattro Advanced Híbrido", "Audi", "Híbrido", "Automática", "SUV", 2021, 650000, "/Backend/Autos/Audi Q5 TFSI QUATTRO ADVANCED HIBRIDO.png"},
	{"Audi E-tron Advanced", "Audi", "Eléctrico", "Automática", "SUV", 2022, 800000, "/Backend/Autos/Audi E-tron Advanced.png"},
	{"Audi A3 1.2 Turbo", "Audi", "Gasolina", "Manual", "Sedán", 2015, 200000, "/Backend/Autos/Audi-A3-Sportback-2014.png"},
}

func getCars(w http.ResponseWriter, r *http.Request) {
	marca := r.URL.Query().Get("marca")
	combustible := r.URL.Query().Get("combustible")
	transmision := r.URL.Query().Get("transmision")
	carroceria := r.URL.Query().Get("carroceria")

	filteredCars := []Car{}

	for _, car := range cars {
		if (marca == "" || car.Marca == marca) &&
			(combustible == "" || car.Combustible == combustible) &&
			(transmision == "" || car.Transmision == transmision) &&
			(carroceria == "" || car.Carroceria == carroceria) {
			filteredCars = append(filteredCars, car)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(filteredCars)
}

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

//url := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
//os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
//db, err := sqlx.Connect("postgres", url)
//if err != nil {
//  log.Printf("Fallo la conexión a PostgreSQL, error: %s", err.Error())
//return nil, err
//}
//log.Printf("Nos conectamos bien a la base de datos: %#v", db)
//return db, nil
//}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/*err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	db, err := ConectarDB()
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	repo := repository.NewRepository[models.Usuario](db)
	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalf("Error al crear el controlador: %v", err)
	}

	userHandler := handlers.NewUserHandler(controller)
	carHandler := handlers.NewCarHandler() // Agregar un nuevo manejador de carros

	r := mux.NewRouter()*/

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Amigo y se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
	repo, err := repositorio.NewRepository[models.Usuario](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}
	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	r := mux.NewRouter()
	// Rutas relacionadas con los usuarios
	/*r.HandleFunc("/usuarios", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/usuarios", userHandler.ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/usuarios/{id}", userHandler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/usuarios/{id}", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/usuarios/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)*/

	r.Handle("/usuarios", http.HandlerFunc(handler.ListarUsuarios)).Methods(http.MethodGet)
	r.Handle("/usuarios", http.HandlerFunc(handler.CrearUsuario)).Methods(http.MethodPost)
	r.Handle("/usuarios/{id}", http.HandlerFunc(handler.TraerUsuario)).Methods(http.MethodGet)
	r.Handle("/usuarios/{id}", http.HandlerFunc(handler.ActualizarUsuario)).Methods(http.MethodPatch)
	r.Handle("/usuarios/{id}", http.HandlerFunc(handler.EliminarUsuario)).Methods(http.MethodDelete)
	// Rutas relacionadas con los carros

	//r.HandleFunc("/cars", carHandler.GetCars).Methods(http.MethodGet)
	r.HandleFunc("/cars", getCars).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
