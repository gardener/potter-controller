package testing

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"

	landscaper "github.com/gardener/landscaper/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	appRepov1 "github.com/gardener/potter-controller/api/external/apprepository/v1alpha1"
	hubv1 "github.com/gardener/potter-controller/api/v1"
)

type ReactorFuncs map[string]func() error

type ReactiveMockClient struct {
	FakeClient   client.Client
	StatusWriter *ReactiveMockStatusWriter
	ReactorFuncs
}

func NewReactiveMockClient(funcs map[string]func() error, initObjs ...runtime.Object) ReactiveMockClient {
	scheme := runtime.NewScheme()

	_ = hubv1.AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)
	_ = appRepov1.AddToScheme(scheme)
	_ = landscaper.AddToScheme(scheme)

	fakeClient := fake.NewFakeClientWithScheme(scheme, initObjs...) // nolint

	return ReactiveMockClient{
		FakeClient:   fakeClient,
		StatusWriter: &ReactiveMockStatusWriter{funcs, fakeClient.Status()},
		ReactorFuncs: funcs,
	}
}

func (t *ReactiveMockClient) Scheme() *runtime.Scheme {
	return nil
}

func (t *ReactiveMockClient) RESTMapper() meta.RESTMapper {
	return nil
}

// Get retrieves an obj for the given object key from the Kubernetes Cluster.
// obj must be a struct pointer so that obj can be updated with the response
// returned by the Server.
func (t *ReactiveMockClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	fun := t.ReactorFuncs[key.String()]
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Get(ctx, key, obj)
}

// List retrieves list of objects for a given namespace and list options. On a
// successful call, Items field in the list will be populated with the
// result returned from the server.
func (t *ReactiveMockClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	fun := getReactorFuncForObjectList(list, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.List(ctx, list, opts...)
}

func (t *ReactiveMockClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Create(ctx, obj, opts...)
}

// Delete deletes the given obj from Kubernetes cluster.
func (t *ReactiveMockClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Delete(ctx, obj, opts...)
}

// Update updates the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (t *ReactiveMockClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Update(ctx, obj, opts...)
}

// Patch patches the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (t *ReactiveMockClient) Patch(ctx context.Context, obj client.Object, patch client.Patch,
	opts ...client.PatchOption) error { // nolint
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Patch(ctx, obj, patch, opts...)
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (t *ReactiveMockClient) DeleteAllOf(ctx context.Context, obj client.Object,
	opts ...client.DeleteAllOfOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.DeleteAllOf(ctx, obj, opts...)
}

type ReactiveMockStatusWriter struct {
	ReactorFuncs
	FakeClient client.StatusWriter
}

func (t *ReactiveMockClient) Status() client.StatusWriter { // nolint
	return t.StatusWriter
}

func (t *ReactiveMockStatusWriter) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Update(ctx, obj, opts...)
}

// Patch patches the given object's subresource. obj must be a struct
// pointer so that obj can be updated with the content returned by the
// Server.
func (t *ReactiveMockStatusWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	fun := getReactorFuncForObject(obj, &t.ReactorFuncs)
	if fun != nil {
		return fun()
	}

	return t.FakeClient.Patch(ctx, obj, patch, opts...)
}

func getReactorFuncForObject(obj client.Object, reactorFuncs *ReactorFuncs) func() error {
	key := client.ObjectKeyFromObject(obj)

	if fun, ok := (*reactorFuncs)[key.String()]; ok {
		return fun
	}
	return nil
}

func getReactorFuncForObjectList(obj client.ObjectList, reactorFuncs *ReactorFuncs) func() error {
	return nil
}
