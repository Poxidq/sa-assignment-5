from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from .database import get_db
from .models import User

app = FastAPI()

@app.post("/register", status_code=201)
def register_user(username: str, db: Session = Depends(get_db)):
    # Check if user exists
    existing_user = db.query(User).filter(User.username == username).first()
    if existing_user:
        raise HTTPException(status_code=400, detail="Username already exists")
    
    # Create new user
    user = User(username=username)
    db.add(user)
    db.commit()
    return {"username": username}

@app.post("/login")
def login_user(username: str, db: Session = Depends(get_db)):
    # Check if user exists
    user = db.query(User).filter(User.username == username).first()
    if not user:
        raise HTTPException(status_code=400, detail="Invalid username")
    
    return {"username": username}

@app.get("/me")
def get_me(username: str, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.username == username).first()
    
    if not user:
        raise HTTPException(status_code=400, detail="User not found")
    
    return {"username": user.username}