step by step:

```
mkdir -p $GOPATH/src/github.com/app-sre/
cd $GOPATH/src/github.com/app-sre/
export GO111MODULE=on
operator-sdk new qontract-operator
cd qontract-operator

operator-sdk add api --api-version=qontract.openshift.io/v1alpha1 --kind=QontractApply

operator-sdk add controller --api-version=qontract.openshift.io/v1alpha1 --kind=QontractApply

go mod vendor
operator-sdk build quay.io/jmelis/qontract-operator
docker push quay.io/jmelis/qontract-operator
sed -i 's|REPLACE_IMAGE|quay.io/jmelis/qontract-operator|g' deploy/operator.yaml
```
