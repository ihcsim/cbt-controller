## CBT Controller

This repository contains an [aggregated API server] prototype used to serve the
[CSI changed block tracking] API.

The primary goal is to explore ways to implement an in-cluster API endpoint
which can be used to retrieve a long list of changed block entries (in the order
of hundreds of MiB), without putting the Kubernetes API server and etcd in the
data retrieval path.

This prototype explores:

* The implementation of a non-etcd [`custom` storage] that supports the
Kubernetes [`rest.Connecter`] interface
* The implementation of a [custom `Option`] that converts `url.Values` to a
`runtime.Object`

The `rest.Connecter` implements a modified `GET` handler to return a collection
of changed block entries, if the `fetchcbd` query parameter is defined.

Essentially, the default API endpoint will return a `VolumeSnapshotDelta` custom
resource:

```sh
curl "http://127.0.0.1:8001/apis/cbt.storage.k8s.io/v1alpha1/namespaces/default/volumesnapshotdelta/test-delta" | jq .
{
  "kind": "VolumeSnapshotDelta",
  "apiVersion": "cbt.storage.k8s.io/v1alpha1",
  "metadata": {
    "name": "test-delta",
    "namespace": "default",
    "creationTimestamp": null
  },
  "spec": {
    "baseVolumeSnapshotName": "base",
    "targetVolumeSnapshotName": "target",
    "mode": "block"
  },
}
```

Appending the API endpoint with the `fetchcbd=true` query parameter will return
the list of changed block entries:

```sh
curl "http://127.0.0.1:8001/apis/cbt.storage.k8s.io/v1alpha1/namespaces/default/volumesnapshotdelta/test-delta?fetchcbd-true&limit=256&offset=0" | jq .
{
  "kind": "VolumeSnapshotDelta",
  "apiVersion": "cbt.storage.k8s.io/v1alpha1",
  "metadata": {
    "name": "test-delta",
    "namespace": "default",
    "creationTimestamp": null
  },
  "spec": {
    "baseVolumeSnapshotName": "base",
    "targetVolumeSnapshotName": "target",
    "mode": "block"
  },
  "status": {
    "ChangedBlockDeltas": [
      {
        "offset": 0,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "ieEEQ9Bj7E6XR",
          "issuanceTime": "2022-07-13T03:19:30Z",
          "ttl": "3h0m0s"
        }
      },
      {
        "offset": 1,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "widvSdPYZCyLB",
          "issuanceTime": "2022-07-13T03:19:30Z",
          "ttl": "3h0m0s"
        }
      },
      {
        "offset": 2,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "VtSebH83xYzvB",
          "issuanceTime": "2022-07-13T03:19:30Z",
          "ttl": "3h0m0s"
        }
      }
    ]
  }
}
```

Most of the setup code of the aggregated API server is generated using the
[`apiserver-builder`] tool.

## Quick Start

Setup and connect to a Kubernetes cluster.

Create the `csi-cbt` namespace:

```sh
kubectl create ns csi-cbt
```

Deploy `etcd`:

```sh
make etcd
```

Deploy the CBT aggregated API server and mock HTTP and GRPC servers:

```sh
make deploy
```

The Docker images are hosted on public repositories at `quay.io/isim`.

## Development

Install the `apiserver-builder` tool following the instructions
[here](https://github.com/kubernetes-sigs/apiserver-builder-alpha#installation).
The `apiserver-boot` tool requires the code to be checked out into the local
`$GOPATH` i.e. `github.com/ihcsim/cbt-aggapi`.

To run the tests:

```sh
go test ./...
```

To work with the Docker images, first define the repository URL and tag for your
images:

```sh
export IMAGE_REPO_AGGAPI=<your_agg_apiserver_image_repo>
export IMAGE_TAG_AGGAPI=<your_agg_apiserver_image_tag>
export IMAGE_REPO_GRPC=<your_agg_apiserver_image_repo>
export IMAGE_TAG_GRPC=<your_agg_apiserver_image_tag>
```

Then use these `make` targets to build and push the images:

```sh
make image

make push
```

### Working With The Custom Resource

Create a `VolumeSnapshotDelta` resource:

```sh
cat<<EOF | kubectl apply -f -
apiVersion: cbt.storage.k8s.io/v1alpha1
kind: VolumeSnapshotDelta
metadata:
  name: test-delta
  namespace: default
spec:
  baseVolumeSnapshotName: vs-00
  targetVolumeSnapshotName: vs-01
  mode: block
EOF
```

Use `kubectl` to `GET` the resources:

```sh
kubectl get volumesnapshotdelta test-delta -oyaml
```

```yaml
apiVersion: cbt.storage.k8s.io/v1alpha1
kind: VolumeSnapshotDelta
metadata:
  creationTimestamp: null
  name: test-delta
  namespace: default
spec:
  baseVolumeSnapshotName: base
  mode: block
  targetVolumeSnapshotName: target
status: {}
```

Use the `kubectl proxy` to start a proxy between localhost and cluster:

```sh
kubectl proxy &
```

Get the changed block entries of the resource:

```sh
curl -k "http://127.0.0.1:8001/apis/cbt.storage.k8s.io/v1alpha1/namespaces/default/volumesnapshotdelta/test-delta?fetchcbd=true&limit=256&offset=0"
```

```json
{
  "kind": "VolumeSnapshotDelta",
  "apiVersion": "cbt.storage.k8s.io/v1alpha1",
  "metadata": {
    "name": "test-delta",
    "namespace": "default",
    "creationTimestamp": null
  },
  "spec": {
    "baseVolumeSnapshotName": "base",
    "targetVolumeSnapshotName": "target",
    "mode": "block"
  },
  "status": {
    "changedBlockDeltas": [
      {
        "offset": 0,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "ieEEQ9Bj7E6XR",
          "issuanceTime": "2022-07-13T03:30:46Z",
          "ttl": "3h0m0s"
        }
      },
      {
        "offset": 1,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "widvSdPYZCyLB",
          "issuanceTime": "2022-07-13T03:30:46Z",
          "ttl": "3h0m0s"
        }
      },
      {
        "offset": 2,
        "blockSizeBytes": 524288,
        "dataToken": {
          "token": "VtSebH83xYzvB",
          "issuanceTime": "2022-07-13T03:30:46Z",
          "ttl": "3h0m0s"
        }
      }
    ]
  }
}
```

### Re-generate Code and YAML

The Kubernetes YAML manifests are found in the `yaml` folder. To re-generate them:

```sh
make yaml
```

This will output all the manifests into the `yaml-generated` folder. Move what is
needed into the `yaml` folder.

## License

Apache License 2.0, see [LICENSE].

[aggregated API server ]:https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/
[CSI changed block tracking]: https://github.com/kubernetes/enhancements/pull/3367
[`rest.Connecter`]: https://pkg.go.dev/k8s.io/apiserver/pkg/registry/rest#Connecter
[`custom` storage]: pkg/storage/custom.go
[custom `Option`]: pkg/apis/cbt/v1alpha1/volumesnapshotdeltaoption_types.go
[`apiserver-builder`]: https://github.com/kubernetes-sigs/apiserver-builder-alpha
[LICENSE]: LICENSE
