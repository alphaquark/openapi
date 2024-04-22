import redis
import json
import config
import telegram
import logging
import requests
from fastapi import FastAPI, Depends
from fastapi.responses import JSONResponse


logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.INFO)


app = FastAPI(title='Alphaquark-api',
            description='Getting Alphaquark information',
            openapi_url='/api/openapi.json')

settings = config.Settings()
currency = ['KRW', 'USD', 'IDR', 'SGD', 'THB']


def get_redis():
    host = settings.redis_host
    port = settings.redis_port
    db = settings.redis_db
    return redis.Redis(host=host, port=port, db=db, decode_responses=True)


# health check
@app.get('/health')
def health_check(redisdb : get_redis = Depends(get_redis)):
    try:
        redisdb.ping()
        logging.info('Successfully connected to redis')
    except redis.exceptions.ConnectionError as e:
        logging.error('Redis connection error')
        config.send_telegram_message(f"alphqqark-api error: path=/health {e}")
        return {"message" : "DB False"}
    return {"message" : "healthy"}


@app.get('/api/aqt/info')
def aqt_info(redisdb : get_redis = Depends(get_redis)):
    results = []
    try:
        for s in currency:
            data = redisdb.hgetall(s)
            new_data = {}
            for key, value in data.items():
                lower_key = key[0].lower() + key[1:]
                try:
                    if lower_key == "lastUpdatedTimestamp":
                        new_data[lower_key] = int(value) * 1000
                    else:
                        new_data[lower_key] = int(value) if float(value) % 1 == 0 else float(value)
                except:
                    new_data[lower_key] = str(value)
            results.append(new_data)
    except Exception as e:
        logging.error(f"/api/aqt/info error: {e}")
        config.send_telegram_message(f"alphqqark-api error: path=/api/aqt/info {e}")
        return e
    logging.info(results)
    return JSONResponse(content=results, status_code=200)


@app.get('/api/aqt/info/circulatingSupply')
def aqt_circulating_supply(redisdb : get_redis = Depends(get_redis)):
    data = redisdb.hget('USD', 'CirculatingSupply')
    if data is None:
        config.send_telegram_message(
            "alphqqark-api error: path=/api/aqt/info/circulatingSupply data is None"
        )
        raise ValueError("circulatingSupply is None")
    return JSONResponse(int(data), status_code=200)


@app.get('/api/aqt/info/totalSupply')
def aqt_total_supply(redisdb : get_redis = Depends(get_redis)):
    try:
        url =f'https://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress={settings.address}&apikey={settings.etherscan_api_key}'
        res = requests.get(url)
        res_json = res.json()
        wei_value = res_json["result"]
        total_supply = int(float(wei_value) / 10**18)
        redisdb.set('TotalSupply', total_supply)
    except Exception as e:
        config.send_telegram_message(
            "alphqqark-api error: TotalSupply error."
        )
        raise ValueError("TotalSupply error.")
    return JSONResponse(total_supply, status_code=200)
