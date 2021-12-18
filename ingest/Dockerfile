FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "./wait-for-it.sh", "db:3306", "--strict" , "--timeout=300" , "--" , "python", "./ingest.py" ]