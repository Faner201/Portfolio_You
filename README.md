# Portfolio_You

REST API to create a portfolio, written using the GIN.

## API Endpoints:
### User interaction.
- [x] `POST /sign-up` - New user registration.
- [x] `POST /sign-in` - New user authorization.
---
- [x] `POST /portfolio/create` - Creating a new portfolio.
- [x] `POST /portfolio/create/menu` - Create a brief description of the portfolio.
- [x] `POST /portfolio/open` - Viewing the portfolio made.
- [x] `GET /portfolio/menu` - Display a list of user-created portfolios.
- [x] `DELETE /portfolio/menu` - Deleting a portfolio.

## Project structure

```bash
portfolio_you
|--- auth
        |--- delivery # Working with handlers.
        |--- repository
                |--- database # Working with database.
                |--- localstorage # Working with localstorage.
                |--- mock # Interface plugs under localstorage
        |--- usecase # Implementation of different business logic.
        |--- error.go # All errors used at the handler level.
        |--- repository.go # Interface repository.
        |--- usecase.go # Interface Use Case.
|--- cmd
        |--- main.go # Initializing configuration files, starting the server.
|--- config
        |--- config.yml # Configuration file of the whole project.
        |--- init.go # Read a configuration file.
|--- models # Main project entities.
|--- portfolios
        |--- .... # Repeats the hierarchy from auth.         
```
      