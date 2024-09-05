package k8s

import (
	"context"
)

// namespace: openiiot, openiiot-dt, openiiot-tenant1
// volume dir: /volumes/openiiot/dt. /volumes/openiiot/tenant1
func (a *K8sService) CreateMysqlComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateServerComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateConsolemanagerComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateNoderedComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}
	/*if err := a.CreateJob(ctx, k8sUns); err != nil {
		return err
	}*/
	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateIngress(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateGrafanaComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}
	/*if err := a.CreateJob(ctx, k8sUns); err != nil {
		return err
	}*/
	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateIngress(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateEmqxComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateIngress(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *K8sService) CreateTdengineComponent(ctx context.Context, k8sUns K8sUns) error {

	k8sUns.Tag = "data"
	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns.Tag = "log"
	if err := a.CreatePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreatePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns.Tag = ""
	if err := a.CreateDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.CreateService(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

// Delete use reverse order of creation, job will auto delete after TTLSecondsAfterFinished time, so no need clean job
func (a *K8sService) DeleteComponent(ctx context.Context, k8sUns K8sUns) error {

	if err := a.DeleteIngress(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.DeleteService(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.DeleteDeployment(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.DeletePersistentVolumeClaim(ctx, k8sUns); err != nil {
		return err
	}
	if err := a.DeletePersistentVolume(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}
