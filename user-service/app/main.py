from fastapi import FastAPI
from .routes import app as user_routes
from .database import init_db

def create_app():
    app = FastAPI()
    app.include_router(user_routes)
    init_db()  # Initialize the database tables on startup
    return app

app = create_app()
