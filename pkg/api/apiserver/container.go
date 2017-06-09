package apiserver

func QueryContainers(containerName string, pageCnt, pageNum int, serviceId uint) (list []*Container, total int) {

	if serviceId == 0 {
		if containerName != "" {
			db.Where("name like ? ", `%`+containerName+`%`).Offset(pageCnt * pageNum).Limit(pageCnt).Order("name desc").Find(&list)
			db.Model(new(Service)).Where("name like ?", containerName).Count(&total)
		} else {
			db.Offset(pageCnt * pageNum).Limit(pageCnt).Order("name desc").Find(&list)
			db.Model(new(Service)).Count(&total)
		}
	} else {
		if containerName != "" {
			db.Where("name like ? and service_id=?", `%`+containerName+`%`, serviceId).Offset(pageCnt * pageNum).Limit(pageCnt).Order("name desc").Find(&list)
			db.Model(new(Service)).Where("name like ? and service_id=?", `%`+containerName+`%`, serviceId).Count(&total)
		} else {
			db.Where("service_id=?", serviceId).Offset(pageCnt * pageNum).Limit(pageCnt).Order("name desc").Find(&list)
			db.Model(new(Service)).Where("service_id=?", serviceId).Count(&total)
		}
	}

	for _, container := range list {
		var (
			config      = &ContainerConfig{}
			base        = &BaseConfig{}
			configmap   = &ConfigMap{}
			superConfig = &SuperConfig{}
			volumes     []*Volume
			envs        []*Env
			ports       []*Port
		)
		db.Find(config, ContainerConfig{ContainerId: container.ID})
		container.Config = config

		db.First(base, BaseConfig{ServiceConfigId: config.ID})
		db.Find(&volumes, Volume{BaseConfigId: base.ID})
		base.Volumes = volumes
		config.BaseConfig = base

		db.First(configmap, ConfigMap{ServiceConfigId: config.ID})
		config.ConfigMap = configmap

		db.First(superConfig, SuperConfig{ServiceConfigId: config.ID})
		db.Find(&envs, Env{SuperConfigId: superConfig.ID})
		db.Find(&ports, Port{SuperConfigId: superConfig.ID})
		superConfig.Envs = envs
		superConfig.Ports = ports
		config.SuperConfig = superConfig
	}
	return
}

func QueryContainerById(id uint) *Container {
	container := &Container{}
	db.First(container, id)
	return container
}

func UpdateContainer(container *Container) {
	db.Model(new(Container)).Update(container)
}

func DeleteContainer(container *Container) {
	db.Delete(container, "service_id=?", container.ServiceId)
}

/*
//we not plan to provide those api of the container
func InsertContainer(container *Container) {}

func DeleteContainer(container *Container) {}

func UpdateContainer(container *Container) {}

func QueryContainerById(id uint) {}
*/
