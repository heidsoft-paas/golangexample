import boto
import boto.s3.connection
#access_key = 'PC8URMEKC3IO1VS752KF'
#secret_key = 'gmcECdJiPkLQU07Ljm/iaWDQaYeLvkWyQAB+EBh3'

#sz
access_key = '9HAGI3A40XXC9DDCLDEW'
secret_key = 'RRNJf2XP64k4L4108v1Fod4Y/I6iy0lFFEYiku3u'

#sh
#access_key = '9N6JSFR97Z2LL3F8KZAT'
#secret_key = 'ncLzrQRFY9qG0TLMglv1UwjpBl+DJohpdJ0ZQL8G'

conn = boto.connect_s3(
        aws_access_key_id = access_key,
        aws_secret_access_key = secret_key,
        host = 'xxxx',
        is_secure=False,
        calling_format = boto.s3.connection.OrdinaryCallingFormat(),
        )
bucket = conn.create_bucket('cn-sz-bucket1')
for bucket in conn.get_all_buckets():
        print "{name}\t{created}".format(
                name = bucket.name,
                created = bucket.creation_date,
        )
