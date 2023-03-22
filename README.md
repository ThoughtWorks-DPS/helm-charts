<div align="center">
	<p>
		<img alt="Thoughtworks Logo" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/thoughtworks_flamingo_wave.png?sanitize=true" width=200 />
    <br />
		<img alt="DPS Title" src="https://raw.githubusercontent.com/ThoughtWorks-DPS/static/master/EMPCPlatformStarterKitsImage.png?sanitize=true" width=350/>
	</p>
  <h3>EMPC Helm Chart Registry</h3>
</div>
<br />

## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to Helm's [documentation](https://helm.sh/docs) to get started.  

Once Helm has been set up correctly, add the repo as follows:  

```bash
$ helm repo add twdps https://thoughtworks-dps.github.io/helm-charts
```

If you had already added this repo earlier, run `helm repo update` to retrieve the latest versions of the packages. You can then run `helm search repo twdps` to see the charts.  

### example  

To install the opa-sidecar-admission-controller:  
```bash
$ helm install opa-sidecar-admission-controller twdps/opa-sidecar-admission-controller --namespace opa-system  
```
To uninstall the chart:  
```bash
$ helm uninstall opa-sidecar-admission-controller --namespace opa-system 
```


#### To Do

https://charts.bitnami.com/bitnami
azure-marketplace https://marketplace.azurecr.io/helm/v1/repo
