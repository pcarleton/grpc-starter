# grpc-starter

This is a starter project that can be used to create a gRPC api server.  The features of this server are:

* Authentication via Google JWT
* Containerized
* Bazel-enabled
* Terraform for a micro server instance
* TODO: Let's Encrypt support


# Getting Started

## Local Development

To run locally, first [install bazel](https://docs.bazel.build/versions/master/install.html).  Then run:

```
# In one shell Start the server
# (This will probably take awhile to build the first time)
bazel run //server/bin/run_server -- -insecure


# Start the client
bazel run //server/bin/run_client -- -insecure -address localhost:5001
```
