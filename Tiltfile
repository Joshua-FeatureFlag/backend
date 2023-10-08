load("ext://helm_resource", "helm_resource")

docker_build("backend", ".")

image_deps = ["backend"]
image_keys = [("image.repository", "image.tag")]

helm_resource(
    "backend",
    "./helm",
    flags=[],
    image_deps=image_deps,
    image_keys=image_keys,
    port_forwards="50051:50051",
    labels=["services"],
)
