# Egresos Hospitalarios - API REST

Este proyecto es una API REST desarrollada en Go para gestionar egresos hospitalarios. Proporciona endpoints para crear, ver y consultar registros de egresos hospitalarios.

## Requisitos Previos

- Go 1.16 o superior
- Docker y Docker Compose (opcional, para despliegue en contenedores)
- SQLite3

## Instalación

### Instalación Local

1. Clonar el repositorio:
```bash
git clone [URL_DEL_REPOSITORIO]
cd Egresos-Hospitalarios
```

2. Instalar dependencias:
```bash
go mod download
```

3. Inicializar la base de datos:
```bash
go run main.go
```

### Despliegue con Docker

1. Construir la imagen:
```bash
docker build -t egresos-hospitalarios .
```

2. Ejecutar el contenedor:
```bash
docker-compose up
```

## Uso de la API

La API está disponible en `http://localhost:8191` y proporciona los siguientes endpoints:

### Crear un nuevo registro
```bash
curl -X POST http://localhost:8191/crear \
  -H "Content-Type: application/json" \
  -d '{
    "campo1": "valor1",
    "campo2": "valor2"
  }'
```

### Ver todos los registros
```bash
curl -X GET http://localhost:8191/ver
```

### Ver un registro específico por ID
```bash
curl -X GET http://localhost:8191/ver/{id}
```

## Estructura del Proyecto

```
.
├── main.go              # Punto de entrada de la aplicación
├── db/                  # Configuración y operaciones de base de datos
├── handlers/            # Manejadores de endpoints
├── models/              # Modelos de datos
├── Dockerfile          # Configuración para contenedor Docker
└── docker-compose.yml  # Configuración de servicios Docker
```

## Base de Datos

El proyecto utiliza SQLite3 como base de datos. El archivo de base de datos se encuentra en `egresosh.db`.

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo [LICENSE](LICENSE) para más detalles.
 
