#include <stdio.h>
#include <stdlib.h>
#include "rbd/librbd.h"

#define IMAGE_BUF_SIZE 4194304

int main(int argc, char* argv[])
{
	int err;
	rados_t cluster;
	rados_ioctx_t io;
	rbd_image_t image;
	char *poolname = "pool100";
	char buf[IMAGE_BUF_SIZE] = {0};

	err = rados_create(&cluster, NULL);
	if (err < 0) {
		fprintf(stderr, "%s: cannot create a cluster handle: %s\n", argv[0], strerror(-err));
		exit(1);
	}

	err = rados_conf_read_file(cluster, "/etc/ceph/ceph.conf");
	if (err < 0) {
		fprintf(stderr, "%s: cannot read config file: %s\n", argv[0], strerror(-err));
		exit(1);
	}

	err = rados_connect(cluster);
	if (err < 0) {
		fprintf(stderr, "%s: cannot connect to cluster: %s\n", argv[0], strerror(-err));
		exit(1);
	}


	err = rados_ioctx_create(cluster, poolname, &io);
	if (err < 0) {
		fprintf(stderr, "%s: cannot open rados pool %s: %s\n", argv[0], poolname, strerror(-err));
		rados_shutdown(cluster);
		exit(1);
	}


	err = rbd_open(io, "user1_image2", &image, NULL);
	if (err < 0){
		fprintf(stderr, "open image failed: %s\n", strerror(-err));
		goto out;
	}
	

	err = rbd_read(image, IMAGE_BUF_SIZE*2, IMAGE_BUF_SIZE, buf);
	if (err < 0) {
		fprintf(stderr, "%s: cannot read image: %s\n",  poolname, strerror(-err));
	}else{
		fprintf(stderr, "read image return :%d\n", err);
	}
out:
	rados_ioctx_destroy(io);
	rados_shutdown(cluster);
	return 0;
}

