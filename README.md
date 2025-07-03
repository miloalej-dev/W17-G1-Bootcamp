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

