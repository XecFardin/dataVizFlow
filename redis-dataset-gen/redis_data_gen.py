import json
import random
import redis
import time

r = redis.Redis(host='127.0.0.1', port=6379, db=0)

import random
from datetime import datetime, timedelta

def generate_birth_data(length):
    birth_data = []
    current_year = datetime.now().year
    for _ in range(length):
        year = random.randint(current_year - 70, current_year - 18)
        birth_rate = round(random.uniform(0, 1), 3)  # Birth rate per woman
        data_entry = {
            "year": year,
            "birth_rate": birth_rate
        }
        birth_data.append(data_entry)
    return birth_data

def generate_internet_tower_data(length):
    internet_data = [
    {"location": "Mumbai, India", "internet_speed": 100.5},
    {"location": "Delhi, India", "internet_speed": 98.2},
    {"location": "Bangalore, India", "internet_speed": 105.8},
    {"location": "Hyderabad, India", "internet_speed": 97.3},
    {"location": "Chennai, India", "internet_speed": 101.1},
    {"location": "Kolkata, India", "internet_speed": 99.9},
    {"location": "Ahmedabad, India", "internet_speed": 96.7},
    {"location": "Pune, India", "internet_speed": 102.4},
    {"location": "Jaipur, India", "internet_speed": 94.6},
    {"location": "Lucknow, India", "internet_speed": 97.9},
    {"location": "Surat, India", "internet_speed": 98.3},
    {"location": "Kanpur, India", "internet_speed": 95.2},
    {"location": "Nagpur, India", "internet_speed": 99.0},
    {"location": "Indore, India", "internet_speed": 97.5},
    {"location": "Thane, India", "internet_speed": 101.2},
    {"location": "Bhopal, India", "internet_speed": 96.4},
    {"location": "Visakhapatnam, India", "internet_speed": 98.7},
    {"location": "Pimpri-Chinchwad, India", "internet_speed": 103.6},
    {"location": "Patna, India", "internet_speed": 96.8},
    {"location": "Vadodara, India", "internet_speed": 99.3}]
    return internet_data

def generate_weather_data(length):
    weather_data = []
    current_date = datetime.now()
    for i in range(length):
        date = (current_date - timedelta(days=i)).strftime('%Y-%m-%d')
        temperature = round(random.uniform(-20, 40), 2)  # Temperature in Celsius
        humidity = round(random.uniform(0, 100), 2)      # Humidity in percentage
        wind_speed = round(random.uniform(0, 30), 2)     # Wind speed in km/h
        precipitation = round(random.uniform(0, 20), 2) # Precipitation in mm

        data_entry = {
            "date": date,
            "temperature": temperature,
            "humidity": humidity,
            "wind_speed": wind_speed,
            "precipitation": precipitation
        }

        weather_data.append(data_entry)

    return weather_data

    

def send_data_to_redis(Key,data):

    json_data = json.dumps(data)

    r.set(Key, json_data)

    print(f" {json_data}")

    time.sleep(3)  

if __name__ == "__main__":
    data_length = 35
    random_birth_data = generate_birth_data(data_length)
    random_internet_data = generate_internet_tower_data(data_length)
    random_weather_data = generate_weather_data(data_length)

    send_data_to_redis("birthRate",random_birth_data)
    send_data_to_redis("internet",random_internet_data)
    send_data_to_redis("weather",random_weather_data)
