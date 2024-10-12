from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy import create_engine, Column, Integer, String, DateTime
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, Session
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
class MessageDB(Base):
    __tablename__ = 'messages'

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, nullable=False)
    content = Column(String, nullable=False)
    created_at = Column(DateTime, default=datetime.utcnow)
    likes = Column(Integer, default=0)

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

# GET /messages/last10 - Get last 10 messages
@app.get("/feed", response_model=List[MessageResponse])
def get_last_10_messages(db: Session = Depends(get_db)):
    messages = db.query(MessageDB).order_by(MessageDB.created_at.desc()).limit(10).all()
    
    if not messages:
        raise HTTPException(status_code=404, detail="No messages found")

    return messages

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=5000)
