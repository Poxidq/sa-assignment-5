from fastapi import FastAPI, HTTPException, Depends
from sqlalchemy import create_engine, Column, Integer, String, DateTime, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, Session
from pydantic import BaseModel
from datetime import datetime
from typing import List
import os

# FastAPI app
app = FastAPI()

# Database setup
DATABASE_URL = os.getenv("DATABASE_URL")

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

# Database models
class MessageDB(Base):
    __tablename__ = 'messages'

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, nullable=False)
    content = Column(String, nullable=False)
    created_at = Column(DateTime, default=datetime.utcnow)
    likes = Column(Integer, default=0)

# Create tables if not exist
Base.metadata.create_all(bind=engine)

# Dependency to get DB session
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Pydantic models for request/response validation
class CreateMessage(BaseModel):
    username: str
    content: str

class MessageResponse(BaseModel):
    id: int
    username: str
    content: str
    created_at: datetime
    likes: int

# POST /messages - Create a new message
@app.post("/messages", response_model=MessageResponse, status_code=201)
def create_message(message_data: CreateMessage, db: Session = Depends(get_db)):
    message = MessageDB(
        username=message_data.username,
        content=message_data.content
    )
    db.add(message)
    db.commit()
    db.refresh(message)
    return message

# GET /messages - Get all messages
@app.get("/messages", response_model=List[MessageResponse])
def get_messages(db: Session = Depends(get_db)):
    messages = db.query(MessageDB).all()
    return messages

# POST /messages/{id}/like - Like a message
@app.post("/messages/{id}/like", response_model=MessageResponse)
def like_message(id: int, db: Session = Depends(get_db)):
    message = db.query(MessageDB).filter(MessageDB.id == id).first()
    
    if not message:
        raise HTTPException(status_code=404, detail="Message not found")

    message.likes += 1
    db.commit()
    db.refresh(message)
    return message
