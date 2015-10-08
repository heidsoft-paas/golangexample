
import os
import platform
import sys
if os.path.exists('/usr/share/pyshared/radosgw_agent'):
    sys.path.insert(0,'/usr/share/pyshared/radosgw_agent')
elif os.path.exists('/usr/share/radosgw-agent'):
    sys.path.insert(0,'/usr/share/radosgw-agent')
elif os.path.exists('/usr/share/pyshared/radosgw-agent'):
    sys.path.insert(0,'/usr/share/pyshared/radosgw-agent')
elif os.path.exists('/usr/lib/python2.6/site-packages/radosgw_agent'):
    sys.path.insert(0,'/usr/lib/python2.6/site-packages/radosgw_agent')

import logging
import urllib
import boto
import boto.s3.connection


from radosgw_agent import client

#sz
access_key = '9HAGI3A40XXC9DDCLDEW'
secret_key = 'RRNJf2XP64k4L4108v1Fod4Y/I6iy0lFFEYiku3u'

#sh
#access_key = '9N6JSFR97Z2LL3F8KZAT'
#secret_key = 'ncLzrQRFY9qG0TLMglv1UwjpBl+DJohpdJ0ZQL8G'

def url_safe(component):
    if isinstance(component, basestring):
        string = component.encode('utf8')
    else:
        string = str(component)
    return urllib.quote(string)

logging.basicConfig(filename="boto.log", level=logging.DEBUG)
host='cn-sz-radosgw-test1'

dest=client.Endpoint(host, 80, False, access_key, secret_key, 'cn', 'cn-sz')
conn=client.connection(dest)
#conn = boto.s3.connection.S3Connection(
#        aws_access_key_id = access_key,
#        aws_secret_access_key = secret_key,
#        host = 'cn-sz-radosgw-test1',
#        is_secure=False,
        #calling_format = boto.s3.connection.OrdinaryCallingFormat(),
#	calling_format = 'boto.s3.connection.OrdinaryCallingFormat'
#        )
bucket_name = 'docker-image-bucket'
obj_name='test/images/34e94e67e63a0f079d9336b3c2a52e814d138e5b3f1f614a0cfe273814ed7c0a/json'
src_zone='cn-sh'
client_id='radosgw-agent'
op_id='cn-sh-radosgw-test1'
#client.sync_object_intra_region(conn, bucket_name, obj_name, src_zone, client_id, op_id)
path = u'{bucket}/{object}'.format(
        bucket=bucket_name,
        object=obj_name,
        )

params = {
        'rgwx-source-zone': src_zone,
        'rgwx-client-id': client_id,
        'rgwx-op-id': op_id,
    }


client.request(conn, 'put', path,
                   params=params,
                   headers={
                       'x-amz-copy-source': url_safe('%s/%s' % (bucket_name, obj_name)),
                       },
                   expect_json=False)
