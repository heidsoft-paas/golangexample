#include <stdio.h>
#include <stdlib.h>
#include "rados/librados.h"
int main(int argc, char* argv[])
{
	int err;
	rados_t cluster;
	rados_ioctx_t io;
	char *poolname = "pool100";

	char *buf, *st;
  	size_t buflen, stlen;
  	char *cmd[2];

  	cmd[1] = NULL;

  	cmd[0] = (char *)"{\"prefix\":\"get_command_descriptions\"}";

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

	err = rados_mon_command(cluster, (const char **)cmd, 1, "", 0, &buf, &buflen, &st, &stlen);
	if (err < 0){
		fprintf(stderr, "run cmd failed");
	}else{
		fprintf(stderr, "%s\n", buf);
	}
  	rados_buffer_free(buf);
  	rados_buffer_free(st);

/*
	err = rados_ioctx_create(cluster, poolname, &io);
	if (err < 0) {
		fprintf(stderr, "%s: cannot open rados pool %s: %s\n", argv[0], poolname, strerror(-err));
		rados_shutdown(cluster);
		exit(1);
	}

	rados_ioctx_destroy(io);*/

	rados_shutdown(cluster);
	return 0;
}

