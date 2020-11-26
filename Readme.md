# CPF

Calcula o DV de um CPF dado os 9 numeros iniciais.

## Usando a API:

```shell
sudo docker build . -t cpfdv:1.0
sudo docker run --env PORT=8080 -p 8080:8080 cpfdv:1.0
```

No seu browser digite a URL:

``http://127.0.0.1:8080/cpfdv/351172282

## Publicando como Google Function

Eu uso docker ao inves de instalar o gcloud. Veja aqui como fazer: 

https://cloud.google.com/sdk/docs/downloads-docker

```shell
sudo docker run --rm --volumes-from gcloud-config -v /CODE_DIR/cpf_dv:/src gcr.io/google.com/cloudsdktool/cloud-sdk:latest gcloud functions deploy CpfFull --runtime go113 --set-env-vars ENTRY_POINT="/" --trigger-http --allow-unauthenticated --source /src/.
```

Se tudo deu certo, use a URL que pode ser encontrada fazendo:

```shell
sudo docker run --rm --volumes-from gcloud-config gcr.io/google.com/cloudsdktool/cloud-sdk:latest gcloud functions describe CpfFull
```
