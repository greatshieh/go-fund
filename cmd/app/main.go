package main

import (
	"context"
	"fmt"
	"gospider/app/fund"
	"gospider/global"
	"gospider/initialize"
	"gospider/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/docker/go-connections/nat"
	"go.uber.org/zap"
)

// 检查容器是否存在
func checkContainer(cli *client.Client, name string) (bool, string) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Names[0] == fmt.Sprintf("/%s", name) {
			return true, container.ID
		}
	}

	return false, ""

}

// 创建容器
func createContainer() (*client.Client, string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	isExited, contID := checkContainer(cli, "proxy-tool")

	if !isExited {
		//创建容器
		var openPort nat.Port = "5050"
		cont, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image:     "proxy-pool:v1", //镜像名称
			Tty:       true,            //docker run命令中的-t选项
			OpenStdin: true,            //docker run命令中的-i选项
			ExposedPorts: nat.PortSet{
				openPort: struct{}{}, //docker容器对外开放的端口
			},
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				openPort: []nat.PortBinding{{
					HostIP:   "127.0.0.1", //docker容器映射的宿主机的ip
					HostPort: "5050",      //docker 容器映射到宿主机的端口
				}},
			},
			Mounts: []mount.Mount{ //docker 容器目录挂在到宿主机目录
				{
					Type:   mount.TypeBind,
					Source: "/home/ivan/workspace/go-proxy-pool/proxy.db",
					Target: "/app/proxy.db",
				},
			},
		}, nil, nil, "proxy-tool")
		if err == nil {
			global.GPA_LOG.Info(fmt.Sprintf("\nsuccess create container:%s\n", cont.ID))
		} else if err == errdefs.Conflict(err) {
			global.GPA_LOG.Info("container is existed")
		} else {
			panic(err)
		}

		cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})

		return cli, cont.ID
	}

	global.GPA_LOG.Info("container is existed")
	return cli, contID

}

func removeContainer(cli *client.Client, conID string) {
	cli.ContainerStop(context.Background(), conID, nil)
	err := cli.ContainerRemove(context.Background(), conID, types.ContainerRemoveOptions{})
	if err != nil {
		panic(err)
	}

	global.GPA_LOG.Info("succe close container.")
}

func main() {
	global.GPA_VP = initialize.Viper() // 初始化Viper
	global.GPA_LOG = initialize.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GPA_LOG)

	cli, conID := createContainer()

	defer removeContainer(cli, conID)

	fileName := fund.Run()

	utils.SendEmail(fileName)
}
