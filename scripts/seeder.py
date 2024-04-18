import random
import string
from datetime import datetime

import psycopg2
from faker import Faker

sql_facility_types_insert_query = """
INSERT INTO facility_types
(uuid, "name", bahasa_name, created_at, updated_at)
VALUES(%s, %s, %s, %s, %s) RETURNING id
"""

sql_location_types_insert_query = """
INSERT INTO location_types
(uuid, "name", created_at, updated_at)
VALUES(%s, %s, %s, %s) RETURNING id
"""

sql_jobs_insert_query = """
INSERT INTO jobs
(uuid, "name", created_at, updated_at)
VALUES(%s, %s, %s, %s) RETURNING id
"""

sql_job_details_insert_query = """
INSERT INTO job_details
(uuid, job_id, "name", created_at, updated_at)
VALUES(%s, %s, %s, %s, %s) RETURNING id
"""

def seed_facility_types():
    facilities_types = {
        "Healthcare Provider": "Fasilitas Pelayan Kesehatan",
        "Hospital Department": "Departemen dalam Rumah Sakit",
        "Organizational Team": "Tenaga Kesehatan yang Menjalankan Fungsi tertentu",
        "Government": "Organisasi Pemerintah",
        "Insurance Company": "Perusahaan Asuransi",
        "Payer": "Badan Penjamin",
        "Educational Institute": "Institusi Pendidikan / Penelitian",
        "Religious Institution": "Organisasi Keagamaan",
        "Clinical Research Sponsor": "Sponsor Penelitian Klinis",
        "Community": "Kelompok",
        "Group": "Masyarakat",
        "Non-Healthcare Business Corporation": "Perusahaan diluar bidang kesehatan",
        "Other": "Lain-lain",
    }

    for name, bahasaName in facilities_types.items():
        uuid = fake.uuid4()
        cur.execute(sql_facility_types_insert_query,
                    (uuid, name, bahasaName, datetime.now(), datetime.now()))

def seed_location_types():
    location_types = [
        "Site",
        "Building",
        "Wing",
        "Ward",
        "Level",
        "Corridor",
        "Room",
        "Bed",
        "Vehicle",
        "House",
        "Cabinet",
        "Road",
        "Area",
        "Jurisdiction",
        "Virtual",
    ]

    for name in location_types:
        uuid = fake.uuid4()
        cur.execute(sql_location_types_insert_query,
                    (uuid, name, datetime.now(), datetime.now()))

# Create Faker instance
fake = Faker()

# Define your PostgreSQL connection parameters
DATABASE_NAME = "databasename"
USER = "user"
PASSWORD = "password"
HOST = "localhost"
PORT = "5432"

try:
    # Establish a connection to the database
    conn = psycopg2.connect(
        dbname=DATABASE_NAME,
        user=USER,
        password=PASSWORD,
        host=HOST,
        port=PORT,
    )

    # Create a cursor object
    cur = conn.cursor()
    seed_facility_types()
    seed_location_types()
    conn.commit()

    # Close the cursor and connection
    cur.close()
    conn.close()

except (Exception, psycopg2.DatabaseError) as error:
    print("Error while inserting data into PostgreSQL table", error)

finally:
    if conn is not None:
        cur.close()
        conn.close()
