from fastapi import FastAPI, Depends, HTTPException
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

# Dependency to get DB session
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Pydantic models for response validation
class MessageResponse(BaseModel):
    id: int
    username: str
    content: str
    created_at: datetime
    likes: int

    class Config:
        orm_mode = True

# GET /feed - Get the last 10 messages
@app.get("/feed", response_model=List[MessageResponse])
def get_last_10_messages(db: Session = Depends(get_db)):
    messages = db.query(MessageDB).join(UserDB).order_by(MessageDB.created_at.desc()).limit(10).all()

    if not messages:
        raise HTTPException(status_code=404, detail="No messages found")

    return [
        MessageResponse(
            id=message.id,
            username=message.user.username,  # Retrieve username from the related user
            content=message.content,
            created_at=message.created_at,
            likes=message.likes,
        )
        for message in messages
    ]

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)
