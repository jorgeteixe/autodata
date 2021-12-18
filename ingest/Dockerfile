FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chmod +x ./wait-for-it.sh

CMD [ "./wait-for-it.sh", "db:3306", "--strict" , "--timeout=300" , "--" , "python", "-u", "./ingest.py" ]