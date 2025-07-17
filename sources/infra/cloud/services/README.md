# K8s Manifests

Considering there are only two deployments, I figured I could keep
this simple.

I use [`envsusbt`](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html)
to populate YAML templates with values from the environment, and then
deploy using `kubectl`.

## Usage

Setup a `.env` file using the [`env.template`](./env.template) reference.

```sh
# Modify values as needed
cp env.template .env
```

Update K8s after template substitution, for example:

```sh
source .env && envsubst < common.yml | kubectl apply -f -
```
