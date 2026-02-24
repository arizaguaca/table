# Table Backend - Clean Architecture

Este es el backend para el proyecto "table", implementado siguiendo los principios de **Clean Architecture** en Go.

## Estructura del Proyecto

- `cmd/api/`: Punto de entrada de la aplicación. Aquí se realiza la inyección de dependencias.
- `internal/domain/`: Contiene las entidades de negocio y las interfaces (contratos) para repositorios y casos de uso. Esta capa no tiene dependencias externas.
- `internal/usecase/`: Implementa la lógica de negocio. Depende únicamente de las interfaces del dominio.
- `internal/repository/`: Implementaciones concretas de acceso a datos (ej. SQL, NoSQL, In-memory).
- `internal/http/`: Capa de transporte. Maneja las peticiones HTTP y formatea las respuestas.

## Cómo ejecutar

1. Asegúrate de tener Go instalado.
2. Ejecuta el servidor:
   ```bash
   go run cmd/api/main.go
   ```
3. El servidor estará disponible en `http://localhost:8080`.

## Endpoints

- `POST /tables`: Crear una nueva tabla.
- `GET /tables`: Listar todas las tablas.

app para realizar pedidos de forma autonoma en los restaurantes
