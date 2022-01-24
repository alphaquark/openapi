import pytest
import fakeredis
import json
import sys
import os
from fastapi.testclient import TestClient


SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
sys.path.append(os.path.dirname(SCRIPT_DIR))

from main import app, currency, Aqt, fredis


def health_check_test():
    client = TestClient(app)
    response = client.get('/health')
    assert response == 'DB OK'
    assert response.status_code == 200  

def store_fake_db():
    insert_data = [
        {
            "KRW":
                {
                "symbol":"AQT",
                "currencyCode":"KRW",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
                }
            },
        {
            "USD":
                {
                "symbol": "AQT",
                "currencyCode":"USD",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
                }
            },
        {
            "IDR":
                {
                "symbol": "AQT",
                "currencyCode": "IDR",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
                }
            },
        {
            "SGD":
                {
                "symbol": "AQT",
                "currencyCode": "SGD",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
                }
            },
        {
            "THB":
                {
                "symbol": "AQT",
                "currencyCode": "THB",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            }
        }
    ]

    for data in insert_data:
        k = next(iter(data))
        print(k)
        print(data[k])
        value = data[k]
        fredis.hmset(k, value)



def delete_fade_db():
    for k in currency:
        fredis.delete(k)


def test_aqt_info():
    store_fake_db()
    client = TestClient(app)
    Aqt.host = None
    response = client.get('/api/aqt/info')
    print("********************")
    print(response.text)
    print("********************")

    assert response.status_code == 200
    assert response.json() == [
            {
                "symbol":"AQT",
                "currencyCode":"KRW",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            },
            {
                "symbol": "AQT",
                "currencyCode": "USD",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            },
            {
                "symbol": "AQT",
                "currencyCode": "IDR",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            },
            {
                "symbol": "AQT",
                "currencyCode": "SGD",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            },
            {
                "symbol": "AQT",
                "currencyCode": "THB",
                "price": 1234,
                "marketCap": 1234.1234,
                "accTradePrice24h": 1234,
                "circulatingSupply": 1234,
                "maxSupply": 3000,
                "provider": "Alpha Quark Limited",
                "lastUpdatedTimestamp": 1234
            }
        ]