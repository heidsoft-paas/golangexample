import logging
import boto
import boto.s3.connection
access_key = 'PC8URMEKC3IO1VS752KF'
secret_key = 'gmcECdJiPkLQU07Ljm/iaWDQaYeLvkWyQAB+EBh3'

logging.basicConfig(filename="boto.log", level=logging.DEBUG)
conn = boto.s3.connection.S3Connection(
        aws_access_key_id = access_key,
        aws_secret_access_key = secret_key,
        host = 'host2',
        is_secure=False,
        calling_format = boto.s3.connection.OrdinaryCallingFormat(),
        )
bucket = conn.delete_bucket("docker-image-bucket")

for bucket in conn.get_all_buckets():
        print "{name}\t{created}".format(
                name = bucket.name,
                created = bucket.creation_date,
        )
