#Download code-generator repository.
git clone https://github.com/kubernetes/code-generator.git

##Use generate-groups.sh to generate 
../code-generator/generate-groups.sh all k8s-programming-tutorial/pkg/client k8s-programming-tutorial/pkg/apis k8s.ovn.org:v1alpha1 --go-header-file ../code-generator/hack/boilerplate.go.txt --output-base $(pwd)/../

## Generate List
1. Deep copy objects
2. clientset
3. informers
4. lister