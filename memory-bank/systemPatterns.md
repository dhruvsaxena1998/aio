# System Patterns

## System Architecture

The Habit Tracking API is designed as a RESTful service built with Go (Golang). The architecture is modular and follows the MVC (Model-View-Controller) pattern, with some adjustments to fit the API nature.

### Key Components

1. **Models**: Represent the data structure and interact with the database.
2. **Handlers**: Handle HTTP requests and responses.
3. **Middlewares**: Provide cross-cutting concerns like authentication and logging.
4. **Routes**: Define the API endpoints and their corresponding handlers.
5. **Database**: Stores user and habit data using PostgreSQL.

### Design Patterns

1. **MVC Pattern**: Separation of concerns between models, handlers, and routes.
2. **Singleton Pattern**: For database connection management.
3. **Middleware Pattern**: For request processing and response handling.
4. **JWT Authentication**: For secure user authentication.

### Component Relationships

- **Models** interact with the **Database** to perform CRUD operations.
- **Handlers** use **Models** to process data and return responses.
- **Middlewares** are applied to **Routes** to handle authentication and logging.
- **Routes** map HTTP requests to **Handlers**.

### Critical Implementation Paths

1. **User Authentication**: Implemented using JWT tokens.
2. **Habit Management**: CRUD operations for habits.
3. **Completion Tracking**: Track habit completions and provide analytics.
