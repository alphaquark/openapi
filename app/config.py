import os
import telegram
from dotenv import load_dotenv
from pydantic import BaseSettings


env = load_dotenv('common.env')

class Settings(BaseSettings):
    redis_url: str = os.getenv('REDIS_URL')
    redis_host: str = os.getenv('REDIS_HOST')
    redis_port: int = os.getenv('REDIS_PORT')
    redis_db: int = os.getenv('REDIS_DB')
    
    telegram_token: str = os.getenv('TELEGRAM_TOKEN')
    telegram_chat_id: str = os.getenv('TELEGRAM_CHAT_ID')

    etherscan_api_key: str = os.getenv('ETHERSCAN_API_KEY')
    address: str = os.getenv('ADDRESS')


def send_telegram_message(message: str):
    bot = telegram.Bot(token=Settings().telegram_token)
    bot.sendMessage(
        chat_id=Settings().telegram_chat_id, 
        text=message
    )
