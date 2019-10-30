package framework

import (
	"fmt"

	"stash.appscode.dev/stash/apis/stash/v1alpha1"
	"stash.appscode.dev/stash/apis/stash/v1beta1"

	"github.com/appscode/go/crypto/rand"
	. "github.com/onsi/ginkgo"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	store "kmodules.xyz/objectstore-api/api/v1"
)

func (f *Invocation) BackupBlueprint(secretName string) *v1beta1.BackupBlueprint {

	return &v1beta1.BackupBlueprint{
		ObjectMeta: metav1.ObjectMeta{
			Name: rand.WithUniqSuffix(f.app),
		},
		Spec: v1beta1.BackupBlueprintSpec{
			RepositorySpec: v1alpha1.RepositorySpec{
				Backend: store.Backend{
					S3: &store.S3Spec{
						Endpoint: f.MinioServiceAddres(),
						Bucket:   "minio-bucket",
						Prefix:   fmt.Sprintf("stash-e2e/%s/%s", f.Namespace(), f.App()),
					},
					StorageSecretName: secretName,
				},
				WipeOut: false,
			},
			Schedule: "*/59 * * * *",
			RetentionPolicy: v1alpha1.RetentionPolicy{
				Name:     "keep-last-5",
				KeepLast: 5,
				Prune:    true,
			},
		},
	}
}

func (f *Framework) CreateBackupBlueprint(backupBlueprint *v1beta1.BackupBlueprint) (*v1beta1.BackupBlueprint, error) {
	return f.StashClient.StashV1beta1().BackupBlueprints().Create(backupBlueprint)
}

func (f *Invocation) DeleteBackupBlueprint(name string) error {
	if name == "" {
		return nil
	}
	err := f.StashClient.StashV1beta1().BackupBlueprints().Delete(name, &metav1.DeleteOptions{})
	if kerr.IsNotFound(err) {
		return nil
	}
	return err
}

func (f *Framework) GetBackupBlueprint(name string) (*v1beta1.BackupBlueprint, error) {
	return f.StashClient.StashV1beta1().BackupBlueprints().Get(name, metav1.GetOptions{})
}

func (f Invocation) CreateBackupBlueprintForWorkload(name string) (*v1beta1.BackupBlueprint, error) {
	// Create Secret for BackupBlueprint
	secret, err := f.CreateBackendSecretForMinio()
	if err != nil {
		return &v1beta1.BackupBlueprint{}, err
	}

	// Generate BackupBlueprint definition
	bb := f.BackupBlueprint(secret.Name)
	bb.Name = name

	By(fmt.Sprintf("Creating BackupBlueprint: %s", bb.Name))
	createdBB, err := f.CreateBackupBlueprint(bb)
	f.AppendToCleanupList(createdBB)
	return createdBB, err
}

func (f Invocation) CreateBackupBlueprintForPVC(name string) (*v1beta1.BackupBlueprint, error) {
	// Create Secret for BackupBlueprint
	secret, err := f.CreateBackendSecretForMinio()
	if err != nil {
		return &v1beta1.BackupBlueprint{}, err
	}

	// Generate BackupBlueprint definition
	bb := f.BackupBlueprint(secret.Name)
	bb.Spec.Task.Name = TaskPVCBackup
	bb.Name = name

	By(fmt.Sprintf("Creating BackupBlueprint: %s", bb.Name))
	createdBB, err := f.CreateBackupBlueprint(bb)
	f.AppendToCleanupList(createdBB)
	return createdBB, err
}
