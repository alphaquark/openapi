import os
from dotenv import load_dotenv
from pydantic import BaseSettings


env = load_dotenv('common.env')

class Settings(BaseSettings):
    redis_url: str = os.getenv('REDIS_URL')
    redis_host: str = os.getenv('REDIS_HOST')
    redis_port: int = os.getenv('REDIS_PORT')
    redis_db: int = os.getenv('REDIS_DB')