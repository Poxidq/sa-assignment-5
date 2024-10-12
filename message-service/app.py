from fastapi import FastAPI, HTTPException, Depends
from sqlalchemy import create_engine, Column, Integer, String, DateTime, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, Session, relationship
from pydantic import BaseModel
from typing import List
from datetime import datetime
import os

# FastAPI app
app = FastAPI()

# Database setup
DATABASE_URL = os.getenv("DATABASE_URL")

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

# Database models
class UserDB(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String(255), nullable=False, unique=True)

class MessageDB(Base):
    __tablename__ = 'messages'

    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(Integer, ForeignKey('users.id'), nullable=False)
    content = Column(String(400), nullable=False)
    created_at = Column(DateTime, default=datetime.utcnow)
    likes = Column(Integer, default=0)

    user = relationship("UserDB", back_populates="messages")

UserDB.messages = relationship("MessageDB", order_by=MessageDB.id, back_populates="user")

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

    class Config:
        orm_mode = True

# POST /messages - Create a new message
@app.post("/messages", response_model=MessageResponse, status_code=201)
def create_message(message_data: CreateMessage, db: Session = Depends(get_db)):
    # Find the user by username
    user = db.query(UserDB).filter(UserDB.username == message_data.username).first()

    # If the user doesn't exist, raise an error
    if not user:
        raise HTTPException(status_code=404, detail="User not found")

    # Create a new message associated with the user
    message = MessageDB(
        user_id=user.id,
        content=message_data.content
    )
    db.add(message)
    db.commit()
    db.refresh(message)

    # Return the response with the message details and the user's username
    return MessageResponse(
        id=message.id,
        username=user.username,
        content=message.content,
        created_at=message.created_at,
        likes=message.likes
    )

# GET /messages - Get all messages
@app.get("/messages", response_model=List[MessageResponse])
def get_messages(db: Session = Depends(get_db)):
    # Join users and messages to get the username for each message
    messages = db.query(MessageDB).join(UserDB).all()

    return [
        MessageResponse(
            id=message.id,
            username=message.user.username,
            content=message.content,
            created_at=message.created_at,
            likes=message.likes
        ) for message in messages
    ]

# POST /messages/{id}/like - Like a message
@app.post("/messages/{id}/like", response_model=MessageResponse)
def like_message(id: int, db: Session = Depends(get_db)):
    # Find the message by its ID
    message = db.query(MessageDB).filter(MessageDB.id == id).first()

    if not message:
        raise HTTPException(status_code=404, detail="Message not found")

    # Increment the likes
    message.likes += 1
    db.commit()
    db.refresh(message)

    # Return the updated message with the user's username
    return MessageResponse(
        id=message.id,
        username=message.user.username,
        content=message.content,
        created_at=message.created_at,
        likes=message.likes
    )


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)
