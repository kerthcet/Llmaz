# ModelLoader

ModelLoader maintains the codes to load model weights with various ways, such as from huggingface or from s3.

## Load Models Via ModelHub

The full list of supported model hubs:

- [Huggingface](https://huggingface.co/welcome)
- [ModelScope](https://www.modelscope.cn/home)

### How to use

The loader will be build as an image to provide services, the logic is handled by the model source respectively if the model source refers to modelHub.

### How to build

- `make modelhub-image-load` will build the image locally.
- `make modelhub-image-push` will push the image to the registry, default to `inftyai/model-loader`.

To build an official image, try to run `make modelhub-image-push LOADER_IMAGE_TAG=v0.0.x-modelhub`.

## Load Models From ObjectStore

The full list of supported object stores:

- AliYun OSS
- AWS S3
- Azure Storage Account
- Google Cloud Storage
- MinIO or other S3-compatible object storages
- Tencent COS

### How to use

The loader will be build as an image to provide services, the logic is handled by the model source respectively if the model source refers to URI starts with `objstore://`

### How to build

- `make objstore-image-load` will build the image locally.
- `make objstore-image-push` will push the image to the registry, default to `inftyai/model-loader:<version>-objstore`.

To build an official image, try to run `make objstore-image-push LOADER_IMAGE_TAG=v0.0.x-objstore`.
