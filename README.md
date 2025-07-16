# W17-G1-Bootcamp (MERCADO LIBRE - FRESCOS)

El objetivo de este proyecto es implementar una API REST, aplicando los conocimientos adquiridos en el BOOTCAMP-GO MELI (gestionando aspectos como control de versiones, desarrollo en Go, almacenamiento y aseguramiento de calidad). La iniciativa simula una nueva expansión de MercadoLibre, líder en e-commerce en LATAM, que busca incluir productos frescos (que requieren refrigeración) en su catálogo.
La incorporación de productos frescos implica nuevos desafíos en la forma de almacenar, manipular, transportar y comercializar este tipo de mercadería, garantizando condiciones óptimas y trazabilidad (como fecha de caducidad y número de lote). Además, se deben considerar distintas áreas de almacenamiento y envío para asegurar la calidad de los productos hasta su llegada al cliente final.

## Estándares

### Idiomas

Para mantener la consistencia y profesionalismo en el desarrollo del proyecto Mercado Libre - FRESCOS, se establecen los siguientes estándares de idioma que deben ser respetados por todos los miembros del equipo:

#### 📝 Código Fuente

* Idioma: Inglés
* Aplica a:
    * Nombres de variables, funciones, métodos y constantes
    * Nombres de archivos y directorios
    * Comentarios dentro del código
    * Mensajes de error y logs
    * Nombres de estructuras, interfaces y tipos

Ejemplo:
```go
package repository

// ✅ Correcto

func GetProductByID(productID int) (*Product, error) {
// Retrieve product from database
return repository.FindProduct(productID)
}

// ❌ Incorrecto

func ObtenerProductoPorID(idProducto int) (*Producto, error) {
// Obtener producto de la base de datos
return repositorio.BuscarProducto(idProducto)
}
```

#### 🔄 Commits

* Idioma: Inglés
* Aplica a:
    * Mensajes de commit siguiendo la convención Conventional Commits
    * Descripciones cortas y cuerpo del commit
    * Referencias a issues o tareas

Ejemplo:

```git
# ✅ Correcto
feat(products): add fresh product validation logic
fix(api): resolve authentication middleware error
docs(readme): update installation instructions

# ❌ Incorrecto
feat(productos): agregar lógica de validación de productos frescos
fix(api): resolver error en middleware de autenticación
```

#### 💻 Documentación de Código

* Idioma: Inglés
* Aplica a:
    * Comentarios de documentación (godoc)
    * Documentación de APIs generada automáticamente
    * Comentarios explicativos en el código
    * Documentación técnica interna

Ejemplo:

```go
package service

// ✅ Correcto

// ProductService handles business logic for fresh products
// It manages product lifecycle, validation, and storage requirements
type ProductService struct {
repository ProductRepository
}

// ❌ Incorrecto
// ProductService maneja la lógica de negocio para productos frescos
// Gestiona el ciclo de vida, validación y requisitos de almacenamiento
```

#### 📖 Documentación de Proyecto

* Idioma: Español
* Aplica a:
    * README.md principal del proyecto
    * Documentación de usuario final
    * Guías de instalación y configuración
    * Manuales de uso
    * Documentación de procesos y metodologías
    * Especificaciones funcionales

Ejemplo `README.md`:

```markdown
# ✅ Correcto
## Instalación
Para instalar el proyecto, sigue estos pasos:1. Clona el repositorio2. Ejecuta `go mod download`
## Uso
La API permite gestionar productos frescos con las siguientes funcionalidades:

# ❌ Incorrecto (en documentación de proyecto)
## Installation
To install the project, follow these steps:1. Clone the repository2. Run `go mod download`
```

#### 🎯 Consideraciones Adicionales

1. Consistencia: Una vez establecido el idioma para cada tipo de contenido, debe mantenerse consistente en todo el proyecto.
2. Colaboración: Todos los miembros del equipo deben seguir estos estándares para facilitar la colaboración y el mantenimiento del código.
3. Herramientas: Configurar el IDE/editor para detectar y sugerir correcciones según estos estándares.
4. Revisión: Durante el proceso de code review, verificar que se cumplan estos estándares de idioma.

Estos estándares aseguran que el proyecto mantenga una estructura profesional y sea accesible tanto para el equipo de desarrollo técnico (inglés) como para los stakeholders del negocio (español).

### Flujo de trabajo

Git flow es una metodología de trabajo para gestionar ramas en proyectos que utilizan Git como sistema de control de versiones. Su propósito principal es proporcionar una estructura clara y ordenada para el desarrollo de software, facilitando la colaboración entre equipos y ayudando a organizar las diferentes fases del ciclo de vida de una aplicación, como el desarrollo de nuevas funcionalidades, la corrección de errores y la publicación de versiones.
En Git flow, hay cinco tipos de ramas diferentes:

1. `Main`
2. `Develop`
3. `Feature`
4. `Release`
5. `Hotfix`

![Git flow](https://www.gitkraken.com/wp-content/uploads/2021/03/git-flow-4.svg)

### Conventional commits

Es una convención para escribir commits en proyectos que utilizan control de versiones, como Git. Su propósito principal es estandarizar la forma en que se redactan estos mensajes, facilitando así la comprensión de los cambios realizados en el código, la automatización de procesos (como el versionado semántico) y la colaboración entre desarrolladores.

```git
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

* `feat`: cuando se añade una nueva funcionalidad.
* `fix`: cuando se arregla un error.
* `chore`: tareas rutinarias que no sean específicas de una feature o un error como por ejemplo añadir contenido al fichero .gitignore o instalar una dependencia.
* `test`: si añadimos o arreglamos tests.
* `docs`: cuando solo se modifica documentación.
* `build`: cuando el cambio afecta al compilado del proyecto.
* `ci`: el cambio afecta a ficheros de configuración y scripts relacionados con la integración continua.
* `style`: cambios de legibilidad o formateo de código que no afecta a funcionalidad.
* `refactor`: cambio de código que no corrige errores ni añade funcionalidad, pero mejora el código.
* `perf`: usado para mejoras de rendimiento.
* `revert`: si el commit revierte un commit anterior. Debería indicarse el hash del commit que se revierte.

## 🔧 Configuración de Variables de Entorno

Este proyecto utiliza variables de entorno para configurar diferentes aspectos de la aplicación y la base de datos. La configuración se gestiona a través de un archivo `.env` que Docker Compose lee automáticamente.

### Archivo .env

El archivo `.env` debe crearse en la raíz del proyecto y contiene todas las variables de configuración necesarias:

```dotenv
# Database Configuration
MYSQL_ROOT_PASSWORD=your_root_password_here
MYSQL_DATABASE=your_database_here
MYSQL_USER=your_user_here
MYSQL_PASSWORD=your_password_here
MYSQL_CHARACTER_SET_SERVER=your_character_set_here
MYSQL_COLLATION_SERVER=your_collation_here
MYSQL_PORT=your_mysql_port_here

# Application Configuration
APP_PORT=your_app_port_here
```

### Descripción de Variables

#### 🗄️ Configuración de Base de Datos

| Variable                     | Descripción                                   | 
|------------------------------|-----------------------------------------------|
| `MYSQL_ROOT_PASSWORD`        | Contraseña del usuario root de MySQL          |
| `MYSQL_DATABASE`             | Nombre de la base de datos que se creará      |
| `MYSQL_USER`                 | Usuario de aplicación para conectarse a MySQL |
| `MYSQL_PASSWORD`             | Contraseña del usuario de aplicación          |
| `MYSQL_CHARACTER_SET_SERVER` | Conjunto de caracteres del servidor MySQL     |
| `MYSQL_COLLATION_SERVER`     | Collation del servidor MySQL                  |
| `MYSQL_PORT`                 | Puerto donde MySQL aceptará las conexiones    |


#### 🌐 Configuración de la Aplicación (FRESCOS)

| Variable | Descripción |
|----------|-------------|
| `APP_PORT` | Puerto donde se expone la aplicación Go |

## Estructura del proyecto

```markdown
W17-G1-Bootcamp
├── README.md
├── cmd
│   └── main.go
├── docs
│   └── db
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   └── application_default.go
│   ├── handler
│   ├── loader
│   │   └── json.go
│   ├── repository
│   └── service
└── pkg
└── models
```

## 🐳 Docker

Este proyecto incluye configuración completa de Docker para facilitar el desarrollo y despliegue. La aplicación utiliza un build multi-etapa para optimizar el tamaño de la imagen final.

### Prerrequisitos

- Docker Engine 20.10+
- Docker Compose 2.0+

### Ejecutar con Docker Compose (Recomendado)

La forma más sencilla de ejecutar el proyecto completo (aplicación + base de datos MySQL):

```bash
# Construir y ejecutar todos los servicios
docker compose up --build

# Ejecutar en segundo plano (detached mode)
docker compose up --build -d

# Ver logs
docker compose logs app
docker compose logs database

# Detener todos los servicios
docker compose down

# Detener y eliminar volúmenes (reinicio completo)
docker compose down --volumes
```

### Construcción manual con Docker

Si prefieres usar Docker directamente sin Compose:

```bash
# Construir la imagen
docker build -t frescos-app .

# Ejecutar el contenedor
docker run -p 8080:8080 frescos-app

# Ejecutar en segundo plano
docker run -d -p 8080:8080 --name frescos-container frescos-app

# Ver logs
docker logs frescos-container

# Detener y eliminar el contenedor
docker stop frescos-container
docker rm frescos-container
```

## Recursos

1. GitFlow
    1. https://www.atlassian.com/es/git/tutorials/comparing-workflows/gitflow-workflow
    2. https://www.gitkraken.com/learn/git/git-flow
    3. https://danielkummer.github.io/git-flow-cheatsheet/index.html
    4. https://nvie.com/posts/a-successful-git-branching-model/
2. Conventional commits
    1. https://www.conventionalcommits.org/en/v1.0.0/
    2. https://dev.to/achamorro_dev/conventional-commits-que-es-y-por-que-deberias-empezar-a-utilizarlo-23an
3. Mockaroo (generar datos de prueba)
    1. https://www.mockaroo.com/
4. Docker
   1. https://www.docker.com/101-tutorial/
   2. https://docs.docker.com/compose/gettingstarted/