# Ejemplo de API de Buscaminas para una arquitectura hexagonal.

Inicia el servidor

```bash
go run cmd/httpserver/main.go
```

Esto iniciará un servidor HTTP en el puerto 8080, donde las siguientes rutas estarán disponibles:

- `POST /games`: Crear un nuevo juego.

- `GET /games/{id}`: Obtener un juego por su ID.

- `PUT /games/{id}/reveal`: Revelar una celda de un juego.

Si quieres saber más acerca de cómo se juega este juego legendario, sus reglas, patrones y más ...
[Haz click](https://minesweepergame.com/strategy/how-to-play-minesweeper.php)
