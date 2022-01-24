import redis
import fakeredis
from fastapi import FastAPI
from fastapi.responses import JSONResponse

import config


app = FastAPI(title='Alphaquark-api',
            description='Getting Alphaquark information',
            openapi_url='/api/openapi.json')

settings = config.Settings()
print(settings)

fredis = fakeredis.FakeStrictRedis(decode_responses=True)
currency = ['KRW', 'USD', 'IDR', 'SGD', 'THB']

class Aqt():
    host = settings.redis_host
    port = settings.redis_port
    db = settings.redis_db
    fredis = fakeredis.FakeStrictRedis()


# health check
@app.get('/health')
def health_check():
    if Aqt.host is None:
        redisdb = fredis
    else:
        redisdb = redis.Redis(host = Aqt.host, port = Aqt.port, db = Aqt.db, decode_responses=True)
    try:
        redisdb.ping()
        print('Successfully connected to redis')
    except redis.exceptions.ConnectionError as e:
        print('Redis connection error')
        return {"message" : "DB False"}
    return {"message" : "healthy"}


@app.get('/api/aqt/info')
def aqt_info():
    if Aqt.host is None:
        redisdb = fredis
    else:
        redisdb = redis.Redis(host = Aqt.host, port = Aqt.port, db = Aqt.db, decode_responses=True)
    results = []
    try:
        for s in currency:
            data = redisdb.hgetall(s)
            for key, value in data.items():
                try:
                    data[key] = int(value) if float(value) % 1 == 0 else float(value)
                except:
                    data[key] = str(value)
            results.append(data)
    except Exception as e:
        print(e)
        return e
    print(results)
    return JSONResponse(content=results, status_code=200)