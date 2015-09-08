import logging
import urllib
import boto
import boto.s3.connection
#access_key = 'PC8URMEKC3IO1VS752KF'
#secret_key = 'gmcECdJiPkLQU07Ljm/iaWDQaYeLvkWyQAB+EBh3'

#sz
#access_key = '9HAGI3A40XXC9DDCLDEW'
#secret_key = 'RRNJf2XP64k4L4108v1Fod4Y/I6iy0lFFEYiku3u'

#sh
access_key = '9N6JSFR97Z2LL3F8KZAT'
secret_key = 'ncLzrQRFY9qG0TLMglv1UwjpBl+DJohpdJ0ZQL8G'

def url_safe(component):
    if isinstance(component, basestring):
        string = component.encode('utf8')
    else:
        string = str(component)
    return urllib.quote(string)

logging.basicConfig(filename="boto.log", level=logging.DEBUG)
conn = boto.s3.connection.S3Connection(
        aws_access_key_id = access_key,
        aws_secret_access_key = secret_key,
        host = 'cn-sz-radosgw-test1',
        is_secure=False,
        #calling_format = boto.s3.connection.OrdinaryCallingFormat(),
	calling_format = 'boto.s3.connection.OrdinaryCallingFormat'
        )
bucket_name = 'docker-image-bucket'
bucket = conn.get_bucket(bucket_name)
#key = bucket.get_key('_multipart_test/images/34e94e67e63a0f079d9336b3c2a52e814d138e5b3f1f614a0cfe273814ed7c0a/layer.2~BLpb_kJvDumawY0IvCH9sJg1oMi1yVJ.1')
#key = bucket.get_key('test/images/34e94e67e63a0f079d9336b3c2a52e814d138e5b3f1f614a0cfe273814ed7c0a/json')
#key.get_contents_to_filename('/tmp/data.txt')
obj_name='test/images/34e94e67e63a0f079d9336b3c2a52e814d138e5b3f1f614a0cfe273814ed7c0a/json'
#headers={ 'x-amz-copy-source': url_safe('%s/%s' % (bucket_name, obj_name)),}
key = bucket.copy_key('new_key', 'docker-image-bucket', obj_name, headers={ 'x-amz-copy-source': url_safe('%s/%s' % (bucket_name, obj_name)),})
key.get_contents_to_filename('/tmp/data.txt')
