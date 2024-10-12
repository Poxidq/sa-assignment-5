from fastapi import FastAPI
from .routes import router as user_router  # Import the router
from .database import init_db

def create_app():
    app = FastAPI()
    app.include_router(user_router)  # Include the router
    init_db()  # Initialize the database tables on startup
    return app

app = create_app()
