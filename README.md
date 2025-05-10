# Ejemplo de API de Buscaminas para una arquitectura hexagonal.

Inicia el servidor

```bash
go run cmd/httpserver/main.go
```

Esto iniciará un servidor HTTP en el puerto 8080, donde las siguientes rutas estarán disponibles:

- `POST /games`: Crear un nuevo juego.

- `GET /games/{id}`: Obtener un juego por su ID.

- `PUT /games/{id}/reveal`: Revelar una celda de un juego.

